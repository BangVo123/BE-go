package presenter

type ResetPasswordReq struct {
	Email     string `json:"username" validate:"email,required"`
	DigitCode string `json:"code" validate:"required"`
	Password  string `json:"password" validate:"required"`
}
