package helpers

import (
	"path/filepath"
	"strings"
)

func FileAndExt(in string) (string, string) {
	base := filepath.Base(in)
	ext := filepath.Ext(in)
	return strings.TrimSuffix(base, ext), ext
}
