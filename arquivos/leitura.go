package arquivos

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	"github.com/jeancarlopolo/baloes-ed-go/estruturas"
	"github.com/jeancarlopolo/baloes-ed-go/formas"
)

func LerGeo(arquivoGeo io.Reader, db **estruturas.Lista, doneGeo *chan bool) {
	// lê o arquivo .geo linha por linha
	scanner := bufio.NewScanner(arquivoGeo)
	var peso, familia, tamanho string
	for scanner.Scan() {
		linha := scanner.Text()
		// separa a linha em palavras
		palavras := strings.Split(linha, " ")
		// verifica se a linha é válida
		if len(palavras) <= 0 {
			continue
		}
		// verifica o tipo da forma
		switch palavras[0] {
		case "c":
			// cria um círculo
			id, _ := strconv.Atoi(palavras[1])
			x, _ := strconv.ParseFloat(palavras[2], 64)
			y, _ := strconv.ParseFloat(palavras[3], 64)
			r, _ := strconv.ParseFloat(palavras[4], 64)
			corBorda := palavras[5]
			corFundo := palavras[6]
			circulo := formas.Circulo{
				Forma: formas.Forma{
					Id:       id,
					X:        x,
					Y:        y,
					CorBorda: corBorda,
					CorFundo: corFundo,
					Rotacao:  0},
				Raio: r}
			(*db).Inserir(circulo)
		case "r":
			// cria um retângulo
			id, _ := strconv.Atoi(palavras[1])
			x, _ := strconv.ParseFloat(palavras[2], 64)
			y, _ := strconv.ParseFloat(palavras[3], 64)
			l, _ := strconv.ParseFloat(palavras[4], 64)
			a, _ := strconv.ParseFloat(palavras[5], 64)
			corBorda := palavras[6]
			corFundo := palavras[7]

			retangulo := formas.Retangulo{
				Forma: formas.Forma{
					Id:       id,
					X:        x,
					Y:        y,
					CorBorda: corBorda,
					CorFundo: corFundo,
					Rotacao:  0},
				Largura: l,
				Altura:  a}
			(*db).Inserir(retangulo)
		case "t":
			// cria um texto
			id, _ := strconv.Atoi(palavras[1])
			x, _ := strconv.ParseFloat(palavras[2], 64)
			y, _ := strconv.ParseFloat(palavras[3], 64)
			corBorda := palavras[4]
			corFundo := palavras[5]
			ancora := palavras[6]
			// junta as palavras restantes em uma só
			texto := ""
			for i := 7; i < len(palavras); i++ {
				texto += palavras[i] + " "
			}
			// remove o espaço no final
			texto = texto[:len(texto)-1]
			textoForma := formas.Texto{
				Forma: formas.Forma{
					Id:       id,
					X:        x,
					Y:        y,
					CorBorda: corBorda,
					CorFundo: corFundo,
					Rotacao:  0},
				Texto:   texto,
				Ancora:  ancora,
				Familia: familia,
				Tamanho: tamanho,
				Peso:    peso}
			(*db).Inserir(textoForma)
		case "l":
			// cria uma linha
			id, _ := strconv.Atoi(palavras[1])
			x, _ := strconv.ParseFloat(palavras[2], 64)
			y, _ := strconv.ParseFloat(palavras[3], 64)
			x2, _ := strconv.ParseFloat(palavras[4], 64)
			y2, _ := strconv.ParseFloat(palavras[5], 64)
			corBorda := palavras[6]
			linhaForma := formas.Linha{
				Forma: formas.Forma{
					Id:       id,
					X:        x,
					Y:        y,
					CorBorda: corBorda,
					Rotacao:  0},
				X2: x2,
				Y2: y2}
			(*db).Inserir(linhaForma)
		}
	}
	*doneGeo <- true
}

//arquivos.LerQry(arquivoQry, arquivoTxt, db, svgStruct, nomeArquivoSvg)

//func LerQry(qry)