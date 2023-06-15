package controllers

import (
	"github.com/gin-gonic/gin"
	"golang-first-api-rest/models"
	"golang-first-api-rest/service"
	"net/http"
)

func ExibeTodosAlunos(c *gin.Context) {
	var alunos []models.Aluno
	//a funcao nao tem retorno
	//vc passa o endereco de memoria
	//ela faz oq tem q fazer e seta na memoria
	service.ExibeTodosAlunos(&alunos)
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

	service.CriaNovoAluno(&aluno)
	c.JSON(http.StatusOK, aluno)
}

func BuscaAlunoPorId(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	service.BuscaAlunoPorId(&aluno, id)

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

	service.BuscaAlunoPorId(&aluno, id)
	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"info": "Aluno não existe"})
		return
	} else {
		service.DeletaAluno(&aluno, id)
		c.JSON(http.StatusNotFound, gin.H{"info": "Aluno de id " + id + " deletado com sucesso"})
		return
	}
}

func EditaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	service.BuscaAlunoPorId(&aluno, id)

	if aluno.ID != 0 {
		error := c.ShouldBindJSON(&aluno)
		if error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
			return
		}

		service.EditaAluno(&aluno)
		c.JSON(http.StatusOK, aluno)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"info": "Aluno não existe"})
		return
	}
}

func BuscaAlunoPorCpf(c *gin.Context) {
	var aluno models.Aluno
	cpf := c.Param("cpf")
	service.BuscaAlunoPorCpf(&aluno, cpf)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"info": "Aluno não existe"})
		return
	}

	c.JSON(http.StatusOK, aluno)
}
