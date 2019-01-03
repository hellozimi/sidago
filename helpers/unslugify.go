package helpers

import "strings"

// Unslugify creates a title of a slug
// Replacing - with ' ' and capitalizes first letter of each word
func Unslugify(in string) string {
	in = strings.Replace(in, "-", " ", -1)
	return strings.Title(in)
}
