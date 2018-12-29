package sida

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

// IsSida checks if diretory is a valid sida
func IsSida(path string) (bool, error) {
	folders := []string{
		filepath.Join(path, "layout"),
		filepath.Join(path, "pages"),
		filepath.Join(path, "posts"),
	}

	for _, f := range folders {
		ff, err := os.Stat(f)
		if os.IsNotExist(err) {
			return false, nil
		}

		if !ff.Mode().IsDir() {
			return false, nil
		}
	}

	v := viper.New()
	v.SetConfigName("config")
	v.AddConfigPath(path)

	err := v.ReadInConfig()
	if err != nil {
		return false, err
	}

	return true, nil
}
