package main

import (
	"fmt"
	"strings"
)

func contarPalavras(str string) map[string]int {
	ocorrencias := make(map[string]int)

	str = strings.ToLower(str)
	str = strings.ReplaceAll(str, ",", "")
	str = strings.ReplaceAll(str, ".", "")

	palavras := strings.Split(str, " ")

	// Contar ocorrências de cada palavra
	for _, palavra := range palavras {
		ocorrencias[palavra]++
	}

	return ocorrencias
}

func main() {
	texto := "O rato roeu a roupa do rei de Roma"

	resultado := contarPalavras(texto)

	fmt.Println("Ocorrências de palavras:")
	for palavra, ocorrencias := range resultado {
		fmt.Printf("%s: %d\n", palavra, ocorrencias)
	}
}
