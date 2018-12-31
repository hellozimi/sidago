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
	KindSingle PageKind = 0
	KindBlog   PageKind = 1
)

type Page struct {
	path       string
	kind       PageKind
	url        string
	sida       Sida
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
	var err error
	p.rawContent, err = ioutil.ReadFile(p.path)
	if err != nil {
		fmt.Println("Error reading raw content")
	}
	p.content, err = mmark.Parse(p.path)
	if err != nil {
		fmt.Println("Error parsing content")
	}

	p.PageMeta = newPageMeta(p.path)
}

func (p *Page) render() string {
	var tpl *template.Template
	var err error
	tpl, err = template.ParseFiles(
		filepath.Join(p.sida.basePath, "layout/base.html"),
		filepath.Join(p.sida.basePath, "layout/single.html"),
		filepath.Join(p.sida.basePath, "layout/partials/header.html"),
		filepath.Join(p.sida.basePath, "layout/partials/footer.html"),
	)

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
	}

	p := &Page{
		path: path,
		kind: kind,
		sida: *sida,
	}

	p.init()

	return p
}
