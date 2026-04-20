package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID           string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Name         string    `gorm:"size:255;not null" json:"name"`
	Icon         string    `gorm:"size:100" json:"icon"`   // Could store the codepoint or name
	Color        string    `gorm:"size:20" json:"color"`   // Hex code
	MonthlyLimit float64   `gorm:"type:decimal(10,2);not null;default:0" json:"monthlyLimit"`
	IsIncome     bool      `gorm:"default:false" json:"isIncome"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}
