package cmd

import (
	"testing"
)

func TestCacheOperations(t *testing.T) {
	// Initialize cache directory
	InitFileForCaching()

	testUser := "testuser"
	testEvents := []Event{
		{
			Type: "PushEvent",
			Repo: struct {
				Name string "json:\"name\""
			}{
				Name: "test/repo",
			},
			Actor: struct {
				Login string "json:\"login\""
			}{
				Login: testUser,
			},
		},
	}

	// Test setting cache
	SetToCacheFileBased(testUser, testEvents)

	// Test getting from cache
	events, found := GetFromCacheFileBased(testUser)
	if !found {
		t.Error("Expected to find cache entry, but none was found")
	}

	if len(events) != len(testEvents) {
		t.Errorf("Expected %d events, got %d", len(testEvents), len(events))
	}

	// Test cache for non-existent user
	_, found = GetFromCacheFileBased("nonexistentuser")
	if found {
		t.Error("Expected no cache entry for non-existent user, but found one")
	}
}
