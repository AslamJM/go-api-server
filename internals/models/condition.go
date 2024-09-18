package models

import "gorm.io/gorm"

type Condition struct {
	gorm.Model

	Name        string `json:"name" gorm:"column:name;uniqueIndex"`
	Description string `json:"description" gorm:"column:description;type:text"`

	// One-to-Many relation with Treatment
	Treatments []Treatment `json:"treatments" gorm:"foreignKey:ConditionID"`
}

type CreateConditionInput struct {
	Name string `json:"name" validate:"required"`
}

func (i *CreateConditionInput) ValidateInput() error {
	return Validate.Struct(i)
}
