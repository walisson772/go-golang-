package main

import (
	"fmt"
	"time"
)
func tabuadaSoma(numero int)  {
	fmt.Printf("Tabuada de soma do numero %d\n", numero)
	for i := 0; i <= 10; i++ {
		result := numero + i
		fmt.Printf("%d + %d = %d\n", numero, i, result)
		time.Sleep(time.Second)
	}
	
}

func main() {
	tabuadaSoma(3)
}