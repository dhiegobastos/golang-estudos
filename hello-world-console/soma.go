package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := Soma(10, 1)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result)
	}

	// Blank identifier
	result2, _ := Soma(2, 1)
	fmt.Println(result2)

	// Soma tudo
	fmt.Println(SomaTudo(1, 2, 3, 4, 5))
}

// Soma realiza a soma de dois números e aplica regra de validação como erro
func Soma(a int, b int) (int, error) {
	result := a + b
	if result > 10 {
		return 0, errors.New("Total maior que 10.")
	}

	return result, nil
}

// SomaTudo com varArgs e retorno nomeado
func SomaTudo(x ...int) (result int) {
	result = 0

	for _, value := range x {
		result += value
	}

	return
}
