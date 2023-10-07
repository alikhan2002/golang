package main

import (
	data "assignment2.alikhan.net/internal"
	"fmt"
	"net/http"
	"time"
)

func (app *application) createStrollerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new movie")
}
func (app *application) showStrollerHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
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
	err = app.writeJSON(w, http.StatusOK, stroller, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}
