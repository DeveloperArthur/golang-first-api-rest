package service

import (
	"golang-first-api-rest/models"
	"testing"
)

func TestDeveRetornarClientesUnificados(t *testing.T) {
	//Arrange
	clientes := []models.Cliente{
		{"12345", "Sofia", 20.0},
		{"54321", "Sami", 40.0},
		{"12345", "Blade", 20.0},
		{"12345", "Sofia", 60.0},
	}

	//Action
	UnificaListaDeClientes(&clientes)

	//Assert
	for _, cliente := range clientes {
		if cliente.Nome == "Blade" {
			assertEqual(t, 20.0, cliente.ValorParaPagar)
		} else if cliente.Nome == "Sami" {
			assertEqual(t, 40.0, cliente.ValorParaPagar)
		} else if cliente.Nome == "Sofia" {
			assertEqual(t, 80.0, cliente.ValorParaPagar)
		}
	}
}

func assertEqual(t *testing.T, expected interface{}, actual interface{}) {
	if expected != actual {
		t.Errorf("Expected [%+v] Got: [%+v]", expected, actual)
	}
}
