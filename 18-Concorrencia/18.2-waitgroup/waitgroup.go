package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go func() {
		escrever("Ola mundo")
		waitGroup.Done() // Faz -1
	}()

	go func() {
		escrever("Programando em Go")
		waitGroup.Done() // Faz -1
	}()

	waitGroup.Wait() // Espera a que as 2 gouroutine finalizem
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
