package builder

import (
	"html/template"
	"strings"
	"time"

	"github.com/hellozimi/sidago/helpers"
)

type PageNameComponents struct {
	Title string
	Year  string
	Month string
	Day   string
}

type PageMeta struct {
	BaseFilename   string
	Slug           string
	Date           time.Time
	NameComponents PageNameComponents
	page           *Page
}

func (p *PageMeta) URL() template.URL {
	return template.URL(p.page.sida.Global.baseURL + p.page.RelOutputPath())
}

func newPageMeta(path string) PageMeta {
	filename, _ := helpers.FileAndExt(path)
	slug, date := slugAndDateFromFile(filename)

	return PageMeta{
		Slug: slug,
		Date: date,
		NameComponents: PageNameComponents{
			Title: helpers.Unslugify(slug),
			Year:  string(date.Year()),
			Month: string(date.Month()),
			Day:   string(date.Day()),
		},
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
