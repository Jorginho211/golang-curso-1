package config

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// StringConexaoBanco é a string de conexao com o baco de dados
	StringConexaoBanco = ""

	// Porta onde a API vai estar rodando
	Porta = 0

	// Chave para gerar o token
	SecretKey []byte
)

// Carregar vai inicializar as variaveis de ambiente
func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Porta, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		Porta = 9000
	}

	StringConexaoBanco = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_NOME"),
	)

	SecretKey, erro = base64.StdEncoding.DecodeString(os.Getenv("SECRET_KEY"))
	if erro != nil {
		log.Fatal(erro)
	}
}
