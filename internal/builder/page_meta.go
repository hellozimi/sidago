package builder

import (
	"strings"
	"time"

	"github.com/hellozimi/sidago/helpers"
)

type PageMeta struct {
	BaseFilename string
	Slug         string
	Date         time.Time
}

func newPageMeta(path string) PageMeta {
	filename, _ := helpers.FileAndExt(path)
	slug, date := slugAndDateFromFile(filename)

	return PageMeta{
		Slug: slug,
		Date: date,
	}
}

func slugAndDateFromFile(filename string) (string, time.Time) {
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
