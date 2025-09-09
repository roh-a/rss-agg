package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {

	if code > 499 {
		log.Printf("Error caught of type 5xx %v", msg)
	}

	type errorResponse struct {
		Error string `json:"error"`
	}

	repondWithJSON(w, code, errorResponse{
		Error : msg,
	})

}

func repondWithJSON(w http.ResponseWriter, code int, payload any) {
	data, err := json.Marshal(payload)

	if err != nil {
		log.Printf("Failed to marshal JSON response : %v", payload)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}


