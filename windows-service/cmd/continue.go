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

// continueCmd represents the continue command
var continueCmd = &cobra.Command{
	Use:   "continue",
	Short: "continue servcie",
	Long: `continue
	bla bla`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("continue called")
		name, _ := cmd.Flags().GetString("name")
		if len(strings.TrimSpace(name)) > 0 {
			service.ControlService(name, svc.Continue, svc.Running)
		} else {
			service.ControlService("default service name", svc.Continue, svc.Running)
		}
	},
}

func init() {
	rootCmd.AddCommand(continueCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// continueCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// continueCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	continueCmd.Flags().StringP("name", "n", "", "set service name.")
}
