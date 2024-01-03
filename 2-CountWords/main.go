package main

import (
	"fmt"
	"strings"
)

func countWords(words string) int32 {
	// Use strings.Fields para dividir a string em palavras
	palavras := strings.Fields(words)

	// Crie um mapa para rastrear a contagem de cada palavra
	contagemPalavras := make(map[string]int)

	// Itere sobre as palavras e atualize a contagem no mapa
	for _, palavra := range palavras {
		contagemPalavras[palavra]++
	}

	// Retorne o número total de palavras únicas
	return int32(len(contagemPalavras))
}

func main() {
	totalPalavrasUnicas := countWords("teste ola teste")
	fmt.Printf("Número total de palavras únicas: %d\n", totalPalavrasUnicas)
}
