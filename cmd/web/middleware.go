package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// adds CSRF protection to all POST requests
func NoSruve(next http.Handler) http.Handler {
	csrHandler := nosurf.New(next)

	csrHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrHandler
}

// load and save session for every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
