package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rachit1707/go-bookstoore-mgmt/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
