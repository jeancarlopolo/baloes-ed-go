// Pacote formas contém tipos que definem formas geométricas e suas propriedades.
package formas

import (
	"math"
)


type Forma interface {
	Area() float64
	Encosta(f Forma) bool
	String() string
}

func Mover(f Forma, dx, dy float64) {
	switch t := f.(type) {
	case *Retangulo:
		t.X += dx
		t.Y += dy
	case *Circulo:
		t.X += dx
		t.Y += dy
	case *Linha:
		t.X1 += dx
		t.Y1 += dy
		t.X2 += dx
		t.Y2 += dy
	case *Texto:
		t.X += dx
		t.Y += dy
	case *Caca:
		t.X += dx
		t.Y += dy
	case *Balao:
		t.X += dx
		t.Y += dy
	}
}

func (r Retangulo) Area() float64 {
	return r.Altura * r.Largura
}

func (c Circulo) Area() float64 {
	return 3.14 * c.Raio * c.Raio
}

func (l Linha) Area() float64 {
	return math.Sqrt(math.Pow(l.X2-l.X1, 2) + math.Pow(l.Y2-l.Y1, 2))
}

func (t Texto) Area() float64 {
	return 0
}

func (c Caca) Area() float64 {
	return 0
}

func (b Balao) Area() float64 {
	return 0
}

func pontoDentroRetangulo(x, y, x1, y1, x2, y2 float64) bool {
	return x >= x1 && x <= x2 && y >= y1 && y <= y2
}

func pontoDentroCirculo(x, y, cx, cy, raio float64) bool {
	return math.Sqrt(math.Pow(x-cx, 2)+math.Pow(y-cy, 2)) <= raio
}

func linhaEncostaLinha(x1, y1, x2, y2, x3, y3, x4, y4 float64) bool {
	// https://en.wikipedia.org/wiki/Line%E2%80%93line_intersection
	den := (x1-x2)*(y3-y4) - (y1-y2)*(x3-x4)
	if den == 0 {
		return false
	}
	t := ((x1-x3)*(y3-y4) - (y1-y3)*(x3-x4)) / den
	u := -((x1-x2)*(y1-y3) - (y1-y2)*(x1-x3)) / den
	return t >= 0 && t <= 1 && u >= 0 && u <= 1
}

