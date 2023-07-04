package routes

import (
	"github.com/gin-gonic/gin"
	"golang-first-api-rest/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/clientes/unificados", controllers.UnificaListaDeClientes)
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.POST("/alunos", controllers.CriaNovoAluno)
	r.GET("/alunos/:id", controllers.BuscaAlunoPorId)
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	r.PUT("/alunos/:id", controllers.EditaAluno)
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCpf)
	//Path Param :nome
	r.GET("/:nome", controllers.Saldacao)
	r.GET("/busca-cep", controllers.BuscaCep)
	r.Run(":5000") //listen and serve on localhost:5000
}
