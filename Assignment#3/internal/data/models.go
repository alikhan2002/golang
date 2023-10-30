package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

type Models struct {
	Strollers StrollerModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Strollers: StrollerModel{DB: db},
	}
}
