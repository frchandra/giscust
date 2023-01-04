package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID           int
	AppId        string
	AvatarUrl    string
	Email        string
	IsNewSession string
	IsResolved   string
	Name         string
	RoomId       string
	Source       string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
}
