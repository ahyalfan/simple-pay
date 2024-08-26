package service

import (
	"ahyalfan/golang_e_money/domain"
	"ahyalfan/golang_e_money/dto"
	"bytes"
	"context"
	"errors"
	"html/template"

	"gorm.io/gorm"
)

type notificationService struct {
	notificationRepository domain.NotificationRepository
	templateRepository     domain.TemplateRepository
	hub                    *dto.Hub
}

func NewNotification(notificationRepository domain.NotificationRepository,
	templateRepository domain.TemplateRepository, hub *dto.Hub) domain.NotificationService {
	return &notificationService{notificationRepository: notificationRepository, templateRepository: templateRepository, hub: hub}
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
	tmpl, err := n.templateRepository.FindByCode(ctx, code)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("template not found")
	}
	if err != nil {
		return err
	}

	body := new(bytes.Buffer)
	t := template.Must(template.New("notif").Parse(tmpl.Body))
	err = t.Execute(body, data)
	if err != nil {
		return err
	}

	notification := domain.Notification{
		UserID: userId,
		Title:  tmpl.Title,
		Body:   body.String(),
		Status: 1,
		IsRead: 0,
	}
	err = n.notificationRepository.Insert(ctx, &notification)
	if err != nil {
		return err
	}

	if channel, ok := n.hub.NotificationChannel[userId]; ok {
		channel <- dto.NotificationData{
			ID:        notification.ID,
			Title:     notification.Title,
			Body:      notification.Body,
			Status:    notification.Status,
			IsRead:    notification.IsRead,
			CreatedAt: notification.CreatedAt,
		}
	}

	return nil
}

// MarkAsRead implements domain.NotificationService.
func (n *notificationService) MarkAsRead(ctx context.Context, notificationId int64) error {
	// err := n.notificationRepository.Update(ctx,)
	panic("unimplemented")
}
