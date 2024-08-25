package domain

import (
	"ahyalfan/golang_e_money/dto"
	"context"
	"time"
)

type Notification struct {
	ID        int64     `gorm:"primary_key;autoIncrement"`
	UserID    int64     `gorm:"column:user_id"`
	Status    int8      `gorm:"column:status"`
	Title     string    `gorm:"column:title"`
	Body      string    `gorm:"column:body"`
	IsRead    int8      `gorm:"column:is_read"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
}

type NotificationRepository interface {
	FindByUserId(ctx context.Context, userId int64) ([]Notification, error)
	Insert(ctx context.Context, notification *Notification) error
	Update(ctx context.Context, notification *Notification) error
}

type NotificationService interface {
	Insert(ctx context.Context, userId int64, code string, data map[string]string) error
	MarkAsRead(ctx context.Context, notificationId int64) error
	FindByUser(ctx context.Context, userId int64) ([]dto.NotificationData, error)
}
