package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/about", aboutHandler)
	return mux
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Welcome!\nLaugh at my misery")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	resp := Response{Message: "About :\n1. Space to cool down / rant\n2. No judgement\n"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)

}
