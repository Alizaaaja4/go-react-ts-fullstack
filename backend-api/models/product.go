package models

import "time"

type Product struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	Price     float64   `json:"price" gorm:"not null;min=0"`
	Stock     int       `json:"stock" gorm:"not null;min=0"`
	CreatedBy uint      `gorm:"not null;type:bigint unsigned"`
	User      *User     `gorm:"foreignKey:CreatedBy;references:Id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
