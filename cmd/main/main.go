// purpose of main: launch the server and point Go (using router from gorilla/mux) to the routes

package main

import (
	"log"
	"net/http"

	"example.com/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	router := mux.NewRouter()

	routes.RegisterBookStoreRoutes(router)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe("localhost:9010", router))
}
