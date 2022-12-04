package main

import "fmt"

func main() {

	// Defer
	defer fmt.Println("Hola")
	fmt.Println("Mundo")

	for i := 0; i < 5; i++ {
		defer fmt.Println(i) // Sigue el principio LIFO (Last in first out, pila)
	}

	// Continue y break
	for i := 0; i < 5; i++ {
		// continue
		if i == 2 {
			continue
		}
		if i == 8 {
			fmt.Println("Se hace el break")
			break
		}
		fmt.Println(i)
	}
}
