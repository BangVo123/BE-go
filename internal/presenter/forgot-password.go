package presenter

type ForgotPasswordReq struct {
	Email string `json:"email" validate:"required,email"`
	Type  string `json:"type" validate:"required"`
}
