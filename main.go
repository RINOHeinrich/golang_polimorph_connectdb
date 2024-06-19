package main

import (
	"fmt"

	"github.com/RINOHeinrich/golang_polimorph_connectdb/database"
	"github.com/RINOHeinrich/golang_polimorph_connectdb/models"
)

func main() {
	PostgreSQLConnector := database.PostgreSQLConnector{
		Host:     "localhost",
		Port:     5432,
		Username: "postgres",
		Password: "postgres",
	}
	var product models.Product
	DB, err := database.Connect(&PostgreSQLConnector)
	if err != nil {
		panic(err)
	}
	DB.Find(&product)
	fmt.Printf("%v", product)
}
