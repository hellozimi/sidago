package helpers

import "time"

var lazyWithoutTZ = "2006-01-02 15:04:05"

// ParseDate tries to parse multiple date formats
// 2006-01-02 15:04:05 or rfc3339
func ParseDate(in string) (time.Time, error) {
	var date = time.Time{}
	var err error
	date, _ = time.Parse(lazyWithoutTZ, in)

	if date.IsZero() {
		date, err = time.Parse(time.RFC3339, in)
		if err != nil {
			return time.Time{}, err
		}
	}

	return date, nil
}
