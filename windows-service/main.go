/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"alin/window-service/cmd"
	service "alin/window-service/internal"

	"github.com/spf13/viper"
	"golang.org/x/sys/windows/svc"
)

func main() {

	logger := service.GetLogger()

	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.ReadInConfig()

	inService, err := svc.IsWindowsService()
	if err != nil {
		logger.Fatal(err)
	}
	if inService {
		logger.Info("SERVICE START")
		service.RunService(viper.GetString("name"), false)
	} else {
		logger.Error("not service context..")
	}

	cmd.Execute()
}
