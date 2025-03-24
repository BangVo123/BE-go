package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ReadRequest(ctx *gin.Context, target any) error {
	if err := ctx.ShouldBind(target); err != nil {
		return err
	}

	if err := validate.StructCtx(ctx.Request.Context(), target); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