func (r Retangulo) Encosta(f Forma) bool {
	encosta := false
	switch t := f.(type) {
	case Retangulo:
		encosta = linhaEncostaLinha(r.X, r.Y, r.X+r.Largura, r.Y, t.X, t.Y, t.X+t.Largura, t.Y)                                                  // linha superior
		encosta = encosta || linhaEncostaLinha(r.X, r.Y+r.Altura, r.X+r.Largura, r.Y+r.Altura, t.X, t.Y+t.Altura, t.X+t.Largura, t.Y+t.Altura)   // linha inferior
		encosta = encosta || linhaEncostaLinha(r.X, r.Y, r.X, r.Y+r.Altura, t.X, t.Y, t.X, t.Y+t.Altura)                                         // linha esquerda
		encosta = encosta || linhaEncostaLinha(r.X+r.Largura, r.Y, r.X+r.Largura, r.Y+r.Altura, t.X+t.Largura, t.Y, t.X+t.Largura, t.Y+t.Altura) // linha direita
		if !encosta {
			encosta = pontoDentroRetangulo(t.X, t.Y, r.X, r.Y, r.X+r.Largura, r.Y+r.Altura)            // nenhuma linha encosta, portanto os 4 ou nenhum ponto está dentro
			encosta = encosta || pontoDentroRetangulo(r.X, r.Y, t.X, t.Y, t.X+t.Largura, t.Y+t.Altura) // nenhuma linha encosta, portanto os 4 ou nenhum ponto está dentro
		}
	case Circulo:
		// Verifica se algum dos pontos do retângulo está dentro do círculo
		encosta = pontoDentroCirculo(r.X, r.Y, t.X, t.Y, t.Raio)
		encosta = encosta || pontoDentroCirculo(r.X+r.Largura, r.Y, t.X, t.Y, t.Raio)
		encosta = encosta || pontoDentroCirculo(r.X, r.Y+r.Altura, t.X, t.Y, t.Raio)
		encosta = encosta || pontoDentroCirculo(r.X+r.Largura, r.Y+r.Altura, t.X, t.Y, t.Raio)
		// Verifica se algum dos lados do retângulo intersecta o círculo
		encosta = encosta || linhaEncostaLinha(r.X, r.Y, r.X+r.Largura, r.Y, t.X, t.Y, t.X+t.Raio, t.Y)                    // linha superior
		encosta = encosta || linhaEncostaLinha(r.X, r.Y+r.Altura, r.X+r.Largura, r.Y+r.Altura, t.X, t.Y, t.X+t.Raio, t.Y)  // linha inferior
		encosta = encosta || linhaEncostaLinha(r.X, r.Y, r.X, r.Y+r.Altura, t.X, t.Y, t.X, t.Y+t.Raio)                     // linha esquerda
		encosta = encosta || linhaEncostaLinha(r.X+r.Largura, r.Y, r.X+r.Largura, r.Y+r.Altura, t.X, t.Y, t.X, t.Y+t.Raio) // linha direita
	case Linha:
		encosta = linhaEncostaLinha(r.X, r.Y, r.X+r.Largura, r.Y, t.X1, t.Y1, t.X2, t.Y2)                               // linha superior
		encosta = encosta || linhaEncostaLinha(r.X, r.Y+r.Altura, r.X+r.Largura, r.Y+r.Altura, t.X1, t.Y1, t.X2, t.Y2)  // linha inferior
		encosta = encosta || linhaEncostaLinha(r.X, r.Y, r.X, r.Y+r.Altura, t.X1, t.Y1, t.X2, t.Y2)                     // linha esquerda
		encosta = encosta || linhaEncostaLinha(r.X+r.Largura, r.Y, r.X+r.Largura, r.Y+r.Altura, t.X1, t.Y1, t.X2, t.Y2) // linha direita
		encosta = encosta || pontoDentroRetangulo(t.X1, t.Y1, r.X, r.Y, r.X+r.Largura, r.Y+r.Altura)                    // nenhuma linha encosta, portanto os 2 ou nenhum ponto está dentro
	}
	return encosta
}

func (c Circulo) Encosta(f Forma) bool {
	encosta := false
	switch t := f.(type) {
	case Retangulo:
		// Verifica se algum dos pontos do círculo está dentro do retângulo
		encosta = pontoDentroRetangulo(c.X, c.Y, t.X, t.Y, t.X+t.Largura, t.Y+t.Altura)
		encosta = encosta || pontoDentroRetangulo(c.X+c.Raio, c.Y, t.X, t.Y, t.X+t.Largura, t.Y+t.Altura)
		encosta = encosta || pontoDentroRetangulo(c.X, c.Y+c.Raio, t.X, t.Y, t.X+t.Largura, t.Y+t.Altura)
		encosta = encosta || pontoDentroRetangulo(c.X+c.Raio, c.Y+c.Raio, t.X, t.Y, t.X+t.Largura, t.Y+t.Altura)
		// Verifica se algum dos lados do retângulo intersecta o círculo
		encosta = encosta || linhaEncostaLinha(t.X, t.Y, t.X+t.Largura, t.Y, c.X, c.Y, c.X+c.Raio, c.Y)                    // linha superior
		encosta = encosta || linhaEncostaLinha(t.X, t.Y+t.Altura, t.X+t.Largura, t.Y+t.Altura, c.X, c.Y, c.X+c.Raio, c.Y)  // linha inferior
		encosta = encosta || linhaEncostaLinha(t.X, t.Y, t.X, t.Y+t.Altura, c.X, c.Y, c.X, c.Y+c.Raio)                     // linha esquerda
		encosta = encosta || linhaEncostaLinha(t.X+t.Largura, t.Y, t.X+t.Largura, t.Y+t.Altura, c.X, c.Y, c.X, c.Y+c.Raio) // linha direita
	case Circulo:
		encosta = math.Sqrt(math.Pow(c.X-t.X, 2)+math.Pow(c.Y-t.Y, 2)) <= c.Raio+t.Raio
	case Linha:
		dx := t.X2 - t.X1
		dy := t.Y2 - t.Y1
		a := dx*dx + dy*dy
		b := 2 * (dx*(t.X1-c.X) + dy*(t.Y1-c.Y))
		c := c.X*c.X + c.Y*c.Y + t.X1*t.X1 + t.Y1*t.Y1 - 2*(c.X*t.X1+c.Y*t.Y1) - c.Raio*c.Raio
		d := b*b - 4*a*c
		if d >= 0 {
			raiz1 := (-b + math.Sqrt(d)) / (2 * a)
			raiz2 := (-b - math.Sqrt(d)) / (2 * a)
			encosta = (raiz1 >= 0 && raiz1 <= 1) || (raiz2 >= 0 && raiz2 <= 1)
		}
	}
	return encosta
}

