package gosnatch

import (
    "encoding/json"
    "fmt"
    "github.com/astaxie/beego/orm"
    "github.com/dustin/go-humanize"
    "github.com/extemporalgenome/slug"
    "github.com/jinzhu/now"
    _ "github.com/mattn/go-sqlite3"
    log "github.com/sirupsen/logrus"
    "github.com/spf13/viper"
    "os"
    "path"
    "strconv"
    "strings"
    "time"
)

// Contains Presets to Qualitys
type QualityPreset struct {
    Id                int    `orm:"auto"`
    Name              string `orm:"null"`
    Qualitys          string `orm:"null"`
    Language          string `orm:"null"`
    RequiredWords     string `orm:"null"`
    IgnoredWords      string `orm:"null"`
    ArchiveFirstmatch bool   `orm:"null"`
}

func (q *QualityPreset) Unmarshal() []QualityDefinitions {
    qualitys := []QualityDefinitions{}
    err := json.Unmarshal([]byte(q.Qualitys), &qualitys)
    if err != nil {
        return []QualityDefinitions{}
    }
    return qualitys
}

type ShowExceptions struct {
    Id       int     `orm:"auto"`
    Series   *TvShow `orm:"null;rel(fk)"`
    TvdbId   int     `orm:"null"`
    ShowName string  `orm:"null"`
}

// Default TvShows holder
type TvShow struct {
    Id                int               `orm:"auto"`
    TvdbId            int               `orm:"null"`
    TvrageId          int               `orm:"null"`
    TraktId           int               `orm:"null"`
    ShowName          string            `orm:"null"`
    Location          string            `orm:"null"`
    Network           string            `orm:"null"`
    Genre             string            `orm:"null"`
    Classification    string            `orm:"null"`
    Overview          string            `orm:"null"`
    Runtime           int               `orm:"null"`
    Airs              string            `orm:"null"`
    AirTime           string            `orm:"null"`
    Status            string            `orm:"null"`
    Paused            string            `orm:"null"`
    FlattenFolders    int               `orm:"null"`
    Startyear         int               `orm:"null"`
    Language          string            `orm:"size(2)"`
    NotifyList        string            `orm:"null"`
    ImdbId            string            `orm:"null"`
    LastUpdateIndexer int               `orm:"null"`
    IsAnime           bool              `orm:"null"`
    IsScene           bool              `orm:"null"`
    DefaultEpStatus   int               `orm:"null"`
    Slug              string            `orm:"null"`
    Path              string            `orm:"null"`
    Monitored         bool              `orm:"null"`
    QualityPreset     *QualityPreset    `orm:"null;rel(fk)"`
    Episodes          []*TvEpisode      `orm:"null;reverse(many)"`
    AltNames          []*ShowExceptions `orm:"null;reverse(many)"`
}

// get Exceptions from database
func (show *TvShow) getExceptions() []ShowExceptions {
    altNames := []ShowExceptions{}
    o := orm.NewOrm()
    o.QueryTable(&ShowExceptions{}).Filter("tvdb_id", show.TvdbId).All(&altNames)
    return altNames
}

// create show direcotry
func (show *TvShow) makeShowDir() error {
    seriesdir, _ := path.Split(show.Path)
    if !fileExists(show.Path) {
        fi, err := os.Stat(seriesdir)
        if err != nil {
            log.Errorf("os.Stat -> %v", err)
            return err
        } else {
            log.WithFields(log.Fields{"path": show.Path, "permissions": fi.Mode()}).Debug("creating show dir")
            os.MkdirAll(show.Path, fi.Mode())
            return nil
        }
    } else {
        log.Debug("show dir already existing")
    }
    return nil
}

