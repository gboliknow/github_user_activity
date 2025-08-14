package cmd

import (
	"testing"
	"time"
)

func TestFormatEvents(t *testing.T) {
	testEvents := []Event{
		{
			Type:      "PushEvent",
			CreatedAt: time.Now().Format(time.RFC3339),
			Repo: struct {
				Name string "json:\"name\""
			}{
				Name: "test/repo",
			},
			Actor: struct {
				Login string "json:\"login\""
			}{
				Login: "testuser",
			},
			Public: true,
		},
	}

	// This is a visual test - it will output to console
	formatEvents(testEvents)
}

func TestFetchUserActivity(t *testing.T) {
	tests := []struct {
		name         string
		username     string
		activityType string
		wantErr      bool
	}{
		{
			name:         "Valid username",
			username:     "gboliknow",
			activityType: "all",
			wantErr:      false,
		},
		{
			name:         "Invalid username",
			username:     "thisusershouldnotexist12345678990",
			activityType: "all",
			wantErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := fetchUserActivity(tt.username, tt.activityType)
			if (err != nil) != tt.wantErr {
				t.Errorf("fetchUserActivity() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
