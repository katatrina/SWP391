package main

import (
	"database/sql"
	"errors"
	"github.com/julienschmidt/httprouter"
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

func (app *application) displayCategoriesPage(w http.ResponseWriter, r *http.Request) {
	categories, err := app.store.ListCategories(r.Context())
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := app.newTemplateData(r)
	data.Categories = categories

	app.render(w, http.StatusOK, "categories.html", data)
}

func (app *application) displayMainSignupPage(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)

	app.render(w, http.StatusOK, "signup.html", data)
}

type customerSignupFormResult struct {
	FullName            string `form:"full_name"`
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
		form.AddFieldError("email", "Địa chỉ email không hợp lệ")
	}

	if !validator.IsMatchRegex(form.Phone, validator.PhoneRX) {
		form.AddFieldError("phone", "Số điện thoại không hợp lệ")
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

	err = app.store.CreateCustomerTx(r.Context(), arg)
	if err != nil {
		var postgreSQLError *pq.Error
		if errors.As(err, &postgreSQLError) {
			code := postgreSQLError.Code.Name()
			if code == "unique_violation" && strings.Contains(postgreSQLError.Message, "users_uc_email") {
				form.AddFieldError("email", "Địa chỉ email đã được sử dụng")
			}

			if code == "unique_violation" && strings.Contains(postgreSQLError.Message, "users_uc_phone") {
				form.AddFieldError("phone", "Số điện thoại đã được sử dụng")
			}

			data := app.newTemplateData(r)
			data.Form = form

			app.render(w, http.StatusUnprocessableEntity, "signup_customer.html", data)
			return
		}

		app.serverError(w, err)
		return
	}

	app.sessionManager.Put(r.Context(), "flash", "Bạn đã đăng ký tài khoản thành công. Vui lòng đăng nhập.")

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

type providerSignupFormResult struct {
	FullName            string `form:"full_name"`
	Email               string `form:"email"`
	Phone               string `form:"phone"`
	Address             string `form:"address"`
	CompanyName         string `form:"company_name"`
	TaxCode             string `form:"tax_code"`
	Password            string `form:"password"`
	validator.Validator `form:"-"`
}

func (app *application) displaySignupProviderPage(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Form = providerSignupFormResult{}

	app.render(w, http.StatusOK, "signup_provider.html", data)
}

func (app *application) doSignupProvider(w http.ResponseWriter, r *http.Request) {
	var form providerSignupFormResult

	err := app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	if !validator.IsMatchRegex(form.Email, validator.EmailRX) {
		form.AddFieldError("email", "Địa chỉ email không hợp lệ")
	}

	if !validator.IsMatchRegex(form.Phone, validator.PhoneRX) {
		form.AddFieldError("phone", "Số điện thoại không hợp lệ")
	}

	// TODO: Validate more detailed.

	if !form.IsNoErrors() {
		data := app.newTemplateData(r)
		data.Form = form

		app.render(w, http.StatusUnprocessableEntity, "signup_provider.html", data)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.DefaultCost)
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = app.store.CreateProviderTx(r.Context(), sqlc.CreateProviderTxParams{
		FullName:    form.FullName,
		Email:       form.Email,
		Phone:       form.Phone,
		Address:     form.Address,
		CompanyName: form.CompanyName,
		TaxCode:     form.TaxCode,
		Password:    string(hashedPassword),
	})
	if err != nil {
		var postgreSQLError *pq.Error
		if errors.As(err, &postgreSQLError) {
			code := postgreSQLError.Code.Name()
			if code == "unique_violation" && strings.Contains(postgreSQLError.Message, "users_uc_email") {
				form.AddFieldError("email", "Địa chỉ email đã được sử dụng")
			}

			if code == "unique_violation" && strings.Contains(postgreSQLError.Message, "users_uc_phone") {
				form.AddFieldError("phone", "Số điện thoại đã được sử dụng")
			}

			data := app.newTemplateData(r)
			data.Form = form

			app.render(w, http.StatusUnprocessableEntity, "signup_provider.html", data)
			return
		}

		app.serverError(w, err)
		return
	}

	app.sessionManager.Put(r.Context(), "flash", "Bạn đã đăng ký tài khoản thành công. Vui lòng đăng nhập.")

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

type userLoginFormResult struct {
	Email               string `form:"email"`
	Password            string `form:"password"`
	validator.Validator `form:"-"`
}

func (app *application) displayUserLoginPage(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Form = userLoginFormResult{}

	app.render(w, http.StatusOK, "login.html", data)
}

func (app *application) doLoginUser(w http.ResponseWriter, r *http.Request) {
	var form userLoginFormResult

	err := app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	if !validator.IsMatchRegex(form.Email, validator.EmailRX) {
		form.AddFieldError("email", "Địa chỉ email không hợp lệ")
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
			form.AddGenericError("Email hoặc mật khẩu không chính xác")

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
			form.AddGenericError("Email hoặc mật khẩu không chính xác")

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
	app.sessionManager.Put(r.Context(), "authenticatedUserID", user.ID)

	cartID, err := app.store.GetCartIDByUserId(r.Context(), user.ID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.sessionManager.Put(r.Context(), "cartID", cartID)

	path := app.sessionManager.PopString(r.Context(), "redirectPathAfterLogin")
	if path != "" {
		http.Redirect(w, r, path, http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *application) doLogoutUser(w http.ResponseWriter, r *http.Request) {
	// Use the RenewToken() method on the current session to change the session
	// ID again.
	err := app.sessionManager.RenewToken(r.Context())
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Remove the authenticatedUserID from the session data so that user is
	// 'logged out'.
	app.sessionManager.Remove(r.Context(), "authenticatedUserID")

	// Add a flash message to the session to confirm to the user that they've been
	// logged out.
	app.sessionManager.Put(r.Context(), "flash", "Bạn đã đăng xuất thành công!")

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func (app *application) viewAccount(w http.ResponseWriter, r *http.Request) {
	userID := app.sessionManager.GetInt32(r.Context(), "authenticatedUserID")

	user, err := app.store.GetUserByID(r.Context(), userID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := app.newTemplateData(r)
	data.User = user

	app.render(w, http.StatusOK, "account.html", data)
}

func (app *application) listProviderServices(w http.ResponseWriter, r *http.Request) {
	providerID := app.sessionManager.GetInt32(r.Context(), "authenticatedUserID")

	services, err := app.store.ListServiceByProvider(r.Context(), providerID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := app.newTemplateData(r)
	data.Services = services

	app.render(w, http.StatusOK, "provider_services.html", data)
}

type createServiceFormResult struct {
	Title               string `form:"title"`
	Price               int32  `form:"price"`
	Description         string `form:"description"`
	CategoryID          int32  `form:"category"`
	validator.Validator `form:"-"`
}

func (app *application) displayCreateServicePage(w http.ResponseWriter, r *http.Request) {
	categories, err := app.store.ListCategories(r.Context())
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := app.newTemplateData(r)
	data.Categories = categories

	app.render(w, http.StatusOK, "create_service.html", data)
}

func (app *application) doCreateService(w http.ResponseWriter, r *http.Request) {
	var form createServiceFormResult

	err := app.decodeMultipartForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("image")
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	defer file.Close()

	// Now, form fields is kept to be not empty.
	// TODO: Validate form fields more detailed and display error messages in HTML.

	if !form.IsNoErrors() {
		data := app.newTemplateData(r)
		data.Form = form

		app.render(w, http.StatusUnprocessableEntity, "create_service.html", data)
		return
	}

	err = app.saveFileToDisk(file, header)
	if err != nil {
		app.serverError(w, err)
		return
	}

	fileName := header.Filename
	imagePath := "/static/img/" + fileName

	providerID := app.sessionManager.GetInt32(r.Context(), "authenticatedUserID")

	err = app.store.CreateService(r.Context(), sqlc.CreateServiceParams{
		Title:             form.Title,
		Description:       form.Description,
		Price:             form.Price,
		CategoryID:        form.CategoryID,
		ImagePath:         imagePath,
		OwnedByProviderID: providerID,
	})
	if err != nil {
		// TODO: Delete uploaded file from disk if creating service fails.
		app.serverError(w, err)
		return
	}

	app.sessionManager.Put(r.Context(), "flash", "Bạn đã tạo dịch vụ thành công!")

	http.Redirect(w, r, "/account/my-services", http.StatusSeeOther)
}

func (app *application) displayServicesByCategoryPage(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	categorySlug := params.ByName("slug")

	services, err := app.store.GetServicesByCategorySlug(r.Context(), categorySlug)
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := app.newTemplateData(r)
	data.Services = services

	app.render(w, http.StatusOK, "services_by_category.html", data)
}

func (app *application) displayCartPage(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)

	app.render(w, http.StatusOK, "cart-demo.html", data)
}

type addItemToCartFormResult struct {
	ServiceID int32 `form:"service_id"`
	Quantity  int32 `form:"quantity"`
}

func (app *application) addItemToCart(w http.ResponseWriter, r *http.Request) {
	var form addItemToCartFormResult

	err := app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	cartID := app.sessionManager.GetInt32(r.Context(), "cartID")

	service, err := app.store.GetServiceByID(r.Context(), form.ServiceID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	isServiceAlreadyInCart, err := app.store.IsServiceAlreadyInCart(r.Context(), sqlc.IsServiceAlreadyInCartParams{
		CartID:    cartID,
		ServiceID: service.ID,
	})
	if err != nil {
		app.serverError(w, err)
		return
	}

	// If the service is already in the cart, update the quantity and sub-total.
	if isServiceAlreadyInCart {
		cartItem, err := app.store.GetCartItemByCartIDAndServiceID(r.Context(), sqlc.GetCartItemByCartIDAndServiceIDParams{
			CartID:    cartID,
			ServiceID: service.ID,
		})
		if err != nil {
			app.serverError(w, err)
			return
		}

		err = app.store.UpdateCartItemQuantity(r.Context(), sqlc.UpdateCartItemQuantityParams{
			Quantity:  cartItem.Quantity + form.Quantity,
			SubTotal:  cartItem.SubTotal + service.Price*form.Quantity,
			CartID:    cartID,
			ServiceID: service.ID,
		})

	} else {
		// If the service is not in the cart, add it to the cart.
		err = app.store.AddServiceToCart(r.Context(), sqlc.AddServiceToCartParams{
			CartID:    cartID,
			ServiceID: service.ID,
			Quantity:  form.Quantity,
			SubTotal:  service.Price * form.Quantity,
		})
		if err != nil {
			app.serverError(w, err)
			return
		}
	}

	http.Redirect(w, r, r.URL.Path, http.StatusSeeOther)
}

func (app *application) pageNotFound(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)

	app.render(w, http.StatusOK, "not_found.html", data)
}
