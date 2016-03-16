package gosnatch

import (
	"fmt"
	"github.com/gosnatch/gosnatch/gosnatch/Godeps/_workspace/src/github.com/astaxie/beego/orm"
	log "github.com/gosnatch/gosnatch/gosnatch/Godeps/_workspace/src/github.com/sirupsen/logrus"
	"github.com/gosnatch/gosnatch/gosnatch/Godeps/_workspace/src/github.com/spf13/viper"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"
)

func startPostProcessing() {

}

type PostProcessor struct {
	ReleaseName string `json:"release" `
	BaseFolder  string `json:"folder" binding:"required"`
	Failed      int    `json:"failed" `
	history     History
	format      NameFormatter
}

func (p *PostProcessor) History() {
	_, reldir := path.Split(p.BaseFolder)
	name := strings.TrimSuffix(reldir, filepath.Ext(p.BaseFolder))
	fmt.Println(p.BaseFolder, reldir, name)
	o := orm.NewOrm()
	rel := History{}
	fmt.Println(reldir)
	o.QueryTable(&History{}).Filter("title__in", p.ReleaseName, reldir, name).One(&rel)
	log.Info("try to load release from history table")
	p.history = rel
}

func (p *PostProcessor) targetPath() string {
	nametemplate := viper.GetString("NamingPattern")
	t := template.Must(template.New("filename").Parse(nametemplate))

	var data = NameFormatter{}
	data.SeasonNumber = 1
	data.ShowName = "The Newsroom"
	err := t.Execute(os.Stdout, data)
	if err != nil {
		fmt.Println("executing template:", err)
	}

	return "ff"
}

// checks all subdirectorys for files
func (p *PostProcessor) checkDirectory() {

}

func (p *PostProcessor) processRelease() {
	p.targetPath()
	_, reldir := path.Split(p.BaseFolder)

	if p.history.Id == 0 {
		log.Debug("Unable to find a matching release in History")
		for _, re := range regexes {
			if re.IsCompiled {

				//found a match now get the info
				if re.Matcher.MatchString(reldir) {
					re.namedMatch(reldir)
					fmt.Println(re.NamedMap["series_name"])
					break
				}

			}
		}
	} else {
		//should check which quality and what to do next...
		o := orm.NewOrm()
		o.LoadRelated(&p.history, "TvShow")
		o.LoadRelated(&p.history, "TvEpisode")
		o.LoadRelated(&p.history, "Quality")
		log.Debug("found release in history table")
		log.WithFields(log.Fields{"season": p.history.Season, "episode": p.history.Episode, "show": p.history.TvShow.ShowName}).Debug(p.history)
		_, reldir := path.Split(p.BaseFolder)
		moveto := path.Join(p.history.TvShow.Path, reldir)
		if moveFile(p.BaseFolder, moveto, "move") {
			p.history.TvEpisode.Status = STATE_DOWNLOADED
			p.history.TvEpisode.Quality = p.history.Quality
			fmt.Println(p.history.TvEpisode)
			p.history.Action = STATE_DOWNLOADED
			o.Update(&p.history)
			o.Update(p.history.TvEpisode)
			log.Info("episode marked as downloaded")
		} else {
			log.Error("unable to move file to wanted location")
		}

	}

}
