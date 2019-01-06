package helpers

import "testing"

func TestIsSlug(t *testing.T) {
	cases := []struct {
		a string
		e bool
	}{
		{a: "this-is-a-slug", e: true},
		{a: "this is not a slug", e: false},
		{a: "Neither-is This", e: false},
	}

	for _, x := range cases {
		res := IsSlug(x.a)
		if x.e != res {
			t.Errorf("expected result to equal '%v' but was '%v'", x.e, res)
		}
	}
}
