package main

import (
	"fmt"
	"hello/math"
)

var a string // Declaração

func main() {
	a = "Dhiego"  // Atribuição
	b := "Fulano" // Declaração + atribuição

	fmt.Printf("Hello, %v! %T\n", a, a)
	fmt.Printf("Hello, %v! %T\n", b, b)

	resultado := math.Soma(10, 20)

	fmt.Printf("Resultado: %v\n", resultado)
}
