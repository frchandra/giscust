package models

import (
	"gorm.io/gorm"
	"time"
)

type Message struct {
	ID           int
	AppId        string
	AvatarUrl    string
	Email        string
	IsNewSession bool
	IsResolved   bool
	IsHandled    bool
	Name         string
	RoomId       string
	Source       string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
}
