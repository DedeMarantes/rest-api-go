package main

import (
	"fmt"

	"github.com/DedeMarantes/rest-api-go/database"
	"github.com/DedeMarantes/rest-api-go/models"
	"github.com/DedeMarantes/rest-api-go/routes"
)

func main() {
	fmt.Println("Iniciando Servidor http://localhost:8500 ")
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "historia1"},
		{Id: 2, Nome: "nome 2", Historia: "historia2"},
	}
	database.ConectaBancoDados()
	routes.HandleRequest()
}
