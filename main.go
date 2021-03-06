package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"

	"cepservice/providers"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.elastic.co/apm/module/apmgorilla"
)

func findPostalCode(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postalCode, err := validatePostalCode(params["postalCode"])
	token := r.Header.Get("X-Google-Api-Key")

	if err != nil {
		errorResponse(w, err, 422)
		return
	}

	result, err := providers.FindPostalCode(postalCode, token)

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

func loggingHandler(h http.Handler) http.Handler {
	return handlers.LoggingHandler(os.Stdout, h)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/cep/{postalCode}", findPostalCode).Methods("GET")

	port := "8000"

	router.Use(loggingHandler)
	router.Use(handlers.CompressHandler)
	router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))
	router.Use(apmgorilla.Middleware())

	log.Println("listening on " + port)
	http.ListenAndServe(":"+port, router)
}
