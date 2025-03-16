package model

import (
	"time"
)

type Message struct {
	ID        uint      `gorm:"primaryKey"`
	User      string    `gorm:"size:255"`
	Content   string    `gorm:"size:255"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}