// update show informations from tvdb
func (show TvShow) UpdateShowInfo() {
    log.Debugf("starting show update for %s", show.ShowName)
    o := orm.NewOrm()

    t := getTvdb()
    s, _ := t.GetFullSeriesData(show.TvdbId, show.Language)

    show.ShowName = s.Series.SeriesName
    show.Overview = s.Series.Overview
    show.Airs = s.Series.AirDay
    show.AirTime = s.Series.AirTime
    show.Network = s.Series.Network
    show.Status = s.Series.Status

    for _, e := range s.Episode {
        ck := TvEpisode{}
        o.QueryTable(&TvEpisode{}).Filter("season", e.SeasonNumber).Filter("episode", e.EpisodeNumber).Filter("tv_show_id", show.Id).One(&ck)

        if ck.Id != 0 {
            ck.Name = e.EpisodeName
            ck.Overview = e.Overview
            airtime, err := time.Parse("2006-01-02", e.FirstAired)
            if err == nil {
                ck.FirstAired = airtime
            }
            o.Update(&ck)
        } else {
            ck.Name = e.EpisodeName
            ck.Episode = e.EpisodeNumber
            ck.Season = e.SeasonNumber
            ck.Overview = e.Overview
            ck.EpisodeId = e.Id
            ck.ShowId = e.SeriesId
            ck.TvShow = &show

            o.Insert(&ck)

            airtime, err := time.Parse("2006-01-02", e.FirstAired)
            if err == nil {
                ck.FirstAired = airtime
            }
        }
    }

    o.Update(&show)
    log.Infof("show update for '%s' complete", show.ShowName)
}

// returns the humanized size of all episodes in a season
func (show TvShow) HumanizeSize() string {
    o := orm.NewOrm()
    o.LoadRelated(&show, "Episodes")
    var all uint64 = 0
    for _, ep := range show.Episodes {
        all = all + ep.FileSize
    }
    return humanize.Bytes(all)
}

// scan show directory for media files
func (show *TvShow) scanDirectory() {

    o := orm.NewOrm()
    showPath := show.Path
    if fileExists(show.Path) {
        log.Infof("scanning directory %s for files", show.Path)
        chann := parseFilesInDir(showPath)
        for x := range chann {
            s, _ := strconv.Atoi(x.Season)
            e, _ := strconv.Atoi(x.Episode)

            ep := TvEpisode{}
            o.QueryTable(&TvEpisode{}).Filter("episode", e).Filter("season", s).Filter("tv_show_id", show.Id).One(&ep)
            ep.FileSize = uint64(x.FileSize)
            ep.Status = 2
            ep.Path = x.Path
            ep.Quality = x.Quality
            o.Update(&ep)
        }
    } else {
        log.Errorf("unable to access directory %s", show.Path)
    }

}

// return int quality values to accept
func (show TvShow) AcceptedQualitys() []int {
    quali := []int{}
    x := []QualityDefinitions{}
    o := orm.NewOrm()
    o.LoadRelated(&show, "QualityPreset")
    _ = json.Unmarshal([]byte(show.QualityPreset.Qualitys), &x)

    for _, y := range x {
        quali = append(quali, y.Quality)
    }
    return quali
}

// calculate the downloaded percent value
func (show TvShow) Percent() float64 {
    downloaded := show.Downloaded()
    wanted := show.EpisodeCount()
    var x float64 = (float64(len(downloaded)) / float64(wanted) * 100.0)
    // var x float64 = float64((float64(len(downloaded)) / (float64(len(downloaded)) + len(wanted))) * 100.0)
    // fmt.Printf("downloaded %d, wanted %d, gesammt %d, prozent %d", len(downloaded), len(wanted), len(downloaded)+len(wanted), x)
    return x
}

// list of downloaded episodes
func (show *TvShow) Downloaded() []TvEpisode {
    downloaded := []TvEpisode{}
    o := orm.NewOrm()
    o.QueryTable(&TvEpisode{}).Filter("tv_show_id", &show.Id).Filter("status__gte", 2).Exclude("first_aired__isnull", true).Filter("first_aired__lt", now.EndOfDay()).All(&downloaded)
    return downloaded
}

// list of snatched episodes
func (show *TvShow) Snatched() []TvEpisode {
    episodes := []TvEpisode{}
    o := orm.NewOrm()
    o.QueryTable(&TvEpisode{}).Filter("tv_show_id", &show.Id).Filter("status", 1).All(&episodes)
    return episodes
}

// count of all Episodes which are already aired and are not in season 0 (Specials)
func (show *TvShow) EpisodeCount() int {
    episodes := []TvEpisode{}
    o := orm.NewOrm()
    o.QueryTable(&TvEpisode{}).Filter("tv_show_id", &show.Id).Exclude("first_aired__isnull", true).Filter("first_aired__lt", now.EndOfDay()).Filter("season__gt", 0).All(&episodes)
    return len(episodes)
}

