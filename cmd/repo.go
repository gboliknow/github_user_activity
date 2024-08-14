/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

// repoCmd represents the repo command
var RepoCmd = &cobra.Command{
	Use:   "repo",
	Short: "Fetch user repo",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide a username.")
			return
		}
		username := args[0]
		repo, err := fetchUserRepos(username)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		formatRepositories(repo)

	},
}

func init() {

	RepoCmd.Flags().StringP("username", "t", "all", "")
}

func fetchUserRepos(username string) ([]Repository, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/repos", username)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching user repositories: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: received status code %d", resp.StatusCode)
	}

	var repos []Repository
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	if err := json.Unmarshal(body, &repos); err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %w", err)
	}

	return repos, nil
}

func formatRepositories(repos []Repository) {
	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug)
	fmt.Fprintln(writer, "Name\tFull Name\tDescription\tFork")
	fmt.Fprintln(writer, "----\t---------\t-----------\t----")
	for _, repo := range repos {
		forkStatus := "No"
		if repo.Fork {
			forkStatus = "Yes"
		}
		fmt.Fprintf(writer, "%s\t%s\t%s\t%s\n",
			repo.Name,
			repo.FullName,
			repo.Description,
			forkStatus,
		)
	}
	writer.Flush()
}