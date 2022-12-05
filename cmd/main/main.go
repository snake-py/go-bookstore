package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/snake-py/go-bookstore/pkg/routes"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterBookStoreRoutes(router)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe("localhost:3000", router))
}
