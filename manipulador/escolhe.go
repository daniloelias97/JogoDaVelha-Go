package manipulador

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
)

func escolherPalavra(tipoDaPalavra string) string {
	nomeArquivo := "palavras/" + tipoDaPalavra + ".csv"
	arquivo, err := os.Open(nomeArquivo)
	if err != nil {
		fmt.Println("[palavras] Houve um erro ao abrir o arquivo. Erro: ", err.Error())
	}
	leitorCsv := csv.NewReader(arquivo)
	conteudo, err := leitorCsv.ReadAll()
	if err != nil {
		fmt.Println("[palavras] Houve um erro ao ler o arquivo com leitor CSV. Erro: ", err.Error())
	}
	numeroAleatorio := rand.Intn(len(conteudo[0]))
	palavraEscolhida := conteudo[0][numeroAleatorio]
	arquivo.Close()

	return palavraEscolhida
}

//EscolherTipoDaPalavra escolhe um tipo de palavra, e depois escolhe a palavra
func EscolherTipoDaPalavra() (palavraEscolhida string, tipoEscolhido string) {
	tiposDePalavras := []string{"cores", "frutas", "paises"}
	tipoEscolhido = tiposDePalavras[rand.Intn(len(tiposDePalavras))]
	palavraEscolhida = escolherPalavra(tipoEscolhido)
	return
}
