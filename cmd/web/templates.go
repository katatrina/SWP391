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
	"plusOne":     plusOne,
	"shortenDesc": shortenDescription,
}

type templateData struct {
	CurrentYear         int
	Form                any
	Flash               string
	IsAuthenticated     bool
	IsProvider          bool
	IsAdmin             bool
	ProviderDetail      sqlc.ProviderDetail
	ProviderInfo        sqlc.GetProviderDetailsByServiceIDRow
	Service             sqlc.Service
	Services            []sqlc.Service
	ServiceFeedbacks    []sqlc.ListServiceFeedbacksRow
	User                sqlc.User
	Categories          []sqlc.Category
	Cart                Cart
	PurchaseOrders      map[string]PurchaseOrder
	SellOrders          map[string]SellOrder
	SortedOrders        []string
	OrderStatuses       []sqlc.OrderStatus
	HighlightedButtonID int32
	HighlightedCategory string
	IsUserUsedService   bool
	AdminDashboard      AdminDashboard
	AdminEmail          string
	InactiveServices    []InactiveService
	Customers           []sqlc.User
	Providers           []sqlc.GetProvidersRow
	ProviderDashBoard   ProviderDashboard
}

type ProviderDashboard struct {
	TotalServices        int64
	TotalCompletedOrders int64
	TotalRevenue         int32
}

type InactiveService struct {
	ProviderCompanyName string
	CategoryName        string
	sqlc.Service
}

type AdminDashboard struct {
	TotalCustomers int64
	TotalProviders int64
	TotalServices  int64
	CategoryStats  []CategoryStat
}

type CategoryStat struct {
	CategoryID        int32
	CategoryName      string
	CategoryImagePath string
	Profit            int32
	TotalService      int64
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
		IsAdmin:         app.isAdmin(r),
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
