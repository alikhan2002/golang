package main

import (
	"assignment2.alikhan.net/internal/data"
	"assignment2.alikhan.net/internal/validator"
	"fmt"
	"net/http"
	"time"
)

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
	v := validator.New()
	if data.ValidateStroller(v, stroller); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) showStrollerHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
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
		app.serverErrorResponse(w, r, err)
	}
}
