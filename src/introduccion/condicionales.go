package main

import "fmt"

func main() {

	number := 100
	number2 := 200

	if number > 0 && number2 > 20 {
		fmt.Println("Se cumple")
	}

	switch modulo := 4 % 2; modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}

	// Sin condición
	value := 200
	switch {
	case value > 100:
		fmt.Println("Es mayor a 100")
	case value < 0:
		fmt.Println("Es menor a 0")
	default:
		fmt.Println("No condición")
	}
}
