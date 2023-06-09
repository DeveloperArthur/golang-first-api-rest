package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang-first-api-rest/models"
	"log"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	stringDeConexao := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open("postgres", stringDeConexao)
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}

	//para criar uma tabela no banco de dados com
	//base na struct Aluno, utilizando gorm
	DB.AutoMigrate(&models.Aluno{})
}
