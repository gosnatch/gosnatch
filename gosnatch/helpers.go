package gosnatch

import (
    "fmt"
    "github.com/astaxie/beego/orm"
    "github.com/gosnatch/gosnatch/go-trakt"
    "github.com/gosnatch/gosnatch/tvdb"
    _ "github.com/mattn/go-sqlite3"
    "github.com/nfnt/resize"
    log "github.com/sirupsen/logrus"
    "github.com/spf13/viper"
    "image/jpeg"
    "io/ioutil"
    "net/http"
    "os"
    "path"
    "path/filepath"
    "regexp"
    "strings"
)

func postUrl(url string) {

}

type SafeLogger struct {
    fmt *log.TextFormatter
}

func (f *SafeLogger) Format(entry *log.Entry) ([]byte, error) {
    fmt.Println(entry.Message)
    fmt.Println((entry.Message))
    entry.Message = cleanLogMessage(entry.Message)
    for z, y := range entry.Data {
        switch y := y.(type) {
        case string:
            entry.Data[z] = cleanLogMessage(y)
        }

    }
    return f.fmt.Format(entry)
}

// move a file from oldname to newname, mode can be any of (move, copy, link)
func moveFile(oldname string, newname string, mode string) bool {
    if mode == "move" {
        err := os.Rename(oldname, newname)
        if err != nil {
            return false
        }
        return true
    } else if mode == "copy" {
        return false
    } else if mode == "link" {
        err := os.Symlink(oldname, newname)
        if err != nil {
            return false
        }
        return true
    }
    return false
}

// scan a directory and parse all containing media files
func parseFilesInDir(location string) chan ParseResult {
    chann := make(chan ParseResult)
    go func() {
        o := orm.NewOrm()
        filepath.Walk(location, func(path string, f os.FileInfo, _ error) (err error) {
            if !f.IsDir() {

                for _, x := range validMovieExtensions {
                    if strings.HasSuffix(f.Name(), x) {
                        _, file := filepath.Split(path)

                        // first check ONLY s00e00 stuff and then if nothing is found, parse the release regexes over it..
                        y := parseFilename(file)

                        if len(y) > 0 {
                            r := Release{Title: file}
                            i, _ := checkQuality(r)
                            if i == 0 {
                                i, _ = sickbeardQuality(r)
                            }
                            qual := QualityDefinitions{}
                            o.QueryTable(&QualityDefinitions{}).Filter("quality", i).One(&qual)
                            z := ParseResult{ShowName: y["series_name"], Season: y["season_num"], Episode: y["ep_num"], Path: path, FileSize: f.Size(), Quality: &qual}
                            chann <- z

                            return
                        } else {
                            log.Debugf("unable to parse '%s' with any regex", file)
                        }

                        return nil
                    }
                }
            }
            return
        })
        defer close(chann)
    }()
    return chann
}

// get contetn of given url as []byte
func getUrl(url string) ([]byte, error) {
    client := &http.Client{}

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        log.Error(fmt.Errorf("http.NewRequest -> %v", err))
        return nil, err
    }

    req.Header.Set("User-Agent", USERAGENT)

    resp, err := client.Do(req)
    if err != nil {
        log.Error(fmt.Errorf("client.Do -> %v", err))
        return nil, err
    }

    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Error(fmt.Errorf("ioutil.ReadAll -> %v", err))
        return nil, err
    }

    return body, nil
}

// return a tvdb.TVDB instance
func getTvdb() *tvdb.TVDB {
    return tvdb.Open(viper.GetString("TvdbKey"), viper.GetString("TvdbUrl"))
}

// return a trakt.Client instance
func getTrakt() *trakt.Client {
    return trakt.NewClient("d4161a7a106424551add171e5470112e4afdaf2438e6ef2fe0548edc75924868", nil)
}

// check if the given path exists
func fileExists(name string) bool {
    if _, err := os.Stat(name); err != nil {
        if os.IsNotExist(err) {
            return false
        }
    }
    return true
}

// split filename and extension
func splitExt(path string) (root, ext string) {
    ext = filepath.Ext(path)
    root = path[:len(path)-len(ext)]
    return
}

