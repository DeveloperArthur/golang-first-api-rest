package controllers

import (
	"github.com/gin-gonic/gin"
	"golang-first-api-rest/models"
	"golang-first-api-rest/service"
	"log"
	"strconv"
)

func CalculaFolhaDePagamento(c *gin.Context) {
	var folhaPagamento models.FolhaPagamento

	idFuncionarioString := c.Params.ByName("idFuncionario")
	idFuncionario, err := strconv.Atoi(idFuncionarioString)
	if err != nil {
		log.Fatal("Erro ao converter a string para int:", err)
	}

	service.CalculaFolhaDePagamento(&folhaPagamento, idFuncionario)
	c.JSON(200, folhaPagamento)
}
