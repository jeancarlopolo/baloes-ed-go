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
func LerQry(arquivoQry *os.File, arquivoTxt *os.File, db *estruturas.Lista, svgStruct *svg.SVG) {
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
				forma := db.Obter(i)
				switch forma.(type) {
				case formas.Retangulo:
					id = forma.(formas.Retangulo).Id
					if id == idProcurado {
						fmt.Fprintf(arquivoTxt, forma + " movido para (%.2f, %.2f)", x, y)
						formas.Mover(forma, x, y)
					}
					break 
				case formas.Circulo:
					id = forma.(formas.Circulo).Id
					if id == idProcurado {
						fmt.Fprintf(arquivoTxt, forma + " movido para (%.2f, %.2f)", x, y)
						formas.Mover(forma, x, y)
					}
					break
				case formas.Linha:
					id = forma.(formas.Linha).Id
					if id == idProcurado {
						fmt.Fprintf(arquivoTxt, forma + " movida para (%.2f, %.2f)", x, y)
						formas.Mover(forma, x, y)
					}
					break
				case formas.Texto:
					id = forma.(formas.Texto).Id
					if id == idProcurado {
						fmt.Fprintf(arquivoTxt, forma + " movido para (%.2f, %.2f)", x, y)
						formas.Mover(forma, x, y)
					}
					break
				case formas.Caca:
					id = forma.(formas.Caca).Id
					if id == idProcurado {
						fmt.Fprintf(arquivoTxt, forma + " movido para (%.2f, %.2f)", x, y)
						formas.Mover(forma, x, y)
					}
					break
				case formas.Balao:
					id = forma.(formas.Balao).Id
					if id == idProcurado {
						fmt.Fprintf(arquivoTxt, forma + " movido para (%.2f, %.2f)", x, y)
						formas.Mover(forma, x, y)
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
					forma := db.Obter(i)
					switch forma.(type) {
					case formas.Retangulo:
						id = forma.(formas.Retangulo).Id
						if id == idProcurado {
							fmt.Fprintf(arquivoTxt, forma + " rotacionado em %.2f graus", graus)
							formas.Girar(forma, graus)
						}
						break 
					case formas.Circulo:
						id = forma.(formas.Circulo).Id
						if id == idProcurado {
							fmt.Fprintf(arquivoTxt, forma + " rotacionado em %.2f graus", graus)
							formas.Girar(forma, graus)
						}
						break
					case formas.Linha:
						id = forma.(formas.Linha).Id
						if id == idProcurado {
							fmt.Fprintf(arquivoTxt, forma + " rotacionada em %.2f graus", graus)
							formas.Girar(forma, graus)
						}
						break
					case formas.Texto:
						id = forma.(formas.Texto).Id
						if id == idProcurado {
							fmt.Fprintf(arquivoTxt, forma + " rotacionado em %.2f graus", graus)
							formas.Girar(forma, graus)
						}
						break
					case formas.Caca:
						id = forma.(formas.Caca).Id
						if id == idProcurado {
							fmt.Fprintf(arquivoTxt, forma + " rotacionado em %.2f graus", graus)
							formas.Girar(forma, graus)
						}
						break
					case formas.Balao:
						id = forma.(formas.Balao).Id
						if id == idProcurado {
							fmt.Fprintf(arquivoTxt, forma + " rotacionado em %.2f graus", graus)
							formas.Girar(forma, graus)
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
							id = forma.(formas.Balao).Id
							if id == idProcurado {
								fmt.Fprintf(arquivoTxt, forma + " alterado para r: %.2f, p: %.2f, h: %.2f", r, p, h)
								balao.RaioCamera = r
								balao.ProfundidadeCamera = p
								balao.DistanciaCamera = h
							}
							break
						default:
							fmt.Println("Erro: tipo de forma desconhecido")
						}
						}

			}
			

		}
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
		}
	}

}
