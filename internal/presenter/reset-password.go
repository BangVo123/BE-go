package presenter

type ResetPasswordReq struct {
	DigitCode string `json:"digit_code" validate:"required"`
	Password  string `json:"password" validate:"required"`
}
