package entity

import "time"

type Server struct {
	ID          uint64
	Name        string
	Description string
	IP          string
	Port        int
	IsActive    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
