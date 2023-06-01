package models

import "gorm.io/gorm"

type Aluno struct {
	// dessa forma o gorm cria uma tabela espelhando
	// a struct Aluno, e adicionando + campos
	// como id etc
	gorm.Model
	Nome string `json:"nome"`
	CPF  string `json:"cpf"`
	RG   string `json:"rg"`
}
