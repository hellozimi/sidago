package helpers

import (
	"path/filepath"
	"strings"
	"time"
)

// FileAndExt returns a tuple with filename and extension of a file
func FileAndExt(in string) (string, string) {
	base := filepath.Base(in)
	ext := filepath.Ext(in)
	return strings.TrimSuffix(base, ext), ext
}

// SlugAndDateFromFile returns a tuple with name and time.Time
// from '2006-01-02_name' format. If no date is present it will
// just return the passed filename with a zero date
func SlugAndDateFromFile(filename string) (string, time.Time) {
	if len(filename) < 10 {
		return filename, time.Time{}
	}

	d, err := time.Parse("2006-01-02", filename[:10])
	if err != nil {
		return filename, time.Time{}
	}

	slug := strings.Trim(filename[10:], " -_")

	return slug, d
}
