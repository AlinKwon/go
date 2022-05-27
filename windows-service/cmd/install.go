/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"strings"
	service "alin/window-service/internal"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "windows service install.",
	Long: `windows service install.
	blabla`,
	Run: func(cmd *cobra.Command, args []string) {
		logger.Info("install called", args)
		name, _ := cmd.Flags().GetString("name")
		desc, _ := cmd.Flags().GetString("desc")
		if len(strings.TrimSpace(name)) > 0 {
			service.InstallService(name, desc)
			viper.Set("name", name)
		} else {
			service.InstallService(viper.GetString("name"), desc)
		}
	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	installCmd.Flags().StringP("name", "n", "", "set service name.")
	installCmd.Flags().StringP("desc", "d", "", "set service description.")
}
