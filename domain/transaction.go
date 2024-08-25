package domain

import (
	"ahyalfan/golang_e_money/dto"
	"context"
	"time"
)

type Transaction struct {
	ID                  int64     `gorm:"primary_key;autoIncrement"`
	SofNumber           string    `gorm:"column:sof_number"` //dari mana uangnya
	DofNumber           string    `gorm:"column:dof_number"` // mau kemana uangnya
	Amount              float64   `gorm:"column:amount"`
	TransactionType     string    `gorm:"column:transaction_type"`
	AccountId           int64     `gorm:"column:account_id"`
	TransactionDatetime time.Time `gorm:"column:transaction_datetime;autoCreateTime"`
}

type TransactionRepository interface {
	FindBySofNumber(ctx context.Context, sofNumber string) (Transaction, error)
	FindByDofNumber(ctx context.Context, dofNumber string) (Transaction, error)
	FindByAccountId(ctx context.Context, accountId int64) ([]Transaction, error)
	Insert(ctx context.Context, transaction *Transaction) error
}

type TransactionService interface {
	TransferInquiry(ctx context.Context, req dto.TransferInQuiryReq) (dto.TransferInQuiryRes, error)
	TransferExecute(ctx context.Context, req dto.TransferExecuteReq) error
}
