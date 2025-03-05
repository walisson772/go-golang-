package main

import "fmt"

func main() {
	fmt.Println("Ponteiro")

	var variavel int = 10
	var variavel2 int = variavel

	fmt.Println(variavel, variavel2)

	//	 PONTEIRO É UMA REFERÊNCIA DE MEMÓRIA
	var variavel3 int
	var ponteiro *int

	variavel3 = 100
	ponteiro = &variavel3

	fmt.Println(variavel3, *ponteiro) // desferenciação

	variavel3 = 150
	fmt.Println(*ponteiro)
}