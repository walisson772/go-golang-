package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int8
	altura    int16
}

type studante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	p1 := pessoa{"walisson", "Santos", 17, 178}

	e1 := studante{p1, "Tec. Informatica", "Ceteps"}
	fmt.Println(p1)
	fmt.Println(e1)

}