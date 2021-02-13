package main

import "fmt"

func main() {
	// ARITMETICOS
	soma := 1 + 2
	substracao := 1 - 2
	var divisao int8 = 10 / 4
	multiplicacao := 10 * 5
	var restoDaDivisao int8 = 10 % 2

	fmt.Println(soma, substracao, divisao, multiplicacao, restoDaDivisao)

	// var numero1 int16 = 10
	// var numero2 int32 = 25
	// var soma2 int16 = numero1 + numero2 //Erro por ser diferentes tipo
	// fmt.Println(soma)

	//FIM DOS ARITMETICOS

	// ATRIBUIÇAO
	var variavel1 string = "String"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)
	//FIM DOS OPERADORES ATRIBUIÇAO

	//OPERADORES RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)
	//FIM DOS OPERADORES RELACIONAIS

	//OPERADORES LOGICOS
	fmt.Println("----------------------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)
	//FIM DOS OPERADORES LOGICOS

	//OPERADORES UNARIOS
	fmt.Println("----------------------------")
	numero := 10
	numero++
	fmt.Println(numero)
	numero += 15
	fmt.Println(numero)

	numero--
	numero -= 20 // numero = numero - 20

	numero *= 3 // numero = numero * 3
	numero /= 10
	numero %= 3

	// --numero //Non existe o prefixado en Go

	fmt.Println(numero)
	//FIM OPERADORES UNARIOS

	// Os operadores ternarios non existen en Go
	// Debese a que Go so quere que as cousas se poidan
	// facer dunha soa forma

	// texto := numero > 5 ? "Maior que 5" : "Menor que 5"
	// fmt.Println(text)

	// Forma alternativa de facelo
	var texto string
	if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}
	fmt.Println(texto)
}
