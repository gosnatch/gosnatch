package tvdb

import (
    "encoding/xml"
    "net/http"
    "strconv"
)

// Episode, as defined by TVDB.
type Episode struct {
    Id                    int    `xml:"id"`
    SeasonId              int    `xml:"seasonid"`
    CombinedEpisodeNumber string `xml:"Combined_episodenumber"`
    CombinedSeason        string `xml:"Combined_season"`
    EpisodeName           string
    EpisodeNumber         int
    SeasonNumber          int
    FirstAired            string
    Director              string
    Writer                string
    GuestStars            string
    Overview              string
    ProductionCode        string
    Flagged               int    `xml:"flagged"`
    DvdDiscId             string `xml:"DVD_discid"`
    DvdSeason             string `xml:"DVD_season"`
    DvdEpisodeNumber      string `xml:"DVD_episodenumber"`
    DvdChapter            string `xml:"DVD_chapter"`
    AbsoluteNumber        string `xml:"absolute_number"`
    Filename              string `xml:"filename"`
    SeriesId              int    `xml:"seriesid"`
    MirrorUpdate          string `xml:"mirrorupdate"`
    ImdbId                string `xml:"IMDB_ID"`
    EpImgFlag             string
    Rating                string
    Language              string
}

// Single episode data.
type SingleEpisodeData struct {
    XMLName xml.Name `xml:"Data"`
    Episode *Episode
}

// Full series data - Series and list of episodes.
type FullSeriesData struct {
    XMLName xml.Name `xml:"Data"`
    Series  Series
    Episode []Episode
}

// For seriesId, season number, and episode number, return SingleEpisodeData.
func (t *TVDB) GetEpisodeBySeasonEp(seriesId, season, episode int, language string) (*SingleEpisodeData, error) {
    url := t.baseUrl() + "/series/" + strconv.Itoa(seriesId) + "/default/" + strconv.Itoa(season) +
        "/" + strconv.Itoa(episode) + "/" + language + ".xml"

    r, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer r.Body.Close()

    var d *SingleEpisodeData
    err = xml.NewDecoder(r.Body).Decode(&d)
    if err != nil {
        return nil, err
    }

    return d, nil
}

// Parse SingleEpisodeData.
func ParseSingleEpisode(src []byte) (*SingleEpisodeData, error) {
    var r SingleEpisodeData
    if err := xml.Unmarshal(src, &r); err != nil {
        return nil, err
    }
    return &r, nil
}

func ParseEpisode(src []byte) (*Episode, error) {
    var r Episode
    if err := xml.Unmarshal(src, &r); err != nil {
        return nil, err
    }
    return &r, nil
}

func (t *TVDB) GetFullSeriesData(seriesId int, language string) (*FullSeriesData, error) {
    url := t.baseUrl() + "/series/" + strconv.Itoa(seriesId) + "/all/" + language + ".xml"
    r, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer r.Body.Close()

    var d *FullSeriesData
    err = xml.NewDecoder(r.Body).Decode(&d)

    if err != nil {
        return nil, err
    }

    // var s *FullSeriesData
    // // if len(d.Series) > 0 {
    // //     s = d.Series[0]
    // //     s.postProc()
    // // }

    return d, nil
}