// list of wanted episodes
func (show *TvShow) Wanted() []*TvEpisode {
    wanted := []*TvEpisode{}
    o := orm.NewOrm()
    o.QueryTable(&TvEpisode{}).Filter("tv_show_id", &show.Id).Filter("status", 0).Exclude("first_aired__isnull", true).Exclude("first_aired", time.Unix(0, 0)).All(&wanted)
    return wanted
}

//check if the WHOLE season is wanted!
func (show *TvShow) SeasonWanted(season_num int) bool {
    wanted := []TvEpisode{}

    o := orm.NewOrm()
    o.QueryTable(&TvEpisode{}).Filter("tv_show_id", &show.Id).Filter("status", 0).Filter("season", season_num).All(&wanted)
    ct, err := o.QueryTable(&TvEpisode{}).Filter("tv_show_id", &show.Id).Filter("season", season_num).Count()
    if err != nil {
        log.Error(err)
        return false
    }
    if int(ct) == len(wanted) {
        return true
    } else {
        return false
    }
}

// return the episode which airs next
func (show TvShow) getNextEpisode() *TvEpisode {
    episode := TvEpisode{}
    o := orm.NewOrm()
    o.QueryTable(&TvEpisode{}).Filter("series_id", show.Id).Filter("air_date").One(&show)
    // SELECT * from tv_episode WHERE series_id = ? and airdate => ? AND airdate <= (SELECT airdate from tv_episode WHERE series_id = ? AND airdate >= ? ORDER BY airdate ASC LIMIT 1) AND status = ?
    return &episode
    //return &quality
}

type Indexers struct {
    Id           int    `orm:"auto"`
    Name         string `orm:"null"`
    Settings     string `orm:"size(255)"`
    EnableRss    bool   `orm:"null"`
    EnableSearch bool   `orm:"null"`
}

type IndexerSettings struct {
    Url        string
    ApiKey     string
    Categories string
}

type QualityDefinitions struct {
    Id      int    `orm:"auto"`
    Quality int    `orm:"null"`
    Title   string `orm:"null"`
    MinSize int    `orm:"null"`
    MaxSize int    `orm:"null"`
}

// check if the given quality is in the given QualutyPreset
func (def *QualityDefinitions) IsInPreset(presetid int) bool {
    preset := QualityPreset{}
    o := orm.NewOrm()
    o.QueryTable(&QualityPreset{}).Filter("id", presetid).One(&preset)

    defs := preset.Unmarshal()
    contains := false
    for _, x := range defs {
        if x.Id == def.Id {
            contains = true
        }
    }
    return contains
}

type Profiles struct {
    Id            int
    Name          string
    Cutoff        int
    RequiredRegex string
    IgnoreRegex   string
    Items         []struct {
        Quality int
        Allowed bool
    }
}

type RootFolders struct {
    Id   int
    Path string
}

type Notifications struct {
    Id         int
    Name       string
    OnGrab     bool
    OnDownload bool
    Settings   string
}

// History Struct containing information about recently processed releases
type History struct {
    Id            int                 `orm:"auto"`
    Action        int                 `orm:"null"` // 1 = snatched, 2 = downloaded, 3 = failed
    Date          time.Time           `orm:"auto_now_add;type(datetime)"`
    ShowId        int                 `orm:"null"`
    Season        int                 `orm:"null"`
    Episode       int                 `orm:"null"`
    Resource      string              `orm:"null"`
    ProviderName  string              `orm:"null"`
    QualityString string              `orm:"null"`
    Title         string              `orm:"null"`
    Link          string              `orm:"null"`
    Version       int                 `orm:"null"`
    Provider      *NewznabProvider    `orm:"null;rel(fk)"`
    Quality       *QualityDefinitions `orm:"null;rel(fk)"`
    TvShow        *TvShow             `orm:"null;rel(fk)"` //optinal series link
    TvEpisode     *TvEpisode          `orm:"null;rel(fk)"` //optinal episode link
}

