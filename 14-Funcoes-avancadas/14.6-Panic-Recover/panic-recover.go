package main

import "fmt"

func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Execuçao recuperada com sucesso!")
	}
}

func alunoEstaAprovado(n1, n2 float64) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A MEDIA E EXATAMENTE 6!")
}

func main() {
	fmt.Println(alunoEstaAprovado(6, 6))
	fmt.Println("Pos execuçao")
}
