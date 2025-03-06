package main

import (
	"fmt"
)

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao - 2) + fibonacci(posicao - 1)

}

func main() {
	fmt.Println("Funções recursivas")

	var posicao uint = 10
	fmt.Println(fibonacci(posicao))
}