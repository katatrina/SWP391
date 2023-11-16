package main

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"github.com/katatrina/SWP391/internal/db/sqlc"
	"github.com/katatrina/SWP391/internal/validator"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"sort"
	"strconv"
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
	if app.isAuthenticated(r) {
		http.Redirect(w, r, "/account/view", http.StatusSeeOther)
		return
	}

	if app.isAdmin(r) {
		http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
		return
	}

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
	if app.isAuthenticated(r) {
		http.Redirect(w, r, "/account/view", http.StatusSeeOther)
		return
	}

	if app.isAdmin(r) {
		http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
		return
	}

	data := app.newTemplateData(r)
	data.Form = customerSignupFormResult{}

	app.render(w, http.StatusOK, "signup_customer.html", data)
}

func (app *application) doSignupCustomer(w http.ResponseWriter, r *http.Request) {
	if app.isAuthenticated(r) {
		http.Redirect(w, r, "/account/view", http.StatusSeeOther)
		return
	}

	if app.isAdmin(r) {
		http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
		return
	}

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
		form.AddFieldError("phone", "Số điện thoại không hợp lệ. VD: 0256789123")
	}

	if validator.IsAddressTooShort(form.Address) {
		form.AddFieldError("address", "Địa chỉ quá ngắn")
	}

	if validator.IsPasswordTooShort(form.Password) {
		form.AddFieldError("password", "Mật khẩu quá ngắn. Tối thiểu 8 ký tự")
	}

	if !form.IsNoErrors() {
		data := app.newTemplateData(r)
		data.Form = form

		app.render(w, http.StatusUnprocessableEntity, "signup_customer.html", data)
		return
	}

	arg := sqlc.CreateCustomerParams{
		FullName: form.FullName,
		Email:    form.Email,
		Phone:    form.Phone,
		Address:  form.Address,
		Password: form.Password,
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
	if app.isAuthenticated(r) {
		http.Redirect(w, r, "/account/view", http.StatusSeeOther)
		return
	}

	if app.isAdmin(r) {
		http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
		return
	}

	data := app.newTemplateData(r)
	data.Form = providerSignupFormResult{}

	app.render(w, http.StatusOK, "signup_provider.html", data)
}

