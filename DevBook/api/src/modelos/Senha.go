package modelos

// Senha representa o formato da requisiçao de alteraçao da senha
type Senha struct {
	Nova  string `json:"nova"`
	Atual string `json:"atual"`
}
