package helpers

import "strings"

func Unslugify(in string) string {
	in = strings.Replace(in, "-", " ", -1)
	return strings.Title(in)
}
