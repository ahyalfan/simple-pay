package repository

import (
	"ahyalfan/golang_e_money/domain"
	"context"

	"gorm.io/gorm"
)

type accountRepository struct {
	db *gorm.DB
}

func NewAccount(db *gorm.DB) domain.AccountRepository {
	return &accountRepository{db: db}
}

// FindByAccountNumber implements domain.AccountRepository.
func (a *accountRepository) FindByAccountNumber(ctx context.Context, accountNumber string) (account domain.Account, err error) {
	err = a.db.WithContext(ctx).Where("account_number = ?", accountNumber).First(&account).Error
	return
}

// FindById implements domain.AccountRepository.
func (a *accountRepository) FindByUserId(ctx context.Context, id int64) (account domain.Account, err error) {
	err = a.db.WithContext(ctx).Where("user_id = ?", id).First(&account).Error
	return
}

// Insert implements domain.AccountRepository.
func (a *accountRepository) Insert(ctx context.Context, account *domain.Account) error {
	return a.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Create(account).Error
	})
}

// Update implements domain.AccountRepository.
func (a *accountRepository) Update(ctx context.Context, account *domain.Account) error {
	return a.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Save(account).Error
	})
}
