package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"golang-first-api-rest/models"
	"net/http"
	"os"
)

var (
	//variavel global para definir erros personalizados
	ErrCEPBadRequest = errors.New("CEP não encontrado")
)

func BuscaCep(endereco *models.Endereco, cep string) error {
	//a estrutura com err é o equivalente ao try catch no golang

	response, err := http.Get(os.Getenv("VIACEP_ENDPOINT") + cep + "/json/")
	if err != nil {
		return fmt.Errorf("Ocorreu um erro ao tentar enviar request para URL " +
			os.Getenv("VIACEP_ENDPOINT") + cep + "/json/ - Erro:" + err.Error())
	}

	if response.StatusCode == http.StatusBadRequest {
		return ErrCEPBadRequest
	}

	err = json.NewDecoder(response.Body).Decode(endereco)
	if err != nil {
		return fmt.Errorf("Ocorreu um erro ao tentar converter response.Body " +
			"para models.Endereco - Erro:" + err.Error())
	}

	return nil
}
