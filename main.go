package main

import (
	"github.com/jimmmisss/go-api-gin-gorm/database"
	"github.com/jimmmisss/go-api-gin-gorm/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
