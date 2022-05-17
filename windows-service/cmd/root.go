/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	service "alin/window-service/internal"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/sys/windows/svc"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "window-service",
	Short: "golang windows service run.",
	Long:  `golang windows service run.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		inService, err := svc.IsWindowsService()
		if err != nil {
			log.Fatal(err)
		}
		if inService {
			name, _ := cmd.Flags().GetString("name")
			if len(strings.TrimSpace(name)) > 0 {
				service.RunService(name, false)
			} else {
				service.RunService("default service name", false)
			}
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.window-service.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().StringP("name", "n", "", "set service name.")
}