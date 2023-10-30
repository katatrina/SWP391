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
	"formatDate":  formatVietnameseDate,
}

type templateData struct {
	CurrentYear         int
	Form                any
	Flash               string
	IsAuthenticated     bool
	IsProvider          bool
	ProviderDetail      sqlc.GetProviderDetailsByServiceIDRow
	Service             sqlc.Service
	Services            []sqlc.Service
	Category            sqlc.Category
	ServiceFeedbacks    []sqlc.ListServiceFeedbacksRow
	User                sqlc.User
	Categories          []sqlc.Category
	Cart                Cart
	PurchaseOrders      map[string]PurchaseOrder
	SellOrders          map[string]SellOrder
	SortedOrders        []string
	OrderStatuses       []sqlc.OrderStatus
	HighlightedButtonID int32
	IsUserUsedService   bool
}

type Cart struct {
	GrandTotal int32
	Items      map[string][]sqlc.GetCartItemsByCartIDRow
}

type PurchaseOrder struct {
	Provider   sqlc.GetFullProviderInfoRow
	Order      sqlc.GetPurchaseOrdersRow
	OrderItems []sqlc.GetFullOrderItemsInformationByOrderIdRow
}

type SellOrder struct {
	Customer   sqlc.User
	Order      sqlc.GetSellOrdersRow
	OrderItems []sqlc.GetFullOrderItemsInformationByOrderIdRow
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
