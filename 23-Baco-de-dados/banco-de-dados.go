package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//usuario:senha@/banco
	stringConexao := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", stringConexao)

	if erro != nil {
		fmt.Println("Dentro do SQL.Open")
		log.Fatal(erro)
	}
	defer db.Close()

	// Para saber se a conexao esta aberta
	if erro = db.Ping(); erro != nil {
		fmt.Println("Dentro do Ping")
		log.Fatal(erro)
	}

	fmt.Println("Conexao está aberta")

	linhas, erro := db.Query("select * from usuarios")
	if erro != nil {
		log.Fatal(erro)
	}
	defer linhas.Close()

	fmt.Println(linhas)
}
