package main

import "fmt"

func tentativaRecupecao() {
	if r := recover(); r != nil {
		fmt.Println("Função recuperada com sucesso")
	}
}

func alunoEstaAprovado(n1, n2, n3 float64) bool {
	defer tentativaRecupecao()
	media := (n1 + n2 + n3) / 3
	
	if media > 6 {
		return true
	}else if media < 6 {
		return false
	} 

	panic("A médoa é exatamente 6!")
}

func main() {
	fmt.Println("Panic Recover")

	fmt.Println(alunoEstaAprovado(6, 6, 6))
	fmt.Println("Pós execução")
}