package main

import (
	"Devini0101/loja/routes"
	"net/http"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil) //sobe o server
}
