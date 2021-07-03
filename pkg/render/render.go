package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/adamhartleb/go_booking_app/pkg/config"
	"github.com/adamhartleb/go_booking_app/pkg/models"
)

var functions = template.FuncMap{}
var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Template does not exist.")
	}

	td = AddDefaultData(td)

	err := t.Execute(w, td)
	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	templateCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return templateCache, err
	}

	layouts, err := filepath.Glob("./templates/*.layout.html")
	if err != nil {
		return templateCache, err
	}

	for _, page := range pages {
		templateSet, err := template.ParseFiles(page)
		if err != nil {
			return templateCache, err
		}

		if len(layouts) > 0 {
			templateSet, err = templateSet.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return templateCache, err
			}
		}

		name := filepath.Base(page)
		templateCache[name] = templateSet
	}

	return templateCache, nil
}
