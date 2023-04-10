package main

import (
	"fmt"
	"github.com/ajstarks/svgo"
	"github.com/jeancarlopolo/baloes-ed-go/arquivos"
	"github.com/jeancarlopolo/baloes-ed-go/estruturas"
	"github.com/jeancarlopolo/baloes-ed-go/formas"
	"os"
)

func main() {
	// -e [path de entrada] (opcional, default: diretório atual)
	// -o [path de saída] (obrigatório)
	// -f [nome do arquivo .geo de entrada] (obrigatório)
	// -q [nome do arquivo .qry de entrada] (opcional, default: não executar queries)

	pathEntrada := "./"
	pathSaida := ""
	nomeArquivoGeo := ""
	nomeArquivoQry := ""

	// lê os argumentos da linha de comando
	for i := 1; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "-e":
			pathEntrada = os.Args[i+1]
			i++
		case "-o":
			pathSaida = os.Args[i+1]
			i++
		case "-f":
			nomeArquivoGeo = os.Args[i+1]
			i++
		case "-q":
			nomeArquivoQry = os.Args[i+1]
			i++
		}
	}
	if pathSaida == "" || nomeArquivoGeo == "" {
		fmt.Println("Argumentos inválidos")
		return
	}

	db := estruturas.NovaLista(1000)
	doneGeo := make(chan bool)

	// arruma os paths
	if pathEntrada[len(pathEntrada)-1] != '/' {
		pathEntrada += "/"
	}
	if pathSaida[len(pathSaida)-1] != '/' {
		pathSaida += "/"
	}

	arquivoGeo, err := os.Open(pathEntrada + nomeArquivoGeo)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer arquivoGeo.Close()

	// lê o arquivo .geo
	go arquivos.LerGeo(arquivoGeo, db, doneGeo)

	nomeArquivoSvg := pathSaida + nomeArquivoGeo[:len(nomeArquivoGeo)-4] + nomeArquivoQry[:len(nomeArquivoQry)-4]
	arquivoSvg, err := os.Create(nomeArquivoSvg + ".svg")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer arquivoSvg.Close()

	// escreve o cabeçalho do svg
	svgStruct := svg.New(arquivoSvg)
	svgStruct.Start(1000, 1000)
	svgStruct.Title("Arquivo SVG")

	var arquivoQry *os.File
	var arquivoTxt *os.File

	// cria o arquivo .qry
	if nomeArquivoQry != "" {
		arquivoQry, err = os.Open(pathEntrada + nomeArquivoQry)
		if err != nil {
			fmt.Println(err)
			return
		}
		arquivoTxt, err = os.Create(pathSaida + nomeArquivoGeo[:len(nomeArquivoGeo)-4] + nomeArquivoQry[:len(nomeArquivoQry)-4] + ".txt")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer arquivoTxt.Close()
		defer arquivoQry.Close()
	}

	for {
		select {
		case <-doneGeo:
			if arquivoQry != nil {
				arquivos.LerQry(arquivoQry, arquivoTxt, db, svgStruct, nomeArquivoSvg)
			}
			for i := 0; i < db.Tamanho(); i++ {
				forma := db.Obter(i).(formas.Forma)
				arquivos.DesenharForma(forma, svgStruct)
			}
		default:
			if db.Tamanho() == 0 {
				svgStruct.End()
				return
			}
		}
	}
}
