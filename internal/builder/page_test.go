package builder

import (
	"testing"
)

func TestPageName(t *testing.T) {
	cases := []struct {
		a string
		e bool
	}{
		{a: "2018-12-14_my_title.md", e: true},
		{a: "2018-12-14_my_title with space.md", e: false},
		{a: "2018-12-14_new_title", e: false},
		{a: "my_other_title_for_page.md", e: false},
		{a: "a failing title", e: false},
	}

	for _, x := range cases {
		res := bre.Match([]byte(x.a))
		if res != x.e {
			t.Errorf("expected '%s' to equal %v but was %v", x.a, x.e, res)
		}
	}
}

func TestPageNameComponents(t *testing.T) {

	cases := []struct {
		a string
		y string
		m string
		d string
		t string
	}{
		{a: "2018-12-14_my_title.md", y: "2018", m: "12", d: "14", t: "my_title"},
		{a: "2018-04-34_my_title.md", y: "", m: "", d: "", t: ""},
		{a: "2018-13-04_my_title.md", y: "", m: "", d: "", t: ""},
		{a: "my_page.md", y: "", m: "", d: "", t: ""},
	}

	for _, x := range cases {
		re := bre.FindAllStringSubmatch(x.a, -1)

		var (
			y  string
			m  string
			d  string
			ti string
		)

		if len(re) >= 1 && len(re[0]) > 5 {
			g := re[0]
			y = g[1]
			m = g[2]
			d = g[3]
			ti = g[4]
		}

		if x.t != ti {
			t.Errorf("expected %s to equal %s but was %s", x.a, x.t, ti)
		}

		if x.y != y {
			t.Errorf("expected %s to equal %s but was %s", x.a, x.y, y)
		}

		if x.m != m {
			t.Errorf("expected %s to equal %s but was %s", x.a, x.m, m)
		}

		if x.d != d {
			t.Errorf("expected %s to equal %s but was %s", x.a, x.d, d)
		}
	}
}
