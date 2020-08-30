package main

import (
	"encoding/json"
	"log"
	"net/http"

	"cepservice/providers"

	"github.com/gorilla/mux"
)

func findPostalCode(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postalCode := params["id"]

	result := providers.FindPostalCode(postalCode)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/{id}", findPostalCode).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}
