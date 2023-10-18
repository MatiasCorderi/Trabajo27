package main

import (
	"testing"
)

func TestValidar(t *testing.T) {
	// Caso de prueba válido
	if !validar("bob", "50135875") {
		t.Error("Se esperaba que el usuario y la clave sean válidos")
	}

	// Caso de prueba inválido
	if validar("alice", "12345678") {
		t.Error("Se esperaba que el usuario y la clave sean inválidos")
	}

	// Caso de prueba con espacios en blanco
	if validar("eve ", "25436928") {
		t.Error("Se esperaba que el usuario y la clave sean inválidos debido a espacios en blanco")
	}
}
