package pkg

import (
	"fmt"

	"github.com/spf13/viper"
)

type config struct {
	App         appConfig `mapstructure:"app,omitempty"`
	IsInstalled bool      `mapstructure:"lock,omitempty"`
}

type appConfig struct {
	ReadTimeOut  int    `mapstructure:"readTimeOut"`
	WriteTimeOut int    `mapstructure:"writeTimeOut"`
	HttpPort     string `mapstructure:"httpPort"`
	RunMode      string `mapstructure:"runMode"`
}

func initConfig() config {
	viper.SetConfigName("config")   // name of config file (without extension)
	viper.SetConfigType("yaml")     // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./config") // path to look for the config file in
	err := viper.ReadInConfig()     // Find and read the config file
	if err != nil {                 // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	c := config{}
	err = viper.Unmarshal(&c)
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	return c
}
