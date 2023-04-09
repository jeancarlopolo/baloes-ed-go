package estruturas

import "fmt"

type Lista struct {
	Primeiro   *Elemento
	Ultimo     *Elemento
	Capacidade int
}

type Elemento struct {
	Valor interface{}
	Prox  *Elemento
	Ant   *Elemento
}

func NovaLista(n int) *Lista {
	return &Lista{Capacidade: n}
}

func (l *Lista) Vazia() bool {
	return l.Primeiro == nil
}

func (l *Lista) Tamanho() int {
	tamanho := 0
	for e := l.Primeiro; e != nil; e = e.Prox {
		tamanho++
	}
	return tamanho
}

func (l *Lista) Cheia() bool {
	return l.Capacidade == l.Tamanho()
}

func (l *Lista) Inserir(e interface{}) {
	if l.Cheia() {
		panic("lista cheia")
	}
	novo := &Elemento{Valor: e}
	if l.Vazia() {
		l.Primeiro = novo
		l.Ultimo = novo
	} else {
		l.Ultimo.Prox = novo
		novo.Ant = l.Ultimo
		l.Ultimo = novo
	}
}

func (l *Lista) Remover() interface{} {
	if l.Vazia() {
		panic("lista vazia")
	}
	e := l.Ultimo.Valor
	l.Ultimo = l.Ultimo.Ant
	if l.Ultimo == nil {
		l.Primeiro = nil
	} else {
		l.Ultimo.Prox = nil
	}
	return e
}

func (l *Lista) Aplicar(f func(*interface{}, ...interface{}), args ...interface{}) { 
	for e := l.Primeiro; e != nil; e = e.Prox {
		f(&e.Valor, args...) 
	}
}

func (l *Lista) Recriar(f func(interface{}, ...interface{}) interface{}, args ...interface{}) *Lista {
	nova := NovaLista(l.Capacidade)
	for e := l.Primeiro; e != nil; e = e.Prox {
		novo := f(e.Valor, args...)
		nova.Inserir(novo)
	}
	return nova
}

func (l *Lista) Filtrar(f func(interface{}, ...interface{}) bool, args ...interface{}) *Lista {
	nova := NovaLista(l.Capacidade)
	for e := l.Primeiro; e != nil; e = e.Prox {
		if f(e.Valor, args...) {
			nova.Inserir(e.Valor)
		}
	}
	return nova
}

func (l *Lista) String() string {
	s := "["
	for e := l.Primeiro; e != nil; e = e.Prox {
		s += fmt.Sprintf("%v ", e.Valor)
	}
	return s + "]"
}

func (l *Lista) Deletar() {
	l.Primeiro = nil
	l.Ultimo = nil
}