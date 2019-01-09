package sitemap

import (
	"bytes"
	"testing"
	"time"
)

type item struct {
	path   string
	edited time.Time
}

func TestSitemap(t *testing.T) {

	date, _ := time.Parse(time.RFC3339, "2018-04-24T15:44:00Z")
	cases := []struct {
		items []item
		e     []byte
		elen  int
		err   error
	}{
		{
			items: []item{},
			elen:  0,
			e: []byte(`<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9"></urlset>`),
		},
		{
			items: []item{{"https://example.com/", date}},
			elen:  1,
			e: []byte(`<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9"><url><loc>https://example.com/</loc><lastmod>2018-04-24T15:44:00Z</lastmod></url></urlset>`),
		},
	}

	for _, x := range cases {
		s := New()
		for _, i := range x.items {
			s.Add(i.path, i.edited)
		}

		l := len(s.Items)
		if l != x.elen {
			t.Errorf("expected items len to equal '%v' but was '%v", x.elen, l)
		}

		res, err := s.Render()
		if err != x.err {
			t.Errorf("expected error to equal '%v' but was '%v'", x.err, err)
		}

		if bytes.Compare(res, x.e) != 0 {
			t.Errorf("expected render to equal '%v' but was '%v'", string(x.e), string(res))
		}
	}
}
