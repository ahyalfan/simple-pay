package dto

type CreatePin struct {
	UserID int64  `json:"user_id"`
	Pin    string `json:"pin"`
}
