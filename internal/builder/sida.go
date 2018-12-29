package builder

import (
	"io/ioutil"
	"path/filepath"

	"github.com/hellozimi/sidago/internal/config"
)

type Sida struct {
	pages    []Page
	config   config.Config
	basePath string
}

func (s *Sida) Sitemap() {

}

func (s *Sida) generatePages() {
	assets := s.scanDirPages("pages")
	assets = append(assets, s.scanDirPages("posts")...)
	pages := []*Page{}

	for _, f := range assets {
		pages = append(pages, newPage(f))
	}
}

func (s *Sida) scanDirPages(dir string) []string {
	files, err := ioutil.ReadDir(filepath.Join(s.basePath, dir))
	out := []string{}
	if err != nil {
		return out
	}

	for _, f := range files {
		if filepath.Ext(f.Name()) == ".md" {
			out = append(out, filepath.Join(s.basePath, dir, f.Name()))
		}
	}

	return out
}

func (s *Sida) Build() {
	s.generatePages()
}

func NewSida(path string, config config.Config) *Sida {
	s := Sida{
		config:   config,
		basePath: path,
	}

	return &s
}
