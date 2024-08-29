package domain

import (
	"ahyalfan/golang_e_money/dto"
	"context"
)

type Factor struct {
	ID     int64  `gorm:"primary_key;autoIncrement"`
	UserID int64  `gorm:"column:user_id"`
	PIN    string `gorm:"column:pin"`
}

type FactorRepository interface {
	FindByUserId(ctx context.Context, userId int64) (Factor, error)
	Insert(ctx context.Context, factor *Factor) error
}

type FactorService interface {
	Verify(ctx context.Context, req dto.ValidatePinReq) error
	CreatePin(ctx context.Context, req dto.CreatePin) error
}
