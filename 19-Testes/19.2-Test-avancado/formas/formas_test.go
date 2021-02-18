package formas

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Retangulo", func(t *testing.T) {
		ret := Retangulo{10, 12}
		areaEsperada := float64(120)
		areaRecibida := ret.Area()

		if areaEsperada != areaRecibida {
			// O Fatal nao continua executando os tests
			t.Fatalf("A area recebida %f e diferente da esperada %f",
				areaRecibida,
				areaEsperada,
			)
		}
	})

	t.Run("Circulo", func(t *testing.T) {
		ret := Circulo{10}
		areaEsperada := float64(math.Pi * 100)
		areaRecibida := ret.Area()

		if areaEsperada != areaRecibida {
			t.Fatalf("A area recebida %f e diferente da esperada %f",
				areaRecibida,
				areaEsperada,
			)
		}
	})
}
