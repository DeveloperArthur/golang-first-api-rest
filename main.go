package main

import (
	"golang-first-api-rest/database"
	"golang-first-api-rest/routes"
	"os"
)

func main() {
	os.Setenv("VIACEP_ENDPOINT", "https://viacep.com.br/ws/")
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
