package manipulador

import (
	"fmt"
)

//Apresentacao será a tela inicial do jogo
func Apresentacao(palavra string, tipoDaPalavra string) {
	tamanhoDaPalavra := len(palavra)
	fmt.Printf("Sua palavra têm %d letras e é do tipo %s\n", tamanhoDaPalavra, tipoDaPalavra)
}

//Partida fará toda a parte de pedir a letra, buscar, e preencher a palavra do jogador
func Partida(palavraEscolhida string) bool {
	var letraEscolhidaDoJogador string
	vitoria := true
	palavraDoJogador := make([]string, len(palavraEscolhida))
	for i := 0; i < len(palavraDoJogador); i++ {
		palavraDoJogador[i] = " "
	}
	for vitoria {
		fmt.Printf("Escolha uma letra: ")
		fmt.Scan(&letraEscolhidaDoJogador)
		palavraDoJogador, achou := BuscarLetra(letraEscolhidaDoJogador, palavraEscolhida, palavraDoJogador)

		if achou {
			impressao(palavraDoJogador)
			vitoria = DecisaoDeVitoria(palavraDoJogador)
		} else {
			fmt.Printf("\nTente outra letra!\n")
		}
	}

	return continua(palavraDoJogador)
}

func impressao(palavraDoJogador []string) {
	for _, conteudo := range palavraDoJogador {
		if conteudo != " " {
			fmt.Printf(" %s ", conteudo)
		} else {
			fmt.Printf(" _ ")
		}
	}
}

func continua(palavraDoJogador []string) bool {
	fmt.Printf("Parabéns, você terminou a palavra: ")
	for _, conteudo := range palavraDoJogador {
		fmt.Printf(conteudo)
	}
	var continua int
	fmt.Printf("\nDeseja tentar outra palavra? Sim[1] ou Não[2]: ")
	fmt.Scan(&continua)

	if continua != 1 {
		return false
	}
	return true
}
