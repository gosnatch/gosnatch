package gosnatch

import (
    "encoding/json"
    "fmt"
    "github.com/astaxie/beego/orm"
    "github.com/dustin/go-humanize"
    assetfs "github.com/elazarl/go-bindata-assetfs"
    "github.com/gin-gonic/contrib/static"
    "github.com/gin-gonic/gin"
    "github.com/gosnatch/gosnatch/tvdb"
    _ "github.com/mattn/go-sqlite3"
    "github.com/nicksnyder/go-i18n/i18n"
    "github.com/nicksnyder/go-i18n/i18n/language"
    "github.com/nicksnyder/go-i18n/i18n/translation"
    log "github.com/sirupsen/logrus"
    "github.com/spf13/viper"
    "html/template"
    "net/http"
    "net/url"
    "path"
    "strconv"
    "strings"
    "time"
)

func ApiKeyRequired() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Try to get an authorization pass.
        // We provide a AuthRequest and we should get AuthPass in response
        key := c.Request.Header.Get("X-Api-Key")
        //should use config api key and send to all requests...
        if key != viper.GetString("ApiKey") {
            c.Fail(401, fmt.Errorf("%s", "Unknown Api Key"))
        }

    }
}

func DumpHeaders() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Try to get an authorization pass.
        // We provide a AuthRequest and we should get AuthPass in response
        fmt.Println(c.Request.Header)
        fmt.Println(c.Request.Host)
        fmt.Println(c.Request.RemoteAddr)
        fmt.Println(c.Request)

    }
}

type binaryFileSystem struct {
    fs http.FileSystem
}

func (b *binaryFileSystem) Open(name string) (http.File, error) {
    return b.fs.Open(name)
}

func (b *binaryFileSystem) Exists(prefix string, filepath string) bool {

    if p := strings.TrimPrefix(filepath, prefix); len(p) < len(filepath) {
        if _, err := b.fs.Open(p); err != nil {
            return false
        }
        return true
    }
    return false
}

func BinaryFileSystem(root string) *binaryFileSystem {
    fs := &assetfs.AssetFS{Asset, AssetDir, root}
    return &binaryFileSystem{
        fs,
    }
}

func getApiKey(s string) string {
    return viper.GetString("ApiKey")
}

//load templates from the asset that go-bindata provides
func loadTemplates(list ...string) *template.Template {

    //load languages from go-bindata assets folder
    for _, curlang := range []string{"de-DE", "en-US"} {
        transString, _ := Asset("assets/translations/" + curlang + ".all.json")
        var transUnmarshal []map[string]interface{}
        err := json.Unmarshal(transString, &transUnmarshal)
        if err != nil {
            fmt.Println(err)
        }

        for _, y := range transUnmarshal {
            trans, _ := translation.NewTranslation(y)
            lang := language.MustParse(curlang)
            i18n.AddTranslation(lang[0], trans)
        }
    }

    T, _ := i18n.Tfunc(viper.GetString("Language"))

    templates := template.New("")

    funcMap := template.FuncMap{
        "ApiKey":       getApiKey,
        "humanizeTime": func(a time.Time) string { return humanize.Time(a) },
        "title":        func(a string) string { return strings.Title(a) },
        "T":            T,
    }

    templates.Funcs(funcMap)

    for _, x := range list {
        templateString, err := Asset("assets/templates/" + x)
        if err != nil {
            log.Fatal(err)
        }

        // get file contents as string
        _, err = templates.New(x).Parse(string(templateString))
        if err != nil {
            log.Fatal(err)
        }
    }

    return templates
}

