package domain

import (
	"ahyalfan/golang_e_money/dto"
	"context"
)

type IpCheckerService interface {
	Query(ctx context.Context, ip string) (dto.IpChecker, error)
}
