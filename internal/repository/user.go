package repository

import (
	"ahyalfan/golang_e_money/domain"
	"context"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) domain.UserRepository {
	return &userRepository{db: db}
}

// FindById implements domain.UserRepository.
func (ur *userRepository) FindById(ctx context.Context, id int64) (user domain.User, err error) {
	err = ur.db.WithContext(ctx).Where("id =?", id).First(&user).Error
	return
}

// FindByUsername implements domain.UserRepository.
func (ur *userRepository) FindByUsername(ctx context.Context, username string) (user domain.User, err error) {
	err = ur.db.WithContext(ctx).Where("username =?", username).First(&user).Error
	return
}

// Insert implements domain.UserRepository.
func (ur *userRepository) Insert(ctx context.Context, user *domain.User) error {
	return ur.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Create(user).Error
	})
}

// Update implements domain.UserRepository.
func (ur *userRepository) Update(ctx context.Context, user *domain.User) error {
	return ur.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Save(user).Error
	})
}
