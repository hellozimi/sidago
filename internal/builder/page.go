package builder

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"time"

	"github.com/hellozimi/sidago/internal/parser"

	"github.com/hellozimi/sidago/helpers"
	"github.com/hellozimi/sidago/internal/mmark"
)

// PageKind type
type PageKind int

var bre = regexp.MustCompile("^(19[0-9]{2}|2[0-9]{3})-(0[1-9]|1[012])-([123]0|[012][1-9]|31)_([-_]*[a-zA-Z0-9]+([-_]*[a-zA-Z0-9]+)*)\\.md$")

const (
	kindIndex PageKind = 0
	kindPage  PageKind = 1
	kindBlog  PageKind = 2
)

// Page struct
type Page struct {
	path       string
	kind       PageKind
	url        string
	sida       *Sida
	Global     *GlobalInfo
	rawContent []byte
	content    []byte
	summary    string
	Title      string
	Date       time.Time
	Draft      bool
	PageMeta
}

// Kind returns the page kind
func (p *Page) Kind() PageKind {
	return p.kind
}

func (p *Page) IsHome() bool {
	return p.kind == kindIndex
}

// Content for template rendering
func (p *Page) Content() template.HTML {
	return template.HTML(string(p.content))
}

// Summary for template rendering
func (p *Page) Summary() template.HTML {
	return template.HTML(p.summary)
}

// Slug returns the slug of the title
func (p *Page) Slug() string {
	return helpers.Slugify(p.Title)
}

func (p *Page) init() {
	pageMeta := newPageMeta(p.path)
	pageMeta.page = p

	p.PageMeta = pageMeta

	filename, _ := helpers.FileAndExt(p.path)
	slug, date := helpers.SlugAndDateFromFile(filename)
	if date.IsZero() {
		date = time.Now()
	}
	p.Date = date
	p.Title = helpers.Unslugify(slug)

	// Skips content initialization for index
	// because index is a dummy file
	if p.Kind() != kindIndex {
		p.initContent()
	}
}

func (p *Page) initContent() {
	var err error
	p.rawContent, err = ioutil.ReadFile(p.path)
	if err != nil {
		fmt.Println("error reading raw content")
		return
	}

	var body []byte

	if parser.HasFrontmatter(p.rawContent) {
		f, b, err := parser.ParseFrontmatterBody(p.rawContent)
		if err != nil {
			fmt.Printf("error parsing frontmatter on %s:\n\t%v", p.path, err)
			return
		}
		body = b

		p.updateWithFrontmatter(f)
	} else {
		body = p.rawContent
	}

	p.content, err = mmark.ParseBytes(body)
	if err != nil {
		fmt.Println("Error parsing content")
		return
	}

	p.summary = summaryFromContent(string(p.content))
}

func (p *Page) render() string {
	var tpl *template.Template
	var err error
	var layout string

	switch p.Kind() {
	case kindBlog:
		layout = "single"
		break
	case kindPage:
		layout = "page"
		break
	case kindIndex:
		layout = "index"
		break
	}

	tpl = template.Must(template.New("").Funcs(template.FuncMap{}).ParseFiles(
		filepath.Join(p.sida.basePath, "layout/base.html"),
		filepath.Join(p.sida.basePath, fmt.Sprintf("layout/%s.html", layout)),
		filepath.Join(p.sida.basePath, "layout/partials/header.html"),
		filepath.Join(p.sida.basePath, "layout/partials/footer.html"),
	))

	if err != nil {
		panic(err)
	}

	var tplb bytes.Buffer
	if err := tpl.ExecuteTemplate(&tplb, "base.html", p); err != nil {
		panic(fmt.Sprintf("Unable to render base.html - error: %v", err))
	}

	return tplb.String()
}

func newPage(path string, sida *Sida) *Page {
	f := filepath.Base(path)
	var kind PageKind
	if bre.MatchString(f) {
		kind = kindBlog
	} else {
		kind = kindPage
	}

	p := &Page{
		path: path,
		kind: kind,
		sida: sida,
	}
	p.Global = &sida.Global

	p.init()

	return p
}

func newIndex(sida *Sida) *Page {
	kind := kindIndex

	p := &Page{
		path: "index",
		kind: kind,
		sida: sida,
	}
	p.Global = &sida.Global

	p.init()

	return p
}

func summaryFromContent(h string) string {
	stripped := helpers.StripHTML(h)
	return helpers.TruncateFull(stripped, 20)
}
