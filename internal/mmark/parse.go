package mmark

import (
	"fmt"
	"io/ioutil"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"github.com/mmarkdown/mmark/mparser"
	"github.com/mmarkdown/mmark/render/mhtml"
)

var opts = html.RendererOptions{
	Comments:       [][]byte{[]byte("//"), []byte("#")},
	RenderNodeHook: mhtml.RenderHook,
	Flags:          html.CommonFlags | html.FootnoteNoHRTag | html.FootnoteReturnLinks,
	Generator:      `  <meta name="GENERATOR" content="github.com/mmarkdown/mmark Mmark Markdown Processor - mmark.nl`,
}

// ParseBytes parses markdown bytes to html
func ParseBytes(in []byte) ([]byte, error) {
	p := parser.NewWithExtensions(mparser.Extensions)
	doc := markdown.Parse(in, p)
	renderer := html.NewRenderer(opts)
	return markdown.Render(doc, renderer), nil
}

// ParseFile parses markdown file to html
func ParseFile(fileName string) ([]byte, error) {
	var (
		d []byte
	)
	d, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Couldn't open %q: %q", fileName, err)
		return d, err
	}

	return ParseBytes(d)
}
