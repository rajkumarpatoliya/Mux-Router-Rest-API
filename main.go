package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

func main() {
	// Init Router

	r := mux.NewRouter()

	// Router Handlers / Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")

}
