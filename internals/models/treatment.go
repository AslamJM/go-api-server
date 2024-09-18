package models

import "gorm.io/gorm"

type Currency string

const (
	USD Currency = "USD"
	EUR Currency = "EUR"
	GBP Currency = "GBP"
)

type DosageSchedule string

const (
	Daily   DosageSchedule = "Daily"
	Weekly  DosageSchedule = "Weekly"
	Monthly DosageSchedule = "Monthly"
)

type Treatment struct {
	gorm.Model

	Name           string         `json:"name" gorm:"column:name"`
	IsNewTreatment bool           `json:"is_new_treatment" gorm:"column:is_new_treatment"`
	Unit           string         `json:"unit" gorm:"column:unit"`
	Currency       Currency       `json:"currency" gorm:"column:currency;type:enum('USD', 'EUR', 'GBP');default:'USD'"`
	Dosage         float64        `json:"dosage" gorm:"column:dosage"`
	DosageSchedule DosageSchedule `json:"dosage_schedule" gorm:"column:dosage_schedule;type:enum('Daily', 'Weekly', 'Monthly')"`

	// Many-to-One relation with Condition
	ConditionID uint      `json:"condition_id" gorm:"column:condition_id"`
	Condition   Condition `json:"condition" gorm:"foreignKey:ConditionID"`
}
