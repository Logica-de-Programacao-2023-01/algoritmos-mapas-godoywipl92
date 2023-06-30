package main

import (
	"fmt"
	"strings"
)

func countLettersInWords(frase string) map[string]map[rune]int {
	wordCounts := make(map[string]map[rune]int)

	words := strings.Fields(frase)

	for _, word := range words {
		letterCount := make(map[rune]int)

		for _, letter := range word {
			letterCount[letter]++
		}

		wordCounts[word] = letterCount
	}

	return wordCounts
}

func main() {
	frase := "O rato roeu a roupa do rei de Roma"

	letterCounts := countLettersInWords(frase)

	for word, count := range letterCounts {
		fmt.Printf("Word: %s\n", word)
		for letter, frequency := range count {
			fmt.Printf("Letter: %c, Frequency: %d\n", letter, frequency)
		}
		fmt.Println()
	}
}
