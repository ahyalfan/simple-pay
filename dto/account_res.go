package dto

type AccountRes struct {
	UserId        int64   `json:"user_id"`
	Balance       float64 `json:"balance"`
	AccountNumber string  `json:"account_number"`
}
