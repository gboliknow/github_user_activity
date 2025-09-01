package cmd

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type RateLimit struct {
	Limit     int
	Remaining int
	Reset     time.Time
}

func checkRateLimit(resp *http.Response) (*RateLimit, error) {
	limit, _ := strconv.Atoi(resp.Header.Get("X-RateLimit-Limit"))
	remaining, _ := strconv.Atoi(resp.Header.Get("X-RateLimit-Remaining"))
	reset, _ := strconv.ParseInt(resp.Header.Get("X-RateLimit-Reset"), 10, 64)

	return &RateLimit{
		Limit:     limit,
		Remaining: remaining,
		Reset:     time.Unix(reset, 0),
	}, nil
}

func handleRateLimit(resp *http.Response) error {
	rateLimit, err := checkRateLimit(resp)
	if err != nil {
		return err
	}

	if rateLimit.Remaining == 0 {
		waitTime := time.Until(rateLimit.Reset)
		return fmt.Errorf("rate limit exceeded. Reset in %v", waitTime)
	}

	return nil
}
