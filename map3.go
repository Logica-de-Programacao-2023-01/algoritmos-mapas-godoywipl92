package main

import (
	"fmt"
)

func sumMapValues(m map[string]int) int {
	sum := 0
	for _, value := range m {
		sum += value
	}
	return sum
}

func main() {
	myMap := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	totalSum := sumMapValues(myMap)

	fmt.Println("Total Sum:", totalSum)
}
