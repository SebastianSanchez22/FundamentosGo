package main

import (
	"fmt"
	"strings"
)

func isPalindromo(text string) bool {
	text = strings.ToLower(text)

	for i := 0; i < len(text)/2; i++ {
		if text[i] != text[len(text)-i-1] {
			return false
		}
	}

	return true

}

func main() {

	slice := []string{"ejercicio", "con", "slices"}

	for i := range slice {
		fmt.Println(i)
	}

	for _, valor := range slice {
		fmt.Println(valor)
	}

	fmt.Println(isPalindromo("level"))
	fmt.Println(isPalindromo("test"))

}
