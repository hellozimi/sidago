package builder

import (
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
	path    string
	kind    PageKind
	url     string
	sida    Sida
	content string
}

func (p *Page) Kind() PageKind {
	return p.kind
}

func (p *Page) init() {
	mmark.Parse(p.path)
}

func newPage(path string) *Page {
	f := filepath.Base(path)
	var kind PageKind
	if bre.MatchString(f) {
		kind = KindBlog
	}

	p := &Page{
		path: path,
		kind: kind,
	}

	p.init()

	return p
}
