// Pacote fila implementa uma estrutura de dados do tipo fila circular estática.
package fila

import (
	"fmt"
)

// Fila é uma estrutura de dados do tipo fila circular estática.
type Fila struct {
	elementos []interface{} // Elementos da fila.
	inicio    int           // Índice do primeiro elemento da fila.
	fim       int           // Índice do último elemento da fila.
	tamanho   int           // Número de elementos da fila.
}

// NovaFila cria uma nova fila com capacidade máxima de n elementos.
func NovaFila(n int) *Fila {
	return &Fila{elementos: make([]interface{}, n)}
}

// Vazia retorna true se a fila não possui elementos.
func (f *Fila) Vazia() bool {
	return f.tamanho == 0
}

// Cheia retorna true se a fila não possui mais espaço para novos elementos.
func (f *Fila) Cheia() bool {
	return f.tamanho == len(f.elementos)
}

// Inserir insere um novo elemento na fila.
func (f *Fila) Inserir(e interface{}) {
	if f.Cheia() {
		panic("fila cheia")
	}
	f.elementos[f.fim] = e
	f.fim = (f.fim + 1) % len(f.elementos)
	f.tamanho++
}

// Remover remove o primeiro elemento da fila e retorna o seu valor.
func (f *Fila) Remover() interface{} {
	if f.Vazia() {
		panic("fila vazia")
	}
	e := f.elementos[f.inicio]
	f.inicio = (f.inicio + 1) % len(f.elementos)
	f.tamanho--
	return e
}

// String retorna uma representação em string da fila.
func (f *Fila) String() string {
	s := "["
	for i := 0; i < f.tamanho; i++ {
		s += fmt.Sprintf("%v ", f.elementos[(f.inicio+i)%len(f.elementos)])
	}
	return s + "]"
}

// Aplicar aplica a função f a cada elemento da fila.
func (f *Fila) Aplicar(f2 func(interface{}, ...interface{})) { 
	for i := 0; i < f.tamanho; i++ {
		f2(f.elementos[(f.inicio+i)%len(f.elementos)])
	}
}
