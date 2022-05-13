package controllers

import "net/http"

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get all books"))
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get a book"))
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a book"))
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete a book"))
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update a book"))
}
