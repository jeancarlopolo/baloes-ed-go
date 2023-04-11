// Package fila implementa uma fila circular estática de fotos.
package fila

import (
	"fmt"
	"github.com/jeancarlopolo/lista"
)

// Fila é uma fila circular estática de listas duplamente encadeadas.
type Fila struct {
	Inicio *lista.Lista
	Fim    *lista.Lista
}