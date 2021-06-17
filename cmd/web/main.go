package main

import (
	"net/http"

	"github.com/adamhartleb/go_booking_app/pkg/handlers"
)

const PORT = ":8080"

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/about", handlers.About)

	http.ListenAndServe(PORT, mux)
}
