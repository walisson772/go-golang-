package main

import "fmt"

type usuario struct {
	nome string
	idade int8
	endereco endereco
}

type endereco struct {
	longradouro string
	numero int8
}

func main() {
	fmt.Println("opa")
	var u usuario 
	u.nome = "walisson"
	u.idade = 17
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua dos otarios", 99}

	usuario1 := usuario{"walisson", 17, enderecoExemplo}
	fmt.Println(usuario1)
}