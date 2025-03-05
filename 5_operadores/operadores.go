package main

import "fmt"

func main() {
	//  ARITIMETICOS
	soma := 1 + 1
	subtracao := 4 - 2
	divisao := 4 / 4
	multiplicacao := 2 * 2
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	var numero1 int16 = 10
	var numero2 int16 = 25
	soma2 := numero1 + numero2
	fmt.Println(soma2)

	// FIM DOS ARITIMETICOS

	// ATRIBUIÇÃO
	var variavel1 string = "String"
	variavel2 := "String"
	fmt.Println(variavel1, variavel2)
	// FIM ATRIBUIÇÃO

	// OPERADORES RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 != 2)
	// FIM OPERADORES RELACIONAIS

	// OPERADORES LOGICOS
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)
	// FIM OPERADORES LOGICOS

	// OPERADORES UNÁRIOS
	numero := 10
	numero += 5 // numero = numero + 5
	
	numero -= 5
	
	numero /= 5

	numero %= 3
	fmt.Println(numero)
	// FIM OPERADORES UNÁRIOS

	var texto string
	if numero >= 5 {
		texto = "Maior que 5"
	}else {
		texto = "Menor que 5"
	}

	fmt.Println(texto)
}
