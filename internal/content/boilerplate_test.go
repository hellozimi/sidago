package content

import (
	"testing"
	"time"
)

func TestBoilerplate(t *testing.T) {
	cases := []struct {
		a   boilerplate
		e   string
		err error
	}{
		{
			a:   newBoilerplate("This Is A Page", time.Time{}, typePage),
			e:   `# This Is A Page`,
			err: nil,
		},
		{
			a: newBoilerplate("This Is A Post", timeF(t, "2019-01-05"), typePost),
			e: `---
title = "This Is A Post"
date = "2019-01-05T00:00:00Z"
---
# This Is A Post`,
			err: nil,
		},
		{
			a:   newBoilerplate("This Is Something Unknown", timeF(t, "2019-01-05"), typeUnknown),
			e:   "",
			err: errNoContent,
		},
	}

	for _, x := range cases {
		res, err := x.a.content()
		if err != x.err {
			t.Errorf("expected error to be '%v' but was '%v'", x.err, err)
		}

		if res != x.e {
			t.Errorf("expected content to be '%s' but was '%s", x.e, res)
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
