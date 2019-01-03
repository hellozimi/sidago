package builder

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"

	"github.com/hellozimi/sidago/fs"
	"github.com/hellozimi/sidago/internal/builder/config"
)

// Sida main struct
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

// Posts returns all blog posts sorted by date desc
func (s *Sida) Posts() []*Page {
	posts := make([]*Page, 0)
	for _, p := range s.allPages {
		if p.Kind() == kindBlog && !p.Draft {
			posts = append(posts, p)
		}
	}
	sort.Slice(posts[:], func(i, j int) bool {
		return posts[i].Date.After(posts[j].Date)
	})
	return posts
}

// Build starts the build procedure and generates
// all the html files and copies static content to
// the build directory
func (s *Sida) Build() error {
	s.generatePages()

	c := len(s.allPages)
	for i, p := range s.allPages {
		fmt.Printf("\r🔨 Generating page %d/%d", i+1, c)
		op := p.OutputPath()
		dir := filepath.Dir(op)
		o := p.render()
		os.MkdirAll(dir, 0777)
		err := ioutil.WriteFile(op, []byte(o), 0777)
		if err != nil {
			return err
		}
	}

	fmt.Println()

	os.RemoveAll(filepath.Join(s.basePath, "build/static"))
	fmt.Println("📁 Copying static assets")
	err := fs.CopyDir(
		filepath.Join(s.basePath, "layout/static"),
		filepath.Join(s.basePath, "build/static"),
	)
	if err != nil {
		return fmt.Errorf("error copying static assets: %v", err)
	}

	fmt.Printf("\n🚀 Build completed...\n\n")

	return nil
}

func (s *Sida) buildInfo() {
	s.Global = GlobalInfo{
		Title:       s.config.GetString("title"),
		Description: s.config.GetString("description"),
		Copyright:   s.config.GetString("copyright"),
		baseURL:     s.config.GetString("base_url"),
		sida:        s,
	}
}

// NewSida creates the main building block to generate
// a static site
func NewSida(path string, config config.Config) *Sida {
	s := Sida{
		config:   config,
		basePath: path,
	}

	s.buildInfo()

	return &s
}
