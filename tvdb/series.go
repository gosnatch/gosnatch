package tvdb

import (
    "encoding/xml"
    "errors"
    "fmt"
    "net/http"
    "net/url"
    "strings"
)

// Series, as defined by TVDB.
type Series struct {
    XMLName       xml.Name `xml:"Series"`
    Id            int      `xml:"id"`
    Actors        string   `xml:"Actors"`
    AirDay        string   `xml:"Airs_DayOfWeek"`
    AirTime       string   `xml:"Airs_Time"`
    ContentRating string   `xml:"ContentRating"`
    FirstAired    string   `xml:"FirstAired"`
    genres        string   `xml:"Genre"`
    Genres        []string `xml:"-"`
    ImdbId        string   `xml:"IMDB_ID"`
    Language      string   `xml:"Language"`
    Network       string   `xml:"Network"`
    NetworkId     string   `xml:"NetworkID"`
    Overview      string   `xml:"Overview"`
    Rating        float64  `xml:"Rating"`
    RatingCount   int      `xml:"RatingCount"`
    Runtime       int      `xml:"Runtime"`
    SeriesName    string   `xml:"SeriesName"`
    Status        string   `xml:"Status"`
    Added         string   `xml:"added"`
    Banner        string   `xml:"banner"`
    FanArt        string   `xml:"fanart"`
    Year          int      `xml:"year"`
    Poster        string   `xml:"poster"`
    LastUpdated   string   `xml:"lastupdated"`
    TmsWantedOld  int      `xml:"tms_wanted_old"`
    Zap2idId      string   `xml:"zap2it_id"`
    AddedBy       string   `xml:"addedBy"`
    TvrageId      int      `xml:"tvrage_id"`
    Banners       []Banner `xml:"-"`
    // SeriesId    int    `xml:"SeriesID"`
}

type Banner struct {
    XMLName    xml.Name `xml:"Banner"`
    Id         int      `xml:"id"`
    BannerPath string   `xml:"BannerPath"`
    BannerType string   `xml:"BannerType"`
    Language   string   `xml:"Language"`
}

type Banners struct {
    Banner []Banner
}

func (s *Series) getBanners(t *TVDB, id int) {
    url := fmt.Sprintf("%s/series/%d/banners.xml", t.baseUrl(), id)

    r, err := http.Get(url)
    if err != nil {
    }
    defer r.Body.Close()

    var d Banners
    err = xml.NewDecoder(r.Body).Decode(&d)
    if err != nil {
    }
    for _, b := range d.Banner {
        if b.BannerType == "poster" {
            s.Poster = b.BannerPath
        }
        if b.BannerType == "fanart" {
            s.FanArt = b.BannerPath
        }
    }
    s.Banners = d.Banner
}

func (s *Series) postProc() {
    s.Genres = strings.Split(s.genres, "|")
}

type SeriesData struct {
    XMLName xml.Name  `xml:"Data"`
    Series  []*Series `xml:"Series"`
}

func (t *TVDB) GetSeriesById(id int, language string) (*Series, error) {
    url := fmt.Sprintf("%s/series/%d/", t.baseUrl(), id)
    if language != "" {
        url = fmt.Sprintf("%s%s.xml", url, language)
    }

    r, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer r.Body.Close()

    var d *SeriesData
    err = xml.NewDecoder(r.Body).Decode(&d)
    if err != nil {
        return nil, err
    }

    var s *Series
    if len(d.Series) > 0 {
        s = d.Series[0]
        s.postProc()
    }

    return s, nil
}

func (t *TVDB) GetSeriesByRemoteId(imdbid string, language string) (*Series, error) {

    args := &url.Values{}
    args.Add("imdbid", imdbid)
    if language != "" {
        args.Add("language", language)
    }
    url := fmt.Sprintf("%s/api/GetSeriesByRemoteID.php?%s", t.Location, args.Encode())
    r, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer r.Body.Close()

    var d *SeriesData
    err = xml.NewDecoder(r.Body).Decode(&d)
    if err != nil {
        return nil, err
    }

    var s *Series
    if len(d.Series) > 0 {
        s = d.Series[0]
        s.postProc()
        s.getBanners(t, s.Id)
    }

    return s, nil
}

func (t *TVDB) GetSeries(name, language string) ([]*Series, error) {
    if name == "" {
        return nil, errors.New("name must not be empty")
    }

    args := &url.Values{}
    args.Add("seriesname", name)
    if language != "" {
        args.Add("language", language)
    }

    r, err := http.Get(fmt.Sprintf("%s/api/GetSeries.php?%s", t.Location, args.Encode()))
    if err != nil {
        return nil, err
    }
    defer r.Body.Close()

    var d *SeriesData
    err = xml.NewDecoder(r.Body).Decode(&d)
    if err != nil {
        return nil, err
    }

    if len(d.Series) > 0 {
        for _, s := range d.Series {
            s.postProc()
            s.getBanners(t, s.Id)
        }
    }

    return d.Series, nil
}
