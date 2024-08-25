package domain

import "errors"

var ErrUserNotFound = errors.New("User not found")
var ErrInquiryNotFound = errors.New("inquiry not found")
var ErrAccountNotFound = errors.New("Account not found")
var ErrAccountAlreadyExists = errors.New("account already")
var ErrUserAlreadyExists = errors.New("User already exists")
var ErrAuthFailed = errors.New("auth failed")
var ErrInvalidOTP = errors.New("otp invalid")
var ErrInsufficientBalance = errors.New("insufficient balance")

var ErrInvalidTransfer = errors.New("invalid transfer")

var ErrCacheMiss = errors.New("cache miss")
