package main

import (
	"encoding/json"
	"log"
	"net/http"

	"cepservice/providers"

	"github.com/gorilla/mux"
)

type errorResponse struct {
	Message string
}

func findPostalCode(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postalCode := params["id"]

	result, err := providers.FindPostalCode(postalCode)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Postal code not found")
	} else {
		json.NewEncoder(w).Encode(result)
	}

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/{id}", findPostalCode).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}
