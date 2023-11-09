package main

type contextKey string

const isAuthenticatedContextKey = contextKey("isAuthenticated")

const isProviderContextKey = contextKey("isProvider")

const isAdminContextKey = contextKey("isAdmin")
