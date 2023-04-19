// Implementa uma fila circular estática.
package estruturas

import (
	"fmt"
)

// Fila é uma fila circular estática.
type Fila struct {
	Inicio     int
	Fim        int
	Capacidade int
	Valores    []interface{}
}

// NewFila cria uma nova fila com a capacidade informada.
func NewFila(capacidade int) *Fila {
	return &Fila{Capacidade: capacidade, Valores: make([]interface{}, capacidade)}
}

// String retorna uma string com os valores da fila.
func (f *Fila) String() string {
	s := ""
	for i := f.Inicio; i != f.Fim; i = (i + 1) % f.Capacidade {
		s += fmt.Sprintf("%v \n", f.Valores[i])
	}
	return s
}

// Inserir insere um valor na fila.
func (f *Fila) Inserir(valor interface{}) {
	if f.Cheia() {
		return
	}
	f.Valores[f.Fim] = valor
	f.Fim = (f.Fim + 1) % f.Capacidade
}

// Remover remove um valor da fila.
func (f *Fila) Remover() interface{} {
	if f.Vazia() {
		return nil
	}
	valor := f.Valores[f.Inicio]
	f.Inicio = (f.Inicio + 1) % f.Capacidade
	return valor
}

// Vazia retorna true se a fila estiver vazia.
func (f *Fila) Vazia() bool {
	return f.Inicio == f.Fim
}

// Cheia retorna true se a fila estiver cheia.
func (f *Fila) Cheia() bool {
	return (f.Fim+1)%f.Capacidade == f.Inicio
}