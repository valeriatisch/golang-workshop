// Mathutil package to multiply numbers

package mathutil

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
)

// Multiply multiplies two numbers and returns the result
func Multiply(a ...int) int {
	res := 1
	for _, v := range a {
		if v == 0 {
			return 0
		}
		res *= v
	}
	return res
}

// MultiplyHandler handles a POST request to multiply two numbers and returns the result as JSON
func HandleMultiply(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	aStr := r.FormValue("a")
	bStr := r.FormValue("b")

	a, erra := strconv.Atoi(aStr)
	b, errb := strconv.Atoi(bStr)
	if erra != nil || errb != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result := Multiply(a, b)

	response := struct {
		Result int `json:"result"`
	}{
		Result: result,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Get Handler
func HandleGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	response := struct {
		Result int `json:"result"`
	}{
		Result: rand.Intn(100),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
