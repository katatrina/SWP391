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

	router.Handler(http.MethodGet, "/", dynamic.ThenFunc(app.home))
	router.Handler(http.MethodGet, "/services", dynamic.ThenFunc(app.displayServicePage))
	router.Handler(http.MethodGet, "/products", dynamic.ThenFunc(app.displayProductPage))
	router.Handler(http.MethodGet, "/blogs", dynamic.ThenFunc(app.displayBlogPage))
	router.Handler(http.MethodGet, "/user/signup", dynamic.ThenFunc(app.displayUserSignupPage))
	//router.Handler(http.MethodPost, "/user/signup", dynamic.ThenFunc(app.doSignupUser))
	router.Handler(http.MethodGet, "/user/login", dynamic.ThenFunc(app.displayUserLoginPage))
	//router.Handler(http.MethodPost, "/user/login", dynamic.ThenFunc(app.doLoginUser))
	router.Handler(http.MethodGet, "/about", dynamic.ThenFunc(app.about))

	standard := alice.New(app.logRequest)
	return standard.Then(router)
}
