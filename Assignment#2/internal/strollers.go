package data

import (
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
