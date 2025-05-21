package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config interface {
	GetInt(key string) int
	GetString(key string) string
	GetBool(key string) bool
}

type config struct {
	v *viper.Viper
}

func New(fileName, fileType, fileDir string) Config {
	v := viper.New()
	v.SetConfigName(fileName) // name of config file (without extension)
	v.SetConfigType(fileType) // REQUIRED if the config file does not have the extension in the name
	v.AddConfigPath(fileDir)  // path to look for the config file in
	err := v.ReadInConfig()   // Find and read the config file
	if err != nil {           // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	return &config{v: v}
}

func (c *config) GetInt(key string) int {
	return c.v.GetInt(key)
}

func (c *config) GetString(key string) string {
	return c.v.GetString(key)
}

func (c *config) GetBool(key string) bool {
	return c.v.GetBool(key)
}
