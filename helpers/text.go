package helpers

import (
	"regexp"
	"strings"
)

var suffixRe = regexp.MustCompile("(\\.|\\?|\\.\"|!)$")

// TruncateFull truncates text to a minmum amount of words
// and tries to find next period after that min word count
// which always results in a full sentence.
func TruncateFull(in string, min int) string {
	words := strings.Fields(in)

	if len(words) < min {
		return in
	}

	for i, w := range words {
		if suffixRe.Match([]byte(w)) && i >= min {
			c := i + 1
			return strings.Join(words[:c], " ")
		}
	}

	return in
}
