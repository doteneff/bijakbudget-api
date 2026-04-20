package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID         string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Title      string    `gorm:"size:255;not null" json:"title"`
	Amount     float64   `gorm:"type:decimal(10,2);not null" json:"amount"`
	Type       string    `gorm:"size:20;not null" json:"type"` // 'income' or 'expense'
	CategoryID string    `gorm:"type:uuid;not null" json:"categoryId"`
	Category   Category  `gorm:"foreignKey:CategoryID" json:"category"`
	Date       time.Time `gorm:"not null" json:"date"`
	Note       *string   `gorm:"type:text" json:"note"`
	MemberID   string    `gorm:"type:uuid;not null" json:"memberId"`
	Member     User      `gorm:"foreignKey:MemberID" json:"member"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}
