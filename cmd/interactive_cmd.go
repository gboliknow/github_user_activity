package cmd

import (
	"github.com/spf13/cobra"
)

var interactiveCmd = &cobra.Command{
	Use:   "interactive",
	Short: "Start interactive mode",
	Long:  `Start an interactive session to browse GitHub user information`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := runInteractiveMode(); err != nil {
			cmd.PrintErrf("Error in interactive mode: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(interactiveCmd)
}
