package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Saldacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	fmt.Println(nome)
	c.JSON(200, gin.H{
		"APi diz:": "E ai " + nome + " tudo beleza?",
	})
}
