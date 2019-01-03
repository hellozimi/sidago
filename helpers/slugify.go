package helpers

import (
	"regexp"
	"strings"
)

var slugRe = regexp.MustCompile("[^a-z0-9]+")

// Slugify creates a slug of a title
func Slugify(in string) string {
	in = strings.ToLower(in)
	return strings.Trim(slugRe.ReplaceAllString(in, "-"), "-")
}