func (l Linha) Encosta(f Forma) bool {
	encosta := false
	switch t := f.(type) {
	case Retangulo:
		encosta = linhaEncostaLinha(t.X, t.Y, t.X+t.Largura, t.Y, l.X1, l.Y1, l.X2, l.Y2)                               // linha superior
		encosta = encosta || linhaEncostaLinha(t.X, t.Y+t.Altura, t.X+t.Largura, t.Y+t.Altura, l.X1, l.Y1, l.X2, l.Y2)  // linha inferior
		encosta = encosta || linhaEncostaLinha(t.X, t.Y, t.X, t.Y+t.Altura, l.X1, l.Y1, l.X2, l.Y2)                     // linha esquerda
		encosta = encosta || linhaEncostaLinha(t.X+t.Largura, t.Y, t.X+t.Largura, t.Y+t.Altura, l.X1, l.Y1, l.X2, l.Y2) // linha direita
		encosta = encosta || pontoDentroRetangulo(l.X1, l.Y1, t.X, t.Y, t.X+t.Largura, t.Y+t.Altura)                    // nenhuma linha encosta, portanto os 2 ou nenhum ponto está dentro
		encosta = encosta || pontoDentroRetangulo(l.X2, l.Y2, t.X, t.Y, t.X+t.Largura, t.Y+t.Altura)
	case Circulo:
		dx := l.X2 - l.X1
		dy := l.Y2 - l.Y1
		a := dx*dx + dy*dy
		b := 2 * (dx*(l.X1-t.X) + dy*(l.Y1-t.Y))
		c := t.X*t.X + t.Y*t.Y + l.X1*l.X1 + l.Y1*l.Y1 - 2*(t.X*l.X1+t.Y*l.Y1) - t.Raio*t.Raio
		d := b*b - 4*a*c
		if d >= 0 {
			raiz1 := (-b + math.Sqrt(d)) / (2 * a)
			raiz2 := (-b - math.Sqrt(d)) / (2 * a)
			encosta = (raiz1 >= 0 && raiz1 <= 1) || (raiz2 >= 0 && raiz2 <= 1)
		}
	case Linha:
		encosta = linhaEncostaLinha(t.X1, t.Y1, t.X2, t.Y2, l.X1, l.Y1, l.X2, l.Y2)
	}
	return encosta
}

func (t Texto) Encosta(f Forma) bool {
	encosta := false
	switch c := f.(type) {
	case Retangulo:
		encosta = pontoDentroRetangulo(t.X, t.Y, c.X, c.Y, c.X+c.Largura, c.Y+c.Altura)
	case Circulo:
		encosta = math.Sqrt(math.Pow(t.X-c.X, 2)+math.Pow(t.Y-c.Y, 2)) <= c.Raio
	}
	return encosta
}

func (ca Caca) Encosta(f Forma) bool {
	encosta := false
	switch c := f.(type) {
	case Retangulo:
		encosta = pontoDentroRetangulo(ca.X, ca.Y, c.X, c.Y, c.X+c.Largura, c.Y+c.Altura)
	case Circulo:
		encosta = math.Sqrt(math.Pow(ca.X-c.X, 2)+math.Pow(ca.Y-c.Y, 2)) <= c.Raio
	}
	return encosta
}

func (b Balao) Encosta(f Forma) bool {
	encosta := false
	switch c := f.(type) {
	case Retangulo:
		encosta = pontoDentroRetangulo(b.X, b.Y, c.X, c.Y, c.X+c.Largura, c.Y+c.Altura)
	case Circulo:
		encosta = math.Sqrt(math.Pow(b.X-c.X, 2)+math.Pow(b.Y-c.Y, 2)) <= c.Raio
	}
	return encosta
}