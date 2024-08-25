package dto

import "time"

type NotificationData struct {
	ID        int64     `json:"id"`
	Status    int8      `json:"status"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	IsRead    int8      `json:"is_read"`
	CreatedAt time.Time `json:"created_at"`
}
