package gosnatch

import (
    log "github.com/sirupsen/logrus"
    // "os"
    "encoding/json"
    "github.com/mitchellh/go-homedir"
    "github.com/spf13/viper"
    "math/rand"
    "os"
    "path"
)

const (
    USERAGENT = "gosnatch v0.1"
    VERSION   = "0.1"

    STATE_WANTED     = 0
    STATE_SNATCHED   = 1
    STATE_FAILED     = 3
    STATE_DOWNLOADED = 2
)

type config struct {
    DataDir         string
    Debug           bool
    Devel           bool
    ApiKey          string
    SnatchFolder    string
    Port            int
    TvdbUrl         string
    TvdbKey         string
    LogJSON         bool
    FirstRun        bool
    UpdateOnStart   bool
    SearchFrequency int
}

var C config

var validMovieExtensions = []string{"mkv", "mp4", "mov", "mpg", "mpeg", "wmv", "ogm", "iso", "img", "divx", "m2ts", "m4v", "ts", "flv", "f4v", "rmvb", "vob", "dvr-ms", "wtv", "ogv", "3gp", "webm", "avi"}
var subtitleExtensions = []string{"srt", "sub", "ass", "idx", "ssa"}

func saveConfig(cfgpath string) error {

    viper.Marshal(&C)

    b, err := json.MarshalIndent(C, "", "    ")
    if err != nil {
        return err
    }

    f, err := os.Create(cfgpath)
    if err != nil {
        return err
    }

    defer f.Close()

    f.WriteString(string(b))

    return nil
}

var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}

func InitConfig() {
    usrdir, err := homedir.Dir()
    if err != nil {
        log.Error(err)
    }
    usrdir, _ = homedir.Expand(usrdir)

    cfgdir := path.Join(usrdir, ".gosnatch")

    viper.SetConfigName("config")
    // viper.AddConfigPath(cfgdir)

    //default values
    viper.SetDefault("DataDir", cfgdir)
    viper.SetDefault("Debug", false)
    viper.SetDefault("FirstRun", true)
    viper.SetDefault("UpdateOnStart", true)
    viper.SetDefault("Port", 1313)
    viper.SetDefault("LogJSON", false)
    viper.SetDefault("Language", "en-US")
    viper.SetDefault("Devel", false)
    viper.SetDefault("TvdbUrl", "https://tvdb.gosnatch.me")
    viper.SetDefault("TvdbKey", "BTLPI2H0LITPW71J") //goSnatch's api key
    viper.SetDefault("ApiKey", randSeq(20))
    viper.SetDefault("SearchFrequency", 20)
    viper.SetDefault("NamingPattern", "Season {{.SeasonNumber}}/{{.ShowName}}")
    viper.SetDefault("AutoUpdte", false)
    viper.SetEnvPrefix("snt") // will be uppercased automatically
    viper.BindEnv("DataDir")
    viper.BindEnv("Devel")
    viper.BindEnv("Debug")

    //now it can be overwritten by env!
    viper.AddConfigPath(viper.GetString("DataDir"))
    viper.SetDefault("SnatchFolder", path.Join(viper.GetString("DataDir"), "snatched"))

    viper.ReadInConfig()

    if viper.GetBool("Debug") {
        log.SetLevel(log.DebugLevel)
    } else {
        log.SetLevel(log.InfoLevel)
    }

    if viper.GetBool("LogJSON") {
        log.SetFormatter(&log.JSONFormatter{})
    } else {
        x := SafeLogger{&log.TextFormatter{}}
        log.SetFormatter(&x)
    }

    os.MkdirAll(viper.GetString("DataDir"), 0777)
    cfgpath := path.Join(viper.GetString("DataDir"), "config.json")

    cachedir := path.Join(viper.GetString("DataDir"), "cache", "images")

    for _, x := range []string{"fanart", "banner", "poster"} {
        curdir := path.Join(cachedir, x)
        if !fileExists(curdir) {
            os.MkdirAll(curdir, 0777)
        }
    }

    GetDatabase()

    insertDefaultValues()

    err = saveConfig(cfgpath)
    if err != nil {
        log.Error(err)
    }

    log.Debug("config initialized")
    log.Info("all config args can be overwritten by env flasg with SNT_ prefix ex: SNT_PORT=8080")

}
