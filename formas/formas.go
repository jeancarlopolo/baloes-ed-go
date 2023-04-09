// Pacote formas contém tipos que definem formas geométricas e suas propriedades.
package formas

import (
	"fmt"
	"github.com/jeancarlopolo/baloes-ed-go/estruturas"
)

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

// String retorna uma representação em string de um retângulo.
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

// String retorna uma representação em string de um círculo.
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

// String retorna uma representação em string de uma linha.
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

// String retorna uma representação em string de um texto.
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

// Caças são textos identificados pelo texto "|-T-|" e que podem atirar.
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

// String retorna uma representação em string de uma caça.
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

// Caças são textos identificados pelo texto "v_O_v" e que podem tirar fotos.
type Balao struct {
	Id                 int
	X                  float64
	Y                  float64
	Texto              string
	Rotacao            float64
	CorBorda           string
	CorFundo           string
	Familia            string
	Tamanho            string
	Peso               string
	RaioCamera         float64
	ProfundidadeCamera float64
	DistanciaCamera    float64
	// 10 filas de 15 fotos
	Filas [9]estruturas.Fila
}

// String retorna uma representação em string de um balão.
func (b Balao) String() string {
	s := fmt.Sprintf("Balão %d: (%.2f, %.2f)"+
		"\n%s"+
		"\n%.2fº"+
		"\nBorda: %s"+
		"\nFundo: %s"+
		"\nFamília: %s"+
		"\nTamanho: %s"+
		"\nPeso: %s"+
		"\nRaio da Câmera: %.2f"+
		"\nProfundidade da Câmera: %.2f"+
		"\nDistância da Câmera: %.2f", b.Id, b.X, b.Y, b.Texto, b.Rotacao, b.CorBorda, b.CorFundo, b.Familia, b.Tamanho, b.Peso, b.RaioCamera, b.ProfundidadeCamera, b.DistanciaCamera)
	// anexa na string s as filas
	for i := 0; i < 10; i++ {
		s += fmt.Sprintf("\nFila %d: %v", i, b.Filas[i])
	}
	return s
}

// NovoBalao retorna um novo balão (é necessário para criar as filas).
func NovoBalao() Balao {
	b := Balao{}
	b.Filas = [9]estruturas.Fila{}
	for i := 0; i < 10; i++ {
		b.Filas[i] = *estruturas.NovaFila(15)
	}
	return b
}
