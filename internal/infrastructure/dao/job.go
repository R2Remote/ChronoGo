package dao

import (
	"time"
)

type Job struct {
	ID          uint64    `gorm:"primaryKey;autoIncrement"`
	Name        string    `gorm:"column:name;type:varchar(255);not null"`
	Command     string    `gorm:"column:command;type:text;not null"`
	CronSpec    string    `gorm:"column:cron_spec;type:varchar(100);not null"`
	Description string    `gorm:"column:description;type:text"`
	RetryCount  int       `gorm:"column:retry_count;default:3"`
	Timeout     int       `gorm:"column:timeout;default:60"`
	IsActive    bool      `gorm:"column:is_active;default:true"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (Job) TableName() string {
	return "jobs"
}
