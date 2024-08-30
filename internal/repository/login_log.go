package repository

import (
	"ahyalfan/golang_e_money/domain"
	"context"

	"gorm.io/gorm"
)

type loginLogRepository struct {
	db *gorm.DB
}

func NewLoginLog(db *gorm.DB) domain.LoginLogRepository {
	return &loginLogRepository{db: db}
}

// FindLastAuthorized implements domain.LoginLogRepository.
func (l *loginLogRepository) FindLastAuthorized(ctx context.Context, userId int64) (loginLog domain.LoginLog, err error) {
	err = l.db.WithContext(ctx).Where("user_id =? AND is_authorized =?", userId, true).Last(&loginLog).Error
	return
}

// Insert implements domain.LoginLogRepository.
func (l *loginLogRepository) Insert(ctx context.Context, loginLog *domain.LoginLog) error {
	return l.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Create(loginLog).Error
	})
}
