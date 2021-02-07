package main

import (
	"errors"
	"fmt"
)

func main() {
	// int8, int16, int32, int64 // Numero de bits que ocupa
	// int - Usa a arquitetura do computador como base 32 bits ou 64 bits

	var numero int16 = 100
	fmt.Println(numero)

	numero2 := -1000000 // Inferiria o tipo int
	fmt.Println(numero2)

	// uint8, uint16, uint32, uint64 // Sem signo
	var numero3 uint32 = 1000
	fmt.Println(numero3)

	// alias
	// INT32 == RUNE
	var numero4 rune = 10000
	fmt.Println(numero4)

	// UINT8 == BYTE
	var numero5 byte = 125
	fmt.Println(numero5)

	// FIM NUMEROS INTEIROS

	// float32, float64
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 123000000000.4512321321312
	fmt.Println(numeroReal2)

	// Infire
	numeroReal3 := 123000000000.4512321321312
	fmt.Println(numeroReal3)

	//FIM NUMEROS REAIS

	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	// O char nao existe, somente mostra o numero ascii
	char := 'B'
	fmt.Println(char)

	//FIM STRINGS

	// Os dados em go sao sempre inicialiçados com o valor 0
	// int, float32... = 0
	// string = \0
	// bool = false
	var texto string
	fmt.Println(texto)

	var booleano1 bool = true
	fmt.Println(booleano1)

	// Muito usado nas funçoes para o tratamento de erros
	var erro error = errors.New("Error interno")
	fmt.Println(erro)
}
