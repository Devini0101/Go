package controllers

import (
	"Devini0101/loja/models"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

var temps = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := models.FindAllProducts()
	temps.ExecuteTemplate(w, "index", products) //quem consegue passar a resposta, quem vai mostrar e qual dados mandar para lá
}

func New(w http.ResponseWriter, r *http.Request) {
	temps.ExecuteTemplate(w, "new", nil) //quem consegue passar a resposta, quem vai mostrar e qual dados mandar para lá
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("nome")
		description := r.FormValue("descricao")
		price := r.FormValue("preco")
		amount := r.FormValue("quantidade")

		formattedPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Erro na conversão de preço vindo do front", err)
		}

		formattedAmount, err := strconv.Atoi(amount)

		if err != nil {
			log.Println("Erro na conversão de quantidade")
		}

		models.CreateNewProduct(name, description, formattedPrice, formattedAmount)
	}
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	product := models.SelectProductById(productId)
	temps.ExecuteTemplate(w, "edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("nome")
		description := r.FormValue("descricao")
		price := r.FormValue("preco")
		amount := r.FormValue("quantidade")

		formattedPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Erro na conversão de preço vindo do front", err)
		}

		formattedAmount, err := strconv.Atoi(amount)

		if err != nil {
			log.Println("Erro na conversão de quantidade")
		}

		models.Update(name, description, formattedPrice, formattedAmount, id)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	models.Delete(productId)
	http.Redirect(w, r, "/", 301)
}