// holds a single Release
type Release struct {
    Title         string
    Link          string
    Quality       int
    QualityString string
    Provider      Provider
    Status        int // 1 = snatched, 2 = downloaded, 3 = failed
    Episode       *TvEpisode
    Season        []*TvEpisode
    SeasonNum     int
}

// orders a list of releases by quality from good to bad
type ByQuality []Release

func (a ByQuality) Len() int {
    return len(a)
}

func (a ByQuality) Swap(i, j int) {
    a[i], a[j] = a[j], a[i]
}

func (a ByQuality) Less(i, j int) bool {
    return a[i].Quality < a[j].Quality
}

// mark release as snatched
// TODO: clean this! for complete seasons move this to a seperate function?
func (rel Release) markSnatched() {
    o := orm.NewOrm()

    if len(rel.Season) > 0 {
        for _, e := range rel.Season {
            if e.Season == rel.SeasonNum {
                e.Status = STATE_SNATCHED
                o.Update(e)
            }
        }
    } else {
        rel.Episode.Status = STATE_SNATCHED
        o.Update(rel.Episode)
    }

    //fmt.Println(rel)

    //hist := History{ReleaseName: rel.Title, Action: 1, Date: 0, ShowId: rel.Episode.TvShow.Id, Quality: rel.Quality, Resource: "dafuq", Provider: rel.Provider.Name, Version: 1}

    if len(rel.Season) > 0 {
        for _, e := range rel.Season {
            if e.Season == rel.SeasonNum {
                hist := History{}
                hist.Season = e.Season
                hist.Episode = e.Episode
                hist.Title = rel.Title
                hist.Action = STATE_SNATCHED
                hist.ProviderName = rel.Provider.Name

                quali := QualityDefinitions{}
                o.QueryTable(&QualityDefinitions{}).Filter("quality", rel.Quality).One(&quali)
                hist.Quality = &quali
                hist.TvEpisode = e
                hist.Link = rel.Link
                hist.QualityString = quali.Title
                hist.TvShow = e.TvShow
                o.Insert(&hist)
            }
        }

    } else {
        hist := History{}
        hist.Season = rel.Episode.Season
        hist.Episode = rel.Episode.Episode

        hist.Title = rel.Title
        hist.Action = STATE_SNATCHED
        hist.ProviderName = rel.Provider.Name

        quali := QualityDefinitions{}
        o.QueryTable(&QualityDefinitions{}).Filter("quality", rel.Quality).One(&quali)
        hist.Quality = &quali
        hist.TvEpisode = rel.Episode
        hist.Link = rel.Link
        hist.QualityString = quali.Title
        hist.TvShow = rel.Episode.TvShow

        o.Insert(&hist)
    }

}

type TvEpisode struct {
    Id           int                 `orm:"auto"`
    EpisodeId    int                 `orm:"null"`
    ShowId       int                 `orm:"null"`
    IndexerId    int                 `orm:"null"`
    Overview     string              `orm:"null"`
    Indexer      int                 `orm:"null"`
    Name         string              `orm:"null"`
    Season       int                 `orm:"null"`
    Episode      int                 `orm:"null"`
    Description  string              `orm:"null"`
    Status       int                 `orm:"null"`
    Path         string              `orm:"null"`
    Location     int                 `orm:"null"`
    FileSize     uint64              `orm:"null"`
    ReleaseName  string              `orm:"null"`
    ReleaseGroup string              `orm:"null"`
    FirstAired   time.Time           `orm:"null;type(date)"`
    Monitored    bool                `orm:"null"`
    IsProper     bool                `orm:"null"`
    TvShow       *TvShow             `orm:"null;rel(fk)"`
    Quality      *QualityDefinitions `orm:"null;rel(fk)"`
}

// return if the episode has aired already
func (ep *TvEpisode) Aired() bool {
    if ep.FirstAired.IsZero() {
        return false
    }

    return ep.FirstAired.Before(now.BeginningOfDay())
}

// humanize the episodes first aired date, returns a string like: 20 Days ago
func (ep *TvEpisode) HumanizeAired() string {
    if ep.FirstAired.IsZero() {
        return "never"
    }
    return humanize.Time(ep.FirstAired)
}

