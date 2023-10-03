package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	"net/http"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	fileServer := http.FileServer(http.Dir("./ui/static"))
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static/", fileServer))

	dynamic := alice.New(app.sessionManager.LoadAndSave, app.authenticate)

	// TODO: Create a NotFound html file.
	router.Handler(http.MethodGet, "/404", dynamic.ThenFunc(app.pageNotFound))

	// Guest permissions.
	router.Handler(http.MethodGet, "/", dynamic.ThenFunc(app.home))

	router.Handler(http.MethodGet, "/services", dynamic.ThenFunc(app.displayServicePage))
	router.Handler(http.MethodGet, "/products", dynamic.ThenFunc(app.displayProductPage))
	router.Handler(http.MethodGet, "/blogs", dynamic.ThenFunc(app.displayBlogPage))
	router.Handler(http.MethodGet, "/about", dynamic.ThenFunc(app.about))

	router.Handler(http.MethodGet, "/signup", dynamic.ThenFunc(app.displayMainSignupPage))

	router.Handler(http.MethodGet, "/signup/customer", dynamic.ThenFunc(app.displaySignupCustomerPage))
	router.Handler(http.MethodPost, "/signup/customer", dynamic.ThenFunc(app.doSignupCustomer))

	router.Handler(http.MethodGet, "/signup/provider", dynamic.ThenFunc(app.displaySignupProviderPage))
	router.Handler(http.MethodPost, "/signup/provider", dynamic.ThenFunc(app.doSignupProvider))

	router.Handler(http.MethodGet, "/login", dynamic.ThenFunc(app.displayUserLoginPage))
	router.Handler(http.MethodPost, "/login", dynamic.ThenFunc(app.doLoginUser))

	// User permissions.
	protected := dynamic.Append(app.requireAuthentication)

	router.Handler(http.MethodGet, "/logout", protected.ThenFunc(app.doLogoutUser))

	router.Handler(http.MethodGet, "/account/view", protected.ThenFunc(app.viewAccount))

	// Provider permissions.
	advanced := protected.Append(app.requireProviderPermission)

	router.Handler(http.MethodGet, "/account/services", advanced.ThenFunc(app.listProviderServices))

	router.Handler(http.MethodGet, "/service/create", advanced.ThenFunc(app.displayCreateServicePage))
	router.Handler(http.MethodPost, "/service/create", advanced.ThenFunc(app.doCreateService))

	router.Handler(http.MethodGet, "/account/products", advanced.ThenFunc(app.listProviderProducts))

	router.Handler(http.MethodGet, "/product/create", advanced.ThenFunc(app.displayCreateProductPage))
	router.Handler(http.MethodPost, "/product/create", advanced.ThenFunc(app.doCreateProduct))

	standard := alice.New(app.logRequest)
	return standard.Then(router)
}
