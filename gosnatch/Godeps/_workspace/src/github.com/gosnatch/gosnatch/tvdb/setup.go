// Package tvdb provides API for thetvdb.com
package tvdb

// https://github.com/rharter/tvdb

import (
    "encoding/xml"
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
)

const (
    // Known TVDB API endpoint.
    TVDB_API = "http://thetvdb.com/api/"
    // Default TVDB API key. If you're using this for anything substantial, use a different one.
    TVDB_DEFAULT_KEY = "CACE3A94B49F1566"
)

// TVDB settings.
type TVDB struct {
    // Location to use for API access.
    Location string

    // Key to use.
    ApiKey string
}

// Create TVDB config with location & key from well-known constants.
func Open(key, url string) *TVDB {
    return &TVDB{
        Location: url,
        ApiKey:   key,
    }
}

// Returns a constructed base url for the api
func (t *TVDB) baseUrl() string {
    return fmt.Sprintf("%s/api/%s", t.Location, t.ApiKey)
}

func (t *TVDB) GetImageUrl(path string) string {
    return fmt.Sprintf("%s/banners/%s", t.Location, path)
}

func (t *TVDB) GetURL(path string, args *url.Values) (string, error) {
    query, err := url.Parse(t.Location)
    if err != nil {
        return "", err
    }

    if query, err = query.Parse(path); err != nil {
        return "", err
    }

    if args != nil {
        query.RawQuery = args.Encode()
    }
    return query.String(), nil
}

// Do an HTTP GET with given URL, return result as []byte.
func httpGet(query string) ([]byte, error) {
    resp, err := http.Get(query)

    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)

    if err != nil {
        return nil, err
    }

    return body, nil
}

// Query given relative API with given args. Real URL is formed by joining t.Location and path.
func (t *TVDB) QueryURL(path string, args *url.Values) ([]byte, error) {
    query, err := t.GetURL(path, args)
    if err != nil {
        return nil, err
    }
    return httpGet(query)
}

// Query given relative API with args, unmarshal xml result.
func (t *TVDB) QueryAndUnmarshal(path string, args *url.Values, result interface{}) error {
    body, err := t.QueryURL(path, args)
    if err != nil {
        return err
    }
    return xml.Unmarshal(body, result)
}
