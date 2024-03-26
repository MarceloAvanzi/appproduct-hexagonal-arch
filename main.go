package main

import (
	"database/sql"
	db2 "github.com/MarceloAvanzi/appproduct-hexagonal-arch/adapters/db"
	_ "github.com/mattn/go-sqlite3"
	"github.com/MarceloAvanzi/appproduct-hexagonal-arch/application"
)

func main() {
	db, _ := sql.Open("sqlite3", "sqlite.db")
	productDbAdapter := db2.NewProductDb(db)
	productService := NewProductService(productDbAdapter)

	product, _ := ProductService.Create("Product Examplo", 30)
	ProductService.Enable(product)
}