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

func (app *application) displayMainSignupPage(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)

	app.render(w, http.StatusOK, "signup.html", data)
}

type customerSignupFormResult struct {
	FullName            string `form:"fullName"`
	Email               string `form:"email"`
	Phone               string `form:"phone"`
	Address             string `form:"address"`
	Password            string `form:"password"`
	validator.Validator `form:"-"`
}

func (app *application) displaySignupCustomerPage(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Form = customerSignupFormResult{}

	app.render(w, http.StatusOK, "signup_customer.html", data)
}

func (app *application) doSignupCustomer(w http.ResponseWriter, r *http.Request) {
	var form customerSignupFormResult

	err := app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	if !validator.IsMatchRegex(form.Email, validator.EmailRX) {
		form.AddFieldError("email", "This field must be a valid email address")
	}

	if !validator.IsMatchRegex(form.Phone, validator.PhoneRX) {
		form.AddFieldError("phone", "This field must be a valid phone number")
	}

	// TODO: Validate other fields more detailed.

	if !form.IsNoErrors() {
		data := app.newTemplateData(r)
		data.Form = form

		app.render(w, http.StatusUnprocessableEntity, "signup_customer.html", data)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.DefaultCost)
	if err != nil {
		app.serverError(w, err)
		return
	}

	arg := sqlc.CreateCustomerParams{
		FullName: form.FullName,
		Email:    form.Email,
		Phone:    form.Phone,
		Address:  form.Address,
		Password: string(hashedPassword),
	}

	err = app.store.CreateCustomer(r.Context(), arg)
	if err != nil {
		var postgreSQLError *pq.Error
		if errors.As(err, &postgreSQLError) {
			code := postgreSQLError.Code.Name()
			if code == "unique_violation" && strings.Contains(postgreSQLError.Message, "users_uc_email") {
				form.AddFieldError("email", "Email address is already in use")
			}

			if code == "unique_violation" && strings.Contains(postgreSQLError.Message, "users_uc_phone") {
				form.AddFieldError("phone", "Phone number is already in use")
			}

			data := app.newTemplateData(r)
			data.Form = form

			app.render(w, http.StatusUnprocessableEntity, "signup_customer.html", data)
			return
		}

		app.serverError(w, err)
		return
	}

	app.sessionManager.Put(r.Context(), "flash", "Your signup was successful. Please log in.")

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

type providerSignupFormResult struct {
	FullName            string `form:"fullName"`
	Email               string `form:"email"`
	Phone               string `form:"phone"`
	Address             string `form:"address"`
	CompanyName         string `form:"companyName"`
	TaxCode             string `form:"taxCode"`
	Password            string `form:"password"`
	validator.Validator `form:"-"`
}

func (app *application) displaySignupProviderPage(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Form = providerSignupFormResult{}

	app.render(w, http.StatusOK, "signup_provider.html", data)
}

func (app *application) doSignupProvider(w http.ResponseWriter, r *http.Request) {

}

//func (app *application) doSignupUser(w http.ResponseWriter, r *http.Request) {
//	var form customerSignupFormResult
//
//	err := app.decodePostForm(r, &form)
//	if err != nil {
//		app.clientError(w, http.StatusBadRequest)
//		return
//	}
//
//	if !validator.IsMatchRegex(form.Email, validator.EmailRX) {
//		form.AddFieldError("email", "This field must be a valid email address")
//	}
//
//	if !validator.IsMatchRegex(form.Phone, validator.PhoneRX) {
//		form.AddFieldError("phone", "This field must be a valid phone number")
//	}
//
//	// TODO: Validate address more detailed.
//
//	if !validator.IsStringNotLessThanLimit(form.Password, 8) {
//		form.AddFieldError("password", "This field must be at least 8 characters long")
//	}
//
//	app.infoLog.Println(form.Role)
//
//	// Register a Provider account.
//	if form.Role == "provider" {
//		form.SelectedRole = "provider"
//		app.infoLog.Println(form)
//
//		if !validator.IsNotBlank(form.CompanyName) {
//			form.AddFieldError("companyName", "This field cannot be blank")
//		}
//
//		if !validator.IsNotBlank(form.TaxCode) {
//			form.AddFieldError("taxCode", "This field cannot be blank")
//		}
//
//		// TODO: Validate company name more detailed.
//
//		// TODO: Validate tax code more detailed.
//		fmt.Println("CCCCC")
//
//		if !form.IsNoErrors() {
//			data := app.newTemplateData(r)
//			data.Form = form
//			fmt.Println(form)
//
//			app.render(w, http.StatusUnprocessableEntity, "signup.html", data)
//			return
//		}
//
//		fmt.Println("WTF")
//		app.errorLog.Println(form)
//
//		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.DefaultCost)
//		if err != nil {
//			app.serverError(w, err)
//			return
//		}
//
//		err = app.store.CreateProviderTx(r.Context(), sqlc.CreateProviderTxParams{
//			FullName:    form.FullName,
//			Email:       form.Email,
//			Phone:       form.Phone,
//			Address:     form.Address,
//			CompanyName: form.CompanyName,
//			TaxCode:     form.TaxCode,
//			Password:    string(hashedPassword),
//		})
//		app.errorLog.Println(err)
//		if err != nil {
//			var postgreSQLError *pq.Error
//			if errors.As(err, &postgreSQLError) {
//				code := postgreSQLError.Code.Name()
//				if code == "unique_violation" && strings.Contains(postgreSQLError.Message, "users_uc_email") {
//					form.AddFieldError("email", "Email address is already in use")
//				}
//
//				if code == "unique_violation" && strings.Contains(postgreSQLError.Message, "users_uc_phone") {
//					form.AddFieldError("phone", "Phone number is already in use")
//				}
//
//				data := app.newTemplateData(r)
//				data.Form = form
//				app.render(w, http.StatusUnprocessableEntity, "signup.html", data)
//				return
//			}
//
//			app.serverError(w, err)
//			return
//		}
//	} else {
//		form.SelectedRole = "customer"
//		// Register a Customer account.
//		if !form.IsNoErrors() {
//			data := app.newTemplateData(r)
//			data.Form = form
//
//			app.render(w, http.StatusUnprocessableEntity, "signup.html", data)
//			return
//		}
//
//		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.DefaultCost)
//		if err != nil {
//			app.serverError(w, err)
//			return
//		}
//
//		arg := sqlc.CreateCustomerParams{
//			FullName: form.FullName,
//			Email:    form.Email,
//			Phone:    form.Phone,
//			Address:  form.Address,
//			Password: string(hashedPassword),
//		}
//
//		err = app.store.CreateCustomer(r.Context(), arg)
//		if err != nil {
//			var postgreSQLError *pq.Error
//			if errors.As(err, &postgreSQLError) {
//				code := postgreSQLError.Code.Name()
//				if code == "unique_violation" && strings.Contains(postgreSQLError.Message, "users_uc_email") {
//					form.AddFieldError("email", "Email address is already in use")
//				}
//
//				if code == "unique_violation" && strings.Contains(postgreSQLError.Message, "users_uc_phone") {
//					form.AddFieldError("phone", "Phone number is already in use")
//				}
//
//				data := app.newTemplateData(r)
//				data.Form = form
//				app.render(w, http.StatusUnprocessableEntity, "signup.html", data)
//				return
//			}
//
//			app.serverError(w, err)
//			return
//		}
//	}
//
//	app.sessionManager.Put(r.Context(), "flash", "Your signup was successful. Please log in.")
//
//	http.Redirect(w, r, "/login", http.StatusSeeOther)
//}

func (app *application) displayLoginPage(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Form = userLoginFormResult{}

	app.render(w, http.StatusOK, "login.html", data)
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

	// Check whether user with the email provided exists.
	user, err := app.store.GetUserByEmail(r.Context(), form.Email)
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

	// Check whether the hashed password and plain-text password that user provided match.
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

	// Until here, the user is authenticated successfully.

	// Use the RenewToken() method on the current session to change the session
	// ID. It's good practice to generate a new session ID when the
	// authentication state or privilege levels changes for the user (e.g. login
	// and logout operations).
	err = app.sessionManager.RenewToken(r.Context())
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Add the ID of the current user to the session, so that they are now
	// 'logged in'.
	app.sessionManager.Put(r.Context(), "authenticatedUserID", int(user.ID))

	path := app.sessionManager.PopString(r.Context(), "redirectPathAfterLogin")
	if path != "" {
		http.Redirect(w, r, path, http.StatusSeeOther)
		return
	}

	// Redirect the user to the create snippet page.
	http.Redirect(w, r, "/snippet/create", http.StatusSeeOther)
}

func (app *application) viewAccount(w http.ResponseWriter, r *http.Request) {
	userID := app.sessionManager.GetInt(r.Context(), "authenticatedUserID")

	user, err := app.store.GetUserByID(r.Context(), int32(userID))
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
