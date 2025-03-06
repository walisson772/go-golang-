package main

import "fmt"

func main() {
	fmt.Println("Estruturas de controle")

	numero := 10

	if numero == 10 {
		fmt.Println("O numero é 10")
	}else if numero == 15 {
		fmt.Println("Esse numero é menor que 15")
	}else{
		fmt.Println("Esse numero é maior que 15")
	}

	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Esse numero é maior que 0")
	}
}