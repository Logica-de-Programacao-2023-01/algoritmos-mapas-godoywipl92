package main

import "fmt"

func fibonacciSequence(n int) map[int]int {
	sequence := make(map[int]int)

	sequence[0] = 0
	sequence[1] = 1

	for i := 2; sequence[i-1] <= n; i++ {
		sequence[i] = sequence[i-1] + sequence[i-2]
	}

	return sequence
}

func main() {
	n := 100 // Limite superior da sequência de Fibonacci

	fibSequence := fibonacciSequence(n)

	fmt.Println("Sequência de Fibonacci até", n)
	for index, value := range fibSequence {
		fmt.Printf("Fib[%d] = %d\n", index, value)
	}
}
