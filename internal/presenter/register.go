package presenter

type RegisterReq struct {
	Email    string `json:"username" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Code     string `json:"code" validate:"required"`
}
