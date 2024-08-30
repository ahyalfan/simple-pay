package service

import (
	"ahyalfan/golang_e_money/domain"
	"ahyalfan/golang_e_money/dto"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ipCheckerService struct {
}

func NewIpChecker() domain.IpCheckerService {
	return &ipCheckerService{}
}

// Query implements domain.IpCheckerService.
func (i *ipCheckerService) Query(ctx context.Context, ip string) (checker dto.IpChecker, err error) {
	// kita bisa cek ip public di ip-api.com
	url := fmt.Sprintf("http://ip-api.com/json/%s?fields=status,lat,lon,timezone,query", ip)
	resp, err := http.Get(url)
	if err != nil {
		return dto.IpChecker{}, err
	}
	defer resp.Body.Close() // kita close karena ini pakai resource

	body, err := io.ReadAll(resp.Body)
	err = json.Unmarshal(body, &checker)
	return
}
