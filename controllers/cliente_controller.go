package controllers

import (
	"github.com/gin-gonic/gin"
	"golang-first-api-rest/models"
	"golang-first-api-rest/service"
	"net/http"
)

func UnificaListaDeClientes(c *gin.Context) {
	var clientes []models.Cliente

	error := c.ShouldBindJSON(&clientes)
	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": error.Error()})
		return
	}

	service.UnificaListaDeClientes(&clientes)
	service.LogaClientesUnificados(&clientes)
	c.JSON(http.StatusOK, clientes)
}
