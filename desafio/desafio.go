// traduzir esse algoritmo em go: https://github.com/DeveloperArthur/algoritmos-guias-anotacoes-uteis/blob/main/listaUnificada.java
// colocar cliente em um arquivo separado
// alterar a lista com ponteiros

package main

import (
	"fmt"
	"golang-first-api-rest/models"
	"strings"
)

func main() {
	var clientes = []models.Cliente{
		{"12345", "Sofia", 20.0},
		{"54321", "Sami", 40.0},
		{"12345", "Blade", 20.0},
		{"12345", "Sofia", 60.0},
	}

	fmt.Println(clientes)
	unificaLista(&clientes)
	mostraClientes(&clientes)
}

func unificaLista(clientes *[]models.Cliente) {
	listaUnificada := make(map[string]models.Cliente)

	for _, cliente := range *clientes {
		chave := strings.ToLower(cliente.Rg) + strings.ToLower(cliente.Nome)
		valorExistente, possuiChave := listaUnificada[chave]
		if possuiChave {
			var valorNovo = valorExistente.ValorParaPagar + cliente.ValorParaPagar
			valorExistente.ValorParaPagar = valorNovo
			delete(listaUnificada, chave)
			listaUnificada[chave] = valorExistente
		} else {
			listaUnificada[chave] = cliente
		}
	}

	//estamos atribuindo uma lista vazia a variavel clientes
	//pois estamos substituindo o valor que o ponteiro "clientes" guarda
	//*clientes faz a desreferencia e acessa o valor q esta no endereço de memoria
	*clientes = []models.Cliente{}

	//cada item da listaUnificada é adicionado na lista de clientes
	for _, cliente := range listaUnificada {
		*clientes = append(*clientes, cliente)
	}
}

func mostraClientes(clientes *[]models.Cliente) {
	//Desreferencia o ponteiro para obter a lista real
	//Quando usamos ponteiro com * na frente, pegamos o valor da variavel
	//Quando usamos ponteiro sem * na frente, pegamos o endereço de memória
	for _, cliente := range *clientes {
		fmt.Println(cliente.Rg, " ", cliente.Nome, " ", cliente.ValorParaPagar)
	}
}
