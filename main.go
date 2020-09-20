package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"cepservice/providers"

	"github.com/gorilla/mux"
)

func findPostalCode(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postalCode, err := validatePostalCode(params["postalCode"])

	if err != nil {
		errorResponse(w, err, 422)
		return
	}

	result, err := providers.FindPostalCode(postalCode)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		errorResponse(w, err, 404)
		return
	}

	json.NewEncoder(w).Encode(result)

}

func errorResponse(w http.ResponseWriter, err error, statusCode int) {
	http.Error(w, err.Error(), statusCode)
}

func validatePostalCode(postalCode string) (string, error) {
	if len(postalCode) != 8 {
		return "", errors.New("Postal code must have 8 characters")
	}

	_, err := strconv.ParseInt(postalCode, 10, 64)
	if err != nil {
		return "", errors.New("Postal code must contain only numbers")
	}

	return postalCode, nil
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/{postalCode}", findPostalCode).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}
