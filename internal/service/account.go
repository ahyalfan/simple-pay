package service

import (
	"ahyalfan/golang_e_money/domain"
	"ahyalfan/golang_e_money/dto"
	"ahyalfan/golang_e_money/internal/util"
	"context"
)

type accountService struct {
	accountRepository domain.AccountRepository
}

func NewAccount(accountRepository domain.AccountRepository) domain.AccountService {
	return &accountService{accountRepository: accountRepository}
}

// Create implements domain.AccountService.
func (a *accountService) Create(ctx context.Context) (dto.AccountRes, error) {
	user := ctx.Value("x-user").(dto.UserData)

	if user.ID == 0 {
		return dto.AccountRes{}, domain.ErrAuthFailed
	}

	_, err := a.accountRepository.FindByUserId(ctx, user.ID)
	if err == nil {
		return dto.AccountRes{}, domain.ErrAccountAlreadyExists
	}

	random := util.GeneratorRandomNumber(16)

	account := domain.Account{
		UserID:        user.ID,
		AccountNumber: random,
		Balance:       0.0,
	}
	err = a.accountRepository.Insert(ctx, &account)
	if err != nil {
		return dto.AccountRes{}, err
	}
	return dto.AccountRes{
		UserId:        account.UserID,
		AccountNumber: account.AccountNumber,
		Balance:       account.Balance,
	}, nil
}
