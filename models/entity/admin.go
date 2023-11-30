package entity

import (
	"basic-trade/helpers"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Admin struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	UUID      string     `json:"uuid" gorm:"type:varchar(191);uniqueIndex" `
	Name      string     `json:"name" gorm:"not null" form:"name" valid:"required~Your name is required"`
	Email     string     `json:"email" gorm:"not null;unique" form:"email" valid:"required~Your email is required, email~Invalid email format"`
	Password  string     `json:"password" gorm:"not null" form:"password" valid:"required~Your password is required, minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Products  []Product  `json:"products" gorm:"foreignKey:AdminUUID;references:UUID"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func (a *Admin) BeforeCreate(tx *gorm.DB) (err error) {

	a.UUID = uuid.NewString()

	_, errCreate := govalidator.ValidateStruct(a)

	if errCreate != nil {
		err = errCreate
		return
	}

	a.Password = helpers.HashPass(a.Password)

	err = nil
	return
}
