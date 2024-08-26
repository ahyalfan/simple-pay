package repository

import (
	"ahyalfan/golang_e_money/domain"
	"context"

	"gorm.io/gorm"
)

type templateRepository struct {
	db *gorm.DB
}

func NewTemplate(db *gorm.DB) domain.TemplateRepository {
	return &templateRepository{db: db}
}

// FindByCode implements domain.TemplateRepository.
func (t *templateRepository) FindByCode(ctx context.Context, code string) (template domain.Template, err error) {
	err = t.db.WithContext(ctx).First(&template, "code = ? ", code).Error
	return
}
