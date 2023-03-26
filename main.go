package main

import (
	"github.com/felipeazsantos/gin-api-rest/database"
	"github.com/felipeazsantos/gin-api-rest/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequest()
}
