package sida

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

// IsSida checks if diretory is a valid sida
func IsSida(path string) bool {
	folders := []string{
		filepath.Join(path, "layout"),
		filepath.Join(path, "pages"),
		filepath.Join(path, "posts"),
	}

	for _, f := range folders {
		ff, err := os.Stat(f)
		if os.IsNotExist(err) || !ff.Mode().IsDir() {
			return false
		}
	}

	v := viper.New()
	v.SetConfigName("config")
	v.AddConfigPath(path)

	err := v.ReadInConfig()
	if err != nil {
		fmt.Printf("err :%v", err)
		return false
	}

	return true
}
