/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	service "alin/window-service/internal"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/sys/windows/svc"
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("stop called")
		name, _ := cmd.Flags().GetString("name")
		if len(strings.TrimSpace(name)) > 0 {
			service.ControlService(name, svc.Stop, svc.Stopped)
		} else {
			service.ControlService("default service name", svc.Stop, svc.Stopped)
		}
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stopCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stopCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	stopCmd.Flags().StringP("name", "n", "", "set service name.")
}
