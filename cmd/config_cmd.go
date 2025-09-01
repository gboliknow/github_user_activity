package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage configuration settings",
	Long:  `View or update configuration settings like GitHub token, default user, and output format.`,
}

var configSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a configuration value",
	Run: func(cmd *cobra.Command, args []string) {
		token, _ := cmd.Flags().GetString("token")
		user, _ := cmd.Flags().GetString("user")
		format, _ := cmd.Flags().GetString("format")
		ttl, _ := cmd.Flags().GetInt("cache-ttl")

		if token != "" {
			config.GitHubToken = token
		}
		if user != "" {
			config.DefaultUser = user
		}
		if format != "" {
			config.OutputFormat = format
		}
		if ttl != 0 {
			config.CacheTTL = ttl
		}

		if err := saveConfig(); err != nil {
			fmt.Printf("Error saving config: %v\n", err)
			return
		}
		fmt.Println("Configuration saved successfully!")
	},
}

var configViewCmd = &cobra.Command{
	Use:   "view",
	Short: "View current configuration",
	Run: func(cmd *cobra.Command, args []string) {
		w := createTableWriter()
		printTableHeader(w, "Setting", "Value")

		// Mask token for security
		token := "not set"
		if config.GitHubToken != "" {
			token = "********" + config.GitHubToken[len(config.GitHubToken)-4:]
		}

		printTableRow(w, "%s\t%s", "GitHub Token", token)
		printTableRow(w, "%s\t%s", "Default User", config.DefaultUser)
		printTableRow(w, "%s\t%s", "Output Format", config.OutputFormat)
		printTableRow(w, "%s\t%d", "Cache TTL (seconds)", config.CacheTTL)
		w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(configSetCmd)
	configCmd.AddCommand(configViewCmd)

	// Add flags for setting configuration
	configSetCmd.Flags().String("token", "", "GitHub API token")
	configSetCmd.Flags().String("user", "", "Default GitHub username")
	configSetCmd.Flags().String("format", "", "Output format (table/json)")
	configSetCmd.Flags().Int("cache-ttl", 0, "Cache TTL in seconds")
}
