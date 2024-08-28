package domain

import (
	"ahyalfan/golang_e_money/dto"
	"context"
	"time"
)

type Topup struct {
	ID        string    `gorm:"primary_key"`
	UserId    int64     `gorm:"column:user_id"`
	Amount    float64   `gorm:"column:amount"`
	Status    int8      `gorm:"column:status"`
	SnapLink  string    `gorm:"column:snap_url"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
}

type TopupRpository interface {
	FindById(ctx context.Context, id string) (Topup, error)
	Insert(ctx context.Context, topup *Topup) (string, error)
	Update(ctx context.Context, topup *Topup) error
}

type TopupService interface {
	ConfimedTopup(ctx context.Context, id string) error
	InitializeTopup(ctx context.Context, req dto.TopupReq) (dto.TopupRes, error)
}
