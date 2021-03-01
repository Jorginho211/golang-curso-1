package modelos

// DadosAutenticacao contem o id e o token do usuario autenticado
type DadosAutenticacao struct {
	Id    string `json:"id"`
	Token string `json:"token"`
}
