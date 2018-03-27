package main

import (
	"daniloelias/JogoDaVelha/manipulador"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	jogarNovamente := true

	for jogarNovamente {
		palavraEscolhida, tipoEscolhido := manipulador.EscolherTipoDaPalavra()
		manipulador.Apresentacao(palavraEscolhida, tipoEscolhido)
		jogarNovamente = manipulador.Partida(palavraEscolhida)
	}
	fmt.Printf("Obrigado por jogar!")
}
