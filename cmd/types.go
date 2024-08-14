package cmd

import (
	"encoding/json"
	"time"
)

type UserProfile struct {
	Login       string `json:"login"`
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Company     string `json:"company"`
	Blog        string `json:"blog"`
	Location    string `json:"location"`
	Email       string `json:"email"`
	Bio         string `json:"bio"`
	PublicRepos int    `json:"public_repos"`
	Followers   int    `json:"followers"`
	Following   int    `json:"following"`
	CreatedAt   time.Time `json:"created_at"` 
	// Add other fields as needed
}

type Repository struct {
	Name        string `json:"name"`
	FullName    string `json:"full_name"`
	Description string `json:"description"`
	Fork        bool   `json:"fork"`
	// Add other fields as needed
}

type RepositoryDetails struct {
	Name        string `json:"name"`
	FullName    string `json:"full_name"`
	Description string `json:"description"`
	Private     bool   `json:"private"`
	Fork        bool   `json:"fork"`
	// Add other fields as needed
}

type Event struct {
	Type      string `json:"type"`
	CreatedAt string `json:"created_at"`
	Repo      struct {
		Name string `json:"name"`
	} `json:"repo"`
	Actor struct {
		Login string `json:"login"`
	} `json:"actor"`
	Payload json.RawMessage `json:"payload"` // RawMessage allows flexible handling of payload data
	Public  bool            `json:"public"`
}
