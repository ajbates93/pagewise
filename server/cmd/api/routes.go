package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	// Waitlist
	router.HandlerFunc(http.MethodPost, "/v1/waitlist", app.registerWaitlistUser)

	// Google Books API
	router.HandlerFunc(http.MethodGet, "/v1/search", app.searchBooksHandler)

	return app.enableCORS(router)
}
