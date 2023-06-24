package main

import (
	"fmt"
)

func sumWordCounts(wordCounts []map[string]int) map[string]int {
	totalCounts := make(map[string]int)

	for _, count := range wordCounts {
		for word, frequency := range count {
			totalCounts[word] += frequency
		}
	}

	return totalCounts
}

func main() {
	text1 := "apple banana apple cherry"
	text2 := "banana cherry cherry"
	text3 := "apple apple apple"

	counts1 := countWords(text1)
	counts2 := countWords(text2)
	counts3 := countWords(text3)

	wordCounts := []map[string]int{counts1, counts2, counts3}

	totalCounts := sumWordCounts(wordCounts)

	for word, frequency := range totalCounts {
		fmt.Printf("Word: %s, Frequency: %d\n", word, frequency)
	}
}

func countWords(text string) map[string]int {
	wordCount := make(map[string]int)
	words := splitIntoWords(text)

	for _, word := range words {
		wordCount[word]++
	}

	return wordCount
}

func splitIntoWords(text string) []string {
	
	return []string{"apple", "banana", "cherry"}
}
