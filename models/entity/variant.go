package entity

import "time"

type Variant struct {
	ID   uint64 `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"not null;unique"`
	// OrderedAt    string `json:"ordered_at" gorm:"not null"`
	// Items        []Item
	CreatedAt time.Time
	UpdatedAt time.Time
}
