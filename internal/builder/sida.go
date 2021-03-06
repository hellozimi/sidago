package builder

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/hellozimi/sidago/internal/sitemap"

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

// VisiblePages returns a list of all pages which
// filter out drafts
func (s *Sida) visiblePages() []*Page {
	allPages := make([]*Page, 0)
	for _, p := range s.allPages {
		if !p.Draft {
			allPages = append(allPages, p)
		}
	}
	return allPages
}

func (s *Sida) makeSitemap() *sitemap.Sitemap {
	sm := sitemap.New()
	for _, p := range s.visiblePages() {
		location := string(p.URL())
		sm.Add(location, p.Date)
	}
	return sm
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
		return fmt.Errorf("❌ error copying static assets: %v", err)
	}

	sm, err := s.makeSitemap().Render()
	if err != nil {
		return err
	}

	sitemapFile := filepath.Join(s.basePath, "build", "sitemap.xml")
	fmt.Printf("🌳 Creating sitemap\n")
	err = ioutil.WriteFile(sitemapFile, sm, 0777)
	if err != nil {
		fmt.Printf("❌ couldn't write sitemap.xml\n ")
		return err
	}

	fmt.Printf("\n🚀 Build completed...\n\n")

	return nil
}

func (s *Sida) buildInfo() {
	baseURL := s.config.GetString("base_url")

	if len(baseURL) > 0 && !strings.HasSuffix(baseURL, "/") {
		baseURL = baseURL + "/"
	}

	menu := newMenu(s.config.Get("menu"))

	s.Global = GlobalInfo{
		Title:       s.config.GetString("title"),
		Description: s.config.GetString("description"),
		Copyright:   s.config.GetString("copyright"),
		baseURL:     s.config.GetString("base_url"),
		Menu:        menu,
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
