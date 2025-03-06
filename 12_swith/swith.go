package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"

	case 2:
		return "Segunda-Feira"

	case 3:
		return "Terça-Feira"

	case 4:
		return "Quarta-Feira"

	case 5:
		return "Quinta-Feira"

	case 6:
		return "Sexta-Feira"

	case 7:
		return "Sabado"

	default:
		return "Numero invalido"
	}
}	

func diaDaSemana2(numero int) string {
	var diaSemana string

	switch {
	case numero == 1:
		diaSemana = "Domingo"

	case numero == 2:
		diaSemana = "Segunda-Feira"

	case numero == 3:
		diaSemana = "Terça-Feira"

	case numero ==  4:
		diaSemana = "Quarta-Feira"

	case numero == 5:
		diaSemana = "Quinta-Feira"

	case numero ==  6:
		diaSemana = "Sexta-Feira"

	case numero ==  7:
		diaSemana = "Sabado"

	default:
		diaSemana = "Numero invalido"
	}

	return diaSemana
}

func main() {
	fmt.Println("Swith")

	dia := diaDaSemana(5)
	fmt.Println(dia)

	dia2 := diaDaSemana2(6)
	fmt.Println(dia2)
	
}