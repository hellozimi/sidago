package builder

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/hellozimi/sidago/fs"
	"github.com/hellozimi/sidago/internal/builder/config"
)

type Sida struct {
	allPages []*Page
	config   config.Config
	basePath string
	Global   GlobalInfo
}

func (s *Sida) generatePages() {
	assets := s.scanDirPages("pages")
	assets = append(assets, s.scanDirPages("posts")...)
	allPages := []*Page{}

	for _, f := range assets {
		p := newPage(f, s)
		allPages = append(allPages, p)
	}

	allPages = append(allPages, newIndex(s))

	s.allPages = allPages
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

func (s *Sida) Posts() []*Page {
	posts := make([]*Page, 0)
	for _, p := range s.allPages {
		if p.Kind() == KindBlog {
			posts = append(posts, p)
		}
	}
	sort.Slice(posts[:], func(i, j int) bool {
		return posts[i].PageMeta.Date.After(posts[j].PageMeta.Date)
	})
	return posts
}

func (s *Sida) copyStatic() {

}

func (s *Sida) Build() {
	s.generatePages()
	for _, p := range s.allPages {
		fmt.Printf("Generating: %s\n", p.PageMeta.NameComponents.Title)
		op := p.OutputPath()
		dir := filepath.Dir(op)
		o := p.render()
		os.MkdirAll(dir, 0777)
		ioutil.WriteFile(op, []byte(o), 0777)
	}

	os.RemoveAll(filepath.Join(s.basePath, "build/static"))
	fmt.Println("Copying static assets")
	err := fs.CopyDir(
		filepath.Join(s.basePath, "layout/static"),
		filepath.Join(s.basePath, "build/static"),
	)
	if err != nil {
		fmt.Printf("Error copying static assets: %v", err)
	}
}

func (s *Sida) buildInfo() {
	s.Global = GlobalInfo{
		Title:     s.config.GetString("title"),
		baseURL:   s.config.GetString("base_url"),
		Copyright: s.config.GetString("copyright"),
		sida:      s,
	}
}

func NewSida(path string, config config.Config) *Sida {
	s := Sida{
		config:   config,
		basePath: path,
	}

	s.buildInfo()

	return &s
}
