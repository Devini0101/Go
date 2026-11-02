package main

import (
	"api-go-rest/database"
	"api-go-rest/routes"
	"fmt"
)

func main() {
	fmt.Println("Initializing go rest server!")
	database.ConnectDatabase()
	routes.HandleRequest()
}
