package main

import (
	"context"
	"math/rand"
	"time"
)

// Context is used to manage the lifetime of a process.
// It's used to pass request-scoped values, cancelation signals, and deadlines across API boundaries between goroutines.

type Response struct {
	val int
	err error
}

func main() {
	// "Send" multiple requests
}

func fetchUserData(ctx context.Context, userID int) (int, error) {
	// Define maximum time to wait for the slow third party data

	// Because we can't return from a goroutine, we use a channel

	return 1, nil
}

func fetchSlowThirdPartyData() (int, error) {
	// This may take different amount of time, it's not deterministic
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(400)))
	return rand.Intn(10), nil
}
