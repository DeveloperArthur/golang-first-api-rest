package controllers

import (
	"github.com/gin-gonic/gin"
	"golang-first-api-rest/models"
	"golang-first-api-rest/service"
	"net/http"
)

func BuscaCep(c *gin.Context) {
	var endereco models.Endereco
	cep := c.Query("cep") //RequestParam

	err := service.BuscaCep(&endereco, cep)
	if err != nil {
		if err == service.ErrCEPBadRequest {
			c.JSON(http.StatusBadRequest, err.Error())
		} else {
			c.JSON(http.StatusInternalServerError, err.Error())
		}
	} else {
		c.JSON(200, endereco)
	}
}
