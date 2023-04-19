package formas

import (
	"fmt"
	"github.com/ajstarks/svgo"
)

// Desenhavel é uma interface que define um método Desenhar
type Desenhavel interface {
	Desenhar(*svg.SVG)
}

func (c Circulo) Desenhar(canvas *svg.SVG) {
	canvas.Circle(int(c.X), int(c.Y), int(c.Raio), fmt.Sprintf("fill:%s;stroke:%s;transform:rotate(%f,%f,%f)", c.CorFundo, c.CorBorda, c.Rotacao, c.X, c.Y))
}

func (r Retangulo) Desenhar(canvas *svg.SVG) {
	canvas.Rect(int(r.X), int(r.Y), int(r.Largura), int(r.Altura), fmt.Sprintf("fill:%s;stroke:%s;transform:rotate(%f,%f,%f)", r.CorFundo, r.CorBorda, r.Rotacao, r.X, r.Y))
}

func (t Texto) Desenhar(canvas *svg.SVG) {
	canvas.Text(int(t.X), int(t.Y), t.Texto, fmt.Sprintf("font-size:%spx;fill:%s;stroke:%s;transform:rotate(%f,%f,%f);text-anchor:%s;font-family:%s;font-weight:%s", t.Tamanho, t.CorFundo, t.CorBorda, t.Rotacao, t.X, t.Y, t.Ancora, t.Familia, t.Peso))
}

func (l Linha) Desenhar(canvas *svg.SVG) {
	canvas.Line(int(l.X), int(l.Y), int(l.X2), int(l.Y2), fmt.Sprintf("stroke:%s;transform:rotate(%f,%f,%f)", l.CorBorda, l.Rotacao, l.X, l.Y))
}