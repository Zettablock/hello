package dao

import (
	"github.com/lib/pq"
)

type Avs struct {
	ID                    string         `json:"id" gorm:"primaryKey;column:id"`
	AvsList               string         `json:"avs_list" gorm:"column:avs_list"`
	RestakeableStrategies pq.StringArray `gorm:"column:restakeable_strategies;type:text[]" json:"restakeable_strategies"`
	MetadataUri           string         `json:"metadata_uri" gorm:"column:metadata_uri"`
}