func (app *application) doSignupProvider(w http.ResponseWriter, r *http.Request) {
	if app.isAuthenticated(r) {
		http.Redirect(w, r, "/account/view", http.StatusSeeOther)
		return
	}

	if app.isAdmin(r) {
		http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
		return
	}

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
		form.AddFieldError("phone", "Số điện thoại không hợp lệ. VD: 0256789123")
	}

	if validator.IsAddressTooShort(form.Address) {
		form.AddFieldError("address", "Địa chỉ quá ngắn")
	}

	if !validator.IsMatchRegex(form.TaxCode, validator.TaxCodeRX) {
		form.AddFieldError("tax_code", "Mã số thuế không hợp lệ. VD: 123-45-67890")
	}

	if validator.IsPasswordTooShort(form.Password) {
		form.AddFieldError("password", "Mật khẩu quá ngắn. Tối thiểu 8 ký tự")
	}

	if !form.IsNoErrors() {
		data := app.newTemplateData(r)
		data.Form = form

		app.render(w, http.StatusUnprocessableEntity, "signup_provider.html", data)
		return
	}

	err = app.store.CreateProviderTx(r.Context(), sqlc.CreateProviderTxParams{
		FullName:    form.FullName,
		Email:       form.Email,
		Phone:       form.Phone,
		Address:     form.Address,
		CompanyName: form.CompanyName,
		TaxCode:     form.TaxCode,
		Password:    form.Password,
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
	if app.isAuthenticated(r) {
		http.Redirect(w, r, "/account/view", http.StatusSeeOther)
		return
	}

	if app.isAdmin(r) {
		http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
		return
	}

	data := app.newTemplateData(r)
	data.Form = userLoginFormResult{}

	app.render(w, http.StatusOK, "login.html", data)
}

func (app *application) doLoginUser(w http.ResponseWriter, r *http.Request) {
	if app.isAuthenticated(r) {
		http.Redirect(w, r, "/account/view", http.StatusSeeOther)
		return
	}

	if app.isAdmin(r) {
		http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
		return
	}

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

	// Check whether user is admin.
	isAdmin, err := app.store.IsAdminByEmail(r.Context(), form.Email)
	if err != nil {
		app.serverError(w, err)
		return
	}

	if isAdmin {
		admin, err := app.store.GetAdminByEmail(r.Context(), form.Email)
		if err != nil {
			app.serverError(w, err)
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(form.Password))
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

		err = app.sessionManager.RenewToken(r.Context())
		if err != nil {
			app.serverError(w, err)
			return
		}

		app.sessionManager.Put(r.Context(), "authenticatedAdminID", admin.ID)

		http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)

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

func (app *application) listProviderServices(w http.ResponseWriter, r *http.Request) {
	providerID := app.sessionManager.GetInt32(r.Context(), "authenticatedUserID")

	services, err := app.store.ListServiceByProvider(r.Context(), providerID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	activeServices, err := app.store.ListActiveServicesByProviderID(r.Context(), providerID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := app.newTemplateData(r)
	data.Services = services

	var providerDashboard ProviderDashboard

	providerDashboard.TotalServices = int64(len(activeServices))

	providerDashboard.TotalCompletedOrders, err = app.store.CountCompletedOrdersByProviderID(r.Context(), providerID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	revenueResult, err := app.store.GetTotalRevenueByProviderID(r.Context(), providerID)
	if err != nil {
		app.serverError(w, err)
		return
	}
	//providerDashboard.TotalRevenue, ok := revenueResult.(int32)
	revenue, ok := revenueResult.(int64)
	if !ok {
		app.serverError(w, errors.New("cannot convert revenueResult to int64"))
		return
	}

	providerDashboard.TotalRevenue = revenue

	data.ProviderDashBoard = providerDashboard

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
	data.Form = createServiceFormResult{}

	app.render(w, http.StatusOK, "create_service.html", data)
}

func (app *application) doCreateService(w http.ResponseWriter, r *http.Request) {
	var form createServiceFormResult

	err := app.decodeMultipartForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	if !validator.IsTitleValid(form.Title) {
		form.AddFieldError("title", "Tiêu đề không hợp lệ. Độ dài phải từ 10 đến 150 ký tự.")
	}

	if !validator.IsDescriptionValid(form.Description) {
		form.AddFieldError("description", "Mô tả không hợp lệ. Độ dài phải từ 10 đến 500 ký tự.")
	}

	if form.Price < 1_000 || form.Price > 1_000_000_000 {
		form.AddFieldError("price", "Giá tiền không hợp lệ. Giá tiền phải từ ₫1.000 đến ₫1.000.000.000.")
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
		categories, err := app.store.ListCategories(r.Context())
		if err != nil {
			app.serverError(w, err)
			return
		}

		data := app.newTemplateData(r)
		data.Form = form
		data.Categories = categories

		fmt.Printf("%+v\n", data.Form)

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

	app.sessionManager.Put(r.Context(), "flash", "Đã yêu cầu tạo dịch vụ thành công. Vui lòng chờ quản trị viên phê duyệt.")

	http.Redirect(w, r, "/account/my-services", http.StatusSeeOther)
}

func (app *application) displayEditServicePage(w http.ResponseWriter, r *http.Request) {
	userID := app.sessionManager.GetInt32(r.Context(), "authenticatedUserID")

	params := httprouter.ParamsFromContext(r.Context())

	serviceID, err := strconv.ParseInt(params.ByName("id"), 10, 32)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	service, err := app.store.GetServiceByID(r.Context(), int32(serviceID))
	if err != nil {
		app.serverError(w, err)
		return
	}

	if service.OwnedByProviderID != userID {
		app.clientError(w, http.StatusForbidden)
		return
	}

	categories, err := app.store.ListCategories(r.Context())
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := app.newTemplateData(r)
	data.Service = service
	data.Categories = categories

	app.render(w, http.StatusOK, "edit_service.html", data)
}

func (app *application) displayServicesByCategoryPage(w http.ResponseWriter, r *http.Request) {
	userID := app.sessionManager.GetInt32(r.Context(), "authenticatedUserID")

	params := httprouter.ParamsFromContext(r.Context())

	categorySlug := params.ByName("slug")
	if categorySlug == "all" {
		categories, err := app.store.ListCategories(r.Context())
		if err != nil {
			app.serverError(w, err)
			return
		}

		services, err := app.store.ListServices(r.Context(), userID)
		if err != nil {
			app.serverError(w, err)
			return
		}

		data := app.newTemplateData(r)
		data.Services = services
		data.Categories = categories
		data.HighlightedCategory = "all"

		app.render(w, http.StatusOK, "services_by_category.html", data)
		return
	}

	isExists, err := app.store.IsCategoryExists(r.Context(), categorySlug)
	if err != nil {
		app.serverError(w, err)
		return
	}

	if !isExists {
		app.clientError(w, http.StatusNotFound)
		return
	}

	services, err := app.store.GetServicesByCategorySlug(r.Context(), sqlc.GetServicesByCategorySlugParams{
		Slug:              categorySlug,
		OwnedByProviderID: userID,
	})
	if err != nil {
		app.serverError(w, err)
		return
	}

	categories, err := app.store.ListCategories(r.Context())
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := app.newTemplateData(r)
	data.Services = services
	data.Categories = categories
	data.HighlightedCategory = categorySlug

	app.render(w, http.StatusOK, "services_by_category.html", data)
}

func (app *application) displayServiceDetailsPage(w http.ResponseWriter, r *http.Request) {
	userID := app.sessionManager.GetInt32(r.Context(), "authenticatedUserID")

	params := httprouter.ParamsFromContext(r.Context())

	serviceID, err := strconv.ParseInt(params.ByName("id"), 10, 32)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	providerDetail, err := app.store.GetProviderDetailsByServiceID(r.Context(), int32(serviceID))
	if err != nil {
		app.serverError(w, err)
		return
	}

	service, err := app.store.GetServiceByID(r.Context(), int32(serviceID))
	if err != nil {
		app.serverError(w, err)
		return
	}

	isUserUsedService, err := app.store.IsUserUsedService(r.Context(), sqlc.IsUserUsedServiceParams{
		BuyerID:   userID,
		ServiceID: int32(serviceID),
	})
	if err != nil {
		app.serverError(w, err)
		return
	}
	fmt.Println(isUserUsedService)

	feedbacks, err := app.store.ListServiceFeedbacks(r.Context(), int32(serviceID))
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := app.newTemplateData(r)
	data.Service = service
	data.ProviderInfo = providerDetail
	data.ServiceFeedbacks = feedbacks
	data.IsUserUsedService = isUserUsedService

	app.render(w, http.StatusOK, "service_details.html", data)
}

func (app *application) displayCart(w http.ResponseWriter, r *http.Request) {
	cartID := app.sessionManager.GetInt32(r.Context(), "cartID")

	cart := Cart{
		GrandTotal: 0,
		Items:      make(map[string][]sqlc.GetCartItemsByCartIDRow),
	}

	items, err := app.store.GetCartItemsByCartID(r.Context(), cartID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	for _, item := range items {
		companyName, err := app.store.GetCompanyNameByServiceID(r.Context(), item.ServiceID)
		if err != nil {
			app.serverError(w, err)
			return
		}

		cart.Items[companyName] = append(cart.Items[companyName], item)

		cart.GrandTotal += item.SubTotal
	}

	data := app.newTemplateData(r)
	data.Cart = cart

	app.render(w, http.StatusOK, "cart.html", data)
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

		err = app.store.UpdateCartItemQuantityAndSubTotal(r.Context(), sqlc.UpdateCartItemQuantityAndSubTotalParams{
			Quantity: form.Quantity + cartItem.Quantity,
			SubTotal: cartItem.SubTotal + service.Price*form.Quantity,
			UUID:     cartItem.UUID,
		})
		if err != nil {
			app.serverError(w, err)
			return
		}

		http.Redirect(w, r, "/cart", http.StatusSeeOther)
		return
	}

	// If the service is not in the cart, add it to the cart.
	itemID, err := uuid.NewRandom()
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = app.store.AddServiceToCart(r.Context(), sqlc.AddServiceToCartParams{
		UUID:      itemID.String(),
		CartID:    cartID,
		ServiceID: service.ID,
		Quantity:  form.Quantity,
		SubTotal:  service.Price * form.Quantity,
	})
	if err != nil {
		app.serverError(w, err)
		return
	}

	// TODO: Display flash message.

	http.Redirect(w, r, "/cart", http.StatusSeeOther)
}

func (app *application) updateCart(w http.ResponseWriter, r *http.Request) {
	cartID := app.sessionManager.GetInt32(r.Context(), "cartID")

	items, err := app.store.GetCartItemsByCartID(r.Context(), cartID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	for _, item := range items {
		quantity, err := strconv.ParseInt(r.PostFormValue(item.UUID), 10, 32)
		if err != nil {
			app.errorLog.Println(err)
			app.clientError(w, http.StatusBadRequest)
			return
		}

		// TODO: Handle the case when quantity is 0.

		err = app.store.UpdateCartItemQuantityAndSubTotal(r.Context(), sqlc.UpdateCartItemQuantityAndSubTotalParams{
			Quantity: int32(quantity),
			SubTotal: item.Price * int32(quantity),
			UUID:     item.UUID,
		})
		if err != nil {
			app.serverError(w, err)
			return
		}
	}

	http.Redirect(w, r, "/cart", http.StatusSeeOther)
}

func (app *application) removeItemFromCart(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	cartItemID := params.ByName("id")

	err := app.store.RemoveItemFromCart(r.Context(), cartItemID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, "/cart", http.StatusSeeOther)
}

func (app *application) displayCheckoutPage(w http.ResponseWriter, r *http.Request) {
	userID := app.sessionManager.GetInt32(r.Context(), "authenticatedUserID")
	user, err := app.store.GetUserByID(r.Context(), userID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	cartID := app.sessionManager.GetInt32(r.Context(), "cartID")

	cart := Cart{
		GrandTotal: 0,
		Items:      make(map[string][]sqlc.GetCartItemsByCartIDRow),
	}

	items, err := app.store.GetCartItemsByCartID(r.Context(), cartID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	for _, item := range items {
		companyName, err := app.store.GetCompanyNameByServiceID(r.Context(), item.ServiceID)
		if err != nil {
			app.serverError(w, err)
			return
		}

		cart.Items[companyName] = append(cart.Items[companyName], item)

		cart.GrandTotal += item.SubTotal
	}

	data := app.newTemplateData(r)
	data.Cart = cart
	data.User = user

	app.render(w, http.StatusOK, "checkout.html", data)
}

func (app *application) doCheckout(w http.ResponseWriter, r *http.Request) {
	paymentMethod := r.PostFormValue("payment_method")

	userID := app.sessionManager.GetInt32(r.Context(), "authenticatedUserID")

	cartID := app.sessionManager.GetInt32(r.Context(), "cartID")

	cart := make(map[int32][]sqlc.GetCartItemsByCartIDRow)

	items, err := app.store.GetCartItemsByCartID(r.Context(), cartID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	for _, item := range items {
		cart[item.OwnedByProviderID] = append(cart[item.OwnedByProviderID], item)
	}

	for providerID, cartItems := range cart {
		err = app.store.CreateOrderTx(r.Context(), sqlc.CreateOrderTxParams{
			BuyerID:       userID,
			SellerID:      providerID,
			PaymentMethod: paymentMethod,
			CartItems:     cartItems,
			StatusID:      1,
		})
		if err != nil {
			app.serverError(w, err)
			return
		}
	}

	app.sessionManager.Put(r.Context(), "flash", "Bạn đã đặt dịch vụ thành công!")

	http.Redirect(w, r, "/my-orders/identity/buyer", http.StatusSeeOther)
}

var categoryStatusMap = map[string]string{
	"1": "pending",
	"2": "confirmed",
	"3": "completed",
	"4": "cancelled",
}

func (app *application) displayPurchaseOrdersPage(w http.ResponseWriter, r *http.Request) {
	userID := app.sessionManager.GetInt32(r.Context(), "authenticatedUserID")

	orders := make(map[string]PurchaseOrder)

	categoryStatusID := r.URL.Query().Get("type")

	var highlightedButtonID int32

	orderStatuses, err := app.store.GetOrderStatuses(r.Context())
	if err != nil {
		app.serverError(w, err)
		return
	}

	if categoryStatusID == "" || categoryStatusID == "0" {
		// Get all purchase orders of the current user.
		myOrders, err := app.store.GetPurchaseOrders(r.Context(), userID)
		if err != nil {
			app.serverError(w, err)
			return
		}

		for _, order := range myOrders {
			provider, err := app.store.GetFullProviderInfo(r.Context(), order.SellerID)
			if err != nil {
				app.serverError(w, err)
				return
			}

			orderItems, err := app.store.GetFullOrderItemsInformationByOrderId(r.Context(), order.UUID)

			orders[order.UUID] = PurchaseOrder{
				Provider:   provider,
				Order:      order,
				OrderItems: orderItems,
			}
		}

		sortedOrders := make([]string, 0, len(orders))
		for k := range orders {
			sortedOrders = append(sortedOrders, k)
		}

		sort.Slice(sortedOrders, func(i, j int) bool {
			return orders[sortedOrders[i]].Order.CreatedAt.After(orders[sortedOrders[j]].Order.CreatedAt)
		})

		data := app.newTemplateData(r)
		data.PurchaseOrders = orders
		data.OrderStatuses = orderStatuses
		data.HighlightedButtonID = highlightedButtonID
		data.SortedOrders = sortedOrders

		app.render(w, http.StatusOK, "don-mua.html", data)
		return
	}

	orderStatusCode := categoryStatusMap[categoryStatusID]

	// Get all purchase orders of the current user with order status code.
	myOrders, err := app.store.GetPurchaseOrdersWithStatusCode(r.Context(), sqlc.GetPurchaseOrdersWithStatusCodeParams{
		BuyerID: userID,
		Code:    orderStatusCode,
	})
	if err != nil {
		app.serverError(w, err)
		return
	}

	for _, order := range myOrders {
		provider, err := app.store.GetFullProviderInfo(r.Context(), order.SellerID)
		if err != nil {
			app.serverError(w, err)
			return
		}

		orderItems, err := app.store.GetFullOrderItemsInformationByOrderId(r.Context(), order.UUID)

		orders[order.UUID] = PurchaseOrder{
			Provider:   provider,
			Order:      sqlc.GetPurchaseOrdersRow(order),
			OrderItems: orderItems,
		}
	}

	parsedInt, _ := strconv.ParseInt(categoryStatusID, 10, 32)

	sortedOrders := make([]string, 0, len(orders))
	for k := range orders {
		sortedOrders = append(sortedOrders, k)
	}

	sort.Slice(sortedOrders, func(i, j int) bool {
		return orders[sortedOrders[i]].Order.CreatedAt.After(orders[sortedOrders[j]].Order.CreatedAt)
	})

	data := app.newTemplateData(r)
	data.PurchaseOrders = orders
	data.OrderStatuses = orderStatuses
	data.HighlightedButtonID = int32(parsedInt)
	data.SortedOrders = sortedOrders

	app.render(w, http.StatusOK, "don-mua.html", data)
}

func (app *application) displaySellOrdersPage(w http.ResponseWriter, r *http.Request) {
	userID := app.sessionManager.GetInt32(r.Context(), "authenticatedUserID")

	orders := make(map[string]SellOrder)

	categoryStatusID := r.URL.Query().Get("type")

	var highlightedButtonID int32

	orderStatuses, err := app.store.GetOrderStatuses(r.Context())
	if err != nil {
		app.serverError(w, err)
		return
	}

	if categoryStatusID == "" || categoryStatusID == "0" {
		// Get all sell orders of the current provider.
		myOrders, err := app.store.GetSellOrders(r.Context(), userID)
		if err != nil {
			app.serverError(w, err)
			return
		}

		for _, order := range myOrders {
			customer, err := app.store.GetUserByID(r.Context(), order.BuyerID)
			if err != nil {
				app.serverError(w, err)
				return
			}

			orderItems, err := app.store.GetFullOrderItemsInformationByOrderId(r.Context(), order.UUID)

			orders[order.UUID] = SellOrder{
				Customer:   customer,
				Order:      order,
				OrderItems: orderItems,
			}
		}

		sortedOrders := make([]string, 0, len(orders))
		for k := range orders {
			sortedOrders = append(sortedOrders, k)
		}

		sort.Slice(sortedOrders, func(i, j int) bool {
			return orders[sortedOrders[i]].Order.CreatedAt.After(orders[sortedOrders[j]].Order.CreatedAt)
		})

		data := app.newTemplateData(r)
		data.SellOrders = orders
		data.OrderStatuses = orderStatuses
		data.HighlightedButtonID = highlightedButtonID
		data.SortedOrders = sortedOrders

		app.render(w, http.StatusOK, "don-ban.html", data)
		return
	}

	orderStatusCode := categoryStatusMap[categoryStatusID]

	// Get all sell orders of the current provider with order status code.
	myOrders, err := app.store.GetSellOrdersWithStatusCode(r.Context(), sqlc.GetSellOrdersWithStatusCodeParams{
		SellerID: userID,
		Code:     orderStatusCode,
	})
	if err != nil {
		app.serverError(w, err)
		return
	}

	for _, order := range myOrders {
		customer, err := app.store.GetUserByID(r.Context(), order.BuyerID)
		if err != nil {
			app.serverError(w, err)
			return
		}

		orderItems, err := app.store.GetFullOrderItemsInformationByOrderId(r.Context(), order.UUID)

		orders[order.UUID] = SellOrder{
			Customer:   customer,
			Order:      sqlc.GetSellOrdersRow(order),
			OrderItems: orderItems,
		}
	}

	parsedInt, _ := strconv.ParseInt(categoryStatusID, 10, 32)

	sortedOrders := make([]string, 0, len(orders))
	for k := range orders {
		sortedOrders = append(sortedOrders, k)
	}

	sort.Slice(sortedOrders, func(i, j int) bool {
		return orders[sortedOrders[i]].Order.CreatedAt.After(orders[sortedOrders[j]].Order.CreatedAt)
	})

	data := app.newTemplateData(r)
	data.SellOrders = orders
	data.OrderStatuses = orderStatuses
	data.HighlightedButtonID = int32(parsedInt)
	data.SortedOrders = sortedOrders

	app.render(w, http.StatusOK, "don-ban.html", data)
}

type updateOrderStatusFormResult struct {
	OrderID           string `form:"order_id"`
	UpdatedStatusID   int32  `form:"updated_status_id"`
	UpdatedStatusCode string `form:"updated_status_code"`
}

func (app *application) updateOrderStatus(w http.ResponseWriter, r *http.Request) {
	var form updateOrderStatusFormResult

	err := app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	err = app.store.UpdateOrderStatus(r.Context(), sqlc.UpdateOrderStatusParams{
		UUID: form.OrderID,
		Code: form.UpdatedStatusCode,
	})
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/my-orders/identity/seller?type=%v", form.UpdatedStatusID), http.StatusSeeOther)
}

type createServiceFeedbackFormResult struct {
	ServiceID int32  `form:"service_id"`
	Content   string `form:"content"`
}

func (app *application) createServiceFeedback(w http.ResponseWriter, r *http.Request) {
	userID := app.sessionManager.GetInt32(r.Context(), "authenticatedUserID")

	var form createServiceFeedbackFormResult

	err := app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	err = app.store.CreateServiceFeedback(r.Context(), sqlc.CreateServiceFeedbackParams{
		ServiceID: form.ServiceID,
		UserID:    userID,
		Content:   form.Content,
	})
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/service/view/%v", form.ServiceID), http.StatusSeeOther)
}

func (app *application) viewAccount(w http.ResponseWriter, r *http.Request) {
	userID := app.sessionManager.GetInt32(r.Context(), "authenticatedUserID")

	user, err := app.store.GetUserByID(r.Context(), userID)

	role, err := app.store.GetUserRoleByID(r.Context(), userID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := app.newTemplateData(r)
	data.User = user

	if role == "provider" {
		providerDetail, err := app.store.GetProviderDetailsByID(r.Context(), userID)
		if err != nil {
			app.serverError(w, err)
			return
		}

		data.ProviderDetail = providerDetail
	}

	app.render(w, http.StatusOK, "account.html", data)
}

type updateAccountFormResult struct {
	FullName    string `form:"full_name"`
	Phone       string `form:"phone"`
	Email       string `form:"email"`
	Address     string `form:"address"`
	CompanyName string `form:"company_name"`
	TaxCode     string `form:"tax_code"`
}

func (app *application) updateAccount(w http.ResponseWriter, r *http.Request) {
	var form updateAccountFormResult

	err := app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	// TODO: Validate form fields more detailed.

	userID := app.sessionManager.GetInt32(r.Context(), "authenticatedUserID")

	role, err := app.store.GetUserRoleByID(r.Context(), userID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = app.store.UpdateCustomerInfo(r.Context(), sqlc.UpdateCustomerInfoParams{
		ID:       userID,
		FullName: form.FullName,
		Email:    form.Email,
		Phone:    form.Phone,
		Address:  form.Address,
	})
	if err != nil {
		app.serverError(w, err)
		return
	}

	if role == "provider" {
		err = app.store.UpdateProviderInfo(r.Context(), sqlc.UpdateProviderInfoParams{
			ProviderID:  userID,
			CompanyName: form.CompanyName,
			TaxCode:     form.TaxCode,
		})
		if err != nil {
			app.serverError(w, err)
			return
		}
	}

	app.sessionManager.Put(r.Context(), "flash", "Bạn đã cập nhật thông tin tài khoản thành công!")

	http.Redirect(w, r, "/account/view", http.StatusSeeOther)
}

func (app *application) pageNotFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

func (app *application) displayAdminDashboardPage(w http.ResponseWriter, r *http.Request) {
	adminID := app.sessionManager.GetInt32(r.Context(), "authenticatedAdminID")
	adminEmail, err := app.store.GetAdminEmailByID(r.Context(), adminID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := app.newTemplateData(r)

	data.AdminEmail = adminEmail

	totalCustomer, err := app.store.GetCustomerNumber(r.Context())
	if err != nil {
		app.serverError(w, err)
		return
	}
	data.AdminDashboard.TotalCustomers = totalCustomer

	totalProvider, err := app.store.GetProviderNumber(r.Context())
	if err != nil {
		app.serverError(w, err)
		return
	}
	data.AdminDashboard.TotalProviders = totalProvider

	totalService, err := app.store.GetServiceNumber(r.Context())
	if err != nil {
		app.serverError(w, err)
		return
	}
	data.AdminDashboard.TotalServices = totalService

	categories, err := app.store.ListCategories(r.Context())
	if err != nil {
		app.serverError(w, err)
		return
	}

	for _, category := range categories {
		var categoryStat CategoryStat

		categoryStat.CategoryID = category.ID
		categoryStat.CategoryName = category.Name
		categoryStat.CategoryImagePath = category.ImagePath

		serviceNumber, err := app.store.GetServiceNumberByCategoryID(r.Context(), category.ID)
		if err != nil {
			app.serverError(w, err)
			return
		}
		categoryStat.TotalService = serviceNumber

		completedOrderItems, err := app.store.GetCompletedOrderItemsByCategoryID(r.Context(), category.ID)
		if err != nil {
			app.serverError(w, err)
			return
		}

		for _, item := range completedOrderItems {
			categoryStat.Profit += item.SubTotal
		}

		data.AdminDashboard.CategoryStats = append(data.AdminDashboard.CategoryStats, categoryStat)
	}

	app.render(w, http.StatusOK, "admin_dashboard.html", data)
}

func (app *application) displayAdminManageAccountPage(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)

	adminID := app.sessionManager.GetInt32(r.Context(), "authenticatedAdminID")
	adminEmail, err := app.store.GetAdminEmailByID(r.Context(), adminID)
	if err != nil {
		app.serverError(w, err)
		return
	}
	data.AdminEmail = adminEmail

	customers, err := app.store.ListCustomers(r.Context())
	if err != nil {
		app.serverError(w, err)
		return
	}
	data.Customers = customers

	providers, err := app.store.GetProviders(r.Context())
	if err != nil {
		app.serverError(w, err)
		return
	}
	data.Providers = providers

	app.render(w, http.StatusOK, "admin_account_management.html", data)
}

func (app *application) displayAdminManageServicePage(w http.ResponseWriter, r *http.Request) {
	adminID := app.sessionManager.GetInt32(r.Context(), "authenticatedAdminID")
	adminEmail, err := app.store.GetAdminEmailByID(r.Context(), adminID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := app.newTemplateData(r)

	data.AdminEmail = adminEmail

	services, err := app.store.ListInactiveServices(r.Context())
	if err != nil {
		app.serverError(w, err)
		return
	}

	inactiveServices := make([]InactiveService, 0, len(services))

	for _, service := range services {
		providerCompanyName, err := app.store.GetCompanyNameByServiceID(r.Context(), service.ID)
		if err != nil {
			app.serverError(w, err)
			return
		}

		categoryName, err := app.store.GetCategoryNameByServiceID(r.Context(), service.ID)
		if err != nil {
			app.serverError(w, err)
			return
		}

		inactiveService := InactiveService{
			Service:             service,
			CategoryName:        categoryName,
			ProviderCompanyName: providerCompanyName,
		}

		inactiveServices = append(inactiveServices, inactiveService)
	}

	data.InactiveServices = inactiveServices

	app.render(w, http.StatusOK, "admin_service_management.html", data)
}

func (app *application) doLogoutAdmin(w http.ResponseWriter, r *http.Request) {
	err := app.sessionManager.RenewToken(r.Context())
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.sessionManager.Remove(r.Context(), "authenticatedAdminID")

	app.sessionManager.Put(r.Context(), "flash", "Bạn đã đăng xuất thành công!")

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

type handleInactiveServiceFormResult struct {
	ServiceID    int32  `form:"service_id"`
	RejectReason string `form:"reject_reason"`
}

func (app *application) handleInactiveService(w http.ResponseWriter, r *http.Request) {
	var form handleInactiveServiceFormResult

	err := app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	if form.RejectReason != "" {
		err = app.store.UpdateServiceStatus(r.Context(), sqlc.UpdateServiceStatusParams{
			Status:       "rejected",
			RejectReason: form.RejectReason,
			ID:           form.ServiceID,
		})
		if err != nil {
			app.serverError(w, err)
			return
		}

		http.Redirect(w, r, "/admin/manage-service", http.StatusSeeOther)
		return
	}

	err = app.store.UpdateServiceStatus(r.Context(), sqlc.UpdateServiceStatusParams{
		Status: "active",
		ID:     form.ServiceID,
	})
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, "/admin/manage-service", http.StatusSeeOther)
}

func (app *application) handleDeleteAccount(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	userID := r.PostForm.Get("user_id")

	deleteUserID, err := strconv.ParseInt(userID, 10, 32)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	fullName, err := app.store.DeleteAccount(r.Context(), int32(deleteUserID))
	if err != nil {
		app.serverError(w, err)
		return
	}

	isExists, err := app.store.IsUserExist(r.Context(), int32(deleteUserID))
	if err != nil {
		app.serverError(w, err)
		return
	}

	if isExists {
		app.serverError(w, errors.New("user still exists"))
		return
	}

	app.sessionManager.Put(r.Context(), "flash", fmt.Sprintf("Đã xóa tài khoản của người dùng %s ✅", fullName))

	http.Redirect(w, r, "/admin/manage-account", http.StatusSeeOther)
}
