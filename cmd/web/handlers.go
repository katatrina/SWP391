package main

import (
	"database/sql"
	"errors"
	"github.com/katatrina/SWP391/internal/db/sqlc"
	"github.com/katatrina/SWP391/internal/validator"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)

	app.render(w, http.StatusOK, "home.html", data)
}

func (app *application) displayUserSignupPage(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Form = userSignupFormResult{}

	app.render(w, http.StatusOK, "signup.html", data)
}

type userSignupFormResult struct {
	Name                string `form:"name"`
	Email               string `form:"email"`
	Password            string `form:"password"`
	validator.Validator `form:"-"`
}

func (app *application) doSignupUser(w http.ResponseWriter, r *http.Request) {
	var form userSignupFormResult

	err := app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	if !validator.IsNotBlank(form.Name) {
		form.AddFieldError("name", "This field cannot be blank")
	}

	if !validator.IsNotBlank(form.Email) {
		form.AddFieldError("email", "This field cannot be blank")
	}

	if !validator.IsMatchRegex(form.Email, validator.EmailRX) {
		form.AddFieldError("email", "This field must be a valid email address")
	}

	if !validator.IsNotBlank(form.Password) {
		form.AddFieldError("password", "This field cannot be blank")
	}

	if !validator.IsStringNotLessThanLimit(form.Password, 8) {
		form.AddFieldError("password", "This field must be at least 8 characters long")
	}

	if !form.IsNoErrors() {
		data := app.newTemplateData(r)
		data.Form = form

		app.render(w, http.StatusUnprocessableEntity, "signup.html", data)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(form.Password), 12)
	if err != nil {
		app.serverError(w, err)
		return
	}

	arg := sqlc.CreateUserParams{
		Name:     form.Name,
		Email:    form.Email,
		Password: string(hashedPassword),
	}

	err = app.CreateUser(r.Context(), arg)
	if err != nil {
		var postgreSQLError *pq.Error
		if errors.As(err, &postgreSQLError) {
			code := postgreSQLError.Code.Name()
			if code == "unique_violation" && strings.Contains(postgreSQLError.Message, "users_uc_email") {
				form.AddFieldError("email", "Email address is already in use")

				data := app.newTemplateData(r)
				data.Form = form
				app.render(w, http.StatusUnprocessableEntity, "signup.html", data)
				return
			}
		}

		app.serverError(w, err)
		return
	}

	app.sessionManager.Put(r.Context(), "flash", "Your signup was successful. Please log in.")

	http.Redirect(w, r, "/user/login", http.StatusSeeOther)
}

func (app *application) displayUserLoginPage(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Form = userLoginFormResult{}

	app.render(w, http.StatusOK, "login.html", data)
}

type userLoginFormResult struct {
	Email               string `form:"email"`
	Password            string `form:"password"`
	validator.Validator `form:"-"`
}

func (app *application) doLoginUser(w http.ResponseWriter, r *http.Request) {
	var form userLoginFormResult

	err := app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	if !validator.IsNotBlank(form.Email) {
		form.AddFieldError("email", "This field cannot be blank")
	}

	if !validator.IsMatchRegex(form.Email, validator.EmailRX) {
		form.AddFieldError("email", "This field must be a valid email address")
	}

	if !validator.IsNotBlank(form.Password) {
		form.AddFieldError("password", "This field cannot be blank")
	}

	if !form.IsNoErrors() {
		data := app.newTemplateData(r)
		data.Form = form

		app.render(w, http.StatusUnprocessableEntity, "login.html", data)
		return
	}

	user, err := app.GetUserByEmail(r.Context(), form.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			form.AddGenericError("Email or password is incorrect")

			data := app.newTemplateData(r)
			data.Form = form
			app.render(w, http.StatusUnprocessableEntity, "login.html", data)
		} else {
			app.serverError(w, err)
		}

		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(form.Password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			form.AddGenericError("Email or password is incorrect")

			data := app.newTemplateData(r)
			data.Form = form
			app.render(w, http.StatusUnprocessableEntity, "login.html", data)
		} else {
			app.serverError(w, err)
		}

		return
	}

	err = app.sessionManager.RenewToken(r.Context())
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.sessionManager.Put(r.Context(), "authenticatedUserID", int(user.ID))

	path := app.sessionManager.PopString(r.Context(), "redirectPathAfterLogin")
	if path != "" {
		http.Redirect(w, r, path, http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *application) doLogoutUser(w http.ResponseWriter, r *http.Request) {
	err := app.sessionManager.RenewToken(r.Context())
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.sessionManager.Remove(r.Context(), "authenticatedUserID")

	app.sessionManager.Put(r.Context(), "flash", "You've been logged out successfully!")

	http.Redirect(w, r, "/user/login", http.StatusSeeOther)
}

func (app *application) about(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)

	app.render(w, http.StatusOK, "about.html", data)
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

func (app *application) displayChangeUserPasswordPage(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Form = changeUserPasswordFormResult{}

	app.render(w, http.StatusOK, "change-password.html", data)
}

type changeUserPasswordFormResult struct {
	CurrentPassword         string `form:"currentPassword"`
	NewPassword             string `form:"newPassword"`
	NewPasswordConfirmation string `form:"newPasswordConfirmation"`
	validator.Validator     `form:"-"`
}

func (app *application) doUpdateUserPassword(w http.ResponseWriter, r *http.Request) {
	var form changeUserPasswordFormResult

	err := app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	if !validator.IsNotBlank(form.CurrentPassword) {
		form.AddFieldError("currentPassword", "This field cannot be blank")
	}

	if !validator.IsNotBlank(form.NewPassword) {
		form.AddFieldError("newPassword", "This field cannot be blank")
	}

	if !validator.IsStringNotLessThanLimit(form.NewPassword, 8) {
		form.AddFieldError("newPassword", "This field must be at least 8 characters long")
	}

	if form.NewPassword == form.CurrentPassword {
		form.AddFieldError("newPassword", "New password and current password cannot be the same")
	}

	if !validator.IsNotBlank(form.NewPasswordConfirmation) {
		form.AddFieldError("newPasswordConfirmation", "This field cannot be blank")
	}

	if form.NewPassword != form.NewPasswordConfirmation {
		form.AddFieldError("newPasswordConfirmation", "Password and confirmation password do not match")
	}

	if !form.IsNoErrors() {
		data := app.newTemplateData(r)
		data.Form = form

		app.render(w, http.StatusUnprocessableEntity, "change-password.html", data)
		return
	}

	userId := app.sessionManager.GetInt(r.Context(), "authenticatedUserID")

	user, err := app.GetUserByID(r.Context(), int32(userId))
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(form.CurrentPassword))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			form.AddGenericError("Current password is incorrect")

			data := app.newTemplateData(r)
			data.Form = form
			app.render(w, http.StatusUnprocessableEntity, "change-password.html", data)
		} else {
			app.serverError(w, err)
		}

		return
	}

	newUserPassword, err := bcrypt.GenerateFromPassword([]byte(form.NewPassword), 12)
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = app.UpdateUserPassword(r.Context(), sqlc.UpdateUserPasswordParams{
		Password: string(newUserPassword),
		ID:       int32(userId),
	})
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.sessionManager.Put(r.Context(), "flash", "Your password has been updated successfully!")

	http.Redirect(w, r, "/account/view", http.StatusSeeOther)
}
