/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "github_user_activity",
	Short: "Fetch information from GitHub",
	Long:  ``,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addSubcommandPalettes()
}

func addSubcommandPalettes() {
	rootCmd.AddCommand(ActivityCmd)
	rootCmd.AddCommand(ProfileCmd)
	rootCmd.AddCommand(RepoCmd)
}
