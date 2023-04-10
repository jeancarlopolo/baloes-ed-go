package arquivos

import (
	"fmt"

	"github.com/ajstarks/svgo"
	"github.com/jeancarlopolo/baloes-ed-go/formas"
)

func DesenharForma(f formas.Forma, c *svg.SVG) {
	switch forma := f.(type) {
	case *formas.Circulo:
		style := fmt.Sprintf("fill=%s, stroke=%s, transform=rotate(%d, %d, %d)", forma.CorFundo, forma.CorBorda, int(forma.Rotacao), int(forma.X), int(forma.Y))
		c.Circle(int(forma.X), int(forma.Y), int(forma.Raio), style)
	case *formas.Retangulo:
		style := fmt.Sprintf("fill=%s, stroke=%s, transform=rotate(%d, %d, %d)", forma.CorFundo, forma.CorBorda, int(forma.Rotacao), int(forma.X), int(forma.Y))
		c.Rect(int(forma.X), int(forma.Y), int(forma.Largura), int(forma.Altura), style)
	case *formas.Linha:
		style := fmt.Sprintf("stroke=%s, transform=rotate(%d, %d, %d)", forma.CorBorda, int(forma.Rotacao), int(forma.X1), int(forma.Y1))
		c.Line(int(forma.X1), int(forma.Y1), int(forma.X2), int(forma.Y2), style)
	case *formas.Texto:
		style := fmt.Sprintf("fill=%s, transform=rotate(%d, %d, %d), font-family=%s, font-size=%s, text-anchor=%s, font-weight=%s", forma.CorBorda, int(forma.Rotacao), int(forma.X), int(forma.Y), forma.Familia, forma.Tamanho, forma.Ancora, forma.Peso)
		c.Text(int(forma.X), int(forma.Y), forma.Texto, style)
	case *formas.Caca:
		style := fmt.Sprintf("fill=%s, transform=rotate(%d, %d, %d), font-family=%s, font-size=%s, text-anchor=%s, font-weight=%s", forma.CorBorda, int(forma.Rotacao), int(forma.X), int(forma.Y), forma.Familia, forma.Tamanho, forma.Ancora, forma.Peso)
		c.Text(int(forma.X), int(forma.Y), forma.Texto, style)
	case *formas.Balao:
		style := fmt.Sprintf("fill=%s, transform=rotate(%d, %d, %d), font-family=%s, font-size=%s, text-anchor=%s, font-weight=%s", forma.CorBorda, int(forma.Rotacao), int(forma.X), int(forma.Y), forma.Familia, forma.Tamanho, forma.Ancora, forma.Peso)
		c.Text(int(forma.X), int(forma.Y), forma.Texto, style)
	}
}

