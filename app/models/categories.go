package models

import "time"

type Category struct {
	ID        int
	Name      string
	Color     string
	CreatedAt time.Time
}