// resize image at the given path (appends _small) to the original filename
func resizeImage(path string) {
    dir, fname := filepath.Split(path)
    fn, ext := splitExt(fname)

    resized := fmt.Sprintf("%s%s_small%s", dir, fn, ext)

    if !fileExists(resized) {

        file, err := os.Open(path)
        if err != nil {
            log.Fatal(err)
        }

        // decode jpeg into image.Image
        img, err := jpeg.Decode(file)
        if err != nil {
            log.Fatal(err)
        }
        file.Close()

        // resize to width 1000 using Lanczos resampling
        // and preserve aspect ratio
        m := resize.Resize(900, 0, img, resize.Lanczos3)
        if strings.Contains(resized, "banner") || strings.Contains(resized, "poster") {
            m = resize.Resize(500, 0, img, resize.Lanczos3)
        }

        out, err := os.Create(resized)
        if err != nil {
            log.Fatal(err)
        }
        defer out.Close()

        // write new image to file
        jpeg.Encode(out, m, nil)
    }

}

// try to cache images from tvdb
func cacheImages(tvdbid int) error {
    log.Info("Cacheing images for new added show")
    t := getTvdb()
    res, err := t.GetSeriesById(tvdbid, "de")
    if err != nil {
        log.Error("Unable to get informations from tvdb")
        return err
    }

    cachedir := path.Join(viper.GetString("DataDir"), "cache", "images")

    imgtypes := []string{"banner", "fanart", "poster"}
    err = nil

    for _, img := range imgtypes {
        url := ""
        switch {
        case img == "fanart":
            if res.FanArt != "" {
                url = t.GetImageUrl(res.FanArt)
            }
        case img == "banner":
            if res.Banner != "" {
                url = t.GetImageUrl(res.Banner)
            }
        case img == "poster":
            if res.Poster != "" {
                url = t.GetImageUrl(res.Poster)
            }
        }
        if url != "" {
            filename := fmt.Sprintf("%d.jpg", res.Id)
            err = downloadToFile(url, path.Join(cachedir, img, filename))
            resizeImage(path.Join(cachedir, img, filename))
            if err != nil {
                log.Errorf("unable to download %s for show: %s", img, res.SeriesName)
            }
        } else {
            log.WithFields(log.Fields{"tvdb_id": res.Id, "type": img}).Debug("unable to find a download")
        }

    }

    return err

}

// returns a list of Directory and Files inside the given fpath, if fpath is not valid, its parent will be used, if there is no parent ROOT will be used ( / )
func scanFilesystem(fpath string) FileSystem {
    if fpath == "" {
        fpath = string(os.PathSeparator)
    }
    fs := FileSystem{}
    files, err := ioutil.ReadDir(fpath)
    //try to load parent directory if i cant access current typed string
    if err != nil {
        rootFolder, _ := path.Split(fpath)
        files, _ = ioutil.ReadDir(rootFolder)
        fpath = rootFolder
    }

    for _, f := range files {
        fullpath := path.Join(fpath, f.Name())
        if f.IsDir() {
            current := FSDirectory{Name: f.Name(), Path: fullpath, Type: "folder"}
            // if strings.HasPrefix(f.Name(), ".") {
            //     current.HiddenFile = true
            // }
            fs.Directories = append(fs.Directories, current)
        } else {
            current := FSFile{Name: f.Name(), Path: fullpath, Type: "file"}
            fs.Files = append(fs.Files, current)
        }
    }

    return fs
}

// download content of url to file
func downloadToFile(url string, filepath string) error {

    log.WithFields(log.Fields{"url": url, "path": filepath}).Debug("downloading file")
    data, _ := getUrl(url)

    ioutil.WriteFile(filepath, data, 0666)

    return nil
}

// replace sensitive informations in message currently only strips username OR password
func cleanLogMessage(message string) string {
    //url styles
    CleaningRules := []string{
        `(\?|&)(apikey|token|passkey|uid|api|r)=(?P<secret>[^&=]+?)(&|$)`,                            //newznab
        `(\?|&)[^=]*?(username|password)=(?P<secret>[^&=]+?)(&|$)`,                                   // url append
        `""Name""\s*:\s*""[^""]*(username|password)""\s*,\s*""Value""\s*:\s*""(?P<secret>[^""]+?)""`, //nzbget
    }

    // var message string = message
    newMessage := message
    for _, regex := range CleaningRules {
        md := map[string]string{}
        re := regexp.MustCompile(regex)
        n1 := re.SubexpNames()
        match := re.FindAllStringSubmatch(message, -1)
        if len(match) > 0 {
            for i, n := range match[0] {

                if n1[i] != "" {
                    md[n1[i]] = n
                }

            }
            newMessage = strings.Replace(newMessage, md["secret"], "<replaced>", -1)
        }

    }

    return newMessage
}
