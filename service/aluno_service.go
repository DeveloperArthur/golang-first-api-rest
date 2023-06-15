package service

import (
	"golang-first-api-rest/database"
	"golang-first-api-rest/models"
)

func ExibeTodosAlunos(alunos *[]models.Aluno) {
	database.DB.Find(alunos)
}

func CriaNovoAluno(aluno *models.Aluno) {
	database.DB.Create(aluno)
}

func BuscaAlunoPorId(aluno *models.Aluno, id string) {
	database.DB.First(aluno, id)
}

func DeletaAluno(aluno *models.Aluno, id string) {
	database.DB.Delete(aluno, id)
}

func EditaAluno(aluno *models.Aluno) {
	database.DB.Model(aluno).UpdateColumns(&aluno)
}

func BuscaAlunoPorCpf(aluno *models.Aluno, cpf string) {
	database.DB.Where(&models.Aluno{CPF: cpf}).First(aluno)
}
