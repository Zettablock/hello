package dao

import (
	"context"
	"fmt"
	"github.com/lib/pq"
	"gorm.io/gorm"
	"math/rand"
)

type Avs struct {
	ID                    string         `json:"id" gorm:"primaryKey;column:id"`
	AvsList               string         `json:"avs_list" gorm:"column:avs_list"`
	RestakeableStrategies pq.StringArray `gorm:"column:restakeable_strategies;type:text[]" json:"restakeable_strategies"`
	MetadataUri           string         `json:"metadata_uri" gorm:"column:metadata_uri"`
}

func (m *Avs) TableName() string {
	return "public.avs"
}

type AvsDao struct {
	sourceDB  *gorm.DB
	replicaDB []*gorm.DB
	m         *Avs
}

func NewAvsDao(ctx context.Context, dbs ...*gorm.DB) *AvsDao {
	dao := new(AvsDao)
	switch len(dbs) {
	case 0:
		panic("database connection required")
	case 1:
		dao.sourceDB = dbs[0]
		dao.replicaDB = []*gorm.DB{dbs[0]}
	default:
		dao.sourceDB = dbs[0]
		dao.replicaDB = dbs[1:]
	}
	return dao
}

func (d *AvsDao) DB() *gorm.DB {
	return d.sourceDB
}

func (d *AvsDao) Upsert(ctx context.Context, obj *Avs) error {
	err := d.sourceDB.Save(&obj).Error
	if err != nil {
		return fmt.Errorf("AvsDao: %w", err)
	}
	return nil
}

func (d *AvsDao) Create(ctx context.Context, obj *Avs) error {
	err := d.sourceDB.Model(d.m).Create(&obj).Error
	if err != nil {
		return fmt.Errorf("AvsDao: %w", err)
	}
	return nil
}

func (d *AvsDao) Get(ctx context.Context, where string) (*Avs, error) {
	items, err := d.List(ctx, where, 1)
	if err != nil {
		return nil, fmt.Errorf("AvsDao: Get where=%s: %w", where, err)
	}
	if len(items) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &items[0], nil
}

func (d *AvsDao) List(ctx context.Context, where string, limit int) ([]Avs, error) {
	var results []Avs
	err := d.replicaDB[rand.Intn(len(d.replicaDB))].Model(d.m).Where(where).Limit(limit).Find(&results).Error
	if err != nil {
		return nil, fmt.Errorf("AvsDao: List where=%s: %w", where, err)
	}
	return results, nil
}

func (d *AvsDao) Update(ctx context.Context, where string, update map[string]interface{}, args ...interface{}) error {
	err := d.sourceDB.Model(d.m).Where(where, args...).
		Updates(update).Error
	if err != nil {
		return fmt.Errorf("AvsDao:Update where=%s: %w", where, err)
	}
	return nil
}

func (d *AvsDao) Delete(ctx context.Context, where string, args ...interface{}) error {
	if len(where) == 0 {
		return fmt.Errorf("AvsDao: Delete where=%s", where)
	}
	if err := d.sourceDB.Where(where, args...).Delete(d.m).Error; err != nil {
		return fmt.Errorf("AvsDao: Delete where=%s: %w", where, err)
	}
	return nil
}
