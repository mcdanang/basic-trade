package entity

import (
	"time"
)

type Variant struct {
	ID          uint64    `json:"id" form:"id" gorm:"primaryKey"`
	UUID        string    `json:"uuid" form:"uuid" gorm:"type:varchar(191);uniqueIndex"`
	VariantName string    `json:"variant_name" form:"variant_name" gorm:"not null;unique"`
	Quantity    int       `json:"quantity" form:"quantity" gorm:"not null"`
	ProductUUID string    `json:"product_uuid" form:"product_uuid" gorm:"type:varchar(191);constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt   time.Time `json:"created_at" form:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at" form:"updated_at,omitempty"`
}
