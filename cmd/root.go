/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
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
	// Load configuration
	if err := loadConfig(); err != nil {
		fmt.Fprintf(os.Stderr, "Error loading config: %v\n", err)
	}

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("output", "o", "", "Output format (table/json)")
	rootCmd.PersistentFlags().StringP("token", "k", "", "GitHub API token") // Changed from 't' to 'k'
	addSubcommandPalettes()

	// Use config values as defaults
	if config.OutputFormat != "" {
		rootCmd.PersistentFlags().Lookup("output").Value.Set(config.OutputFormat)
	}
	if config.GitHubToken != "" {
		rootCmd.PersistentFlags().Lookup("token").Value.Set(config.GitHubToken)
	}
}

func addSubcommandPalettes() {
	rootCmd.AddCommand(ActivityCmd)
	rootCmd.AddCommand(ProfileCmd)
	rootCmd.AddCommand(RepoCmd)
}
