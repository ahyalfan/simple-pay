package util

import (
	"ahyalfan/golang_e_money/domain"
	"errors"
)

func GetHttpStatus(err error) int {
	if err == nil {
		return 200
	}

	switch {
	case errors.Is(err, domain.ErrAuthFailed):
		return 401
	case errors.Is(err, domain.ErrUserAlreadyExists):
		return 400
	case errors.Is(err, domain.ErrInvalidOTP):
		return 400
	default:
		return 500
	}
}
