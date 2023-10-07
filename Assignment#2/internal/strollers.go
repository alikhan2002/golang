package data

import (
	"time"
)

type Stroller struct {
	ID        int64
	CreatedAt time.Time
	Title     string
	Brand     string
	Ages      string
	Color     string
	Version   int32
}
