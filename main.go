package main

import (
	"crud-go/config"
	homecontroller "crud-go/controllers/homecontroller"
	categorycontroller "crud-go/controllers/categorycontroller"
	"log"
	"net/http"
)

func main() {
	config.ConnectDb()

	//home page
	http.HandleFunc("/", homecontroller.Welcome)

	//category
	http.HandleFunc("/category", categorycontroller.Index)
	http.HandleFunc("/category/add", categorycontroller.Add)
	http.HandleFunc("/category/edit", categorycontroller.Edit)
	http.HandleFunc("/category/delete", categorycontroller.Delete)


	log.Println("Starting server on 8080")
	http.ListenAndServe(":8080", nil)
}