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
	Users     UserModel  // Add a new Users field.
	Tokens    TokenModel // Add a new Tokens field.

}

func NewModels(db *sql.DB) Models {
	return Models{
		Strollers: StrollerModel{DB: db},
		Users:     UserModel{DB: db},  // Initialize a new UserModel instance.
		Tokens:    TokenModel{DB: db}, // Initialize a new TokenModel instance.

	}
}
