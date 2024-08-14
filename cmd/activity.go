package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"text/tabwriter"
	"time"

	"github.com/spf13/cobra"
)

// activityCmd represents the activity command
var ActivityCmd = &cobra.Command{
	Use:   "activity [username]",
	Short: "Fetches user activity based on the username",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide a username.")
			return
		}

		activityType, _ := cmd.Flags().GetString("type")
		username := args[0]
		events, err := fetchUserActivity(username, activityType)
		if err != nil {
			log.Fatalf("Error fetching user activity: %v , %v", err, activityType)
		}
		formatEvents(events)
	},
}

func init() {
	ActivityCmd.Flags().StringP("type", "t", "all", "Type of activities to fetch (e.g., login, logout)")
}

func fetchUserActivity(username string, activityType string) ([]Event, error) {

	data, found := GetFromCacheFileBased(username)
	if found {
		fmt.Printf("Fetching activity from cache %s for user : %s\n", activityType, username)
		return data, nil
	}

	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching data: %w", err)
	} else {
		fmt.Printf("Fetching activity of type %s for user : %s\n", activityType, username)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: received status code %d", resp.StatusCode)
	}
	var events []Event
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	if err := json.Unmarshal(body, &events); err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %w", err)
	}

	SetToCacheFileBased(username, events)
	return events, nil
}



func formatEvents(events []Event) {
	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug)
	fmt.Fprintln(writer, "Type\tCreated At\tRepository\tActor\tPublic")
	fmt.Fprintln(writer, "----\t-----------\t----------\t------\t------")
	for _, event := range events {
		createdAt, _ := time.Parse(time.RFC3339, event.CreatedAt)
		actor := event.Actor.Login
		public := "No"
		if event.Public {
			public = "Yes"
		}
		fmt.Fprintf(writer, "%s\t%s\t%s\t%s\t%s\n",
			event.Type,
			createdAt.Format("2006-01-02 15:04:05"),
			event.Repo.Name,
			actor,
			public,
		)
	}
	writer.Flush()
}
