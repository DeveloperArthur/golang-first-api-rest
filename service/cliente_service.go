package service

import (
	"fmt"
	"golang-first-api-rest/models"
	"strings"
)

func UnificaListaDeClientes(clientes *[]models.Cliente) {
	listaUnificada := make(map[string]models.Cliente)

	/* "clientes" é um slice, pois é uma estrutura
	de dados que contém uma referência a uma lista subjacente*/
	//conceito melhor explicado aqui: https://github.com/DeveloperArthur/introducao-golang/blob/main/src/slice/slice.go

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

func LogaClientesUnificados(clientes *[]models.Cliente) {
	//Desreferencia o ponteiro para obter a lista real
	//Quando usamos ponteiro com * na frente, pegamos o valor da variavel
	//Quando usamos ponteiro sem * na frente, pegamos o endereço de memória
	for _, cliente := range *clientes {
		fmt.Println(cliente.Rg, " ", cliente.Nome, " ", cliente.ValorParaPagar)
	}
}
