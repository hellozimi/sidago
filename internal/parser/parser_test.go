package parser

import (
	"bytes"
	"testing"
)

func TestParseFrontmatterBody(t *testing.T) {
	cases := []struct {
		a   []byte
		err error
		eb  []byte
		ef  map[string]interface{}
	}{
		{
			a: []byte(`---
title = "My title"
visible = true
---
# Header

Paragraph here`),
			err: nil,
			eb:  []byte("# Header\n\nParagraph here"),
			ef:  map[string]interface{}{"title": "My title", "visible": true},
		},
		{
			a:   []byte(`# HeaderParagraph here`),
			err: NoFrontmatter,
			eb:  make([]byte, 0),
			ef:  map[string]interface{}{},
		},
		{
			a: []byte(`---
title = "Test title"
---`),
			err: nil,
			eb:  make([]byte, 0),
			ef:  map[string]interface{}{"title": "Test title"},
		},
		{
			a: []byte(`---
title = "My title"
+++
# Header

Paragraph here`),
			err: MalformedFrontmatter,
			eb:  make([]byte, 0),
			ef:  map[string]interface{}{},
		},
	}

	for _, x := range cases {
		f, b, err := ParseFrontmatterBody(x.a)
		if err != x.err {
			t.Errorf("expected err to be '%v' but was '%v'", x.err, err)
		}

		if bytes.Compare(b, x.eb) != 0 {
			t.Errorf("expected body to equal '%s' but was '%s'", string(x.eb), string(b))
		}

		if x.ef["title"] != f["title"] {
			t.Errorf("expected frontmatter title to be '%s' but was '%s", x.ef["title"], f["title"])
		}

		if x.ef["visible"] != f["visible"] {
			t.Errorf("expected frontmatter visible to be '%v' but was '%v", x.ef["visible"], f["visible"])
		}
	}
}
