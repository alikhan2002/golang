package data

import (
	"assignment2.alikhan.net/internal/validator"
	"time"
)

type Stroller struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Brand     string    `json:"brand"`
	Color     string    `json:"color"`
	Ages      string    `json.:"ages"`
	Version   int32     `json:"version"`
}

func ValidateStroller(v *validator.Validator, stroller *Stroller) {
	v.Check(stroller.Title != "", "title", "must be provided")
	v.Check(len(stroller.Title) <= 500, "title", "must not be more than 500 bytes long")
	v.Check(stroller.Brand != "", "brand", "must be provided")
	v.Check(len(stroller.Brand) <= 500, "brand", "must not be more than 500 bytes long")
	v.Check(stroller.Color != "", "color", "must be provided")
	v.Check(len(stroller.Brand) <= 100, "color", "must not be more than 100 bytes long")
	v.Check(stroller.Ages != "", "ages", "must be provided")
	v.Check(len(stroller.Ages) <= 10, "ages", "must not be more than 10 bytes long")
}