// return the quality of this episode
func (ep *TvEpisode) QualityString() string {
    o := orm.NewOrm()
    o.LoadRelated(ep, "Quality")
    if ep.Quality != nil {
        return ep.Quality.Title
    }
    return "Unknown"
}

type NewznabProvider struct {
    Id         int `orm:"auto"`
    Name       string
    Url        string
    ApiKey     string
    Categories string
}

//register database models
func initDatabase() {
    dbpath := path.Join(viper.GetString("DataDir"), "gosnatch.db")
    orm.RegisterDataBase("default", "sqlite3", dbpath, 30)
    if viper.GetBool("Devel") {
        orm.Debug = true
    }
    orm.RegisterModel(new(TvShow))
    orm.RegisterModel(new(TvEpisode))
    orm.RegisterModel(new(NewznabProvider))
    orm.RegisterModel(new(History))
    orm.RegisterModel(new(QualityPreset))
    orm.RegisterModel(new(Indexers))
    orm.RegisterModel(new(QualityDefinitions))
    orm.RegisterModel(new(Notifications))
    orm.RegisterModel(new(RootFolders))
    orm.RegisterModel(new(ShowExceptions))

    err := orm.RunSyncdb("default", false, viper.GetBool("Devel"))
    if err != nil {
        log.Panic(err)
    }
}

//TODO: Make sure show only gets added once ...
func addShow(tvdbid int, language string, presetid int, rootfolder string) {
    t := getTvdb()
    o := orm.NewOrm()

    //check if show exists
    ck := TvShow{}
    o.QueryTable(&TvShow{}).Filter("tvdb_id", tvdbid).One(&ck)

    if ck.Id != 0 {
        log.Infof("Already in database %s", ck.ShowName)
        return
    }

    result, err := t.GetFullSeriesData(tvdbid, language)

    if err != nil {
        log.Errorf("tvdb error -> %s", err)
        return
    } else {
        log.Infof("loading informations for %s from tvdb", result.Series.SeriesName)
    }

    //add the tvshow to database
    show := new(TvShow)
    show.TvdbId = result.Series.Id
    show.ShowName = result.Series.SeriesName
    o.Insert(show)

    quality := QualityPreset{}
    err = o.QueryTable(&QualityPreset{}).Filter("id", presetid).One(&quality)
    if err != nil {
        log.Error(err)
    }

    //load slug and tvrage id from trakt
    client := getTrakt()

    //TODO: repeat till success!
    shows, err := client.Shows().Tvdb(result.Series.Id)
    if err != nil {
        log.Errorf("Trakt.tv -> %v", err)
    }
    rageID := 0
    traktId := 0
    showSlug := ""
    for _, show := range shows {
        if show.Type == "show" && show.Show.IDs.Tvdb == result.Series.Id {
            rageID = show.Show.IDs.Tvrage
            showSlug = show.Show.IDs.Slug
            traktId = show.Show.IDs.Trakt
        }
    }

    //should never happen!!
    if showSlug == "" {
        splitYear := strings.Split(result.Series.FirstAired, "-")
        showSlug = fmt.Sprintf("%s-%s", strings.ToLower(slug.SlugAscii(result.Series.SeriesName)), splitYear[0])
        log.Errorf("Unable to load Slug... generating one: %s", showSlug)

    }

    sdir := path.Join(rootfolder, result.Series.SeriesName)

    //update show stuff
    show.Network = result.Series.Network
    show.Genre = strings.Join(result.Series.Genres, "|")
    show.Overview = result.Series.Overview
    show.Runtime = result.Series.Runtime
    show.Airs = result.Series.AirDay
    show.Language = language
    show.ImdbId = result.Series.ImdbId
    show.AirTime = result.Series.AirTime
    show.Startyear = result.Series.Year
    show.Slug = showSlug
    show.TvrageId = rageID
    show.QualityPreset = &quality
    show.Path = sdir
    show.TraktId = traktId

    eps := []*TvEpisode{}
    o.Update(show)

    //load and insert exceptions!
    altNames, _ := getShowExceptions(tvdbid)

    exceptions := []ShowExceptions{}

    for _, x := range altNames {
        ex := ShowExceptions{ShowName: x, TvdbId: tvdbid, Series: show}
        exceptions = append(exceptions, ex)
        log.WithFields(log.Fields{"show_name": show.ShowName, "exception": x}).Debug("added exception")
        o.Insert(&ex)
    }

    for _, episode := range result.Episode {
        ep := &TvEpisode{}
        ep.EpisodeId = episode.Id
        ep.Episode = episode.EpisodeNumber
        ep.Season = episode.SeasonNumber
        ep.Name = episode.EpisodeName
        ep.ShowId = episode.SeriesId
        ep.TvShow = show

        // Use UTC Dates in Database!
        airtime, err := time.Parse("2006-01-02", episode.FirstAired)
        if err == nil {
            ep.FirstAired = airtime
        }

        ep.Overview = episode.Overview
        eps = append(eps, ep)

        o.Insert(ep)
    }

    show.makeShowDir()
    go show.scanDirectory()
    log.WithField("show", result.Series.SeriesName).Debug("sucessfully added to database")
    //return fmt.Sprintf("%s (%s)", show.ShowName, quality.Name), nil

}

