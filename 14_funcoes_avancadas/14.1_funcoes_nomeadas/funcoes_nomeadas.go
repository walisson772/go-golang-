package main

import "fmt"

func funcNomeada(n1 int, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}

func main() {
	soma, subtracao := funcNomeada(10, 20)
	fmt.Println(soma, subtracao)
}