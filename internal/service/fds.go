package service

import (
	"ahyalfan/golang_e_money/domain"
	"ahyalfan/golang_e_money/dto"
	"ahyalfan/golang_e_money/internal/util"
	"context"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type fdsService struct {
	ipCheckerService domain.IpCheckerService
	loginLog         domain.LoginLogRepository
}

func NewFds(ipCheckerService domain.IpCheckerService, loginLog domain.LoginLogRepository) domain.FdsService {
	return &fdsService{ipCheckerService: ipCheckerService, loginLog: loginLog}
}

// IsAuthorized implements domain.FdsService.
func (f *fdsService) IsAuthorized(ctx context.Context, ip string, userId int64) bool {
	locationCheck, err := f.ipCheckerService.Query(ctx, ip)
	if err != nil || locationCheck == (dto.IpChecker{}) {
		return false
	}

	newAccess := domain.LoginLog{
		UserID:       userId,
		IpAddress:    ip,
		TimeZone:     locationCheck.TimeZone,
		Latitude:     locationCheck.Lat,
		Longitude:    locationCheck.Lon,
		IsAuthorized: false,
		AccessTime:   time.Now(),
	}

	lastLogin, err := f.loginLog.FindLastAuthorized(ctx, userId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		newAccess.IsAuthorized = true
		_ = f.loginLog.Insert(ctx, &newAccess)
		return true
	}
	if err != nil {
		_ = f.loginLog.Insert(ctx, &newAccess)
		return false
	}

	// cek jarak selisih jamnya
	distanceHour := lastLogin.AccessTime.Sub(newAccess.AccessTime)
	// jarak lokasi
	distanceChange := util.GetDistance(lastLogin.Latitude, lastLogin.Longitude, newAccess.Latitude, newAccess.Longitude)
	fmt.Printf("hour: %f, distance: %f", distanceHour.Hours(), distanceChange)
	// jika jarak lokasi lebih besar 400 km/jam, maka tidak di perbolehkan
	if (distanceChange / distanceHour.Hours()) > 400 {
		_ = f.loginLog.Insert(ctx, &newAccess)
		return false
	}
	newAccess.IsAuthorized = true
	_ = f.loginLog.Insert(ctx, &newAccess)
	return true
}
