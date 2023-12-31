package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	router.HandlerFunc(http.MethodPost, "/v1/strollers", app.createStrollerHandler)
	router.HandlerFunc(http.MethodGet, "/v1/strollers", app.listStrollerHandler)
	router.HandlerFunc(http.MethodGet, "/v1/strollers/:id", app.showStrollerHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/strollers/:id", app.updateStrollerHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/stroller/:id", app.deleteStrollerHandler)

	return app.recoverPanic(app.rateLimit(router))
}
