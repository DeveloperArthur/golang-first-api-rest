package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang-first-api-rest/database"
	"golang-first-api-rest/models"
	"net/http"
)

func Saldacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	fmt.Println(nome)
	c.JSON(200, gin.H{
		"APi diz:": "E ai " + nome + " tudo beleza?",
	})
}

func ExibeTodosAlunos(c *gin.Context) {
	var alunos []models.Aluno
	//a funcao nao tem retorno
	//vc passa o endereco de memoria
	//ela faz oq tem q fazer e seta na memoria
	database.DB.Find(&alunos)
	c.JSON(200, alunos)
}

func CriaNovoAluno(c *gin.Context) {
	var aluno models.Aluno

	error := c.ShouldBindJSON(&aluno)
	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": error.Error()})
		return
	}

	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}

func BuscaAlunoPorId(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"info": "Aluno não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, aluno)
}

func DeletaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")

	database.DB.First(&aluno, id)
	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"info": "Aluno não existe"})
		return
	} else {
		database.DB.Delete(&aluno, id)
		c.JSON(http.StatusNotFound, gin.H{"info": "Aluno de id " + id + " deletado com sucesso"})
		return
	}
}

func EditaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)

	if aluno.ID != 0 {
		error := c.ShouldBindJSON(&aluno)
		if error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
			return
		}

		database.DB.Model(&aluno).UpdateColumns(aluno)
		c.JSON(http.StatusOK, aluno)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"info": "Aluno não existe"})
		return
	}
}

func BuscaAlunoPorCpf(c *gin.Context) {
	var aluno models.Aluno
	cpf := c.Param("cpf")
	database.DB.Where(&models.Aluno{CPF: cpf}).First(&aluno)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"info": "Aluno não existe"})
		return
	}

	c.JSON(http.StatusOK, aluno)
}
