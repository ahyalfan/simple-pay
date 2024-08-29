package service

import (
	"ahyalfan/golang_e_money/domain"
	"ahyalfan/golang_e_money/dto"
	"ahyalfan/golang_e_money/internal/util"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userService struct {
	repo  domain.UserRepository
	cache domain.CacheRepository
	mail  domain.EmailService
	pin   domain.FactorService
}

func NewUserService(repo domain.UserRepository, cache domain.CacheRepository, mail domain.EmailService, pin domain.FactorService) domain.UserService {
	return &userService{repo: repo, cache: cache, mail: mail, pin: pin}
}

// Authenticate implements domain.UserService.
func (u *userService) Authenticate(ctx context.Context, req dto.AuthReq) (dto.AuthRes, error) {
	user, err := u.repo.FindByUsername(ctx, req.Username)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return dto.AuthRes{}, domain.ErrAuthFailed
	}
	if err != nil {
		return dto.AuthRes{}, err
	}
	if !user.EmailVerifiedAtDB.Valid {
		return dto.AuthRes{}, domain.ErrAuthFailed
	}
	// validasi use bcrypt
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return dto.AuthRes{}, domain.ErrAuthFailed
	}
	token := util.GeneratorRandomString(32)

	// ini kita simpan tokenya di sebuah cache
	// jika mau pakai jwt pun boleh aja
	jsonUser, _ := json.Marshal(user)
	_ = u.cache.Set("user:"+token, jsonUser)

	return dto.AuthRes{Token: token}, nil

}

// ValidateToken implements domain.UserService.
func (u *userService) ValidateToken(ctx context.Context, token string) (dto.UserData, error) {
	data, err := u.cache.Get("user:" + token)
	if err != nil {
		return dto.UserData{}, domain.ErrAuthFailed
	}
	var user domain.User
	err = json.Unmarshal(data, &user)
	if err != nil {
		return dto.UserData{}, domain.ErrAuthFailed
	}
	return dto.UserData{
		ID:       user.ID,
		Username: user.Username,
		FullName: user.FullName,
		Phone:    user.Phone,
	}, nil
}

// Register implements domain.UserService.
func (u *userService) Register(ctx context.Context, req dto.UserRegisterReg) (dto.UserRegisterRes, error) {
	exist, err := u.repo.FindByUsername(ctx, req.Username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return dto.UserRegisterRes{}, err
	}
	if exist.ID > 0 {
		return dto.UserRegisterRes{}, domain.ErrUserAlreadyExists
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), 12) //Ini adalah parameter yang dikenal sebagai "cost" atau "work factor" yang menentukan seberapa kuat proses hashing dilakukan.
	user := domain.User{
		Username: req.Username,
		Password: string(hashedPassword),
		FullName: req.FullName,
		Phone:    req.Phone,
		Email:    req.Email,
	}
	err = u.repo.Insert(ctx, &user)
	if err != nil {
		return dto.UserRegisterRes{}, err
	}
	err = u.pin.CreatePin(ctx, dto.CreatePin{
		UserID: user.ID,
		Pin:    req.PIN,
	})
	if err != nil {
		return dto.UserRegisterRes{}, err
	}

	otpCode := util.GeneratorRandomNumber(4)
	referenceId := util.GeneratorRandomString(32)

	err = u.mail.Send(req.Email, "E-Money OTP", "Your OTP is "+otpCode)
	if err != nil {
		log.Println("Error sending email:", err)
	}
	_ = u.cache.Set("otp:"+referenceId, []byte(otpCode))
	_ = u.cache.Set("user:"+referenceId, []byte(user.Username))
	return dto.UserRegisterRes{ReferenceID: referenceId}, nil
}

// ValidateOTP implements domain.UserService.
func (u *userService) ValidateOTP(ctx context.Context, req dto.ValidateOtpReq) error {
	val, err := u.cache.Get("otp:" + req.ReferenceID)
	if err != nil {
		return domain.ErrInvalidOTP
	}
	if string(val) != req.OTP {
		return domain.ErrInvalidOTP
	}
	username, err := u.cache.Get("user:" + req.ReferenceID)
	if err != nil {
		return domain.ErrInvalidOTP
	}
	user, err := u.repo.FindByUsername(ctx, string(username))
	if err != nil {
		return err
	}
	user.EmailVerifiedAtDB = sql.NullTime{
		Valid: true,
		Time:  time.Now(),
	}
	err = u.repo.Update(ctx, &user)
	if err != nil {
		return err
	}
	return nil
}
