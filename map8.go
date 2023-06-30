package main

import (
	"fmt"
)

func equalizeExpenses(expenses map[string]float64) map[string]float64 {
	total := 0.0

	for _, value := range expenses {
		total += value
	}

	average := total / float64(len(expenses))

	equalizedExpenses := make(map[string]float64)

	for name, expense := range expenses {
		equalizedExpenses[name] = average - expense
	}

	return equalizedExpenses
}

func main() {
	expenses := map[string]float64{
		"João":   50.0,
		"Maria":  100.0,
		"Carlos": 75.0,
	}

	equalized := equalizeExpenses(expenses)

	fmt.Println("Equalized expenses:")
	for name, value := range equalized {
		fmt.Printf("%s: %.2f\n", name, value)
	}
}
