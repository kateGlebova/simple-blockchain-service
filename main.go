package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"fmt"
	"strconv"
	"encoding/json"
)

type data_input struct {
	Data string
}

var data = NewData(5)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/add_data", AddData).Methods("POST")
	router.HandleFunc("/last_blocks/{n}", LastBlocks).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}


func AddData(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var input data_input
	err := decoder.Decode(&input)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	data.AddData(input.Data)
}

func LastBlocks(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	n, err := strconv.Atoi(params["n"])
	if err != nil {
		http.NotFound(w, r)
	}
	jsonBlocks, _ := json.MarshalIndent(data.GetNLastBlocks(n), "", "	")
	fmt.Fprintf(w, string(jsonBlocks))
}
