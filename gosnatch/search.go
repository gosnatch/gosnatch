package gosnatch

import (
	"github.com/gosnatch/gosnatch/gosnatch/Godeps/_workspace/src/github.com/astaxie/beego/orm"
	rss "github.com/gosnatch/gosnatch/gosnatch/Godeps/_workspace/src/github.com/jteeuwen/go-pkg-rss"
	_ "github.com/gosnatch/gosnatch/gosnatch/Godeps/_workspace/src/github.com/mattn/go-sqlite3"
	log "github.com/gosnatch/gosnatch/gosnatch/Godeps/_workspace/src/github.com/sirupsen/logrus"
	"github.com/gosnatch/gosnatch/gosnatch/Godeps/_workspace/src/github.com/spf13/viper"
	"sort"
	"time"
)

// returns a list of rss.item objects
func getRSSContent(uri string, timeout int) []*rss.Item {
	feed := rss.New(timeout, true, chanHandler, itemHandler)

	for {
		if err := feed.Fetch(uri, nil); err != nil {
			log.Errorf("[e] %s: %s", uri, err)
			return []*rss.Item{}
		}
		break
		//return feed.Channels[0].Items
		//<-time.After(time.Duration(feed.SecondsTillUpdate() * 1e9))
	}
	return feed.Channels[0].Items
}

func chanHandler(feed *rss.Feed, newchannels []*rss.Channel) {
	//fmt.Printf("%d new channel(s) in %s\n", len(newchannels), feed.Url)
}
func itemHandler(feed *rss.Feed, ch *rss.Channel, newitems []*rss.Item) {
	//fmt.Printf("%d new item(s) in %s\n", len(newitems), feed.Url)
}

// gets RSS Items and returns a list of valid Releases
func getReleasesFromRss(newitems []*rss.Item) []Release {
	releases := []Release{}

	for _, v := range newitems {
		cur := Release{}
		cur.Id = v.Id
		cur.Title = v.Title
		cur.Link = v.Enclosures[0].Url
		cur.Size = v.Enclosures[0].Length
		releases = append(releases, cur)
	}
	// fmt.Println(releases)
	return releases
}

// search for an specific episode by id, automatically downloads the best found release
func EpisodeSearch(episodeid int) Release {
	releases := []Release{}
	o := orm.NewOrm()
	ep := TvEpisode{}
	o.QueryTable(&TvEpisode{}).Filter("id", episodeid).One(&ep)
	o.LoadRelated(&ep, "TvShow")

	//mark as ep as WANTED !
	ep.Status = STATE_WANTED
	o.Update(&ep)

	for _, p := range loadProviders("./provider.txt") {
		curFeed := []*rss.Item{}

		// if rageid is unknown, search by name
		if ep.TvShow.TvrageId == 0 {
			curFeed = getRSSContent(getSearchUrl(p, ep.TvShow.ShowName, ep.Season, ep.Episode), 2)
		} else {
			curFeed = getRSSContent(getRageSearchUrl(p, ep.TvShow.TvrageId, ep.Season, ep.Episode), 2)
		}

		curRels := getReleasesFromRss(curFeed)

		// if no results are returned, try to search by showname
		if len(curRels) == 0 && ep.TvShow.TvrageId != 0 {
			newFeed := getRSSContent(getSearchUrl(p, ep.TvShow.ShowName, ep.Season, ep.Episode), 2)
			curRels = getReleasesFromRss(newFeed)
		}

		for _, r := range curRels {
			r.Provider = p
			r.Quality, r.QualityString = checkQuality(r)

			if isValidShowRelease(r, ep.TvShow) && isNeededRelease(&r, ep.TvShow) {
				if r.checkFilesize() {
					releases = append(releases, r)
				}

			}
		}

		if len(releases) > 0 {
			bestRelease := getBestRelease(releases, ep.TvShow.AcceptedQualitys())
			if bestRelease.Link != "" {
				_ = downloadNZB(bestRelease)
				return bestRelease
			}

		}
	}
	return Release{}
}

