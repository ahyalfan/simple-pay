package service

import (
	"ahyalfan/golang_e_money/domain"
	"ahyalfan/golang_e_money/dto"
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type topupService struct {
	notificationService   domain.NotificationService
	midtransService       domain.MidtransService
	topupRepository       domain.TopupRpository
	accountRepository     domain.AccountRepository
	transactionRepository domain.TransactionRepository
}

func NewTopupService(notificationService domain.NotificationService,
	topupRepository domain.TopupRpository,
	midtransService domain.MidtransService,
	accountRepository domain.AccountRepository,
	transactionRepository domain.TransactionRepository) domain.TopupService {
	return &topupService{
		notificationService:   notificationService,
		topupRepository:       topupRepository,
		midtransService:       midtransService,
		accountRepository:     accountRepository,
		transactionRepository: transactionRepository,
	}
}

// ConfimedTopup implements domain.topupService.
func (t *topupService) ConfimedTopup(ctx context.Context, id string) error {
	topup, err := t.topupRepository.FindById(ctx, id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("topup request not found")
	}
	if err != nil {
		return err
	}
	account, err := t.accountRepository.FindByUserId(ctx, topup.UserId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("user id not found")
	}
	if err != nil {
		return err
	}

	err = t.transactionRepository.Insert(ctx, &domain.Transaction{
		AccountId:       account.ID,
		SofNumber:       "00",
		DofNumber:       account.AccountNumber,
		TransactionType: "C",
		Amount:          topup.Amount,
	})
	if err != nil {
		return err
	}

	account.Balance += topup.Amount
	err = t.accountRepository.Update(ctx, &account)
	if err != nil {
		return err
	}

	data := map[string]string{
		"amount": fmt.Sprintf("%.2f", topup.Amount),
	}

	_ = t.notificationService.Insert(ctx, account.UserID, "TOPUP_SUCCESS", data)
	return nil
}

// InitializeTopup implements domain.TopupService.
func (t *topupService) InitializeTopup(ctx context.Context, req dto.TopupReq) (dto.TopupRes, error) {
	topup := domain.Topup{
		ID:     uuid.NewString(),
		UserId: req.UserId,
		Amount: req.Amount,
		Status: 0,
	}
	err := t.midtransService.GenerateSnapURL(ctx, &topup)
	if err != nil {
		return dto.TopupRes{}, err
	}

	id, err := t.topupRepository.Insert(ctx, &topup)
	if err != nil {
		return dto.TopupRes{}, err
	}
	return dto.TopupRes{
		Id:      id,
		SnapUrl: topup.SnapLink,
	}, nil
}
