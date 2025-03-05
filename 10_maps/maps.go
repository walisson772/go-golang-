package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string {
		"nome": "walisson",
		"sobrenome": "Santos",
	}

	fmt.Println(usuario["nome"])
	fmt.Println(usuario["sobrenome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "walisson",
			"ultimo": "santos",
		},
	}

	fmt.Println(usuario2["nome"]["primeiro"])
	// delete(usuario2, "nome") Serve para deletar 
}