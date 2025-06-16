package main

import (
	"github.com/eswarlearn/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
	// "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookstoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r)) // Start the server on port 9010

}
