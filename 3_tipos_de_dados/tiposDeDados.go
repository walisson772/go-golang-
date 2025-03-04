package main

import (
	"errors"
	"fmt"
)

func main() {
	// numeros inteiros
	var numero int = 1000
	fmt.Println(numero)

	var numero2 uint = 1000
	fmt.Println(numero2)

	//alias 
	// int32 = rune
	var numero3 rune = 123456
	fmt.Println(numero3)

	//BYTE = UINT8
	var numero4 byte = 123
	fmt.Println(numero4)
	//Fim inteiros

	// numeros reais
	var numeroreal float32 = 123.43
	fmt.Println(numeroreal)

	var numeroreal2 float64  = 123123.45
	fmt.Println(numeroreal2)
	// Fim reais

	// Strings
	var str string = "walisson"
	fmt.Println(str)
	// Fim String

	var booleano bool 
	fmt.Println(booleano)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
}