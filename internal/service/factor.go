package service

import (
	"ahyalfan/golang_e_money/domain"
	"ahyalfan/golang_e_money/dto"
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type factorService struct {
	factorRepo domain.FactorRepository
}

func NewFactor(factorRepo domain.FactorRepository) domain.FactorService {
	return &factorService{factorRepo: factorRepo}
}

// Verify implements domain.FactorService.
func (f *factorService) Verify(ctx context.Context, req dto.ValidatePinReq) error {
	factor, err := f.factorRepo.FindByUserId(ctx, req.UserId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return domain.ErrInvalidPin
	}
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(factor.PIN), []byte(req.Pin))
	if err != nil {
		return domain.ErrInvalidPin
	}
	return nil
}

// Generate implements domain.FactorService.
func (f *factorService) CreatePin(ctx context.Context, req dto.CreatePin) error {
	exists, err := f.factorRepo.FindByUserId(ctx, req.UserID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if exists.ID > 0 {
		return domain.ErrPinAlreadyExists
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte(req.Pin), 12)
	factor := domain.Factor{
		UserID: req.UserID,
		PIN:    string(hash),
	}
	err = f.factorRepo.Insert(ctx, &factor)
	return err

}
