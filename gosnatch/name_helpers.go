package gosnatch

import (
	"encoding/json"
	"fmt"
	"github.com/gosnatch/gosnatch/gosnatch/Godeps/_workspace/src/github.com/astaxie/beego/orm"
	log "github.com/gosnatch/gosnatch/gosnatch/Godeps/_workspace/src/github.com/sirupsen/logrus"
	"regexp"
	"strconv"
	"strings"
)

// remove unneeded stuff from series names
func cleanSeriesName(name string) string {
	re := regexp.MustCompile(`(\D)\.(\D)`)
	name = re.ReplaceAllString(name, "$1 $2")
	re = regexp.MustCompile(`(\d)\.(\d{4})`)
	name = re.ReplaceAllString(name, "$1 $2")
	re = regexp.MustCompile(`(\D)\.`)
	name = re.ReplaceAllString(name, "$1")
	re = regexp.MustCompile(`\.(\D)`)
	name = re.ReplaceAllString(name, " $1")
	re = regexp.MustCompile(`_`)
	name = re.ReplaceAllString(name, " ")
	re = regexp.MustCompile(`-$`)
	name = re.ReplaceAllString(name, " ")
	return name

}

// takes a name and returns a slice of possible names
func normalizeName(name string) []string {

	if name == "" {
		return []string{}
	}

	nameList := []string{}
	re := regexp.MustCompile(`(?i)([\. ])and([\. ])`)
	newName := re.ReplaceAllString(name, "\\1&\\2")

	nameList = append(nameList, newName)

	for _, curName := range nameList {
		re := regexp.MustCompile(`(\D)(\d{4})$`)
		nameList = append(nameList, re.ReplaceAllString(curName, "\\1(\\2)"))
	}

	return nameList

}

// reload exceptions for all shows in database
func UpadeShowExceptions() {
	log.Debug("looking for new scene exceptions")
	o := orm.NewOrm()

	exceptions, _ := getAllExceptions()

	for id, ex := range exceptions {

		tvdbid, _ := strconv.Atoi(id)
		for i := 0; i < len(ex); i++ {
			cur := ShowExceptions{}
			o.QueryTable(&ShowExceptions{}).Filter("tvdb_id", tvdbid).Filter("show_name", ex[i]).One(&cur)
			if cur.Id == 0 {
				show := TvShow{}
				o.QueryTable(&TvShow{}).Filter("tvdb_id", tvdbid).All(&show)
				if show.Id != 0 {
					cur.ShowName = ex[i]
					cur.TvdbId = tvdbid
					cur.Series = &show
					o.Insert(&cur)
				}
			}
		}

	}
}

// get all exceptions from server
func getAllExceptions() (map[string][]string, error) {
	data := map[string][]string{}

	body, _ := getUrl("http://tvdb.cytec.us/v2/names")

	err := json.Unmarshal(body, &data)

	if err != nil {
		fmt.Println(err)
		return data, err
	}

	return data, nil
}

//TODO: use getAll and try to access tvdbid in map if not in map return empty without an error!
func getShowExceptions(tvdbid int) ([]string, error) {
	url := fmt.Sprintf("http://tvdb.cytec.us/v2/names/%d", tvdbid)

	body, _ := getUrl(url)

	data := map[string][]string{}
	err := json.Unmarshal(body, &data)

	if err != nil {
		return []string{}, err
	}

	return data[strconv.Itoa(tvdbid)], nil
}

func sanitizeSceneName(name string) string {
	badChars := []string{",", ":", "(", ")", "'", "!", "?", "’"}

	for _, x := range badChars {
		name = strings.Replace(name, x, "", -1)
	}

	name = strings.Replace(name, "- ", ".", -1)
	name = strings.Replace(name, " ", ".", -1)
	name = strings.Replace(name, "&", "and", -1)
	name = strings.Replace(name, "/", ".", -1)

	return name

}
