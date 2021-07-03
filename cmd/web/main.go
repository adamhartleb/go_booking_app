package main

import (
	"log"
	"net/http"

	"github.com/adamhartleb/go_booking_app/pkg/config"
	"github.com/adamhartleb/go_booking_app/pkg/handlers"
	"github.com/adamhartleb/go_booking_app/pkg/render"
)

const PORT = ":8080"

func main() {
	var app config.AppConfig

	mux := http.NewServeMux()

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Unable to create Template Cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	render.NewTemplates(&app)
	handlers.NewHandlers(repo)

	mux.HandleFunc("/", handlers.Repo.Home)
	mux.HandleFunc("/about", handlers.Repo.About)

	http.ListenAndServe(PORT, mux)
}
