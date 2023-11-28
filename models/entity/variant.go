package entity

import (
	"time"

	"github.com/google/uuid"
)

type Variant struct {
	ID          uint64    `form:"id" gorm:"primaryKey"`
	UUID        uuid.UUID `form:"uuid" gorm:"type:varchar(100);uniqueIndex"`
	VariantName string    `form:"variant_name" gorm:"not null;unique"`
	Quantity    int       `form:"quantity" gorm:"not null"`
	ProductUUID string    `form:"product_id" gorm:"type:varchar(100)"`
	CreatedAt   time.Time `form:"created_at,omitempty"`
	UpdatedAt   time.Time `form:"updated_at,omitempty"`
}
