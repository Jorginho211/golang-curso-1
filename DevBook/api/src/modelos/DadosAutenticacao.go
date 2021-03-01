package modelos

// DadosAutenticacao conten o token e o Id do usuario autenticado
type DadosAutenticacao struct {
	Id    string `json:"id"`
	Token string `json:"token"`
}
