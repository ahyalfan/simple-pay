package domain

import "errors"

var ErrUserNotFound = errors.New("User not found")
var ErrUserAlreadyExists = errors.New("User already exists")
var ErrAuthFailed = errors.New("auth failed")
var ErrInvalidOTP = errors.New("otp invalid")
