package routes

import (
	"github.com/gorilla/mux"
	"github.com/rhaqim/gp-bookstore/pkg/controllers"
)

var registerBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/api/books", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", controllers.GetBook).Methods("GET")
	router.HandleFunc("/api/books", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", controllers.DeleteBook).Methods("DELETE")
}
