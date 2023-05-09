package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// Random Data API Docs: https://random-data-api.com/documentation

// TODO: Create constant for the base URL: https://random-data-api.com/api/v2/
const baseURL = "https://random-data-api.com/api/v2/"

// TODO: Create a struct any resource you'd like to fetch
// TODO: Add json tags to the struct fields to unmarshal the response
// For example
type someResource struct {
	ID int `json:"id"`
}

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

// fetchRandomData fetches random data of the specified size from the provided resource.
// It takes a context for cancellation and timeout, the resource name, and the desired size.
// The function returns the fetched data as a byte slice and any error encountered.
func fetchRandomData(ctx context.Context, resource string, size int) ([]byte, error) {
	// Create a new HTTP GET request for the resource
	req, err := http.NewRequest("GET", baseURL+resource, nil)
	if err != nil {
		return nil, err
	}

	// Add the 'size' query parameter to the request
	query := req.URL.Query()
	query.Add("size", fmt.Sprintf("%d", size))
	req.URL.RawQuery = query.Encode()

	// Create an HTTP client to execute the request
	client := &http.Client{}
	// Execute the request with the provided context
	resp, err := client.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	// Ensure the response body is closed after the function returns
	defer resp.Body.Close()

	// Read the response body and return it as a byte slice
	return ioutil.ReadAll(resp.Body)
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

	// Write the fetched user data to the file
	_, _ = file.WriteString("ID, Name, Email\n")
	// General TODO: always handle errors
	// TODO: Create a context with a timeout of 800 milliseconds
	
	n := 0
	for n < 10 {
		n++

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		// TODO: Fetch random data from the API with random size
		data, err := fetchRandomData(ctx, "users", 100)
		if err != nil {
			fmt.Printf("Error fetching data: %v\n", err)
			return
		}

		// TODO: Unmarshal the response data into your struct with json.Unmarshal()
		var users []User
		if err := json.Unmarshal(data, &users); err != nil {
			fmt.Printf("Error unmarshalling data: %v\n", err)
			return
		}

		// Put data into a file
		for _, user := range users {
			_, _ = file.WriteString(fmt.Sprintf("%d, %s %s, %s\n", user.ID, user.FirstName, user.LastName, user.Email))
		}
	}
	println(time.Since(now).String())
}
