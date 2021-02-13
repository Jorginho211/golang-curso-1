package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

// Se coloca o tipo do struct para fazer "herança"
type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança")

	p1 := pessoa{"Joao", "Pedro", 20, 178}
	fmt.Println(p1)

	// É preciso passar uma pessoa, podesse criar tambem direitamente sem
	// variavel intermedia
	e1 := estudante{p1, "Engenharia", "USP"}
	fmt.Println(e1)
	fmt.Println(e1.nome) //Accedese direitamente ao atributo herdado de pessoa
}
