package service

import (
	"ahyalfan/golang_e_money/domain"
	"ahyalfan/golang_e_money/dto"
	"context"
	"errors"

	"gorm.io/gorm"
)

type notificationService struct {
	notificationRepository domain.NotificationRepository
}

func NewNotification(notificationRepository domain.NotificationRepository) domain.NotificationService {
	return &notificationService{notificationRepository: notificationRepository}
}

// FindByUser implements domain.NotificationService.
func (n *notificationService) FindByUser(ctx context.Context, userId int64) ([]dto.NotificationData, error) {
	notifications, err := n.notificationRepository.FindByUserId(ctx, userId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, domain.ErrNoNotificationFound
	}
	if err != nil {
		return []dto.NotificationData{}, err
	}
	var res []dto.NotificationData
	for _, notification := range notifications {
		res = append(res, dto.NotificationData{
			ID:        notification.ID,
			Status:    notification.Status,
			Title:     notification.Title,
			Body:      notification.Body,
			IsRead:    notification.IsRead,
			CreatedAt: notification.CreatedAt,
		})
	}
	return res, nil
}

// Insert implements domain.NotificationService.
func (n *notificationService) Insert(ctx context.Context, userId int64, code string, data map[string]string) error {
	panic("unimplemented")
}

// MarkAsRead implements domain.NotificationService.
func (n *notificationService) MarkAsRead(ctx context.Context, notificationId int64) error {
	// err := n.notificationRepository.Update(ctx,)
	panic("unimplemented")
}
