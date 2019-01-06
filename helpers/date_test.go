package helpers

import (
	"errors"
	"testing"
	"time"
)

func TestParseDate(t *testing.T) {
	cases := []struct {
		a   string
		e   time.Time
		err error
	}{
		{a: "2017-04-05T15:00:01Z", e: dateTimeF(t, "2017-04-05 15:00:01"), err: nil},
		{a: "2017-12-24T15:00:00+01:00", e: dateTimeF(t, "2017-12-24 14:00:00"), err: nil},
		{a: "2017-12-24T15:00:00-07:00", e: dateTimeF(t, "2017-12-24 22:00:00"), err: nil},
		{a: "2017-12-24 15:00:00", e: dateTimeF(t, "2017-12-24 15:00:00"), err: nil},
		{a: "2017-12-24T15:00:00", e: time.Time{}, err: errors.New(`parsing time "2017-12-24T15:00:00" as "2006-01-02T15:04:05Z07:00": cannot parse "" as "Z07:00"`)},
	}

	for _, x := range cases {
		res, err := ParseDate(x.a)
		if err != nil {
			if err.Error() != x.err.Error() {
				t.Errorf("expected error to equal '%v' but was '%v'", x.err, err)
			}
		}

		if res.Unix() != x.e.Unix() {
			t.Errorf("expected parsed date to equal '%v' but was '%v'", x.e, res)
		}
	}
}

func dateTimeF(t *testing.T, date string) time.Time {
	d, err := time.Parse("2006-01-02 15:04:05", date)
	if err != nil {
		t.Fatalf("error parsing date %s", date)
	}
	return d
}
