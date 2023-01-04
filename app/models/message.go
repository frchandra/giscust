package models

import (
	"gorm.io/gorm"
	"time"
)

type Message struct {
	ID           int
	TestTest     string
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
