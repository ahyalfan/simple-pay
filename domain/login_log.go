package domain

import (
	"context"
	"time"
)

// sebenarnya bisa banyak yg bisa kita ambil informasinya, tapi sebagai permulaan ini saja sudah bagus
type LoginLog struct {
	ID           int64     `gorm:"primary_key;auto_increment"`
	UserID       int64     `gorm:"column:user_id"`
	IsAuthorized bool      `gorm:"column:is_authorized"`
	IpAddress    string    `gorm:"column:ip_address"`
	TimeZone     string    `gorm:"column:time_zone"`
	Latitude     float64   `gorm:"column:latitude"`
	Longitude    float64   `gorm:"column:longitude"`
	AccessTime   time.Time `gorm:"column:access_time"`
}

type LoginLogRepository interface {
	FindLastAuthorized(ctx context.Context, userId int64) (LoginLog, error) //kita cek diama dia login terakhir
	// bisa jadi dia dari jakarta tiba tiba luar negeri kan, gk masuk akal, biasanya ini yg pakai vpn. jadi bisa kita block

	Insert(ctx context.Context, loginLog *LoginLog) error
}
