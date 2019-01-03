package helpers

import (
	"path/filepath"
	"strings"
)

// FileAndExt returns a tuple with filename and extension of a file
func FileAndExt(in string) (string, string) {
	base := filepath.Base(in)
	ext := filepath.Ext(in)
	return strings.TrimSuffix(base, ext), ext
}
