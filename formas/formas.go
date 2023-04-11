// Classes e interfaces de formas geométricas
package formas

import (
	"github.com/jeancarlopolo/baloes-ed-go/estruturas/fila"
	"math"
)

type Forma struct {
	Id       int
	X        float64
	Y        float64
	CorBorda string
	CorFundo string
	Rotacao  float64
}

type Retangulo struct {
	Forma
	Largura float64
	Altura  float64
}

type Circulo struct {
	Forma
	Raio float64
}

type Linha struct {
	Forma
	X2 float64
	Y2 float64
}

type Texto struct {
	Forma
	Texto   string
	Ancora  string
	Familia string
	Tamanho float64
	Peso    string
}

type Caca struct {
	Texto
	IdsAtingidos []int
	Disparos     int
}

type Balao struct {
	Texto
	// 10 filas circulares estáticas
	// Cada fila pode ter até 15 listas duplamente encadeadas de formas
	Filas              [9]*fila.Fila
	RaioCamera         float64
	ProfundidadeCamera float64
	AlturaCamera       float64
}

type Pontuavel interface {
	Area() float64
	Pontos() int
}

func (r Retangulo) Area() float64 {
	return r.Largura * r.Altura
}

func (r Retangulo) Pontos() int {
	pontuacao := r.Area() / 4
	switch r.CorBorda {
	case "#80080":
		pontuacao += 10
	case "#AA0088":
		pontuacao += 15
	case "#008033":
		pontuacao += 20
	case "#FFCC00":
		pontuacao += 30
	}
	return int(pontuacao)
}

func (c Circulo) Area() float64 {
	return math.Pi * math.Pow(c.Raio, 2)
}

func (c Circulo) Pontos() int {
	pontuacao := c.Area() / 2
	switch {
	case c.CorBorda == "#FFFFFF" && c.CorFundo == "#FFFF00":
		pontuacao *= 8
	case c.CorBorda == "#D45500" && c.CorFundo == "#FF7F2A":
		pontuacao *= 2
	case c.CorBorda == "#AA0000" && c.CorFundo == "#DE8787":
		pontuacao *= 4
	case c.CorBorda == "#FFFFFF" && c.CorFundo == "#B3B3B3":
		pontuacao = 0
	}
	return int(pontuacao)
}

func (l Linha) Area() float64 {
	// Calcula o comprimento da linha
	return math.Sqrt(math.Pow(l.X2-l.X, 2) + math.Pow(l.Y2-l.Y, 2))
}

func (l Linha) Pontos() int {
	pontuacao := l.Area()
	switch l.CorBorda {
	case "#FFFF00":
		pontuacao *= 3
	case "#DDFF55":
		pontuacao *= 2
	case "#0000FF":
		pontuacao *= 4
	}
	return int(pontuacao)
}

func (t Texto) Area() float64 {
	return float64(len(t.Texto))
}

func (t Texto) Pontos() int {
	pontuacao := t.Area()
	return int(pontuacao)
}

func (c Caca) Pontos() int {
	return 100
}
