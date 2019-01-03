package helpers

import "testing"

func TestSlugiy(t *testing.T) {
	cases := []struct {
		a string
		e string
	}{
		{a: "My title", e: "my-title"},
		{a: "New years 2019", e: "new-years-2019"},
		{a: "Nästa Månad", e: "n-sta-m-nad"},
		{a: "🚀 Rocket", e: "rocket"},
		{a: "The yellow 🚜 tractor", e: "the-yellow-tractor"},
	}

	for _, x := range cases {
		res := Slugify(x.a)
		if res != x.e {
			t.Errorf("expected string to equal '%s' but was '%s'", x.e, res)
		}
	}
}
