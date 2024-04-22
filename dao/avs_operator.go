package dao

import (
	"github.com/lib/pq"
)

type AvsOperator struct {
	ID                    string         `json:"id" gorm:"column:id"`
	Avs                   string         `json:"avs" gorm:"primaryKey;column:avs"`
	Operator              string         `json:"operator" gorm:"primaryKey;column:operator"`
	RestakeableStrategies pq.StringArray `gorm:"column:restakeable_strategies;type:text[]" json:"restakeable_strategies"`
	RegistrationStatus    string         `json:"registration_status" gorm:"column:registration_status"`
}
