// Pacote formas contém tipos que definem formas geométricas e suas propriedades.
package formas

import (
	"fmt"
	"math"
	"github.com/jeancarlopolo/baloes-ed-go/src/estruturas"
)

// Forma é uma interface que define uma forma geométrica.
type Forma interface {
	Area() float64
	String() string
}

// InverteCores inverte as cores de borda e fundo de uma forma geométrica.
func InverteCores(f *Forma) {
	switch t := (*f).(type) {
	case *Retangulo:
		t.CorBorda, t.CorFundo = t.CorFundo, t.CorBorda
	case *Circulo:
		t.CorBorda, t.CorFundo = t.CorFundo, t.CorBorda
	case *Texto:
		t.CorBorda, t.CorFundo = t.CorFundo, t.CorBorda
	case *Caca:
		t.CorBorda, t.CorFundo = t.CorFundo, t.CorBorda
	}
}

// Retangulo é uma forma geométrica que representa um retângulo.
type Retangulo struct {
	Id       int
	X        float64
	Y        float64
	Altura   float64
	Largura  float64
	Rotacao  float64
	CorBorda string
	CorFundo string
}

func (r Retangulo) Area() float64 {
	return r.Altura * r.Largura
}

func (r Retangulo) String() string {
	return fmt.Sprintf("Retângulo %d: (%.2f, %.2f)"+
		"\n%.2f x %.2f"+
		"\n%.2fº"+
		"\nBorda: %s"+
		"\nFundo: %s", r.Id, r.X, r.Y, r.Altura, r.Largura, r.Rotacao, r.CorBorda, r.CorFundo)
}

// Circulo é uma forma geométrica que representa um círculo.
type Circulo struct {
	Id       int
	X        float64
	Y        float64
	Raio     float64
	Rotacao  float64
	CorBorda string
	CorFundo string
}

func (c Circulo) Area() float64 {
	return 3.14 * c.Raio * c.Raio
}

func (c Circulo) String() string {
	return fmt.Sprintf("Círculo %d: (%.2f, %.2f)"+
		"\nRaio: %.2f"+
		"\n%.2fº"+
		"\nBorda: %s"+
		"\nFundo: %s", c.Id, c.X, c.Y, c.Raio, c.Rotacao, c.CorBorda, c.CorFundo)
}

// Linha é uma forma geométrica que representa uma linha.
type Linha struct {
	Id       int
	X1       float64
	Y1       float64
	X2       float64
	Y2       float64
	Rotacao  float64
	CorBorda string
}

// Area é usada para calcular o comprimento da linha (não é uma área).
func (l Linha) Area() float64 {
	return math.Sqrt(math.Pow(l.X2-l.X1, 2) + math.Pow(l.Y2-l.Y1, 2))
}

func (l Linha) String() string {
	return fmt.Sprintf("Linha %d: (%.2f, %.2f) -> (%.2f, %.2f)"+
		"\n%.2fº"+
		"\nBorda: %s", l.Id, l.X1, l.Y1, l.X2, l.Y2, l.Rotacao, l.CorBorda)
}

// Texto é uma forma geométrica que representa um texto.
type Texto struct {
	Id       int
	X        float64
	Y        float64
	Texto    string
	Rotacao  float64
	CorBorda string
	CorFundo string
	Familia  string
	Tamanho  string
	Peso     string
}

func (t Texto) Area() float64 {
	return 0
}

func (t Texto) String() string {
	return fmt.Sprintf("Texto %d: (%.2f, %.2f)"+
		"\n%s"+
		"\n%.2fº"+
		"\nBorda: %s"+
		"\nFundo: %s"+
		"\nFamília: %s"+
		"\nTamanho: %s"+
		"\nPeso: %s", t.Id, t.X, t.Y, t.Texto, t.Rotacao, t.CorBorda, t.CorFundo, t.Familia, t.Tamanho, t.Peso)
}

type Caca struct {
	Id                int
	X                 float64
	Y                 float64
	Texto             string
	Rotacao           float64
	CorBorda          string
	CorFundo          string
	Familia           string
	Tamanho           string
	Peso              string
	DisparosEfetuados int
	IdsAlvosAtingidos []int
}

func (c Caca) Area() float64 {
	return 0
}

func (c Caca) String() string {
	return fmt.Sprintf("Caca %d: (%.2f, %.2f)"+
		"\n%s"+
		"\n%.2fº"+
		"\nBorda: %s"+
		"\nFundo: %s"+
		"\nFamília: %s"+
		"\nTamanho: %s"+
		"\nPeso: %s"+
		"\nDisparos Efetuados: %d"+
		"\nIds Atingidos: %v", c.Id, c.X, c.Y, c.Texto, c.Rotacao, c.CorBorda, c.CorFundo, c.Familia, c.Tamanho, c.Peso, c.DisparosEfetuados, c.IdsAlvosAtingidos)
}

type Balao struct {
	Id                int
	X                 float64
	Y                 float64
	Texto             string
	Rotacao           float64
	CorBorda          string
	CorFundo          string
	Familia           string
	Tamanho           string
	Peso              string
	RaioCamera 	  float64
	// 10 filas de 15 fotos
	// 150 fotos
	Filas [][]Foto