// manual search for an epiosde
// this is similar to EpisodeSearch but doesn't snatch or filter the results in any way
func ManualSearch(episodeid int) []Release {
	releases := []Release{}
	o := orm.NewOrm()
	ep := TvEpisode{}
	o.QueryTable(&TvEpisode{}).Filter("id", episodeid).One(&ep)
	o.LoadRelated(&ep, "TvShow")

	for _, p := range loadProviders("./provider.txt") {
		curFeed := []*rss.Item{}

		if ep.TvShow.TvrageId == 0 {
			curFeed = getRSSContent(getSearchUrl(p, ep.TvShow.ShowName, ep.Season, ep.Episode), 10)
		} else {
			curFeed = getRSSContent(getRageSearchUrl(p, ep.TvShow.TvrageId, ep.Season, ep.Episode), 10)
		}

		curRels := getReleasesFromRss(curFeed)

		//if no releases are found with rage id, try search for showname
		if len(curRels) == 0 && ep.TvShow.TvrageId != 0 {
			newFeed := getRSSContent(getSearchUrl(p, ep.TvShow.ShowName, ep.Season, ep.Episode), 10)
			curRels = getReleasesFromRss(newFeed)
		}

		for _, r := range curRels {
			r.Provider = p
			r.Quality, r.QualityString = checkQuality(r)

			releases = append(releases, r)
		}

	}

	//sort releases by Quality from highest to lowest!
	sort.Sort(ByQuality(releases))

	if len(releases) > 15 {
		releases = releases[0:16]
	}

	return releases
}

// make a backlog search for all that is missing
func BacklogSearch(show TvShow) []Release {
	return []Release{}
}

// returns a list of Proper/repack releases
func ProperSearch() []Release {
	return []Release{}
}

// get last releases from rss cache
func UpdateCache() []Release {
	releases := []Release{}
	for _, p := range loadProviders("./provider.txt") {
		curFeed := getRSSContent(getProviderUrl(p), 5)
		curRels := getReleasesFromRss(curFeed)
		for _, r := range curRels {
			r.Provider = p
			releases = append(releases, r)
		}

	}

	return releases
}

// fetch RSS Feeds and check for needed episodes
func DailySearch(force bool) {

	for {
		if !force {
			log.Debugf("Next RSS search will run in %d minutes", viper.GetInt("SearchFrequency"))
			time.Sleep(time.Duration(viper.GetInt("SearchFrequency")) * time.Minute)
		}
		log.Info("starting RSS search...")
		releases := []Release{}
		found := 0

		for _, p := range loadProviders("./provider.txt") {
			curFeed := getRSSContent(getProviderUrl(p), 5)
			curRels := getReleasesFromRss(curFeed)
			for _, r := range curRels {
				r.Provider = p
				releases = append(releases, r)
			}

		}

		//load all shows and check every release...
		shows := []TvShow{}
		o := orm.NewOrm()
		o.QueryTable(&TvShow{}).All(&shows)

		for _, show := range shows {
			wanted := map[int][]Release{}
			for _, rel := range releases {
				rel.Quality, rel.QualityString = checkQuality(rel)
				if isValidShowRelease(rel, &show) && isNeededRelease(&rel, &show) {
					if isGoodRelease(&rel, &show) {
						if rel.SeasonNum != 0 {
							wanted[rel.SeasonNum] = append(wanted[rel.SeasonNum], rel)
						} else {
							wanted[rel.Episode.Id] = append(wanted[rel.Episode.Id], rel)
						}
					}
				} else {
					//fmt.Println(rel.Title)
				}
			}

			for k, _ := range wanted {
				bestRelease := getBestRelease(wanted[k], show.AcceptedQualitys())
				if bestRelease.Link != "" {
					log.Debugf("best release: %s URL: %s", bestRelease.Title, bestRelease.Link)
					_ = downloadNZB(bestRelease)
					found = found + 1
				}
			}

		}

		if found == 0 {
			log.Info("RSS search finished with no new releases")
		} else {
			log.Infof("RSS search finished, found %d new releases", found)
		}

		if force {
			break
		}
	}

}
