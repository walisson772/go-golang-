package main

import (
	"fmt"
	"time"
)

func main() {
	// fmt.Println("Goroutines")
	go escrever("Olá mundo")
	escrever("Programando em go")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
