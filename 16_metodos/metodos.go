package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do usuario %s no banco de dados\n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	if u.idade >= 18 {
		return true
	}
	return false
}

func main() {
	usuario1 := usuario{"walisson", 17}
	fmt.Println(usuario1)
	usuario1.salvar()

	usuario2 := usuario{"Santos", 12}
	usuario2.salvar()

	maiorDeIdade := usuario1.maiorDeIdade()
	fmt.Println(maiorDeIdade)
}