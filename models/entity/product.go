package entity

import (
	"time"
)

type Product struct {
	ID        uint64 `json:"id" gorm:"primaryKey"`
	UUID      string `json:"uuid" gorm:"type:varchar(191);uniqueIndex"`
	Name      string `json:"name" gorm:"not null;unique"`
	ImageURL  string `json:"image_url"`
	AdminID   uint
	Variants  []Variant `gorm:"foreignKey:ProductUUID;references:UUID"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
