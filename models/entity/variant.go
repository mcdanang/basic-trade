package entity

import (
	"time"
)

type Variant struct {
	ID          uint64    `form:"id" gorm:"primaryKey"`
	UUID        string    `form:"uuid" gorm:"type:varchar(191);uniqueIndex"`
	VariantName string    `form:"variant_name" gorm:"not null;unique"`
	Quantity    int       `form:"quantity" gorm:"not null"`
	ProductUUID string    `form:"product_uuid" gorm:"type:varchar(191)"`
	CreatedAt   time.Time `form:"created_at,omitempty"`
	UpdatedAt   time.Time `form:"updated_at,omitempty"`
}
