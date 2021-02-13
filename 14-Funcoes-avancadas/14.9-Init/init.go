package main

import "fmt"

var n int

func main() {
	fmt.Println("Funçao main sendo executada")
	fmt.Println(n)
}

// Pode ter uma funçao init por arquivo e sera executada
// antes de qualquer coisa, usada para fazer um setup inicial
func init() {
	n = 10
	fmt.Println("Executando a funçao init")
}
