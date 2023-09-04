package main

import (
	"github.com/katatrina/SWP391/internal/db/sqlc"
	"html/template"
	"net/http"
	"path/filepath"
	"time"
)

var functionTemplates = template.FuncMap{
	"humanDate": humanDate,
}

type templateData struct {
	CurrentYear     int
	Form            any
	Flash           string
	User            sqlc.User
	IsAuthenticated bool
}

func (app *application) newTemplateData(r *http.Request) *templateData {
	return &templateData{
		CurrentYear:     time.Now().Year(),
		IsAuthenticated: app.isAuthenticated(r),
	}
}

func initializeTemplateCache() (map[string]*template.Template, error) {
	caches := make(map[string]*template.Template)

	pages, err := filepath.Glob("./ui/html/pages/*.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {

		name := filepath.Base(page)

		ts := template.New(name).Funcs(functionTemplates)

		ts, err = ts.ParseGlob("./ui/html/partials/*.html")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		caches[name] = ts
	}

	return caches, nil
}
