package mmark

import (
	"fmt"
	"io/ioutil"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"github.com/mmarkdown/mmark/mast"
	"github.com/mmarkdown/mmark/mparser"
	"github.com/mmarkdown/mmark/render/mhtml"
)

func Parse(fileName string) ([]byte, error) {
	var (
		d []byte
	)
	d, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Couldn't open %q: %q", fileName, err)
		return d, err
	}

	p := parser.NewWithExtensions(mparser.Extensions)
	init := mparser.NewInitial(fileName)
	documentTitle := "" // hack to get document title from TOML title block and then set it here.
	p.Opts = parser.ParserOptions{
		ParserHook: func(data []byte) (ast.Node, []byte, int) {
			node, data, consumed := mparser.Hook(data)
			if t, ok := node.(*mast.Title); ok {
				if !t.IsTriggerDash() {
					documentTitle = t.TitleData.Title
				}
			}
			return node, data, consumed
		},
		ReadIncludeFn: init.ReadInclude,
		Flags:         parser.FlagsNone,
	}
	doc := markdown.Parse(d, p)
	opts := html.RendererOptions{
		Comments:       [][]byte{[]byte("//"), []byte("#")}, // used for callouts.
		RenderNodeHook: mhtml.RenderHook,
		Flags:          html.CommonFlags | html.FootnoteNoHRTag | html.FootnoteReturnLinks,
		Generator:      `  <meta name="GENERATOR" content="github.com/mmarkdown/mmark Mmark Markdown Processor - mmark.nl`,
	}
	opts.Title = documentTitle // hack to add-in discovered title

	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer), nil
}
