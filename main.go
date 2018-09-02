package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	//"math/rand"
	"net/http"
	//"strconv"
)

// Get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Get Single Book
func getBook(w http.ResponseWriter, r *http.Request) {

}

// Create Single Book
func createBook(w http.ResponseWriter, r *http.Request) {

}

// Update Single Book
func updateBook(w http.ResponseWriter, r *http.Request) {

}

// Delete Single Book
func deleteBook(w http.ResponseWriter, r *http.Request) {

}

// Book Struct (Model)
type Book struct {
	ID     string  `json:id`
	Isbn   string  `json:isbn`
	Title  string  `json:title`
	Author *Author `json:author`
}

// Author Struct
type Author struct {
	Firstname string `json:firstname`
	Lastname  string `json:lastname`
}

// Init Books vas as slice Book struct
var books []Book

func main() {
	// Init Router

	r := mux.NewRouter()

	// Mock Data -todo -Implement DB
	books = append(books, Book{
		ID:    "1",
		Isbn:  "3241843618958",
		Title: "Book Of Life",
		Author: &Author{
			Firstname: "Rabindranath",
			Lastname:  "Tagore",
		},
	})

	// Router Handlers / Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	if err := http.ListenAndServe(":4000", r); err != nil {
		log.Fatal(err)
	}

}
