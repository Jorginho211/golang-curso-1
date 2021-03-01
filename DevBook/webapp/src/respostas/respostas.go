package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

// ErroAPi representa a resposta de erro da API
type ErroAPI struct {
	Erro string `json:"error"`
}

// JSON retorna uma resposta em formato JSON para a requisiçao
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if erro := json.NewEncoder(w).Encode(dados); erro != nil {
		log.Fatal(erro)
	}
}

// TratarStatusCodeDeErro trata as requisiçoes com status code 400 ou superior
func TratarStatusCodeDeErro(w http.ResponseWriter, r *http.Response) {
	var erro ErroAPI
	json.NewDecoder(r.Body).Decode(&erro)
	JSON(w, r.StatusCode, erro)
}
