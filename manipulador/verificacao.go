package manipulador

import (
	"strings"
)

//BuscarLetra procura a letra que o jogador escolhe na palavra
func BuscarLetra(letra string, palavraEscolhida string, palavraDoJogador []string) ([]string, bool) {
	palavraEscolhida = strings.ToLower(palavraEscolhida)
	var achou bool

	for i := 0; i < len(palavraEscolhida); i++ {
		letraDaPalavra := string(palavraEscolhida[i])
		if letraDaPalavra == letra {
			palavraDoJogador[i] = strings.ToUpper(letra)
			achou = true
		}
	}
	return palavraDoJogador, achou
}

//DecisaoDeVitoria verifica quando se a palavra ja foi completada
func DecisaoDeVitoria(palavraDoJogador []string) bool {
	for _, conteudo := range palavraDoJogador {
		if conteudo == " " {
			return true
		}
	}
	return false
}
