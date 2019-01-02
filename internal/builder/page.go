package builder

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"path/filepath"
	"regexp"

	"github.com/hellozimi/sidago/internal/mmark"
)

type PageKind int

var bre = regexp.MustCompile("^(19[0-9]{2}|2[0-9]{3})-(0[1-9]|1[012])-([123]0|[012][1-9]|31)_([-_]*[a-zA-Z0-9]+([-_]*[a-zA-Z0-9]+)*)\\.md$")

const (
	KindIndex PageKind = 0
	KindPage  PageKind = 1
	KindBlog  PageKind = 2
)

type Page struct {
	path       string
	kind       PageKind
	url        string
	sida       *Sida
	Global     *GlobalInfo
	rawContent []byte
	content    []byte
	PageMeta
}

func (p *Page) Kind() PageKind {
	return p.kind
}

func (p *Page) ContentString() template.HTML {
	return template.HTML(string(p.content))
}

func (p *Page) init() {

	pageMeta := newPageMeta(p.path)
	pageMeta.page = p

	p.PageMeta = pageMeta

	// Skips content initialization for index
	// because index is a dummy file
	if p.Kind() != KindIndex {
		p.initContent()
	}
}

func (p *Page) initContent() {
	var err error
	p.rawContent, err = ioutil.ReadFile(p.path)
	if err != nil {
		fmt.Println("Error reading raw content")
		return
	}
	p.content, err = mmark.Parse(p.path)
	if err != nil {
		return
		fmt.Println("Error parsing content")
	}
}

func (p *Page) render() string {
	var tpl *template.Template
	var err error
	var layout string

	switch p.Kind() {
	case KindBlog:
		layout = "single"
		break
	case KindPage:
		layout = "page"
		break
	case KindIndex:
		layout = "index"
		break
	}

	tpl, err = template.ParseFiles(
		filepath.Join(p.sida.basePath, "layout/base.html"),
		filepath.Join(p.sida.basePath, fmt.Sprintf("layout/%s.html", layout)),
		filepath.Join(p.sida.basePath, "layout/partials/header.html"),
		filepath.Join(p.sida.basePath, "layout/partials/footer.html"),
	)
	tpl = tpl.Funcs(template.FuncMap{})

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
		kind = KindBlog
	} else {
		kind = KindPage
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
	kind := KindIndex

	p := &Page{
		path: "index",
		kind: kind,
		sida: sida,
	}
	p.Global = &sida.Global

	p.init()

	return p
}
