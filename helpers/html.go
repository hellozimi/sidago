package helpers

import (
	"github.com/microcosm-cc/bluemonday"
)

var p = bluemonday.StrictPolicy()

func StripHTML(html string) string {
	return p.Sanitize(html)
}
