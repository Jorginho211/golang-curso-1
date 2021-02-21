package autenticacao

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// CriarToken retorna um token assinado com as permissoes do usuario
func CriarToken(usuarioId uint64) (string, error) {
	permissoes := jwt.MapClaims{}
	permissoes["autorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissoes["usuarioId"] = usuarioId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)
	return token.SignedString(config.SecretKey)
}

// ValidarToken verifica se o token passado na requisiçao é valido
func ValidarToken(r *http.Request) error {
	tokenString := extrairToken(r)

	token, erro := jwt.Parse(tokenString, retornarChaveDeVerificacao)
	if erro != nil {
		return erro
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("Token invalido!")
}

// ExtrairUsuarioId retorna o usuarioId que está salvo no token
func ExtrairUsuarioId(r *http.Request) (uint64, error) {
	tokenString := extrairToken(r)

	token, erro := jwt.Parse(tokenString, retornarChaveDeVerificacao)
	if erro != nil {
		return 0, erro
	}

	if permissoes, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		usuarioId, erro := strconv.ParseUint(fmt.Sprintf("%.0f", permissoes["usuarioId"]), 10, 64)

		if erro != nil {
			return 0, erro
		}

		return usuarioId, nil
	}

	return 0, errors.New("Token invalido!")
}

func extrairToken(r *http.Request) string {
	authorization := r.Header.Get("Authorization")

	//Bearer <token>
	token := strings.Split(authorization, " ")
	if len(token) == 2 {
		return token[1]
	}

	return ""
}

func retornarChaveDeVerificacao(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Metodo de asinatura inexperado! %v", token.Header["alg"])
	}

	return config.SecretKey, nil
}
