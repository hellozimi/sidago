package helpers

import (
	"regexp"
	"strings"
)

var suffixRe = regexp.MustCompile("(\\.|\\?|\\.\"|!)$")

func TruncateFull(in string, min int) string {
	if len(in) < min {
		return in
	}

	words := strings.Fields(in)

	for i, w := range words {
		if suffixRe.Match([]byte(w)) && i >= min {
			c := i + 1
			return strings.Join(words[:c], " ")
		}
	}

	return in
}
