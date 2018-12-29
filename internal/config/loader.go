package config

import (
	"bytes"

	"github.com/spf13/viper"
)

// FromString reads toml config from string
func FromString(config string) (Config, error) {
	v := newViper()
	err := v.ReadConfig(bytes.NewBuffer([]byte(config)))
	if err != nil {
		return nil, err
	}

	return v, nil
}

// FromFile reads config from path
func FromFile(configPath, configName string) (Config, error) {
	v := newViper()
	v.SetConfigName(configName)
	v.AddConfigPath(configPath)

	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return v, nil
}

func newViper() *viper.Viper {
	v := viper.New()
	v.SetConfigType("toml")
	return v
}
