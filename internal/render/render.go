package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/IDOMATH/bookings/internal/config"
	"github.com/IDOMATH/bookings/internal/models"
)

var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// renderTemplate renders html templates for pages
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {
		// Get the template cache from app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// Get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	err := t.Execute(buf, td)
	if err != nil {
		log.Println(err)
	}

	// Render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// Get all of the files name *.page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	// Range through all files found above
	for _, page := range pages {
		// Get name of the file, minus full path
		name := filepath.Base(page)
		// Give template a name and parse the file "page"
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		// Look for layout files
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			// If we find a layout file, parse it using the template found above
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
		}

		// Add parsed template to cache.
		myCache[name] = ts
	}

	return myCache, nil
}
