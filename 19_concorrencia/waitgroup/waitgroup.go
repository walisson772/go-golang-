package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	func () {
		escrever("Olá mundo")
		waitGroup.Done()
	}()
	
	func () {
		escrever("Programando em go")
		waitGroup.Done()
	}()

	waitGroup.Wait()
}

func escrever(texto string) {
	for i := 0; i < 5; i += 1{
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
