package repository

import (
	"ahyalfan/golang_e_money/domain"
	"context"

	"gorm.io/gorm"
)

type factorRepository struct {
	db *gorm.DB
}

func NewFactor(db *gorm.DB) domain.FactorRepository {
	return &factorRepository{db: db}
}

// FindByUserId implements domain.FactorRepository.
func (f *factorRepository) FindByUserId(ctx context.Context, userId int64) (factor domain.Factor, err error) {
	err = f.db.WithContext(ctx).First(&factor, "user_id = ?", userId).Error
	return
}

// Insert implements domain.FactorRepository.
func (f *factorRepository) Insert(ctx context.Context, factor *domain.Factor) error {
	return f.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Create(factor).Error
	})
}
