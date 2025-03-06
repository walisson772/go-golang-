package main

import "fmt"

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}

	return total
}

func escrever(texto string, numero... int) {
	for _, num := range numero {
		fmt.Println(texto, num)
	}
}

func main() {
	totalsoma := soma(10, 22, 11)
	fmt.Println(totalsoma)

	escrever("Olá mundo", 1, 2, 3, 4,)
} 