package domain

import (
	"ahyalfan/golang_e_money/dto"
	"context"
	"database/sql"
)

type User struct {
	// jika mau bikin uuid juga boleh
	ID                int64        `gorm:"primary_key;autoIncrement"`
	FullName          string       `gorm:"column:full_name"`
	Phone             string       `gorm:"column:phone"`
	Email             string       `gorm:"column:email"`
	Username          string       `gorm:"column:username"`
	Password          string       `gorm:"column:password"`
	EmailVerifiedAtDB sql.NullTime `gorm:"column:email_verified_at"`
	// EmailVerifiedAt   time.Time    `gorm:"-"`
}

type UserRepository interface {
	FindById(ctx context.Context, id int64) (User, error)
	FindByUsername(ctx context.Context, username string) (User, error)
	Insert(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
}

type UserService interface {
	Authenticate(ctx context.Context, req dto.AuthReq) (dto.AuthRes, error)
	ValidateToken(ctx context.Context, token string) (dto.UserData, error)
	Register(ctx context.Context, req dto.UserRegisterReg) (dto.UserRegisterRes, error)
	ValidateOTP(ctx context.Context, req dto.ValidateOtpReq) error
}
