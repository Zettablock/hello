package dao

import (
	"context"
	"fmt"
	"github.com/lib/pq"
	"gorm.io/gorm"
	"math/rand"
)

type AvsOperator struct {
	ID                    string         `json:"id" gorm:"column:id"`
	Avs                   string         `json:"avs" gorm:"primaryKey;column:avs"`
	Operator              string         `json:"operator" gorm:"primaryKey;column:operator"`
	RestakeableStrategies pq.StringArray `gorm:"column:restakeable_strategies;type:text[]" json:"restakeable_strategies"`
	RegistrationStatus    string         `json:"registration_status" gorm:"column:registration_status"`
}

func (m *AvsOperator) TableName() string {
	return "public.avs_operator"
}

type AvsOperatorDao struct {
	sourceDB  *gorm.DB
	replicaDB []*gorm.DB
	m         *AvsOperator
}

func NewAvsOperatorDao(ctx context.Context, dbs ...*gorm.DB) *AvsOperatorDao {
	dao := new(AvsOperatorDao)
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

func (d *AvsOperatorDao) DB() *gorm.DB {
	return d.sourceDB
}

func (d *AvsOperatorDao) Create(ctx context.Context, obj *AvsOperator) error {
	err := d.sourceDB.Model(d.m).Create(&obj).Error
	if err != nil {
		return fmt.Errorf("AvsOperatorDao: %w", err)
	}
	return nil
}

func (d *AvsOperatorDao) Get(ctx context.Context, fields, where string) (*AvsOperator, error) {
	items, err := d.List(ctx, fields, where, 0, 1)
	if err != nil {
		return nil, fmt.Errorf("AvsOperatorDao: Get where=%s: %w", where, err)
	}
	if len(items) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &items[0], nil
}

func (d *AvsOperatorDao) List(ctx context.Context, fields, where string, offset, limit int) ([]AvsOperator, error) {
	var results []AvsOperator
	err := d.replicaDB[rand.Intn(len(d.replicaDB))].Model(d.m).
		Select(fields).Where(where).Offset(offset).Limit(limit).Find(&results).Error
	if err != nil {
		return nil, fmt.Errorf("AvsOperatorDao: List where=%s: %w", where, err)
	}
	return results, nil
}

func (d *AvsOperatorDao) Update(ctx context.Context, where string, update map[string]interface{}, args ...interface{}) error {
	err := d.sourceDB.Model(d.m).Where(where, args...).
		Updates(update).Error
	if err != nil {
		return fmt.Errorf("AvsOperatorDao:Update where=%s: %w", where, err)
	}
	return nil
}

func (d *AvsOperatorDao) Delete(ctx context.Context, where string, args ...interface{}) error {
	if len(where) == 0 {
		return fmt.Errorf("AvsOperatorDao: Delete where=%s", where)
	}
	if err := d.sourceDB.Where(where, args...).Delete(d.m).Error; err != nil {
		return fmt.Errorf("AvsOperatorDao: Delete where=%s: %w", where, err)
	}
	return nil
}
