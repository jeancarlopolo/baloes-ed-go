// Package fila implementa uma fila circular estática.
package fila

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

// New cria uma nova fila com a capacidade informada.
func New(capacidade int) *Fila {
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