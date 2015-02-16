package gosnatch

import (
    "fmt"
    log "github.com/sirupsen/logrus"
    "github.com/spf13/viper"
    "os"
    "path"
    "regexp"
    "strconv"
)

var regexes = CompileRegexes()

// checks if name contains ALL of the given slices
func checkName(check []string, title string) bool {
    ret := false
    for _, v := range check {
        //log.Info("check for " , v)
        re := regexp.MustCompile("(?i).*" + v + ".*")
        if re.MatchString(title) {
            ret = true
        } else {
            return false
        }
    }

    return ret
}

func getExcetptions(tvdbid int) {
    //http://tvdb.cytec.us/v2/names/95451
    // {
    //     "95451":[
    //         "The Good Wife",
    //         "Good Wife"
    //     ]
    // }
}

// support for SickBeards quality namings
func sickbeardQuality(r Release) (quality int, qualityname string) {
    if checkName([]string{"sd.tv"}, r.Title) {
        return 1, "SD TV"
    }

    if checkName([]string{"sd.dvd"}, r.Title) {
        return 2, "SD DVD"
    }

    if checkName([]string{"hd.tv"}, r.Title) {
        return 3, "HD TV"
    }

    if checkName([]string{"rawhd.tv"}, r.Title) {
        return 4, "RawHD TV"
    }

    if checkName([]string{"1080p", "hd.tv"}, r.Title) {
        return 5, "1080p HD TV"
    }

    if checkName([]string{"720p", "bluray"}, r.Title) {
        return 8, "720 BluRay"
    }

    if checkName([]string{"1080p", "bluray"}, r.Title) {
        return 9, "1080p BluRay"
    }

    return 0, "Unknown"
}

// check quality of a given release
// TODO: return a Proper QualityDefinition struct
func checkQuality(r Release) (quality int, quaityname string) {

    if (checkName([]string{"(pdtv|hdtv|dsr|hdtvrip|webrip|webhdrip)(.repack)?.(xvi-?d|(h|x).?264)"}, r.Title) && !checkName([]string{"(720|1080)[pi]"}, r.Title)) || (checkName([]string{"videomann"}, r.Title) && !checkName([]string{"(720|1080)[pi]"}, r.Title)) {
        return 1, "SD TV"
    }

    if checkName([]string{"(dvdrip|bdrip|blurayrip|ituneshd)(.repack)?(.ws)?.(xvi-?d|divx|(h|x).?264)"}, r.Title) && !checkName([]string{"(720|1080)[pi]"}, r.Title) {
        return 2, "SD DVD"
    }

    if (checkName([]string{"720p", "hdtv", "(h|x).?264"}, r.Title) && !checkName([]string{"(1080)[pi]"}, r.Title)) || checkName([]string{"videomann", "720p"}, r.Title) {
        return 3, "HD TV"
    }

    if checkName([]string{"720p|1080i", "hdtv", "mpeg-?2"}, r.Title) {
        return 4, "RawHD TV"
    }

    if checkName([]string{"1080p", "hdtv", "(h|x)264"}, r.Title) {
        return 5, "1080p HD TV"
    }

    if checkName([]string{"720p", "web.?dl|webrip|ituneshd|webhd(rip)?"}, r.Title) || checkName([]string{"720p", "(webhd|itunes)", "(h|x).?264"}, r.Title) {
        return 6, "720p WEB-DL"
    }

    if checkName([]string{"1080p", "(webhd|itunes)", "(h|x).?264"}, r.Title) || checkName([]string{"1080p", "web.?dl|webrip|webhd(rip)?"}, r.Title) {
        return 7, "1080p WEB-DL"
    }

    if checkName([]string{"720p", "bluray|hddvd|b[r|d](rip)?", "(h|x)264"}, r.Title) {
        return 8, "720p BluRay"
    }

    if checkName([]string{"1080p", "bluray|hddvd|b[r|d](rip)?", "(h|x)264"}, r.Title) {
        return 9, "1080p BluRay"
    }

    return 0, "Unknown"
}

// download nzb for this release to the snatchfolder
func downloadNZB(rel Release) error {
    fname := fmt.Sprintf("%s.nzb", rel.Title)

    if !fileExists(viper.GetString("SnatchFolder")) {
        os.MkdirAll(viper.GetString("SnatchFolder"), 0777)
    }

    err := downloadToFile(rel.Link, path.Join(viper.GetString("SnatchFolder"), fname))
    if err != nil {
        return err
    }
    rel.markSnatched()
    log.Info("Snatched Release ", rel.Title, " from Indexer ", rel.Provider.Name)
    return nil
}

