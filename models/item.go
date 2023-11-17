package models

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Item struct {
	ID          uint64 `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"not null;unique"`
	Description string `json:"description" gorm:"not null"`
	Quantity    int    `json:"quantity" gorm:"not null"`
	OrderID     uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (i *Item) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("Item before create()")

	if len(i.Name) < 4 {
		err = errors.New("Item name is too short")
	}

	return
}
