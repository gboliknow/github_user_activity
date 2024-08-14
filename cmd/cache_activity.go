package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"
)

const cacheDir = "./cache"

func InitFileForCaching() {
	if _, err := os.Stat(cacheDir); os.IsNotExist(err) {
		err := os.Mkdir(cacheDir, 0755)
		if err != nil {
			fmt.Printf("Error creating cache directory: %v\n", err)
		}
	}
}

func getCacheFilePath(username string) string {
	return fmt.Sprintf("%s/%s_cache.json", cacheDir, username)
}

func GetFromCacheFileBased(username string) ([]Event, bool) {
	filePath := getCacheFilePath(username)
	file, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("Cache miss for %s: file not found\n", username)
			return nil, false
		}
		fmt.Printf("Error opening cache file for %s: %v\n", username, err)
		return nil, false
	}
	defer file.Close()

	var entry CacheEntry
	if err := json.NewDecoder(file).Decode(&entry); err != nil {
		fmt.Printf("Error decoding cache file for %s: %v\n", username, err)
		return nil, false
	}

	if time.Since(entry.Timestamp) > cacheDuration {
		fmt.Printf("Cache miss for %s: expired\n", username)
		return nil, false
	}

	return entry.Data, true
}

func SetToCacheFileBased(username string, data []Event) {
	filePath := getCacheFilePath(username)
	entry := CacheEntry{
		Data:      data,
		Timestamp: time.Now(),
	}

	file, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("Error creating cache file for %s: %v\n", username, err)
		return
	}
	defer file.Close()

	if err := json.NewEncoder(file).Encode(entry); err != nil {
		fmt.Printf("Error encoding cache data for %s: %v\n", username, err)
		return
	}

	fmt.Printf("Saving to cache for %s at %v\n", username, time.Now())
}

type CacheEntry struct {
	Data      []Event
	Timestamp time.Time
}

var cache = struct {
	sync.RWMutex
	entries map[string]CacheEntry
}{
	entries: make(map[string]CacheEntry),
}

const cacheDuration = 1 * time.Minute

func GetFromCache(username string) ([]Event, bool) {
	cache.RLock()
	defer cache.RUnlock()

	entry, found := cache.entries[username]
	if !found {
		fmt.Printf("Cache miss for %s: not found\n", username)
		return nil, false
	}

	if time.Since(entry.Timestamp) > cacheDuration {
		fmt.Printf("Cache miss for %s: expired\n", username)
		return nil, false
	}

	fmt.Printf("Cache hit for %s: %v\n", username, entry)
	return entry.Data, true
}

func SetToCache(username string, data []Event) {
	cache.Lock()
	defer cache.Unlock()

	cache.entries[username] = CacheEntry{
		Data:      data,
		Timestamp: time.Now(),
	}
	fmt.Printf("Saving to cache for %s at %v\n", username, time.Now())
}
