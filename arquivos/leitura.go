package arquivos

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ajstarks/svgo"
	"github.com/jeancarlopolo/baloes-ed-go/comandos"
	"github.com/jeancarlopolo/baloes-ed-go/estruturas"
	"github.com/jeancarlopolo/baloes-ed-go/formas"
	"golang.org/x/text/cases"
)

// LerGeo lê o arquivo .geo e adiciona as formas na lista
func LerGeo(arquivoGeo *os.File, db *estruturas.Lista, doneGeo chan bool) {
	defer close(doneGeo)
	scanner := bufio.NewScanner(arquivoGeo)
	for scanner.Scan() {
		linha := scanner.Text()
		if linha == "" {
			continue
		}
		vetorPalavras := strings.Split(linha, " ")
		if vetorPalavras[0] == "ts" {
			// se for ts, atualiza as variáveis globais
			formas.TextoFamilia = vetorPalavras[1]
			formas.TextoPeso = vetorPalavras[2]
			formas.TextoTamanho = vetorPalavras[3]
			continue
		}
		db.Inserir(formas.CriaForma(linha))
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	doneGeo <- true
}

// LerQry lê o arquivo .qry e executa as queries
func LerQry(arquivoQry *os.File, arquivoTxt *os.File, db *estruturas.Lista, svgStruct *svg.SVG, nomeSvg string) {
	scanner := bufio.NewScanner(arquivoQry)
	for scanner.Scan() {
		linha := scanner.Text()
		if linha == "" {
			continue
		}
		vetorPalavras := strings.Split(linha, " ")
		switch vetorPalavras[0] {
		case "mv":
			// mv id x y
			idProcurado, _ := strconv.Atoi(vetorPalavras[1])
			x, _ := strconv.ParseFloat(vetorPalavras[2], 64)
			y, _ := strconv.ParseFloat(vetorPalavras[3], 64)
			var id int
			// procura a forma na lista
			for i := 0; i < db.Tamanho(); i++ {
				formaInterface := db.Obter(i)
				switch forma := formaInterface.(type) {
				case formas.Retangulo:
					id = forma.Id
					if id == idProcurado {
						fmt.Fprintf(arquivoTxt, forma.String()+" movido em (%.2f, %.2f)", x, y)
						forma.X += x
						forma.Y += y
					}
					break
				case formas.Circulo:
					id = forma.Id
					if id == idProcurado {
						fmt.Fprintf(arquivoTxt, forma.String()+" movido em (%.2f, %.2f)", x, y)
						forma.X += x
						forma.Y += y
					}
					break
				case formas.Linha:
					id = forma.Id
					if id == idProcurado {
						fmt.Fprintf(arquivoTxt, forma.String()+" movida em (%.2f, %.2f)", x, y)
						forma.X1 += x
						forma.Y1 += y
						forma.X2 += x
						forma.Y2 += y
					}
					break
				case formas.Texto:
					id = forma.Id
					if id == idProcurado {
						fmt.Fprintf(arquivoTxt, forma.String()+" movido em (%.2f, %.2f)", x, y)
						forma.X += x
						forma.Y += y
					}
					break
				case formas.Caca:
					id = forma.Id
					if id == idProcurado {
						fmt.Fprintf(arquivoTxt, forma.String()+" movido em (%.2f, %.2f)", x, y)
						forma.X += x
						forma.Y += y
					}
					break
				case formas.Balao:
					id = forma.Id
					if id == idProcurado {
						fmt.Fprintf(arquivoTxt, forma.String()+" movido em (%.2f, %.2f)", x, y)
						forma.X += x
						forma.Y += y
					}
					break
				default:
					fmt.Println("Erro: tipo de forma desconhecido")
				}
			}
		case "g":
			// g i graus
			idProcurado, _ := strconv.Atoi(vetorPalavras[1])
			graus, _ := strconv.ParseFloat(vetorPalavras[2], 64)
			var id int
			// procura a forma na lista
			for i := 0; i < db.Tamanho(); i++ {
				formaInterface := db.Obter(i)
				switch forma := formaInterface.(type) {
				case formas.Retangulo:
					id = forma.Id
					if id == idProcurado {
						fmt.Fprintf(arquivoTxt, forma.String()+" rotacionado em %.2f graus", graus)
						forma.Rotacao += graus
					}
					break
				case formas.Circulo:
					id = forma.Id
					if id == idProcurado {
						fmt.Fprintf(arquivoTxt, forma.String()+" rotacionado em %.2f graus", graus)
						forma.Rotacao += graus
					}
					break
				case formas.Linha:
					id = forma.Id
					if id == idProcurado {
						fmt.Fprintf(arquivoTxt, forma.String()+" rotacionada em %.2f graus", graus)
						forma.Rotacao += graus
					}
					break
				case formas.Texto:
					id = forma.Id
					if id == idProcurado {
						fmt.Fprintf(arquivoTxt, forma.String()+" rotacionado em %.2f graus", graus)
						forma.Rotacao += graus
					}
					break
				case formas.Caca:
					id = forma.Id
					if id == idProcurado {
						fmt.Fprintf(arquivoTxt, forma.String()+" rotacionado em %.2f graus", graus)
						forma.Rotacao += graus
					}
					break
				case formas.Balao:
					id = forma.Id
					if id == idProcurado {
						fmt.Fprintf(arquivoTxt, forma.String()+" rotacionado em %.2f graus", graus)
						forma.Rotacao += graus
					}
					break
				default:
					fmt.Println("Erro: tipo de forma desconhecido")
				}
			}
		case "ff":
			// i r p h
			// usado apenas para o balao
			idProcurado, _ := strconv.Atoi(vetorPalavras[1])
			r, _ := strconv.ParseFloat(vetorPalavras[2], 64)
			p, _ := strconv.ParseFloat(vetorPalavras[3], 64)
			h, _ := strconv.ParseFloat(vetorPalavras[4], 64)
			var id int
			// procura a forma na lista
			for i := 0; i < db.Tamanho(); i++ {
				forma := db.Obter(i)
				switch balao := forma.(type) {
				case formas.Balao:
					id = balao.Id
					if id == idProcurado {
						fmt.Fprintf(arquivoTxt, balao.String()+" alterado para r: %.2f, p: %.2f, h: %.2f", r, p, h)
						balao.RaioCamera = r
						balao.ProfundidadeCamera = p
						balao.DistanciaCamera = h
					}
					break
				default:
					fmt.Println("Erro: tipo de forma desconhecido")
				}
			}
		case "tf":
			// tf i f
			// balão i tira foto e coloca na fila f
			idProcurado, _ := strconv.Atoi(vetorPalavras[1])
			f, _ := strconv.Atoi(vetorPalavras[2])
			var id int
			// procura a forma na lista
			for i := 0; i < db.Tamanho(); i++ {
				forma := db.Obter(i)
				switch balao := forma.(type) {
				case formas.Balao:
					id = balao.Id
					if id == idProcurado {
						fmt.Fprintf(arquivoTxt, balao.String()+" tirou foto e colocou na fila %d", f)
						foto := balao.TirarFoto()
						balao.Filas[f].Inserir(foto)
					}
					break
				default:
					fmt.Println("Erro: tipo de forma desconhecido")
				}
			}
		case "df":
			// df i f sfx
			// gera um svg com o nome do arquivo svg + sfx + .svg
			// svg contém as fotos da fila f do balão i
			idProcurado, _ := strconv.Atoi(vetorPalavras[1])
			f, _ := strconv.Atoi(vetorPalavras[2])
			sfx := vetorPalavras[3]
			var id int
			// procura a forma na lista
			for i := 0; i < db.Tamanho(); i++ {
				forma := db.Obter(i)
				switch balao := forma.(type) {
				case formas.Balao:
					id = balao.Id
					if id == idProcurado {
						fmt.Fprintf(arquivoTxt, balao.String()+" gerou svg com as fotos da fila %d", f)
						svgFoto := balao.GerarSvgFotos(f)
						nomeArquivo := nomeArquivoSvg + sfx + ".svg"
						arquivo, err := os.Create(nomeArquivo)
						if err != nil {
							fmt.Println("Erro ao criar arquivo")
						}
						defer arquivo.Close()
						fmt.Fprintf(arquivo, svgFoto)
					}
					break
				default:
					fmt.Println("Erro: tipo de forma desconhecido")
				}
			}
		case "d":
			// d i capac dist j dx
			// caça i dispara bomba de capacidade capac, alcance dist
			// se atingir um balão, clona todos os elementos de todas as fotos que não foram imprimidas usando o comando df
			// o id deles começa em j, as cores de borda e de fundo são invertidas e eles são transladados em dx
			
		}

	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
