package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
	"time"
)

const baseURL = "https://random-data-api.com/api/v2/"

// User struct
type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

// fetchRandomData fetches random data of a specified size from the provided resource.
// The result is sent to the result channel and any error is sent to the error channel.
func fetchRandomData(ctx context.Context, resource string, size int, result chan []byte, errChan chan error) {
	// Create a new HTTP GET request for the resource
	req, err := http.NewRequest("GET", baseURL+resource, nil)
	if err != nil {
		errChan <- err
		return
	}

	// Add the 'size' query parameter to the request
	query := req.URL.Query()
	query.Add("size", fmt.Sprintf("%d", size))
	req.URL.RawQuery = query.Encode()

	// Create an HTTP client and execute the request with the provided context
	client := &http.Client{}
	resp, err := client.Do(req.WithContext(ctx))
	if err != nil {
		errChan <- err
		return
	}
	defer resp.Body.Close()

	// Read the response body
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		errChan <- err
		return
	}

	// Send the response data to the result channel
	result <- data
}

func main() {
	now := time.Now()
	// Open a file for writing the fetched user data
	file, err := os.Create("random_users.csv")
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer file.Close()
	file.WriteString("id,name,email\n")

	// Initialize a WaitGroup to synchronize the goroutines
	var wg sync.WaitGroup

	// Define the number of goroutines to run
	numRoutines := 10

	// Use buffered channels to collect the results from all goroutines
	dataChan := make(chan []byte, numRoutines)
	errChan := make(chan error, numRoutines)

	// Launch x goroutines fetching data concurrently
	for i := 0; i < numRoutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// Create a context with a timeout of x seconds for each goroutine
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()

			fetchRandomData(ctx, "users", 100, dataChan, errChan)
		}()
	}

	// Wait for all goroutines to finish and close the channels
	go func() {
		wg.Wait()
		close(dataChan)
		close(errChan)
	}()

	// Collect and process results from all goroutines
	for data := range dataChan {
		var users []User
		if err := json.Unmarshal(data, &users); err != nil {
			fmt.Printf("Error unmarshalling data: %v\n", err)
			continue
		}

		// Write the fetched user data to the file
		for _, user := range users {
			_, _ = file.WriteString(fmt.Sprintf("%d, %s %s, %s\n", user.ID, user.FirstName, user.LastName, user.Email))
		}
	}

	// Check for errors from all goroutines
	for err := range errChan {
		if err != nil {
			fmt.Printf("Error fetching data: %v\n", err)
		}
	}
	println(time.Since(now).String())
}
