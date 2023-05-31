package mathutil

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func checkErr(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}

func TestMultiply(t *testing.T) {
	res := Multiply(2, 3)
	if res != 6 {
		t.Errorf("Multiply was incorrect, got: %d, exptected: %d", res, 6)
	}
}

func TestMultiplyTablte(t *testing.T) {
	testCases := []struct {
		x   int
		y   int
		res int
	}{
		{2, 3, 6},
		{0, 3, 0},
		{20, 50, 1000},
	}

	for _, c := range testCases {
		res := Multiply(c.x, c.y)
		if res != c.res {
			t.Errorf("Multiply was incorrect, got: %d, exptected: %d", res, c.res)
		}
	}
}

func ExampleMultiply() {
	res := Multiply(2, 3)
	fmt.Println(res)
	// Output: 6
}

func BenchmarkNoMake(b *testing.B) {
	n := 100
	for i := 0; i < b.N; i++ {
		// The slice is created with capacity 0.
		ints := []int{}
		for j := 0; j < n; j++ {
			ints = append(ints, j)
			// When capacity is reached, a new slice with doubled capacity is being created,
			// elements are copied and the old slice is discarded.
		}
	}
}

func BenchmarkMake(b *testing.B) {
	n := 100
	for i := 0; i < b.N; i++ {
		ints := make([]int, n)
		for j := 0; j < n; j++ {
			ints[j] = j
		}
	}
}

func TestHandleGet(t *testing.T) {
	// Create a test server with the HandleGet handler
	server := httptest.NewServer(http.HandlerFunc(HandleGet))
	defer server.Close()

	// Send GET request to our test server
	res, err := http.Get(server.URL)
	checkErr(t, err)
	defer res.Body.Close()

	// Check the status code
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected %v, go %v", http.StatusOK, res.StatusCode)
	}

	// Decode the response body
	var response struct {
		Result int `json:"result"`
	}
	err = json.NewDecoder(res.Body).Decode(&response)
	checkErr(t, err)

	// Check if result is between 0 and 100
	if response.Result < 0 || response.Result > 100 {
		t.Errorf("Expected result to be between 0 and 100, got %d", response.Result)
	}
}

func TestHandleGetR(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	checkErr(t, err)

	rr := httptest.NewRecorder()

	HandleGet(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected 200, got %v", rr.Code)
	}

	cType := "application/json"
	if rr.Header().Get("Content-Type") != cType {
		t.Errorf("Expected application/json, got something else")
	}

	var response struct {
		Result int `json:"result"`
	}

	err = json.NewDecoder(rr.Body).Decode(&response)
	checkErr(t, err)

	// Check if result is between 0 and 100
	if response.Result < 0 || response.Result > 100 {
		t.Errorf("Expected result to be between 0 and 100, got %d", response.Result)
	}
}
