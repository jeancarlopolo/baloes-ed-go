// Implementa uma lista duplamente encadeada.
package estruturas

import (
	"fmt"
)

// Elemento representa um elemento da lista.
type Elemento struct {
	Valor interface{}
	Ant   *Elemento
	Prox  *Elemento
}

// Lista representa uma lista duplamente encadeada.
type Lista struct {
	Inicio *Elemento
	Fim    *Elemento
	Capacidade int
}

// Cheia retorna se a lista está cheia.
func (l *Lista) Cheia() bool {
	return l.Capacidade == l.Tamanho()
}

// Tamanho retorna o tamanho da lista.
func (l *Lista) Tamanho() int {
	tamanho := 0
	for e := l.Inicio; e != nil; e = e.Prox {
		tamanho++
	}
	return tamanho
}

// Vazia retorna se a lista está vazia.
func (l *Lista) Vazia() bool {
	return l.Inicio == nil
}

// Buscar retorna o elemento que contém o valor informado.
func (l *Lista) Buscar(valor interface{}) *Elemento {
	for e := l.Inicio; e != nil; e = e.Prox {
		if e.Valor == valor {
			return e
		}
	}
	return nil
}

// Inserir insere um novo elemento no final da lista.
func (l *Lista) Inserir(valor interface{}) {
	if l.Inicio == nil {
		l.Inicio = &Elemento{Valor: valor}
		l.Fim = l.Inicio
	} else {
		l.Fim.Prox = &Elemento{Valor: valor, Ant: l.Fim}
		l.Fim = l.Fim.Prox
	}
}

// Remover remove o elemento da lista.
func (l *Lista) Remover(valor interface{}) {
	if l.Inicio == nil {
		return
	}

	if l.Inicio.Valor == valor {
		l.Inicio = l.Inicio.Prox
		if l.Inicio != nil {
			l.Inicio.Ant = nil
		}
		return
	}

	for e := l.Inicio.Prox; e != nil; e = e.Prox {
		if e.Valor == valor {
			if e.Prox != nil {
				e.Prox.Ant = e.Ant
			}
			e.Ant.Prox = e.Prox
			return
		}
	}
}

// String retorna uma string com os elementos da lista.
func (l *Lista) String() string {
	s := ""
	for e := l.Inicio; e != nil; e = e.Prox {
		s += fmt.Sprintf("%v \n", e.Valor)
	}
	return s
}

// New cria uma nova lista.
func New(capacidade int) *Lista {
	return &Lista{Capacidade: capacidade}
}