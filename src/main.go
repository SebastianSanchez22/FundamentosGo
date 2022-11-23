package main

import "fmt"

func main() {
	helloMessage := "Hello"
	worldMessage := "World"

	// Println
	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)

	// Printf
	nombre := "Sebastian"
	edad := 20

	fmt.Printf("Soy %s y tengo %d años\n", nombre, edad)
	fmt.Printf("Soy %v y tengo %v años\n", nombre, edad) // %v se usa si no se sabe el tipo de dato que se va imprimir

	// Sprintf
	message := fmt.Sprintf("Soy %s y tengo %d años", nombre, edad)
	fmt.Println(message)

	// Tipo de dato de una variable
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("edad: %T", edad)

}
