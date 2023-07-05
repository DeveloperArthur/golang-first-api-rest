package service

import (
	"fmt"
	"golang-first-api-rest/models"
	"strconv"
	"sync"
	"time"
)

func CalculaFolhaDePagamento(folhaDePagamento *models.FolhaPagamento, idFuncionario int) {
	//Utilização de sync.WaitGroup necessária para garantir
	//que a regra de negócio só seja executada após todos
	//os requests sejam finalizados
	var wg sync.WaitGroup

	var funcionario models.Funcionario
	var horasExtras models.HorasExtras
	var beneficios models.Beneficios

	wg.Add(3)

	go func() {
		defer wg.Done()
		BuscaDadosDeFuncionario(&funcionario, idFuncionario)
	}()

	go func() {
		defer wg.Done()
		BuscaDadosDeHorasExtras(&horasExtras, idFuncionario)
	}()

	go func() {
		defer wg.Done()
		BuscaDadosDeBenecifios(&beneficios, idFuncionario)
	}()

	wg.Wait()

	fmt.Println("[LOG] Processando lógica de negócio ficticia - " + time.Now().String())
	totalParaReceber := funcionario.SalarioBase
	if funcionario.SalarioBase < 2.000 {
		totalParaReceber += 300
	} else if beneficios.DiasFerias > 15 {
		totalParaReceber += (horasExtras.TaxaHorasExtras * float32(horasExtras.HorasExtras))
	} else if beneficios.DiasFerias < 15 {
		totalParaReceber += (horasExtras.TaxaHorasExtras * float32(horasExtras.HorasExtras)) + 300
	}

	fmt.Printf("[LOG] Total a Receber: %f - %s\n", totalParaReceber, time.Now().String())
	folhaDePagamento.IdFuncionario = idFuncionario
	folhaDePagamento.TotalParaReceber = totalParaReceber
}

func BuscaDadosDeFuncionario(funcionario *models.Funcionario, idFuncionario int) {
	fmt.Println("[LOG] Finding funcionario in URL https://rh.7894-2376-5423.com/funcionarios/9876/info?id=" +
		strconv.Itoa(idFuncionario) + " - " + time.Now().String())
	time.Sleep(200 * time.Millisecond)
	funcionario.Id = idFuncionario
	funcionario.Nome = "Iago Filipe Fernandes"
	funcionario.Cargo = "Analista Pleno"
	funcionario.SalarioBase = 3.000
	fmt.Println("[LOG] Request finding funcionario finished - " + time.Now().String())
}

func BuscaDadosDeBenecifios(beneficios *models.Beneficios, idFuncionario int) {
	fmt.Println("[LOG] Finding beneficios in URL https://benefits.gcp.appengine.com/funcionario/" +
		strconv.Itoa(idFuncionario) + " - " + time.Now().String())
	time.Sleep(200 * time.Millisecond)
	beneficios.IdFuncionario = idFuncionario
	beneficios.DiasFerias = 30
	beneficios.PlanoSaude = "Plano Bronze"
	fmt.Println("[LOG] Request finding beneficios finished - " + time.Now().String())
}

func BuscaDadosDeHorasExtras(horasExtras *models.HorasExtras, idFuncionario int) {
	fmt.Println("[LOG] Finding horas extras in URL https://company.s3.server.com/horasextras?key=92847&emplooyeId=" +
		strconv.Itoa(idFuncionario) + " - " + time.Now().String())
	time.Sleep(200 * time.Millisecond)
	horasExtras.IdFuncinario = idFuncionario
	horasExtras.TaxaHorasExtras = 76.5
	horasExtras.HorasExtras = 12
	fmt.Println("[LOG] Request finding horas extras finished - " + time.Now().String())
}
