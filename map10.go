package main

import "fmt"

func countPairs(slice []int) map[[2]int]int {
	pairCount := make(map[[2]int]int)

	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			pair := [2]int{slice[i], slice[j]}
			pairCount[pair]++
		}
	}

	return pairCount
}

func main() {
	numbers := []int{1, 2, 3, 2, 4, 5, 1, 3, 4, 2} // Exemplo de slice de inteiros

	pairFrequency := countPairs(numbers)

	fmt.Println("Frequência dos pares:")
	for pair, frequency := range pairFrequency {
		fmt.Printf("%v: %d\n", pair, frequency)
	}
}
