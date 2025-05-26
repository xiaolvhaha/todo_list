package main

import (
	"github.com/spf13/viper"
)

func main() {
	initViper()

	engine := InitApp()

	engine.Run() // listen and serve on 0.0.0.0:8080
}

func initViper() {
	viper.SetConfigName("dev")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
