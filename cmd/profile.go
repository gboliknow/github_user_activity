/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

// ProfileCmd represents the profile command
var ProfileCmd = &cobra.Command{
	Use:   "profile",
	Short: "Fetch user profile",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide a username.")
			return
		}
		username := args[0]
		profile, err := fetchUserProfile(username)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		formatUserProfile(profile)

	},
}

func init() {

}

func fetchUserProfile(username string) (UserProfile, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s", username)
	resp, err := http.Get(url)
	if err != nil {
		return UserProfile{}, fmt.Errorf("error fetching user profile: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return UserProfile{}, fmt.Errorf("error: received status code %d", resp.StatusCode)
	}

	var profile UserProfile
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return UserProfile{}, fmt.Errorf("error reading response body: %w", err)
	}

	if err := json.Unmarshal(body, &profile); err != nil {
		return UserProfile{}, fmt.Errorf("error unmarshalling JSON: %w", err)
	}

	return profile, nil
}

func formatUserProfile(profile UserProfile) {
	fmt.Println("Profile Information:")
	fmt.Printf("Username: %s\n", profile.Login)
	fmt.Printf("Name: %s\n", profile.Name)
	fmt.Printf("Bio: %s\n", profile.Bio)
	fmt.Printf("Company: %s\n", profile.Company)
	fmt.Printf("Location: %s\n", profile.Location)
	fmt.Printf("Email: %s\n", profile.Email)
	fmt.Printf("Website: %s\n", profile.Blog)
	fmt.Printf("Public Repositories: %d\n", profile.PublicRepos)
	fmt.Printf("Followers: %d\n", profile.Followers)
	fmt.Printf("Following: %d\n", profile.Following)
	fmt.Printf("Created At: %s\n", profile.CreatedAt.Format("2006-01-02 15:04:05"))
}
