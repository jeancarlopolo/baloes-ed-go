// Package lista implementa uma lista duplamente encadeada de forma genérica.
package lista

import (
	"fmt"

	"github.com/jeancarlopolo/formas"
)

// Lista é uma lista duplamente encadeada.
type Lista struct {
	Capacidade int
	Inicio     *Elemento
	Fim        *Elemento
}

// Elemento é um elemento da lista.
type Elemento struct {
	Valor formas.Forma
	Prox  *Elemento
	Ant   *Elemento
}

// New cria uma nova lista.
func New(capacidade int) *Lista {
	return &Lista{Capacidade: capacidade}
}

// Inserir insere um novo elemento na lista.
func (l *Lista) Inserir(f formas.Forma) {
	if l.Inicio == nil {
		l.Inicio = &Elemento{Valor: f}
		l.Fim = l.Inicio
	} else {
		l.Fim.Prox = &Elemento{Valor: f, Ant: l.Fim}
		l.Fim = l.Fim.Prox
	}
}

// Remover remove um elemento da lista.
func (l *Lista) Remover(f formas.Forma) {
	if l.Inicio == nil {
		return
	}

	if l.Inicio.Valor == f {
		l.Inicio = l.Inicio.Prox
		if l.Inicio != nil {
			l.Inicio.Ant = nil
		}
		return
	}

	for e := l.Inicio.Prox; e != nil; e = e.Prox {
		if e.Valor == f {
			if e.Prox != nil {
				e.Prox.Ant = e.Ant
			}
			e.Ant.Prox = e.Prox
			return
		}
	}
}

// Buscar busca um elemento na lista.
func (l *Lista) Buscar(f formas.Forma) *Elemento {
	for e := l.Inicio; e != nil; e = e.Prox {
		if e.Valor == f {
			return e
		}
	}
	return nil
}

// String retorna uma representação em string da lista.
func (l *Lista) String() string {
	s := ""
	for e := l.Inicio; e != nil; e = e.Prox {
		s += fmt.Sprintf("%v \n", e.Valor)
	}
	return s
}
