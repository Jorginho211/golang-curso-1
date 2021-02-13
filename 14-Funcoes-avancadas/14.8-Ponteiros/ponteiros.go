package main

import "fmt"

func inverterSinal(numero int) int {
	numero *= -1
	return numero
}

func inverterSinalComPonteiro(numero *int) {
	*numero *= -1
}

func main() {
	numero := 20
	numeroInvertido := inverterSinal(numero)

	fmt.Println(numeroInvertido)
	fmt.Println(numero)

	novoNumero := 40
	fmt.Println(novoNumero)
	inverterSinalComPonteiro(&novoNumero)
	fmt.Println(novoNumero)
}
