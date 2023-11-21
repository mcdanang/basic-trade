package entity

import "time"

type Product struct {
	ID       uint64 `json:"id" gorm:"primaryKey"`
	Name     string `json:"name" gorm:"not null;unique"`
	ImageURL string `json:"image_url"`
	// AdminID  uint
	// Items        []Item
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
