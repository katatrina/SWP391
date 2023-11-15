package main

import (
	"context"
	"net/http"
)

func (app *application) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.infoLog.Printf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())
		next.ServeHTTP(w, r)
	})
}

func (app *application) requireAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !app.isAuthenticated(r) {
			app.sessionManager.Put(r.Context(), "redirectPathAfterLogin", r.URL.Path)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		w.Header().Add("Cache-Control", "no-store")

		next.ServeHTTP(w, r)
	})
}

func (app *application) authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		adminID := app.sessionManager.GetInt32(r.Context(), "authenticatedAdminID")

		if adminID != 0 {
			isAdmin, err := app.store.IsAdminByID(r.Context(), adminID)
			if err != nil {
				app.serverError(w, err)
				return
			}

			if isAdmin {
				ctx := context.WithValue(r.Context(), isAdminContextKey, true)
				r = r.WithContext(ctx)
			}

			next.ServeHTTP(w, r)

			return
		}

		id := app.sessionManager.GetInt32(r.Context(), "authenticatedUserID")
		if id == 0 {
			next.ServeHTTP(w, r)
			return
		}

		exists, err := app.store.IsUserExist(r.Context(), id)
		if err != nil {
			app.serverError(w, err)
			return
		}

		if exists {
			ctx := context.WithValue(r.Context(), isAuthenticatedContextKey, true)
			r = r.WithContext(ctx)
		}

		isProvider, err := app.store.IsProvider(r.Context(), id)
		if err != nil {
			app.serverError(w, err)
			return
		}

		if isProvider {
			ctx := context.WithValue(r.Context(), isProviderContextKey, true)
			r = r.WithContext(ctx)
		}

		next.ServeHTTP(w, r)
	})
}

func (app *application) authenticateAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := app.sessionManager.GetInt32(r.Context(), "authenticatedAdminID")
		if id == 0 {
			next.ServeHTTP(w, r)
			return
		}

		isAdmin, err := app.store.IsAdminByID(r.Context(), id)
		if err != nil {
			app.serverError(w, err)
			return
		}

		if isAdmin {
			ctx := context.WithValue(r.Context(), isAdminContextKey, true)
			r = r.WithContext(ctx)
		}

		next.ServeHTTP(w, r)
	})
}

func (app *application) requireProviderPermission(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !app.isProvider(r) {
			http.Redirect(w, r, "/404", http.StatusSeeOther)
			return
		}

		w.Header().Add("Cache-Control", "no-store")

		next.ServeHTTP(w, r)
	})
}

func (app *application) requireAdminPermission(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !app.isAdmin(r) {
			http.Redirect(w, r, "/404", http.StatusSeeOther)
			return
		}

		w.Header().Add("Cache-Control", "no-store")

		next.ServeHTTP(w, r)
	})
}
