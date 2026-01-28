package entity

import "time"

type Job struct {
	ID          uint64
	Name        string
	Command     string
	CronSpec    string
	Description string
	RetryCount  int
	Timeout     int
	IsActive    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
