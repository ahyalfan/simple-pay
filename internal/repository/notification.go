package repository

import (
	"ahyalfan/golang_e_money/domain"
	"context"

	"gorm.io/gorm"
)

type notificationRepository struct {
	db *gorm.DB
}

func NewNotification(db *gorm.DB) domain.NotificationRepository {
	return &notificationRepository{db: db}
}

// FindByUserId implements domain.NotificationRepository.
func (n *notificationRepository) FindByUserId(ctx context.Context, userId int64) (notifications []domain.Notification, err error) {
	err = n.db.WithContext(ctx).Find(&notifications, "user_id = ?", userId).Order("id desc").Limit(15).Error
	return
}

// Insert implements domain.NotificationRepository.
func (n *notificationRepository) Insert(ctx context.Context, notification *domain.Notification) error {
	return n.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Create(notification).Error
	})
}

// Update implements domain.NotificationRepository.
func (n *notificationRepository) Update(ctx context.Context, notification *domain.Notification) error {
	return n.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Save(notification).Error
	})
}
