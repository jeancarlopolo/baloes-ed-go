package formas

import "math"

// Formas que podem estar dentro uma outra forma
type Interceptavel interface {
	InterceptaRetangulo(x float64, y float64, largura float64, altura float64) bool
	InterceptaCirculo(x float64, y float64, raio float64) bool
}

func (t Texto) InterceptaRetangulo(x float64, y float64, largura float64, altura float64) bool {
	return x <= t.X && t.X <= x+largura && y <= t.Y && t.Y <= y+altura
}

func (t Texto) InterceptaCirculo(x float64, y float64, raio float64) bool {
	return math.Sqrt(math.Pow(t.X-x, 2)+math.Pow(t.Y-y, 2)) <= raio
}

func (c Circulo) InterceptaRetangulo(x float64, y float64, largura float64, altura float64) bool {
	// Calcula a distância do centro do círculo para o lado mais próximo do retângulo
	testX := c.X
	testY := c.Y
	switch {
	case c.X < x:
		testX = x
	case c.X > x+largura:
		testX = x + largura
	}
	switch {
	case c.Y < y:
		testY = y
	case c.Y > y+altura:
		testY = y + altura
	}
	distX := c.X - testX
	distY := c.Y - testY
	distancia := math.Sqrt((distX * distX) + (distY * distY))
	return distancia <= c.Raio
}

func (c Circulo) InterceptaCirculo(x float64, y float64, raio float64) bool {
	return math.Sqrt(math.Pow(c.X-x, 2)+math.Pow(c.Y-y, 2)) <= c.Raio+raio
}

func (l Linha) interceptaLinha(x1 float64, y1 float64, x2 float64, y2 float64) bool {
	// Calcula a distância entre a linha e o ponto
	distancia := math.Abs((y2-y1)*l.X-(x2-x1)*l.Y+x2*y1-y2*x1) / math.Sqrt(math.Pow(y2-y1, 2)+math.Pow(x2-x1, 2))
	return distancia <= 1
}

func (l Linha) InterceptaRetangulo(x float64, y float64, largura float64, altura float64) bool {
	// Verifica se a linha intercepta algum dos lados do retângulo
	return l.interceptaLinha(x, y, x+largura, y) ||
		l.interceptaLinha(x+largura, y, x+largura, y+altura) ||
		l.interceptaLinha(x+largura, y+altura, x, y+altura) ||
		l.interceptaLinha(x, y+altura, x, y) ||
		l.X >= x && l.X <= x+largura && l.Y >= y && l.Y <= y+altura ||
		l.X2 >= x && l.X2 <= x+largura && l.Y2 >= y && l.Y2 <= y+altura
}

func (l Linha) InterceptaCirculo(x float64, y float64, raio float64) bool {
	if l.X >= x-raio && l.X <= x+raio && l.Y >= y-raio && l.Y <= y+raio ||
		l.X2 >= x-raio && l.X2 <= x+raio && l.Y2 >= y-raio && l.Y2 <= y+raio {
		return true
	}
	dot := ((x-l.X)*(l.X2-l.X) + (y-l.Y)*(l.Y2-l.Y)) / math.Pow(l.Area(), 2)
	closestX := l.X + dot*(l.X2-l.X)
	closestY := l.Y + dot*(l.Y2-l.Y)
	// Verifica se o ponto mais próximo faz parte da linha
	// Se a distância do ponto 1 e do ponto mais próximo + a distância do ponto 2 e do ponto mais próximo for igual a distância entre os pontos 1 e 2, então o ponto mais próximo faz parte da linha
	if math.Sqrt(math.Pow(l.X-closestX, 2)+math.Pow(l.Y-closestY, 2))+math.Sqrt(math.Pow(l.X2-closestX, 2)+math.Pow(l.Y2-closestY, 2))-l.Area() <= 0.1 {
		return math.Sqrt(math.Pow(closestX-x, 2)+math.Pow(closestY-y, 2)) <= raio
	}
	return false
}

func (r Retangulo) InterceptaRetangulo(x float64, y float64, largura float64, altura float64) bool {
	return r.X <= x+largura && r.X+r.Largura >= x && r.Y <= y+altura && r.Y+r.Altura >= y
}

func (r Retangulo) InterceptaCirculo(x float64, y float64, raio float64) bool {
	// Calcula a distância do centro do círculo para o lado mais próximo do retângulo
	testX := x
	testY := y
	switch {
	case x < r.X:
		testX = r.X
	case x > r.X+r.Largura:
		testX = r.X + r.Largura
	}
	switch {
	case y < r.Y:
		testY = r.Y
	case y > r.Y+r.Altura:
		testY = r.Y + r.Altura
	}
	distX := x - testX
	distY := y - testY
	distancia := math.Sqrt((distX * distX) + (distY * distY))
	return distancia <= raio
}
