package domain

import (
	"ahyalfan/golang_e_money/dto"
	"context"
)

type Account struct {
	ID            int64   `gorm:"primary_key;autoIncrement"`
	UserID        int64   `gorm:"column:user_id"`
	Balance       float64 `gorm:"column:balance"`
	AccountNumber string  `gorm:"column:account_number"`
}

type AccountRepository interface {
	FindByUserId(ctx context.Context, id int64) (Account, error)
	FindByAccountNumber(ctx context.Context, accountNumber string) (Account, error)
	Insert(ctx context.Context, account *Account) error
	Update(ctx context.Context, account *Account) error
}

type AccountService interface {
	Create(ctx context.Context) (dto.AccountRes, error)
}
