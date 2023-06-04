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
	/*clientes é um slice literal
	Um slice literal em Go é uma forma de criar uma fatia
	diretamente sem fazer referência a uma lista subjacente.
	Ele permite especificar os elementos da fatia entre chaves {}.*/
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

	/* "clientes" é um slice, pois é uma estrutura
	de dados que contém uma referência a uma lista subjacente*/

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

	//estamos atribuindo um slice vazio (uma lista vazia) a variavel clientes
	//pois estamos substituindo o valor que o ponteiro "clientes" guarda
	//*clientes faz a desreferencia e acessa o valor q esta no endereço de memoria
	*clientes = []models.Cliente{}

	//cada item da listaUnificada é adicionado no slice de clientes
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
