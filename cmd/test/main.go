package main

import (
	"fmt"
	"net/http"

	"github.com/brutally-Honest/test/internal/routes"
)

func main() {
	router := routes.NewRouter()

	port := 7171
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Server running on http://localhost:%d\n", port)
	err := http.ListenAndServe(addr, router)
	if err != nil {
		panic(err)
	}
}
