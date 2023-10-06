package main

import (
	"assignment2.alikhan.net/internal/data"
	"fmt"
	"net/http"
	"time"
)

// Add a createMovieHandler for the "POST /v1/movies" endpoint. For now we simply
// return a plain-text placeholder response.

func (app *application) createStrollerHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title string `json:"title"`
		Brand string `json:"brand"`
		Color string `json:"color"`
		Ages  string `json:"ages"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	// Copy the values from the input struct to a new Movie struct.
	stroller := &data.Stroller{
		Title: input.Title,
		Brand: input.Brand,
		Color: input.Color,
		Ages:  input.Ages,
	}
	// Initialize a new Validator.
	//v := validator.New()
	// Call the ValidateMovie() function and return a response containing the errors if
	// any of the checks fail.
	//if data.ValidateStroller(v, stroller); !v.Valid() {
	//	app.failedValidationResponse(w, r, v.Errors)
	//	return
	//}
	fmt.Fprintf(w, "%+v\n", input)
}

// Add a showMovieHandler for the "GET /v1/movies/:id" endpoint. For now, we retrieve
// the interpolated "id" parameter from the current URL and include it in a placeholder
// response.
func (app *application) showStrollerHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		// Use the new notFoundResponse() helper.
		app.notFoundResponse(w, r)
		return
	}
	stroller := data.Stroller{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "CoolShare Baby Stroller for Toddler",
		Brand:     "CoolShare",
		Color:     "Gray",
		Ages:      "1-3",
		Version:   1,
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"stroller": stroller}, nil)
	if err != nil {
		// Use the new serverErrorResponse() helper.
		app.serverErrorResponse(w, r, err)
	}
}
