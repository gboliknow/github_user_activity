package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// fetchUserIssues fetches both created and assigned issues
func fetchUserIssues(username string) ([]Issue, error) {
	url := fmt.Sprintf("https://api.github.com/search/issues?q=author:%s+is:issue", username)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching issues: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: received status code %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var response struct {
		Items []Issue `json:"items"`
	}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %w", err)
	}

	return response.Items, nil
}

// fetchUserStars fetches repositories starred by the user
func fetchUserStars(username string) ([]StarredRepo, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/starred", username)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching stars: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: received status code %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var stars []StarredRepo
	if err := json.Unmarshal(body, &stars); err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %w", err)
	}

	return stars, nil
}
