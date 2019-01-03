package helpers

import "testing"

func TestTruncateFull(t *testing.T) {
	cases := []struct {
		a string
		b int
		e string
	}{
		{a: "This is a sentence.", b: 0, e: "This is a sentence."},
		{a: "This is a paragraph. The paragraph has two sentences.", b: 0, e: "This is a paragraph."},
		{a: "This is a paragraph! The paragraph has two sentences.", b: 0, e: "This is a paragraph!"},
		{a: "This is a paragraph? The paragraph has two sentences.", b: 0, e: "This is a paragraph?"},
		{a: "This is a paragraph. The paragraph has two sentences.", b: 4, e: "This is a paragraph. The paragraph has two sentences."},
		{a: "This is a paragraph. The paragraph has two sentences but without ending period", b: 6, e: "This is a paragraph. The paragraph has two sentences but without ending period"},
	}

	for _, x := range cases {
		res := TruncateFull(x.a, x.b)
		if res != x.e {
			t.Errorf("expected '%s' to equal '%s' but was '%s'", x.a, x.e, res)
		}
	}
}
