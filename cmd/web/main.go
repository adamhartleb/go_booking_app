package main

import (
	"log"
	"net/http"
	"time"

	"github.com/adamhartleb/go_booking_app/pkg/config"
	"github.com/adamhartleb/go_booking_app/pkg/handlers"
	"github.com/adamhartleb/go_booking_app/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const PORT = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Unable to create Template Cache")
	}

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.TemplateCache = tc
	app.UseCache = false
	app.InProduction = false
	app.Session = session

	repo := handlers.NewRepo(&app)
	render.NewTemplates(&app)
	handlers.NewHandlers(repo)

	mux := Routes(&app)

	http.ListenAndServe(PORT, mux)
}
