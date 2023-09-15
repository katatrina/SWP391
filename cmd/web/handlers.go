package main

import (
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)

	app.render(w, http.StatusOK, "home.html", data)
}

func (app *application) displayServicePage(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)

	app.render(w, http.StatusOK, "services.html", data)
}

func (app *application) displayProductPage(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)

	app.render(w, http.StatusOK, "products.html", data)
}

func (app *application) displayBlogPage(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)

	app.render(w, http.StatusOK, "blogs.html", data)
}

func (app *application) about(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)

	app.render(w, http.StatusOK, "about.html", data)
}

func (app *application) displayUserSignupPage(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)

	app.render(w, http.StatusOK, "signup.html", data)
}

func (app *application) displayUserLoginPage(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)

	app.render(w, http.StatusOK, "login.html", data)
}

func (app *application) viewAccount(w http.ResponseWriter, r *http.Request) {
	userID := app.sessionManager.GetInt(r.Context(), "authenticatedUserID")

	user, err := app.GetUserByID(r.Context(), int32(userID))
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := app.newTemplateData(r)
	data.User = user

	app.render(w, http.StatusOK, "account.html", data)
}

func (app *application) doCreateService(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)

	app.render(w, http.StatusOK, "provider-assets.html", data)
}

func (app *application) doCreateProduct(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)

	app.render(w, http.StatusOK, "provider-assets.html", data)
}

func (app *application) pageNotFound(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)

	app.render(w, http.StatusOK, "notfound.html", data)
}
