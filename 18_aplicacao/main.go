package main

import (
	"fmt"
	"linha_de_comando/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Ponto de partida")
	aplicacao := app.Gerar()
	
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}