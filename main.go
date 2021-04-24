package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	db2 "github.com/rfaguiar/ports-adapter-pattern/adapters/db"
	"github.com/rfaguiar/ports-adapter-pattern/application"
)

func main() {
	db, _ := sql.Open("sqlite3", "sqlite.db")
	productDbAdapter := db2.NewProductDb(db)
	productService := application.NewProductService(productDbAdapter)
	product, _ := productService.Create("Product Exempt", 30)

	productService.Enable(product)

}
