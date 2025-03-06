package main

import (
	"fmt"
	// "time"
)

func main() {
	fmt.Println("Loops")

	// i := 0

	// for i < 10 {
	// 	i++
	// 	fmt.Println("Incrementando ", i)
	// 	time.Sleep(time.Second)
	// }
	// fmt.Println(i)

	// for j := 0; j < 10; j++ {
	// 	fmt.Println("Incrementando ", j)
	// 	time.Sleep(time.Second)
	// }

	nomes := [3]string {"Walisson", "Davi", "joão"}
	for indece, nome := range(nomes) {
		fmt.Println(indece, nome)
	}

	for i, letra := range "PALAVRA" {
		fmt.Println(i, string(letra))
	}

	usuario := map[string]string {
		"nome": "walisson",
		"Sobrenome": "Santos",
	}

	for chave, valor := range(usuario) {
		fmt.Println(chave, valor)
	}
}