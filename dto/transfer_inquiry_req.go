package dto

type TransferInQuiryReq struct {
	AccountNumber string  `json:"account_number" validate:"required,min=6,max=100"`
	Ammount       float64 `json:"ammount" validate:"required,min1"` // custom validate min1
}
