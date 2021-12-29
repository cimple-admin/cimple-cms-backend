package pkg

import (
	"fmt"

	"github.com/spf13/viper"
)

type config struct {
	App appConfig `json:"app,omitempty"`
}

type appConfig struct {
	ReadTimeOut  int    `json:"readTimeOut"`
	WriteTimeOut int    `json:"writeTimeOut"`
	HttpPort     string `json:"httpPort"`
	RunMode      string `json:"runMode"`
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
