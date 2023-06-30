package main

import (
	"fmt"
)

func countCharacters(str string) map[rune]int {
	charCount := make(map[rune]int)

	for _, char := range str {
		charCount[char]++
	}

	return charCount
}

func main() {
	text := "hello world"

	charFrequency := countCharacters(text)

	for char, frequency := range charFrequency {
		fmt.Printf("Character: %c, Frequency: %d\n", char, frequency)
	}
}
