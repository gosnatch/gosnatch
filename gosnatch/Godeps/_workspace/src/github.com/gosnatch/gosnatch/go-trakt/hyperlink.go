package trakt

import (
	"net/url"

	"github.com/gosnatch/gosnatch/gosnatch/Godeps/_workspace/src/github.com/jingweno/go-sawyer/hypermedia"
)

type M map[string]interface{}

type Hyperlink string

func (l Hyperlink) Expand(m M) (u *url.URL, err error) {
	sawyerHyperlink := hypermedia.Hyperlink(string(l))
	u, err = sawyerHyperlink.Expand(hypermedia.M(m))
	return
}
