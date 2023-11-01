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

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.pageNotFound(w)
	})

	// Guest permissions.
	router.Handler(http.MethodGet, "/", dynamic.ThenFunc(app.home))

	router.Handler(http.MethodGet, "/service", dynamic.ThenFunc(app.displayCategoriesPage))
	router.Handler(http.MethodGet, "/service/view/:id", dynamic.ThenFunc(app.displayServiceDetailsPage))
	router.Handler(http.MethodGet, "/service/category/:slug", dynamic.ThenFunc(app.displayServicesByCategoryPage))

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

	router.Handler(http.MethodGet, "/cart", protected.ThenFunc(app.displayCart))
	router.Handler(http.MethodPost, "/cart/add", protected.ThenFunc(app.addItemToCart))
	router.Handler(http.MethodPost, "/cart/update", protected.ThenFunc(app.updateCart))
	router.Handler(http.MethodGet, "/cart/remove/:id", protected.ThenFunc(app.removeItemFromCart))

	router.Handler(http.MethodGet, "/checkout", protected.ThenFunc(app.displayCheckoutPage))
	router.Handler(http.MethodPost, "/checkout", protected.ThenFunc(app.doCheckout))

	router.Handler(http.MethodGet, "/my-orders/identity/buyer", protected.ThenFunc(app.displayPurchaseOrdersPage))
	router.Handler(http.MethodGet, "/my-orders/identity/seller", protected.ThenFunc(app.displaySellOrdersPage))

	router.Handler(http.MethodPost, "/feedback/create", protected.ThenFunc(app.createServiceFeedback))

	router.Handler(http.MethodGet, "/account/view", protected.ThenFunc(app.viewAccount))
	router.Handler(http.MethodPost, "/account/update", protected.ThenFunc(app.updateAccount))
	//router.Handler(http.MethodPost, "/account/change-password", protected.ThenFunc(app.changePasswordAccount))

	// Provider permissions.
	advanced := protected.Append(app.requireProviderPermission)

	router.Handler(http.MethodGet, "/account/my-services", advanced.ThenFunc(app.listProviderServices))

	router.Handler(http.MethodGet, "/service/create", advanced.ThenFunc(app.displayCreateServicePage))
	router.Handler(http.MethodPost, "/service/create", advanced.ThenFunc(app.doCreateService))

	router.Handler(http.MethodPost, "/order/update", advanced.ThenFunc(app.updateOrderStatus))

	standard := alice.New(app.logRequest)
	return standard.Then(router)
}
