package models

import (
	"Devini0101/loja/db"
)

type Product struct {
	Name, Description string
	Price             float64
	Id, Amount        int
}

func FindAllProducts() []Product {
	db := db.Conn()
	defer db.Close()
	selectProducts, err := db.Query("SELECT * FROM public.produtos")
	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for selectProducts.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		//scannear linha a linha os valor determinados

		err = selectProducts.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Name = nome
		p.Description = descricao
		p.Price = preco
		p.Amount = quantidade

		products = append(products, p)
	}

	return products

}

func CreateNewProduct(name string, description string, price float64, amount int) {
	db := db.Conn()
	defer db.Close()
	insertData, err := db.Prepare("INSERT INTO produtos (nome, descricao, preco, quantidade) VALUES ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insertData.Exec(name, description, price, amount)
}

func Delete(id string) {
	db := db.Conn()
	defer db.Close()
	deleteData, err := db.Prepare("DELETE FROM produtos WHERE id = $1")
	if err != nil {
		panic(err.Error())
	}
	deleteData.Exec(id)
}
