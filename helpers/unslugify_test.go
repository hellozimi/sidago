package helpers

import "testing"

func TestUnslugify(t *testing.T) {
	cases := []struct {
		a string
		b string
	}{
		{a: "my-little-slug", b: "My Little Slug"},
		{a: "my-summer-2018", b: "My Summer 2018"},
	}

	for _, x := range cases {
		res := Unslugify(x.a)
		if res != x.b {
			t.Errorf("Expected %s to equal %s but was %s", x.a, x.b, res)
		}
	}
}
