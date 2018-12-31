package builder

import (
	"testing"
	"time"
)

func TestSlugAndDate(t *testing.T) {
	cases := []struct {
		a string
		d time.Time
		s string
	}{
		{a: "2018-12-24_my-slug", d: timeF(t, "2018-12-24"), s: "my-slug"},
		{a: "my-slug-other-slug", d: time.Time{}, s: "my-slug-other-slug"},
	}

	for _, x := range cases {
		s, d := slugAndDateFromFile(x.a)
		if s != x.s {
			t.Errorf("expected %s slug to equal %s but was %s", x.a, x.s, s)
		}
		if d != x.d {
			t.Errorf("expected %s date to equal %v but was %v", x.a, x.d, d)
		}
	}
}

func timeF(t *testing.T, date string) time.Time {
	d, err := time.Parse("2006-01-02", date)
	if err != nil {
		t.Fatalf("error parsing date %s", date)
	}
	return d
}
