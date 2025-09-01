package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/manifoldco/promptui"
)

// Handler for profile command
func handleProfile(username string) error {
	profile, err := fetchUserProfile(username)
	if err != nil {
		return fmt.Errorf("failed to fetch profile: %v", err)
	}

	fmt.Printf("\n%s\n", formatHeader("Profile Information"))
	fmt.Printf("%s %s\n", formatTitle("Username:"), formatPrimary(profile.Login))
	fmt.Printf("%s %s\n", formatTitle("Name:"), formatPrimary(profile.Name))
	fmt.Printf("%s %s\n", formatTitle("Bio:"), formatPrimary(profile.Bio))
	fmt.Printf("%s %s\n", formatTitle("Company:"), formatPrimary(profile.Company))
	fmt.Printf("%s %s\n", formatTitle("Location:"), formatPrimary(profile.Location))
	fmt.Printf("%s %s\n", formatTitle("Email:"), formatPrimary(profile.Email))
	fmt.Printf("%s %s\n", formatTitle("Public Repositories:"), formatSuccess("%d", profile.PublicRepos))
	fmt.Printf("%s %s\n", formatTitle("Followers:"), formatSecondary("%d", profile.Followers))
	fmt.Printf("%s %s\n", formatTitle("Following:"), formatSecondary("%d", profile.Following))
	fmt.Printf("%s %s\n", formatTitle("Created At:"), formatPrimary(profile.CreatedAt.Format("2006-01-02 15:04:05")))

	return nil
}

// Handler for activity command
func handleActivity(username string) error {
	events, err := fetchUserActivity(username, "all")
	if err != nil {
		return fmt.Errorf("failed to fetch activities: %v", err)
	}

	fmt.Printf("\nRecent Activities:\n")
	formatEvents(events)
	return nil
}

// Handler for repository command
func handleRepo(username string) error {
	repos, err := fetchUserRepos(username)
	if err != nil {
		return fmt.Errorf("failed to fetch repositories: %v", err)
	}

	fmt.Printf("\n%s\n", formatHeader("Repositories"))
	w := createTableWriter()
	printTableHeader(w, "Name", "Full Name", "Description", "Fork")

	for _, repo := range repos {
		fmt.Fprintf(w, "%s\t%s\t%s\t%v\n",
			repo.Name,
			repo.FullName,
			repo.Description,
			repo.Fork,
		)
	}
	w.Flush()
	return nil
}

func handleIssues(username string) error {
	issues, err := fetchUserIssues(username)
	if err != nil {
		return fmt.Errorf("failed to fetch issues: %v", err)
	}

	fmt.Printf("\nIssues:\n")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug)
	fmt.Fprintln(w, "Number\tTitle\tState\tCreated\tUpdated")
	fmt.Fprintln(w, "------\t-----\t-----\t-------\t-------")

	for _, issue := range issues {
		fmt.Fprintf(w, "#%d\t%s\t%s\t%s\t%s\n",
			issue.Number,
			issue.Title,
			issue.State,
			issue.CreatedAt.Format("2006-01-02"),
			issue.UpdatedAt.Format("2006-01-02"),
		)
	}
	w.Flush()
	return nil
}

func handleStars(username string) error {
	stars, err := fetchUserStars(username)
	if err != nil {
		return fmt.Errorf("failed to fetch stars: %v", err)
	}

	fmt.Printf("\nStarred Repositories:\n")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug)
	fmt.Fprintln(w, "Name\tDescription\tLanguage\tStars")
	fmt.Fprintln(w, "----\t-----------\t--------\t-----")

	for _, repo := range stars {
		fmt.Fprintf(w, "%s\t%s\t%s\t%d\n",
			repo.FullName,
			repo.Description,
			repo.Language,
			repo.Stars,
		)
	}
	w.Flush()
	return nil
}

func runInteractiveMode() error {
	prompt := promptui.Select{
		Label: "Select action",
		Items: []string{
			"View profile",
			"Check activities",
			"List repositories",
			"View issues",
			"Check stars",
			"Exit",
		},
	}

	_, result, err := prompt.Run()
	if err != nil {
		return err
	}

	userPrompt := promptui.Prompt{
		Label: "Enter GitHub username",
	}

	username, err := userPrompt.Run()
	if err != nil {
		return err
	}

	switch result {
	case "View profile":
		return handleProfile(username)
	case "Check activities":
		return handleActivity(username)
	case "List repositories":
		return handleRepo(username)
	case "View issues":
		return handleIssues(username)
	case "Check stars":
		return handleStars(username)
	case "Exit":
		fmt.Println("Goodbye!")
		return nil
	}

	return fmt.Errorf("invalid option: %s", result)
}
