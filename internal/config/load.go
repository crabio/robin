package config

import (
	"github.com/jinzhu/configor"
)

func LoadConfig(filePath string) (config Config, err error) {
	err = configor.Load(&config, filePath)
	return
}
