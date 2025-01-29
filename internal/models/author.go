package models

import "time"

type Author struct {
	ID        uint64
	UserID    uint64
	Name      string
	Bio       string
	CreatedAt time.Time
	UpdatedAt time.Time
}
