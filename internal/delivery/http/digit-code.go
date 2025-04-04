package http

import (
	"net/http"
	"project/internal/handlers"
	"project/internal/usecase"

	"github.com/gin-gonic/gin"
	mail "gopkg.in/gomail.v2"
)

type CodeHandler struct {
	CodeUseCase usecase.CodeUseCase
	dialer      *mail.Dialer
	Email       string
}

func NewCodeHandler(CodeUseCase usecase.CodeUseCase, dialer *mail.Dialer, Email string) handlers.CodeHandler {
	return &CodeHandler{CodeUseCase: CodeUseCase, dialer: dialer, Email: Email}
}

func (ch *CodeHandler) GetCode(c *gin.Context) {
	// var Email presenter.ForgotPasswordReq
	// err := c.ShouldBind(&Email)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, "Can not get content of request")
	// 	return
	// }

	// Code, err := ch.CodeUseCase.GetCode(c.Request.Context(), Email.Email)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, map[string]any{"error: ": err})
	// 	return
	// }

	// err = utils.SendMail(ch.dialer, ch.Email, Email.Email, strconv.Itoa(Code.Code))
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, map[string]any{"error: ": err})
	// 	return
	// }

	c.JSON(http.StatusOK, "Please check mail to get code")
}

func (ch *CodeHandler) CreateCode(c *gin.Context) {

}

func (ch *CodeHandler) DeleteCode(c *gin.Context) {

}
