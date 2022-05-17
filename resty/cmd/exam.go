/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// examCmd represents the exam command
var examCmd = &cobra.Command{
	Use:   "exam",
	Short: "rest api client test example",
	Long: `use resty project
GET, POST, PUT, DELETE, PATCH examples.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("exam called")
	},
}

func init() {
	rootCmd.AddCommand(examCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// examCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// examCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
