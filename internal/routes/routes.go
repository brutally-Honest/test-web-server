package routes

import (
	"fmt"
	"net/http"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/about", helpHandler)
	return mux
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Welcome!\nLaugh at my misery")
}

func helpHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("About :\n1. Space to cool down / rant\n2. No judgement\n")
}
