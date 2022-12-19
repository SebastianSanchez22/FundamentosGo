package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Sebastian"] = 20
	m["Persona"] = 21

	fmt.Println(m)

	// Recorrer map
	for i, v := range m {
		fmt.Println(i, v)
	}

	// Encontrar valor
	value1, ok1 := m["Sebastian"]
	value2, ok2 := m["Sebastian2"]
	fmt.Println(value1, ok1)
	fmt.Println(value2, ok2)
}
