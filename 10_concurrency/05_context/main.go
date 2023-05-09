package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"runtime"
	"time"
)

type Response struct {
	val int
	err error
}

func main() {
	for i := 0; i < 10; i++ {
		start := time.Now()

		ctx := context.Background()

		val, err := fetchUserData(ctx, i)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(val)

		fmt.Println("Time taken:", time.Since(start))
	}
}

func fetchUserData(ctx context.Context, userID int) (int, error) {
	// Define maximum time to wait for the slow third party data
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*300)
	defer cancel()

	responseCh := make(chan Response)

	// Because we can't return from a goroutine, we use a channel
	go func() {
		val, err := fetchSlowThirdPartyData()
		responseCh <- Response{
			val: val,
			err: err,
		}
	}()

	fmt.Println("Number of Goroutines: ", runtime.NumGoroutine())

	for {
		select {
		case <-ctx.Done(): // empty struct
			return 0, fmt.Errorf("Fetching data from 3rd party exceeded deadline.")

		case response := <-responseCh: // If it's on time, we get a response
			return response.val, response.err
		}
	}
}

func fetchSlowThirdPartyData() (int, error) {
	// This may take different amount of time, it's not deterministic
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(400)))
	return rand.Intn(10), nil
}