func GetDatabase() {
    initDatabase()
}

type Event struct {
    Id    int
    Title string `json:"title"`
    Url   string
    Class string `json:"className"`
    Start string `json:"start"`
    End   string `json:"end"`
}

// return a list of Events for the calendar view
func getCalendarEpisodes(start string, end string) []Event {
    o := orm.NewOrm()
    eps := []TvEpisode{}
    begin, _ := time.Parse("2006-01-02", start)
    till, _ := time.Parse("2006-01-02", end)

    o.QueryTable(&TvEpisode{}).Filter("first_aired__gte", begin).Filter("first_aired__lte", till).All(&eps)
    result := []Event{}

    for _, ep := range eps {

        o.LoadRelated(&ep, "TvShow")
        title := fmt.Sprintf("%s - %s", ep.TvShow.ShowName, ep.Name)
        //calculate start and end time
        // get just the day!
        y := ep.FirstAired.Local().Format("2006-01-02")

        x, _ := time.Parse("2006-01-02 15:04", fmt.Sprintf("%s %s", y, ep.TvShow.AirTime))
        start := x.Format(time.RFC3339)
        endhour := time.Duration(ep.TvShow.Runtime) * time.Minute
        end := x.Add(endhour).Format(time.RFC3339)

        class := "event-unaired"
        // already was running and is missed
        if x.Before(now.BeginningOfMinute()) && ep.Status == 0 {
            class = "event-missed"
        }

        //currently running!
        // we need current time parsed as UTC... so basically well do a now format it and then parse it back in?
        utcBeginningOfMinute := now.BeginningOfMinute().Format("2006-01-02 15:04:05")
        utcBegin, _ := time.Parse("2006-01-02 15:04:05", utcBeginningOfMinute)

        if x.Before(utcBegin) && ep.Status == 0 && x.Add(endhour).After(utcBegin) {
            class = "event-running"
        }

        if ep.Episode == 1 && !x.Before(now.EndOfHour()) && ep.Status == 0 {
            class = "event-premiere"
        }

        //nevermind i got the episode already
        if ep.Status == 1 {
            class = "event-snatched"
        }

        if ep.Status > 1 {
            class = "event-downloaded"
        }

        cur := Event{Title: title, Class: class, Start: start, End: end, Id: ep.Id}
        result = append(result, cur)
    }

    return result
}

