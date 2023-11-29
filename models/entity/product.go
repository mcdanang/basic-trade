package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID        uint64    `json:"id" gorm:"primaryKey"`
	UUID      string    `json:"uuid" gorm:"type:varchar(191);uniqueIndex"`
	Name      string    `json:"name" gorm:"not null;unique"`
	ImageURL  string    `json:"image_url"`
	AdminUUID string    `json:"admin_uuid" form:"admin_uuid" gorm:"type:varchar(191);constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Variants  []Variant `json:"variants" gorm:"foreignKey:ProductUUID;references:UUID"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {

	p.UUID = uuid.NewString()

	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
