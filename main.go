package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	db2 "github.com/raffreitas/fc-hex-go/adapters/db"
	"github.com/raffreitas/fc-hex-go/application"
)

func main() {
	db, _ := sql.Open("sqlite3", "db.sqlite")
	productDbAdapter := db2.NewProductDb(db)
	productService := application.NewProductService(productDbAdapter)
	product, _ := productService.Create("Example Product", 30)

	productService.Enable(product)
}
