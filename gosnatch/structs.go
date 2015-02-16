package gosnatch

import "fmt"

// Result for Parsed Files from Directory
type ParseResult struct {
    ShowName string
    Season   string
    Episode  string
    ExtraEp  int
    FileSize int64
    Path     string
    Quality  *QualityDefinitions
}

// all available in formatting options
type NameFormatter struct {
    ShowName      string
    SeasonNumber  int
    EpisodeNumber int
    EpisodeName   string
    Quality       string
    ReleaseName   string
    ReleaseGroup  string
}

// Newznab Provider Settings
type Provider struct {
    Name      string
    URL       string
    APIKey    string
    Categorys []string
}

type Notification struct {
    Message string
    Type    string
}

type Season struct {
    Season   int64
    Episodes []TvEpisode
}

func (s *Season) DownloadStatus() string {
    downloaded := 0
    for _, e := range s.Episodes {
        if e.Status > 1 {
            downloaded = downloaded + 1
        }
    }

    return fmt.Sprintf("%d/%d", downloaded, len(s.Episodes))
}

func (s *Season) Percent() int {
    downloaded := 0
    for _, e := range s.Episodes {
        if e.Status > 1 {
            downloaded = downloaded + 1
        }
    }
    x := float64(downloaded) / float64(len(s.Episodes)) * 100
    return int(x)
}

type Seasons []Season

type FSDirectory struct {
    Type string
    Name string
    Path string
    Size int
}

type FSFile struct {
    Type string
    Name string
    Path string
    Size string
}

type FileSystem struct {
    Directories []FSDirectory
    Files       []FSFile
}

type response struct {
    Channel struct {
        Item []Release
    }
}
