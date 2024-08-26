package domain

import "context"

type Template struct {
	Code  string `gorm:"primary_key;column:code"`
	Title string `gorm:"column:title"`
	Body  string `gorm:"column:body"`
}

type TemplateRepository interface {
	FindByCode(ctx context.Context, code string) (Template, error)
}
