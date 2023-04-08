package main

import (
	"database/sql"

	db2 "github.com/Lukiita/go-hexagonal/src/adapters/db"
	"github.com/Lukiita/go-hexagonal/src/application"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "db.sqlite")
	productDb := db2.NewProductDb(db)
	productService := application.NewProductService(productDb)
	product, _ := productService.Create("Product Example", 30)

	productService.Enable(product)
}
