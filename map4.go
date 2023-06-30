package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortString(s string) string {
	chars := strings.Split(s, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}

func findAnagrams(words []string) map[string][]string {
	anagrams := make(map[string][]string)

	for _, word := range words {
		sortedWord := sortString(word)
		anagrams[sortedWord] = append(anagrams[sortedWord], word)
	}

	return anagrams
}

func main() {
	wordList := []string{"listen", "silent", "enlist", "inlets", "banana", "nabana"}

	anagramGroups := findAnagrams(wordList)

	for key, group := range anagramGroups {
		fmt.Println("Anagram Group:", group)
		fmt.Println("Key:", key)
		fmt.Println()
	}
}
