package enderecos_test

import (
	. "introducao-testes/enderecos" // O . faz nao ter de usar enderecos.TipoDeEndereco
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {
	t.Parallel() // Indica que este pode ser lançado em paralelo com outros

	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABD", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia das Rosas", "Rodovia"},
		{"Praça das Rosas", "Tipo invalido"},
		{"Estrada Qualquer", "Estrada"},
		{"RUA DOS BOBOS", "Rua"},
		{"AVENIDA REBOUÇAS", "Avenida"},
		{"", "Tipo invalido"},
	}

	for _, cenario := range cenariosDeTeste {
		tipoDeEnderecoRecevedo := TipoDeEndereco(cenario.enderecoInserido)

		if tipoDeEnderecoRecevedo != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s",
				tipoDeEnderecoRecevedo,
				cenario.retornoEsperado,
			)
		}
	}
}

func TestQualquer(t *testing.T) {
	t.Parallel()

	if 1 > 2 {
		t.Errorf("Teste erro!")
	}
}
