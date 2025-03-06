package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a função 1")
}

func funcao2() {
	fmt.Println("Executando a função 2")
}

func alunoEstaAprovado(n1, n2, n3 float32) bool {
	defer fmt.Println("Média calculada. Retornando o resultado!")
	fmt.Println("Entrando na função para calcular a media")
	media := (n1 + n2 + n3) / 3

	if media >= 6 {
		return true
	}

	return false
}

func main() {
	fmt.Println("Defer")

	defer funcao1()
	funcao2()
	fmt.Println(alunoEstaAprovado(6, 6, 9))
}