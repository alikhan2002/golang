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
	Strollers   StrollerModel
	Users       UserModel
	Tokens      TokenModel
	Permissions PermissionModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Strollers:   StrollerModel{DB: db},
		Users:       UserModel{DB: db},
		Tokens:      TokenModel{DB: db},
		Permissions: PermissionModel{DB: db},
	}
}
