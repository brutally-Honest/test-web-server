package routes

import (
	"fmt"
	"net/http"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", rootHandler)
	return mux
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Welcome!\nLaugh at my misery")
}
