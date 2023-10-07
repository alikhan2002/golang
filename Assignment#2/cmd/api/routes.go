package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/strollers", app.createStrollerHandler)
	router.HandlerFunc(http.MethodGet, "/v1/strollers/:id", app.showStrollerHandler)
	return router
}
