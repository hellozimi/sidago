package helpers

import (
	"github.com/microcosm-cc/bluemonday"
)

var p = bluemonday.StrictPolicy()

// StripHTML removes all html tags.
func StripHTML(html string) string {
	return p.Sanitize(html)
}
