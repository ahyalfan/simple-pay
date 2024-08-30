package domain

import "context"

// Fault Detection System (FDS)
type FdsService interface {
	IsAuthorized(ctx context.Context, ip string, userId int64) bool
}
