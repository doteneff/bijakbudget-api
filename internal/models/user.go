package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Name        string    `gorm:"size:255;not null" json:"name"`
	Email       string    `gorm:"size:100;unique;not null" json:"email"`
	Role        string    `gorm:"size:50;not null;default:'member'" json:"role"` // 'head' or 'member'
	AvatarColor int       `gorm:"not null" json:"avatarColor"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
