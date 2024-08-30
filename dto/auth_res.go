package dto

type AuthRes struct {
	Token  string `json:"token"`
	UserID int64  `json:"user_id"`
}