// check if release contains required and/or ignored regex patterns
func isGoodRelease(release *Release, show *TvShow) bool {
    Required := []string{}
    Ignore := []string{}
    Ignore = append(Ignore, ".*(sub|subbed).*")
    if show.Language == "de" {
        Required = append(Required, ".*(german|videomann).*")
    } else {
        Ignore = append(Ignore, ".*(german|videomann).*")
    }

    if len(Required) != 0 {
        for _, w := range Required {
            re := regexp.MustCompile("(?i)" + w)
            if !re.MatchString(release.Title) {
                log.WithFields(log.Fields{"release": release.Title, "reason": w}).Debug("release is missing required word. skipping")
                return false
            }
        }
    }

    if len(Ignore) != 0 {
        for _, w := range Ignore {
            re := regexp.MustCompile("(?i)" + w)
            if re.MatchString(release.Title) {
                log.WithFields(log.Fields{"release": release.Title, "reason": w}).Debug("release contains ignore word. skipping")
                return false
            }
        }
    }
    return true
}

// iter over a list of releases and return the best possible result (quality)
func getBestRelease(rels []Release, quali []int) Release {
    bestRelease := Release{}
    for _, r := range rels {

        // r.Quality, r.QualityString = checkQuality(r)

        // go over accepted quali
        wantQuality := false
        for _, a := range quali {
            if a == r.Quality {
                wantQuality = true
            }
        }

        if !wantQuality {
            log.WithFields(log.Fields{"release": r.Title, "quality": r.QualityString, "qualityValue": r.Quality, "need": quali}).Debug("skipping release because we dont want this quality")
        } else {
            if bestRelease.Quality < r.Quality {
                bestRelease = r
            } else {
                log.WithFields(log.Fields{"current": r.Title, "best": bestRelease.Title, "qualityValue": r.Quality}).Debug("skipping release because we already found a better quality")
            }
        }

    }

    return bestRelease
}

// checks if current release is a needed release for given show
func isNeededRelease(release *Release, show *TvShow) bool {
    wanted := show.Wanted()

    md := map[string]string{}
    for _, curregex := range regexes {
        if curregex.IsCompiled {

            n1 := curregex.Matcher.SubexpNames()
            match := curregex.Matcher.FindAllStringSubmatch(release.Title, -1)

            if len(match) > 0 {
                fields := log.Fields{}
                md["regex_name"] = curregex.Name
                fields["release"] = release.Title

                for i, n := range match[0] {

                    if n1[i] != "" {
                        md[n1[i]] = n
                        fields[n1[i]] = n
                    }

                }
                log.WithFields(fields).Debug(curregex.Name, " regex matched")
                break
            }

        }
    }

    if len(md) > 0 {
        ep_num, _ := strconv.Atoi(md["ep_num"])
        season_num, _ := strconv.Atoi(md["season_num"])

        if md["regex_name"] == "season_only" && show.SeasonWanted(season_num) {
            log.WithFields(log.Fields{"release": release.Title, "season": season_num}).Error("whole season is wanted to im going to accept this release")
            release.Season = wanted
            release.SeasonNum = season_num
            return true
        }

        for _, ep := range wanted {
            if ep.Season == season_num && ep.Episode == ep_num {
                log.WithFields(log.Fields{"show": md["series_name"], "season": season_num, "episode": ep_num}).Infof("%s matches", release.Title)
                release.Episode = ep
                return true
            }
        }
    }

    return false
}

// is the release valid for the given show?
func isValidShowRelease(release Release, show *TvShow) bool {

    md := map[string]string{}
    for _, curregex := range regexes {
        if curregex.IsCompiled {

            n1 := curregex.Matcher.SubexpNames()

            match := curregex.Matcher.FindAllStringSubmatch(release.Title, -1)

            if len(match) > 0 {
                fields := log.Fields{}
                md["regex_name"] = curregex.Name
                fields["release"] = release.Title

                for i, n := range match[0] {

                    if n1[i] != "" {
                        md[n1[i]] = n
                        fields[n1[i]] = n
                    }

                }
                break
            }

        }
    }
    if len(md) > 0 {
        showNames := []string{}
        showNames = append(showNames, show.ShowName)

        for _, v := range show.getExceptions() {
            showNames = append(showNames, v.ShowName)
        }

        //check all available names of a show
        for _, name := range showNames {

            if md["series_name"] == name || cleanSeriesName(md["series_name"]) == name {
                log.WithFields(log.Fields{"series": show.ShowName}).Debugf("%s seems to be a valid release for ", release.Title)
                return true
            }
        }

    }
    return false
}

// returns a named map of the string to match
func parseFilename(filename string) map[string]string {
    md := map[string]string{}
    for _, curregex := range regexes {
        if curregex.IsCompiled {

            n1 := curregex.Matcher.SubexpNames()

            match := curregex.Matcher.FindAllStringSubmatch(filename, -1)

            if len(match) > 0 {
                fields := log.Fields{}
                md["regex_name"] = curregex.Name
                fields["release"] = filename

                for i, n := range match[0] {

                    if n1[i] != "" {
                        md[n1[i]] = n
                        fields[n1[i]] = n
                    }

                }
                log.WithFields(fields).Debug(curregex.Name, " regex matched")
                break
            }

        }
    }
    return md
}
