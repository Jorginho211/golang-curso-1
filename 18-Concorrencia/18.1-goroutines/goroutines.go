package main

import (
	"fmt"
	"time"
)

func main() {
	go escrever("Ola mundo")
	escrever("Programando em Go")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
