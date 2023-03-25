package main

import (
	"fmt"

	"github.com/DedeMarantes/rest-api-go/database"
	"github.com/DedeMarantes/rest-api-go/routes"
)

func main() {
	fmt.Println("Iniciando Servidor http://localhost:8500 ")
	database.ConectaBancoDados()
	routes.HandleRequest()
}
