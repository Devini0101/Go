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

func SelectProductById(id string) Product {
	db := db.Conn()
	defer db.Close()

	row := db.QueryRow("SELECT id, nome, descricao, preco, quantidade FROM public.produtos WHERE id = $1", id)

	product := Product{}
	err := row.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Amount)

	if err != nil {
		panic(err.Error())
	}

	return product
}

func Update(name string, description string, price float64, amount int, id string) {
	db := db.Conn()
	defer db.Close()
	updateData, err := db.Prepare("UPDATE produtos SET nome = $1, descricao = $2, preco = $3, quantidade = $4 WHERE id = $5")
	if err != nil {
		panic(err.Error())
	}
	updateData.Exec(name, description, price, amount, id)
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
