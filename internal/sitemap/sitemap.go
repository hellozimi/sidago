package sitemap

import (
	"encoding/xml"
	"time"
)

type sitemapItem struct {
	XMLName      xml.Name  `xml:"url"`
	Location     string    `xml:"loc"`
	LastModified time.Time `xml:"lastmod"`
}

type Sitemap struct {
	XMLName xml.Name      `xml:"http://www.sitemaps.org/schemas/sitemap/0.9 urlset"`
	Items   []sitemapItem `xml:"urlset`
}

func (s *Sitemap) Add(location string, lastModified time.Time) {
	s.Items = append(s.Items, sitemapItem{
		Location:     location,
		LastModified: lastModified,
	})
}

func (s *Sitemap) Render() ([]byte, error) {
	x, err := xml.Marshal(s)
	if err != nil {
		return []byte{}, err
	}

	return append([]byte(xml.Header), x...), nil
}

func New() *Sitemap {
	return &Sitemap{}
}
