package main

import (
	httphandler "Golang-Rest-API/product/https/product"
	"fmt"
	"net/http"

	middleware "Golang-Rest-API/product/middleware"
	services "Golang-Rest-API/product/services/product"
	stores "Golang-Rest-API/product/stores/product"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	// Data source name (DSN) format: user:password@tcp(host:port)/dbname
	dsn := "root:root@tcp(127.0.0.1:3306)/testdb"

	// Open the connection to the database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	pStore := stores.NewStore(db)
	ps := services.NewProduct(pStore)
	httpClient := httphandler.Handler{
		ProductService: ps,
	}

	fmt.Println("Server starting at 8081....")
	r := mux.NewRouter().StrictSlash(true)
	r.Use(middleware.AuthenticationMiddleware)
	r.HandleFunc("/products", httpClient.Get).Methods("GET")
	r.HandleFunc("/products/{id}", httpClient.GetProduct).Methods("GET")
	r.HandleFunc("/products/{id}", httpClient.DeleteProduct).Methods("DELETE")
	r.HandleFunc("/products/{id}", httpClient.UpdateProduct).Methods("PUT")
	r.HandleFunc("/products", httpClient.CreateProduct).Methods("POST")
	http.ListenAndServe(":8081", r)
}
