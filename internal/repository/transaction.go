package repository

import (
	"ahyalfan/golang_e_money/domain"
	"context"

	"gorm.io/gorm"
)

type transactionRepository struct {
	db *gorm.DB
}

func NewTransaction(db *gorm.DB) domain.TransactionRepository {
	return &transactionRepository{db: db}
}

// FindByAccountId implements domain.TransactionRepository.
func (t *transactionRepository) FindByAccountId(ctx context.Context, accountId int64) (transactions []domain.Transaction, err error) {
	err = t.db.WithContext(ctx).Where("account_id = ?", accountId).Find(&transactions).Error
	return
}

// FindByDofNumber implements domain.TransactionRepository.
func (t *transactionRepository) FindByDofNumber(ctx context.Context, dofNumber string) (transaction domain.Transaction, err error) {
	err = t.db.WithContext(ctx).Where("dof_number = ?", dofNumber).First(&transaction).Error
	return
}

// FindBySofNumber implements domain.TransactionRepository.
func (t *transactionRepository) FindBySofNumber(ctx context.Context, sofNumber string) (transaction domain.Transaction, err error) {
	err = t.db.WithContext(ctx).Where("sof_number = ?", sofNumber).First(&transaction).Error
	return
}

// Insert implements domain.TransactionRepository.
func (t *transactionRepository) Insert(ctx context.Context, transaction *domain.Transaction) error {
	return t.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Create(transaction).Error
	})
}
