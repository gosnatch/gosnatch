package gosnatch

import (
    "fmt"
    "io/ioutil"
    "strconv"
    "strings"

    log "github.com/sirupsen/logrus"
)

// load providers from file
func loadProviders(filepath string) []Provider {
    r, err := ioutil.ReadFile(filepath)
    if err != nil {
        log.Error(err)
    }

    f := []Provider{}

    lines := strings.Split(string(r), "\n")

    for _, line := range lines {

        if len(line) != 0 {
            cur := Provider{}

            split := strings.Split(line, "|")
            cur.Name = split[0]
            cur.URL = split[1]
            cur.APIKey = split[2]
            cats := []string{}
            // convert string to int for categories
            for _, v := range strings.Split(split[3], ",") {
                _, err := strconv.Atoi(v)
                if err == nil {
                    cats = append(cats, v)
                }

            }
            cur.Categorys = cats

            f = append(f, cur)
        }

    }

    return f
}

// generate all search urls for this provider
func getSearchUrls(p Provider, rageid int, showname string, season int, episode int) []string {
    urls := []string{}
    search := getSearchUrl(p, showname, season, episode)
    urls = append(urls, search)

    if rageid > 0 {
        rage := getRageSearchUrl(p, rageid, season, episode)
        urls = append(urls, rage)
    }

    return urls
}

// generate a url for name search
func getSearchUrl(provider Provider, showname string, season int, episode int) string {
    url := fmt.Sprintf("%sapi/?t=tvsearch&apikey=%s&cat=%s&season=%d&ep=%d&q=%s", provider.URL, provider.APIKey, strings.Join(provider.Categorys, ", "), season, episode, showname)
    return url
}

// generate a url for tvrage search
func getRageSearchUrl(provider Provider, rageid int, season int, episode int) string {
    url := fmt.Sprintf("%sapi/?t=tvsearch&apikey=%s&cat=%s&rid=%d&season=%d&ep=%d", provider.URL, provider.APIKey, strings.Join(provider.Categorys, ", "), rageid, season, episode)
    return url
}

// generate a url for rss cache update
func getProviderUrl(provider Provider) string {
    url := fmt.Sprintf("%sapi/?t=search&apikey=%s&cat=%s", provider.URL, provider.APIKey, strings.Join(provider.Categorys, ", "))
    return url
}
