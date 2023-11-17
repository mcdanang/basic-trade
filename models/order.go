package models

import "time"

type Order struct {
	ID           uint64 `json:"id" gorm:"primaryKey"`
	CustomerName string `json:"customer_name" gorm:"not null;unique"`
	OrderedAt    string `json:"ordered_at" gorm:"not null"`
	Items        []Item
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