// populate database with default quality definitions and quality presets
func insertDefaultValues() {
    //should only be done on first run!!!!
    DefaultQualitys := []QualityDefinitions{}

    UNKNOWN := QualityDefinitions{}
    UNKNOWN.Title = "Unknown"
    UNKNOWN.Quality = 0
    DefaultQualitys = append(DefaultQualitys, UNKNOWN)

    SDTV := QualityDefinitions{}
    SDTV.Title = "SD TV"
    SDTV.Quality = 1
    DefaultQualitys = append(DefaultQualitys, SDTV)

    SDDVD := QualityDefinitions{}
    SDDVD.Title = "SD DVD"
    SDDVD.Quality = 2
    DefaultQualitys = append(DefaultQualitys, SDDVD)

    HDTV := QualityDefinitions{}
    HDTV.Title = "HD TV"
    HDTV.Quality = 3
    DefaultQualitys = append(DefaultQualitys, HDTV)

    RAWHDTV := QualityDefinitions{}
    RAWHDTV.Title = "RawHD TV"
    RAWHDTV.Quality = 4
    DefaultQualitys = append(DefaultQualitys, RAWHDTV)

    HDTV1080 := QualityDefinitions{}
    HDTV1080.Title = "HD TV 1080p"
    HDTV1080.Quality = 5
    DefaultQualitys = append(DefaultQualitys, HDTV1080)

    WEBDL720 := QualityDefinitions{}
    WEBDL720.Title = "WEB-DL 720p"
    WEBDL720.Quality = 6
    DefaultQualitys = append(DefaultQualitys, WEBDL720)

    WEBDL1080 := QualityDefinitions{}
    WEBDL1080.Title = "WEB-DL 1080p"
    WEBDL1080.Quality = 7
    DefaultQualitys = append(DefaultQualitys, WEBDL1080)

    BLURAY720 := QualityDefinitions{}
    BLURAY720.Title = "BluRay 720p"
    BLURAY720.Quality = 8
    DefaultQualitys = append(DefaultQualitys, BLURAY720)

    BLURAY1080 := QualityDefinitions{}
    BLURAY1080.Title = "BluRay 1080p"
    BLURAY1080.Quality = 9
    DefaultQualitys = append(DefaultQualitys, BLURAY1080)

    o := orm.NewOrm()
    for _, x := range DefaultQualitys {
        ck := QualityDefinitions{}
        o.QueryTable(&QualityDefinitions{}).Filter("quality", x.Quality).One(&ck)
        if ck.Id == 0 {
            o.Insert(&x)
        }

    }

    ck := QualityPreset{}
    o.QueryTable(&QualityPreset{}).Filter("name", "HD 720p").One(&ck)

    if ck.Id == 0 {
        HD720 := []QualityDefinitions{}
        o.QueryTable(&QualityDefinitions{}).Filter("quality__in", 3, 6, 8).All(&HD720)
        x, _ := json.Marshal(HD720)

        HD720Preset := QualityPreset{}
        HD720Preset.Name = "HD 720p"
        HD720Preset.ArchiveFirstmatch = true
        HD720Preset.Qualitys = string(x)

        o.Insert(&HD720Preset)
    }

    ck = QualityPreset{}
    o.QueryTable(&QualityPreset{}).Filter("name", "HD 1080p").One(&ck)

    if ck.Id == 0 {
        HD1080 := []QualityDefinitions{}
        o.QueryTable(&QualityDefinitions{}).Filter("quality__in", 5, 7, 9).All(&HD1080)
        x, _ := json.Marshal(HD1080)

        HD1080Preset := QualityPreset{}
        HD1080Preset.Name = "HD 1080p"
        HD1080Preset.ArchiveFirstmatch = true
        HD1080Preset.Qualitys = string(x)

        o.Insert(&HD1080Preset)
    }

    ck = QualityPreset{}
    o.QueryTable(&QualityPreset{}).Filter("name", "HD All").One(&ck)

    if ck.Id == 0 {
        HDAny := []QualityDefinitions{}
        o.QueryTable(&QualityDefinitions{}).Filter("quality__in", 3, 5, 6, 7, 8, 9).All(&HDAny)
        x, _ := json.Marshal(HDAny)

        HDAnyPreset := QualityPreset{}
        HDAnyPreset.Name = "HD All"
        HDAnyPreset.ArchiveFirstmatch = true
        HDAnyPreset.Qualitys = string(x)

        o.Insert(&HDAnyPreset)
    }

    ck = QualityPreset{}
    o.QueryTable(&QualityPreset{}).Filter("name", "SD All").One(&ck)

    if ck.Id == 0 {
        SDAll := []QualityDefinitions{}
        o.QueryTable(&QualityDefinitions{}).Filter("quality__in", 1, 2).All(&SDAll)
        x, _ := json.Marshal(SDAll)

        SDAllPreset := QualityPreset{}
        SDAllPreset.Name = "SD All"
        SDAllPreset.ArchiveFirstmatch = true
        SDAllPreset.Qualitys = string(x)

        o.Insert(&SDAllPreset)
    }

    viper.Set("FirstRun", false)

}
