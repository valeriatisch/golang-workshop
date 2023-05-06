package main

import (
	"fmt"
)

// Create type "Currency"
type Currency float64

// Create constants USD = 1.0, EUR = 0.91, GBP = 0.79
const (
	USD Currency = 1.0
	EUR Currency = 0.91
	GBP Currency = 0.79
)

func convert(amount Currency, from Currency, to Currency) Currency {
	return (amount / from) * to
}

func main() {

	// Ask the user for the amount, its currency and the target currency
	var amount Currency
	var source, target string

	fmt.Print("Enter the amount: ")
	fmt.Scan(&amount)

	fmt.Print("Enter the source currency (USD, EUR, GBP): ")
	fmt.Scan(&source)

	fmt.Print("Enter the target currency (USD, EUR, GBP): ")
	fmt.Scan(&target)

	var sourceCurrency, targetCurrency Currency
	switch source {
	case "USD":
		sourceCurrency = USD
	case "EUR":
		sourceCurrency = EUR
	case "GBP":
		sourceCurrency = GBP
	default:
		fmt.Println("Invalid source currency.")
		return
	}

	switch target {
	case "USD":
		targetCurrency = USD
	case "EUR":
		targetCurrency = EUR
	case "GBP":
		targetCurrency = GBP
	default:
		fmt.Println("Invalid target currency.")
		return
	}

	convertedAmount := convert(amount, sourceCurrency, targetCurrency)
	fmt.Printf("%.2f %s is equal to %.2f %s.\n", amount, source, convertedAmount, target)
}