func StartWebserver() {
    o := orm.NewOrm()
    t := getTvdb()

    router := gin.New()
    router.Use(gin.Recovery())
    router.Use(static.Serve("/assets", BinaryFileSystem("assets")))

    if viper.GetBool("Debug") {
        gin.SetMode(gin.DebugMode)
        router.Use(gin.Logger())
    } else {
        gin.SetMode(gin.ReleaseMode)
    }

    // load templates from bin data!
    tmpl := loadTemplates("script.tmpl", "calendar.tmpl", "addSeries.tmpl", "history.tmpl", "shows.tmpl", "show.tmpl", "settings.tmpl", "presets.tmpl")
    // router.LoadHTMLGlob("assets/templates/*")
    router.SetHTMLTemplate(tmpl)

    // router.Static("/assets", "assets/")
    cachedir := path.Join(viper.GetString("DataDir"), "cache")
    router.Static("/cache", cachedir)

    api := router.Group("/api")
    api.Use(ApiKeyRequired())
    {
        api.GET("/series/lookup/:term", func(c *gin.Context) {
            showname := c.Params.ByName("term")
            result := []*tvdb.Series{}
            if strings.HasPrefix(showname, "tvdb:") {
                str := strings.Split(showname, ":")
                id, _ := strconv.Atoi(str[1])
                x, _ := t.GetSeriesById(id, "en")
                if x != nil {
                    result = append(result, x)
                }
            } else if strings.HasPrefix(showname, "imdb:") {
                str := strings.Split(showname, ":")
                x, _ := t.GetSeriesByRemoteId(str[1], "en")
                if x != nil {
                    result = append(result, x)
                }
            } else if strings.HasPrefix(showname, "tt") {
                x, _ := t.GetSeriesByRemoteId(showname, "en")
                if x != nil {
                    result = append(result, x)
                }
            } else {
                result, _ = t.GetSeries(showname, "en")
            }

            if len(result) == 0 {
                msg := fmt.Sprintf("unable to find result for %s", showname)
                result := gin.H{"error": msg}
                c.JSON(200, result)
            } else {
                c.JSON(200, result)
            }

        })

        //refresh all shows!
        api.GET("/refreshshows", func(c *gin.Context) {
            myshows := []TvShow{}
            o.QueryTable(&TvShow{}).All(&myshows)

            for _, show := range myshows {
                go show.UpdateShowInfo()
                go show.scanDirectory()
                go cacheImages(show.TvdbId)
            }

        })

        api.GET("/refreshshow/:id", func(c *gin.Context) {
            ids := c.Params.ByName("id")
            id, _ := strconv.Atoi(ids)
            myshow := TvShow{}
            o.QueryTable(&TvShow{}).Filter("id", id).One(&myshow)

            if myshow.Id != 0 {
                go myshow.UpdateShowInfo()
                go myshow.scanDirectory()
                go cacheImages(myshow.TvdbId)
            }

            c.JSON(200, gin.H{"status": "success", "message": fmt.Sprintf("Force refresh for \"%s\" activated!", myshow.ShowName)})
        })

        api.GET("/deleteshow/:id", func(c *gin.Context) {
            ids := c.Params.ByName("id")
            id, _ := strconv.Atoi(ids)
            myshow := TvShow{}
            o.QueryTable(&TvShow{}).Filter("id", id).One(&myshow)

            o.Delete(&myshow)

            c.JSON(200, gin.H{"satus": "success", "message": fmt.Sprintf("Series \"%s\" successfully deleted", myshow.ShowName)})
        })

        api.GET("/release", func(c *gin.Context) {
            epid := c.Request.URL.Query().Get("episodeId")
            // sort := c.Request.URL.Query().Get("sort_by")
            // order := c.Request.URL.Query().Get("order")
            c.JSON(200, epid)
        })

        api.POST("/addfolder", func(c *gin.Context) {
            type Path struct {
                Path string `json:"path" binding:"required"`
            }
            var newpath Path
            c.Bind(&newpath)
            fmt.Println(newpath)
            ck := RootFolders{}
            o.QueryTable(&RootFolders{}).Filter("path", newpath.Path).One(&ck)
            if ck.Id == 0 && newpath.Path != "" && fileExists(newpath.Path) {
                ck.Path = newpath.Path
                o.Insert(&ck)
            }
            allfolders := []RootFolders{}
            o.QueryTable(&RootFolders{}).All(&allfolders)
            c.JSON(200, allfolders)
        })

        api.GET("/folder", func(c *gin.Context) {
            folders := []RootFolders{}
            o.QueryTable(&RootFolders{}).All(&folders)
            c.JSON(200, folders)
        })

        api.GET("/filesystem", func(c *gin.Context) {
            path := c.Request.URL.Query().Get("path")
            fpath, err := url.QueryUnescape(path)
            if err != nil {
                fmt.Println(err)
            }

            folders := scanFilesystem(fpath)
            c.JSON(200, folders)

        })

        api.POST("/postprocess", func(c *gin.Context) {
            //http POST http://127.0.0.1:1313/api/postprocess X-Api-Key:123 Content-Type:application/json failed:=1 folder=/xyz/downloads/Arrow.S01e02
            var postproc PostProcessor
            c.Bind(&postproc)
            postproc.History()
            postproc.processRelease()
            c.JSON(200, postproc)
        })

        api.POST("/addshow", func(c *gin.Context) {
            //http POST http://127.0.0.1:1313/api/addshow X-Api-Key:123 Content-Type:application/json tvdbid:=80348 quality:=1 language=de folder=/xyz/downloads/
            //http POST http://127.0.0.1:1313/api/addshow X-Api-Key:123 Content-Type:application/json tvdbid:=79349 quality:=1 language=de folder=/xyz/downloads/

            type AddShow struct {
                TvdbId        int    `json:"tvdbid" binding:"required"`
                Language      string `json:"language" binding:"required"`
                QualityPreset int    `json:"quality" binding:"required"`
                Folder        string `json:"folder" binding:"required"`
            }

            var json AddShow
            c.Bind(&json)

            check := TvShow{}
            o.QueryTable(&check).Filter("tvdb_id", json.TvdbId).One(&check)

            if check.TvdbId != 0 {
                c.JSON(200, gin.H{"satus": "success", "message": fmt.Sprintf("Series \"%s\" already in Database", check.ShowName)})
            } else {
                go cacheImages(json.TvdbId)
                go addShow(json.TvdbId, json.Language, json.QualityPreset, json.Folder)
                c.JSON(200, gin.H{"satus": "success", "message": fmt.Sprintf("Series will be added")})
            }

        })

        api.POST("/snatch", func(c *gin.Context) {
            type ManualSnatch struct {
                Title   string `json:"title" binding:"required"`
                Url     string `json:"url" binding:"required"`
                Season  int    `json:"season" binding:"required"`
                Episode int    `json:"episode" binding:"required"`
                ShowId  int    `json:"show_id" binding:"required"`
            }

            //http POST http://127.0.0.1:1313/api/postprocess X-Api-Key:123 Content-Type:application/json failed:=1 folder=/xyz/downloads/Arrow.S01e02
            var json ManualSnatch
            c.Bind(&json)
            c.JSON(200, json)
        })

        api.GET("/quality", func(c *gin.Context) {
            quals := []QualityDefinitions{}
            o.QueryTable(&QualityDefinitions{}).All(&quals)
            c.JSON(200, quals)
        })

        api.GET("/calendar", func(c *gin.Context) {
            start := c.Request.URL.Query().Get("start")
            end := c.Request.URL.Query().Get("end")

            stuff := getCalendarEpisodes(start, end)

            c.JSON(200, stuff)
        })

        api.GET("/preset", func(c *gin.Context) {
            quals := []QualityPreset{}
            o.QueryTable(&QualityPreset{}).All(&quals)
            c.JSON(200, quals)
        })

        api.GET("/manual/episode/:episodeid", func(c *gin.Context) {
            epId := c.Params.ByName("episodeid")
            epint, _ := strconv.Atoi(epId)
            releases := ManualSearch(epint)
            c.JSON(200, releases)

        })

        api.GET("/search/episode/:episodeid", func(c *gin.Context) {
            epId := c.Params.ByName("episodeid")
            epint, _ := strconv.Atoi(epId)
            release := EpisodeSearch(epint)
            c.JSON(200, release)

        })

        api.GET("/settings", func(c *gin.Context) {
            c.JSON(200, &C)
        })

        api.GET("/profiles", func(c *gin.Context) {

        })
    }

    router.GET("/", func(c *gin.Context) {
        go DailySearch(true)
        c.HTML(200, "addSeries.tmpl", nil)
    })

    router.GET("/addseries", func(c *gin.Context) {
        folders := []RootFolders{}
        o.QueryTable(&RootFolders{}).All(&folders)
        c.HTML(200, "addSeries.tmpl", gin.H{"folders": folders})
    })

    router.GET("/shows", func(c *gin.Context) {
        shows := []TvShow{}
        o.QueryTable(&TvShow{}).Exclude("slug", nil).All(&shows)

        c.HTML(200, "shows.tmpl", gin.H{"shows": shows})

    })

    router.GET("/history", func(c *gin.Context) {
        history := []History{}
        ret := []History{}
        o.QueryTable(&History{}).OrderBy("-date").All(&history)
        for _, x := range history {
            o.LoadRelated(&x, "TvShow")
            o.LoadRelated(&x, "TvEpisode")
            o.LoadRelated(&x, "Quality")
            ret = append(ret, x)
        }
        c.HTML(200, "history.tmpl", gin.H{"history": ret})
    })

    router.GET("/presets", func(c *gin.Context) {
        presets := []QualityPreset{}
        qualitys := []QualityDefinitions{}
        o.QueryTable(&QualityPreset{}).All(&presets)
        o.QueryTable(&QualityDefinitions{}).All(&qualitys)
        // for _, preset := range presets {
        //     preset.Qualitys =
        // }
        c.HTML(200, "presets.tmpl", gin.H{"presets": presets, "qualitys": qualitys})
    })

    router.GET("/calendar", func(c *gin.Context) {

        c.HTML(200, "calendar.tmpl", nil)

    })

    router.GET("/show/:slug", func(c *gin.Context) {
        slug := c.Params.ByName("slug")

        curShow := TvShow{}
        o.QueryTable("tv_show").Filter("slug", slug).One(&curShow)
        o.LoadRelated(&curShow, "QualityPreset")

        seasons := Seasons{}
        var maps []orm.Params
        num, err := o.Raw("SELECT id FROM tv_episode WHERE tv_show_id = ? AND season >= 1 GROUP BY season", curShow.Id).Values(&maps)
        if err != nil {
            fmt.Println(err)
        }

        // x := num
        for i := num; i > 0; i-- {
            episodes := []TvEpisode{}
            o.QueryTable(&TvEpisode{}).Filter("tv_show_id", curShow.Id).Filter("season", i).OrderBy("-season", "-episode").All(&episodes)
            // seasons[i] = episodes
            season := Season{i, episodes}
            seasons = append(seasons, season)
        }

        obj := gin.H{"show": curShow, "seasons": seasons}
        c.HTML(200, "show.tmpl", obj)

    })

    if !viper.GetBool("Debug") {
        log.Infof("webserver up and running on 0.0.0.0:%d", viper.GetInt("Port"))
    }

    router.Run(fmt.Sprintf(":%d", viper.GetInt("Port")))
}
