package dto

type ValidateOtpReq struct {
	ReferenceID string `json:"reference_id" validate:"required"`
	OTP         string `json:"otp" validate:"required,numeric"`
}
