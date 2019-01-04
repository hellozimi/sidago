package builder

import (
	"html/template"

	"github.com/hellozimi/sidago/helpers"
)

// DateComps struct
type DateComps struct {
	Year  string
	Month string
	Day   string
}

// PageMeta struct
type PageMeta struct {
	BaseFilename   string
	DateComponents DateComps
	page           *Page
}

// URL returns the permalink for a page
func (p *PageMeta) URL() template.URL {
	return template.URL(p.page.sida.Global.baseURL + p.page.RelOutputPath())
}

func newPageMeta(path string) PageMeta {
	filename, _ := helpers.FileAndExt(path)
	_, date := helpers.SlugAndDateFromFile(filename)

	return PageMeta{
		DateComponents: DateComps{
			Year:  string(date.Year()),
			Month: string(date.Month()),
			Day:   string(date.Day()),
		},
	}
}
