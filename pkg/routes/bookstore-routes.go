package routes

import (
	"example.com/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

// defining a function to initialise the API routes
var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book/{book_id}", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{book_id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{book_id}", controllers.DeleteBook).Methods("DELETE")
}
