package repository

import (
	"ahyalfan/golang_e_money/domain"
	"context"

	"gorm.io/gorm"
)

type topupRepository struct {
	db *gorm.DB
}

func NewTopup(db *gorm.DB) domain.TopupRpository {
	return &topupRepository{db: db}
}

// FindById implements domain.TopupRpository.
func (t *topupRepository) FindById(ctx context.Context, id string) (topup domain.Topup, err error) {
	err = t.db.WithContext(ctx).Where("id =?", id).First(&topup).Error
	return
}

// Insert implements domain.TopupRpository.
func (t *topupRepository) Insert(ctx context.Context, topup *domain.Topup) (string, error) {
	err := t.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Create(topup).Error
	})
	if err != nil {
		return "", err
	}
	return topup.ID, nil
}

// Update implements domain.TopupRpository.
func (t *topupRepository) Update(ctx context.Context, topup *domain.Topup) error {
	return t.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Save(topup).Error
	})
}
