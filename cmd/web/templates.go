package main

import (
	"fmt"
	"github.com/katatrina/SWP391/internal/db/sqlc"
	"html/template"
	"net/http"
	"path/filepath"
	"time"
)

var functionTemplates = template.FuncMap{
	"humanDate":   humanDate,
	"formatPrice": formatVietnamesePrice,
}

type templateData struct {
	CurrentYear     int
	Form            any
	Flash           string
	IsAuthenticated bool
	IsProvider      bool
	Service         sqlc.Service
	Services        []sqlc.Service
	User            sqlc.User
	Categories      []sqlc.Category
	Cart            Cart
}

type Cart struct {
	GrandTotal int32
	Items      map[string][]sqlc.GetCartItemsByCartIDRow
}

func (app *application) newTemplateData(r *http.Request) *templateData {
	return &templateData{
		CurrentYear:     time.Now().Year(),
		Flash:           app.sessionManager.PopString(r.Context(), "flash"),
		IsAuthenticated: app.isAuthenticated(r),
		IsProvider:      app.isProvider(r),
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
			//fmt.Println(err)
			return nil, err
		}

		ts, err = ts.ParseFiles(page)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		caches[name] = ts
	}

	return caches, nil
}
