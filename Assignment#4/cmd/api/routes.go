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

	router.HandlerFunc(http.MethodGet, "/v1/strollers", app.requirePermission("strollers:read", app.listStrollerHandler))
	router.HandlerFunc(http.MethodPost, "/v1/strollers", app.requirePermission("strollers:write", app.createStrollerHandler))
	router.HandlerFunc(http.MethodGet, "/v1/strollers/:id", app.requirePermission("strollers:read", app.showStrollerHandler))
	router.HandlerFunc(http.MethodPatch, "/v1/strollers/:id", app.requirePermission("strollers:write", app.updateStrollerHandler))
	router.HandlerFunc(http.MethodDelete, "/v1/strollers/:id", app.requirePermission("strollers:write", app.deleteStrollerHandler))

	router.HandlerFunc(http.MethodPost, "/v1/users", app.registerUserHandler)

	router.HandlerFunc(http.MethodPut, "/v1/users/activated", app.activateUserHandler)

	router.HandlerFunc(http.MethodPost, "/v1/tokens/authentication", app.createAuthenticationTokenHandler)
	return app.recoverPanic(app.rateLimit(app.authenticate(router)))
}

//-- Set the activated field for alice@example.com to true.
//UPDATE users SET activated = true WHERE email = 'alice@example.com';
//-- Give all users the 'strollers:read' permission
//INSERT INTO users_permissions
//SELECT id, (SELECT id FROM permissions WHERE code = 'strollers:read') FROM users;
//-- Give faith@example.com the 'strollers:write' permission
//INSERT INTO users_permissions
//VALUES (
//(SELECT id FROM users WHERE email = 'faith@example.com'),
//(SELECT id FROM permissions WHERE code = 'strollers:write')
//);
//-- List all activated users and their permissions.
//SELECT email, array_agg(permissions.code) as permissions
//FROM permissions
//INNER JOIN users_permissions ON users_permissions.permission_id = permissions.id
//INNER JOIN users ON users_permissions.user_id = users.id
//WHERE users.activated = true
//GROUP BY email;
