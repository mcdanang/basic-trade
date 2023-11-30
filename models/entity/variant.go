package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Variant struct {
	ID          uint64    `json:"id" form:"id" gorm:"primaryKey"`
	UUID        string    `json:"uuid" form:"uuid" gorm:"type:varchar(191);uniqueIndex"`
	VariantName string    `json:"variant_name" form:"variant_name" gorm:"not null;unique" valid:"required~variant name is required"`
	Quantity    int       `json:"quantity" form:"quantity" gorm:"not null" valid:"required~quantity is required"`
	ProductUUID string    `json:"product_uuid" form:"product_uuid" gorm:"type:varchar(191);constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt   time.Time `json:"created_at" form:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at" form:"updated_at,omitempty"`
}

func (v *Variant) BeforeCreate(tx *gorm.DB) (err error) {

	v.UUID = uuid.NewString()

	_, errCreate := govalidator.ValidateStruct(v)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
