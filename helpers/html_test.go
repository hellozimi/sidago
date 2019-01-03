package helpers

import "testing"

func TestStripHTML(t *testing.T) {
	cases := []struct {
		a string
		e string
	}{
		{a: "<b>Stripping my html</b>!", e: "Stripping my html!"},
		{a: "Hello <STYLE>.XSS{background-image:url(\"javascript:alert('XSS')\");}</STYLE><A CLASS=XSS></A>World", e: "Hello World"},
	}

	for _, x := range cases {
		res := StripHTML(x.a)
		if res != x.e {
			t.Errorf("expected '%s' to equal '%s' but was '%s'", x.a, x.e, res)
		}
	}
}
