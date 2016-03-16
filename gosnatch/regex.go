package gosnatch

import (
	"fmt"
	"regexp"

	log "github.com/gosnatch/gosnatch/gosnatch/Godeps/_workspace/src/github.com/sirupsen/logrus"
)

// List containig all compiled regexes
type RegexList struct {
	Pattern    string
	Name       string
	IsCompiled bool
	Example    string
	Matcher    *regexp.Regexp
	NamedMap   map[string]string
}

func (r *RegexList) namedMatch(release string) {
	md := map[string]string{}
	n1 := r.Matcher.SubexpNames()
	match := r.Matcher.FindAllStringSubmatch(release, -1)

	if len(match) > 0 {
		fields := log.Fields{}
		md["regex_name"] = r.Name
		fields["release"] = release

		for i, n := range match[0] {

			if n1[i] != "" {
				md[n1[i]] = n
				fields[n1[i]] = n
			}

		}
		log.WithFields(fields).Debug(r.Name, " regex matched")
	}
	r.NamedMap = md
}

// compile the default regexes
func CompileRegexes() []RegexList {
	var err error
	regx := []RegexList{}

	stdRepRegex := RegexList{}
	stdRepRegex.Pattern = `(?i)^(?P<series_name>.+?)[. _-]+s(?P<season_num>\d+)[. _-]*e(?P<ep_num>\d+)([. _-]+s(?P<extra_season_num>)[. _-]*e(?P<extra_ep_num>\d+))+[. _-]*((?P<extra_info>.+?)-(?P<release_group>[^- ]+([. _-]\[.*\])?))?$`
	stdRepRegex.Name = "standart_repeat"
	stdRepRegex.Example = "Show.Name.S01E02.S01E03.Source.Quality.Etc-Group"
	stdRepRegex.Matcher, err = regexp.Compile(stdRepRegex.Pattern)
	if err == nil {
		stdRepRegex.IsCompiled = true
	} else {
		fmt.Println(err)
	}

	regx = append(regx, stdRepRegex)

	fovRepRegex := RegexList{}
	fovRepRegex.Pattern = `(?i)^(?P<series_name>.+?)[. _-]+(?P<season_num>\d+)x(?P<ep_num>\d+)([. _-]+(?P<extra_season_num>\d+)x(?P<extra_ep_num>\d+))+[. _-]*((?P<extra_info>.+?)-(?P<release_group>[^- ]+))?$`
	fovRepRegex.Name = "fov_repeat"
	fovRepRegex.Example = "Show.Name.1x02.1x03.Source.Quality.Etc-Group"
	fovRepRegex.Matcher, err = regexp.Compile(fovRepRegex.Pattern)
	if err == nil {
		fovRepRegex.IsCompiled = true
	} else {
		fmt.Println(err)
	}

	regx = append(regx, fovRepRegex)

	stdRegex := RegexList{}
	//stdRegex.Pattern = `(?i)^((?P<series_name>.+?)[. _-]+)?(\()?s(?P<season_num>\d+)[. _-]*e(?P<ep_num>\d+)(\))?(([. _-]*e|-)(?P<extra_ep_num>\d+)(\))?)*[. _-]*((?P<extra_info>.+?)-(?P<release_group>[^- ]+([. _-]\[.*\])?))?$`
	stdRegex.Pattern = `(?i)^((?P<series_name>.+?)[. _-]+)?(\()?s(?P<season_num>\d+)[. _-]*e(?P<ep_num>\d+)(\))?(([. _-]*e|-)(?P<extra_ep_num>\d+)(\))?)*[. _-]*((?P<extra_info>.+?)[ -.](?P<release_group>[^- ]+([. _-]\[.*\])?))?$`
	stdRegex.Name = "standard"
	stdRegex.Example = "Show.Name.S01E02.Source.Quality.Etc-Group"
	stdRegex.Matcher, err = regexp.Compile(stdRegex.Pattern)
	if err == nil {
		stdRegex.IsCompiled = true
	} else {
		fmt.Println(err)
	}

	regx = append(regx, stdRegex)

	fovRegex := RegexList{}
	fovRegex.Pattern = `(?i)^((?P<series_name>.+?)[\[. _-]+)?(?P<season_num>\d+)x(?P<ep_num>\d+)(([. _-]*x|-)(?P<extra_ep_num>\d+))*[\]. _-]*((?P<extra_info>.+?)-(?P<release_group>[^- ]+))?$`
	fovRegex.Name = "fov"
	fovRegex.Example = "Show_Name.1x02.Source_Quality_Etc-Group"
	fovRegex.Matcher, err = regexp.Compile(fovRegex.Pattern)
	if err == nil {
		fovRegex.IsCompiled = true
	} else {
		fmt.Println(err)
	}

	regx = append(regx, fovRegex)

	stupidRegex := RegexList{}
	stupidRegex.Pattern = `(?i)(?P<release_group>.+?)-\w+?[\. ]?(?P<season_num>\d{1,2})(?P<ep_num>\d{2})$`
	stupidRegex.Name = "stupid"
	stupidRegex.Example = "tpz-abc102"
	stupidRegex.Matcher, err = regexp.Compile(stupidRegex.Pattern)
	if err == nil {
		stupidRegex.IsCompiled = true
	} else {
		fmt.Println(err)
	}

	regx = append(regx, stupidRegex)

	verboseRegex := RegexList{}
	verboseRegex.Pattern = `(?i)^(?P<series_name>.+?)[. _-]+season[. _-]+(?P<season_num>\d+)[. _-]+episode[. _-]+(?P<ep_num>\d+)[. _-]+(?P<extra_info>.+)$`
	verboseRegex.Name = "verbose"
	verboseRegex.Example = "tpz-abc102"
	verboseRegex.Matcher, err = regexp.Compile(verboseRegex.Pattern)
	if err == nil {
		verboseRegex.IsCompiled = true
	} else {
		fmt.Println(err)
	}

	regx = append(regx, verboseRegex)

	seasonRegex := RegexList{}
	seasonRegex.Pattern = `(?i)^((?P<series_name>.+?)[. _-]+)?s(eason[. _-])?(?P<season_num>\d+)[. _-]*[. _-]*((?P<extra_info>.+?)-(?P<release_group>[^- ]+([. _-]\[.*\])?))?$`
	seasonRegex.Name = "season_only"
	seasonRegex.Example = "Show.Name.S01.Source.Quality.Etc-Group"
	seasonRegex.Matcher, err = regexp.Compile(seasonRegex.Pattern)
	if err == nil {
		seasonRegex.IsCompiled = true
	} else {
		fmt.Println(err)
	}

	regx = append(regx, seasonRegex)

	for _, r := range regx {
		if r.IsCompiled {
			log.WithFields(log.Fields{"name": r.Name, "pattern": r.Pattern}).Debug("regex compiled")

			if testRegex(r) {
				log.WithFields(log.Fields{"name": r.Example, "pattern": r.Pattern}).Debug("regex test matched")
			}
		}
	}

	return regx

}

// test a regex against its example
func testRegex(regx RegexList) bool {
	if regx.Example != "" {
		matched := regx.Matcher.MatchString(regx.Example)
		if matched {
			return true
		}
	}

	return false
}
