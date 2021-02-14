package main

import (
	"fmt"
)

func main() {
	canal := make(chan string, 2)

	canal <- "Ola Mundo"
	canal <- "Programando em Go"
	//canal <- "Terceiro valor!" // Deadlock porque o buffer esta cheio

	menssagem := <-canal
	menssagem2 := <-canal

	fmt.Println(menssagem)
	fmt.Println(menssagem2)
}
