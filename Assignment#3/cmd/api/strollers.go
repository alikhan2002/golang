package main

import (
	"assignment2.alikhan.net/internal/data"
	"assignment2.alikhan.net/internal/validator"
	"errors"
	"fmt"
	"net/http"
)

func (app *application) createStrollerHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title string `json:"title"`
		Price int32  `json:"price"`
		Brand string `json:"brand"`
		Color string `json:"color"`
		Ages  string `json:"ages"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	stroller := &data.Stroller{
		Title: input.Title,
		Brand: input.Brand,
		Color: input.Color,
		Ages:  input.Ages,
		Price: input.Price,
	}
	v := validator.New()
	if data.ValidateStroller(v, stroller); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	//fmt.Fprintf(w, "%+v\n", input)
	err = app.models.Strollers.Insert(stroller)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/stroolers/%d", stroller.ID))
	err = app.writeJSON(w, http.StatusCreated, envelope{"stroller": stroller}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

// BODY='{"title":"Moana","year":2016,"runtime":"107 mins", "genres":["animation","adventure"]}'
//curl.exe -i -d "{"title":"Forbaby","brand":"mercedes","price":100,"ages":"1-3","color":"blue"}" localhost:4000/v1/strollers
//set BODY={"title":"Forbaby","brand":"mercedes","price":100,"ages":"1-3","color":"blue"}

func (app *application) showStrollerHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	stroller, err := app.models.Strollers.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"stroller": stroller}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) updateStrollerHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	// Retrieve  movie record as normal.
	stroller, err := app.models.Strollers.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	// Declare an input struct to hold the expected data from the client.
	var input struct {
		Title *string `json:"title"`
		Brand *string `json:"brand"`
		Price *int32  `json:"price"`
		Color *string `json:"color"`
		Ages  *string `json:"ages"`
	}
	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	if input.Title != nil {
		stroller.Title = *input.Title
	}
	if input.Brand != nil {
		stroller.Brand = *input.Brand
	}
	if input.Price != nil {
		stroller.Price = *input.Price
	}
	if input.Color != nil {
		stroller.Color = *input.Color // Note that we don't need to dereference a slice.
	}
	if input.Ages != nil {
		stroller.Ages = *input.Ages
	}
	v := validator.New()
	if data.ValidateStroller(v, stroller); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	err = app.models.Strollers.Update(stroller)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrEditConflict):
			app.editConflictResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"stroller": stroller}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) deleteStrollerHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	err = app.models.Strollers.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"message": "stroller successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
