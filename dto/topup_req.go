package dto

type TopupReq struct {
	Amount float64 `json:"amount" validate:"required"`
	UserId int64   `json:"-"`
}
