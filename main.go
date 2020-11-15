package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"titFle"`
	Author *Author `json:"author"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var books []Book

//Get All books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

//Get Single Book

func main() {
	r := mux.NewRouter()

	books = append(books, Book{
		ID:    "1",
		Isbn:  "448743",
		Title: "BookOne",
		Author: &Author{
			Firstname: "John",
			Lastname:  "Doe",
		}})

	//Route Handler
	r.HandleFunc("/api/books", getBooks).Methods("GET")

	fmt.Printf("Listening 8888")

	log.Fatal(http.ListenAndServe(":8888", r))
}
