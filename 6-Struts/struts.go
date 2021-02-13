package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint
}

func main() {
	fmt.Println("Arquivo struts")

	var u usuario
	fmt.Println(u) // Devolve o valor zero de cada campo
	u.nome = "Davi"
	u.idade = 21
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua dos Bobos", 0}
	usuario2 := usuario{"Davi", 21, enderecoExemplo}
	fmt.Println(usuario2)

	//Somente queremos crialo com 1 parametro
	usuario3 := usuario{nome: "Davi"}
	fmt.Println(usuario3)
	usuario4 := usuario{idade: 21}
	fmt.Println(usuario4)
}
