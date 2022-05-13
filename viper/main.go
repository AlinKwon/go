package main

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	fmt.Println(viper.Get("id"))
	fmt.Println(viper.Get("ID"))
	fmt.Println(viper.Get("name"))

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})

	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.ReadInConfig()

	fmt.Println(viper.Get("prod.port"))
}
