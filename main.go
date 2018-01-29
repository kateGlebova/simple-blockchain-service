package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"fmt"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", Hello).Methods("GET")
	router.HandleFunc("/add_data", AddData).Methods("POST")
	router.HandleFunc("/last_blocks/{N}", LastBlocks).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func AddData(w http.ResponseWriter, r *http.Request) {}

func LastBlocks(w http.ResponseWriter, r *http.Request) {}