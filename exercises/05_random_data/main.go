package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

// Random Data API Docs: https://random-data-api.com/documentation

// TODO: Create constant baseURL for: https://random-data-api.com/api/v2/

// TODO: Create a struct any resource you'd like to fetch
// TODO: Add json tags to the struct fields to unmarshal the response
// For example
type someResource struct {
	ID        int    `json:"id"`
}

func fetchRandomData(ctx context.Context, resource string, size int) ([]byte, error) {
	req, err := http.NewRequest("GET", baseURL+resource, nil)
	if err != nil {
		return nil, err
	}

	query := req.URL.Query()
	query.Add("size", fmt.Sprintf("%d", size))
	req.URL.RawQuery = query.Encode()

	client := &http.Client{}
	resp, err := client.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func main() {
	// General TODO: always handle errors
	// TODO: Create a context with a timeout of 800 milliseconds

	// TODO: Fetch random data from the API with random size

	// TODO: Unmarshal the response data into your struct with json.Unmarshal()

	// TODO: Print the data
}
