package main

import "fmt"

// Fibonacci: 1 1 2 3 5 8 13
func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {
	posicao := uint(12) // Convertido se non infeririase int

	for i := uint(0); i <= posicao; i++ {
		fmt.Println(fibonacci(i))
	}
}
