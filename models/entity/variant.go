package entity

import "time"

type Variant struct {
	ID          uint64 `json:"id" gorm:"primaryKey"`
	UUID        string `json:"uuid" gorm:"not null"`
	VariantName string `json:"variant_name" gorm:"not null;unique"`
	Quantity    int    `json:"quantity" gorm:"not null"`
	ProductID   string
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}
