package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	var array1 [5]string
	array1[0] = "Posiçao 1"
	fmt.Println(array1)

	array2 := [5]string{"Posiçao 1", "Posiçao 2", "Posiçao 3", "Posiçao 4", "Posiçao 5"}
	fmt.Println(array2)

	// Tamanho do array em base ao numero de items
	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	// Corchetes sem valor para criar un slice
	slice := []int{10, 11, 12, 13, 14, 15, 16, 17}

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))
	fmt.Println(slice)

	//Adiciona o 18 e retorna o novo slice
	slice = append(slice, 18)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	//O slice2 aponta aos mesmos lementos do array2
	array2[1] = "Posiçao alterada"
	fmt.Println(slice2)

	// ARRAYS INTERNO
	fmt.Println("--------------------")

	// make guarda um espaço em memoria
	// 1º Argumento: Tipo
	// 2º Argumento: Tamanho
	// 3º Argumento: Capacidade (Opcional)
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // length
	fmt.Println(cap(slice3)) // capacidade

	slice3 = append(slice3, 5)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice3 = append(slice3, 6)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3)) // Dobra a capacidade 12 * 2

	slice4 := make([]float32, 5) // Capacidade = Tamanho
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

	slice4 = append(slice4, 4)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4)) // 6 * 2 = 12
}
