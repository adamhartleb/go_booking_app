package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions template.FuncMap

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	templateCache, err := CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	t, ok := templateCache[tmpl]
	if !ok {
		log.Fatal("Template does not exist.")
	}

	err = t.Execute(w, nil)
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
