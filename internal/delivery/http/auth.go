package http

import (
	"fmt"
	"log"
	"net/http"
	"project/config"
	auth "project/internal/handlers"
	"project/internal/presenter"
	"project/internal/usecase"
	"project/utils"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthHandler struct {
	UserCase usecase.UserCase
	cfg      *config.Configuration
}

func NewAuthHandler(userCase usecase.UserCase, config *config.Configuration) auth.Handler {
	return &AuthHandler{UserCase: userCase, cfg: config}
}

func (ah *AuthHandler) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		creds := presenter.LoginReq{}
		err := utils.ReadRequest(c, &creds)
		if err != nil {
			c.JSON(404, gin.H{"error0": err.Error()})
			return
		}

		user, err := ah.UserCase.Login(c.Request.Context(), creds)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error1": err.Error()})
			return
		}

		Payload := map[string]string{
			"id":    user.ID.Hex(),
			"email": user.Email,
		}

		token, err := utils.GenToken(Payload, ah.cfg.JWTAccessTokenSecret)
		if err != nil {
			c.JSON(404, gin.H{"error2": err.Error()})
		}

		c.SetCookie("jwt", token, 86400, "/", "", false, true)
		c.JSON(http.StatusOK, gin.H{"message": "Login success"})
	}
}

func (ah *AuthHandler) SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func (ah *AuthHandler) GoogleOauth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		provider := ctx.Param("provider")
		q := ctx.Request.URL.Query()
		q.Set("provider", provider)
		ctx.Request.URL.RawQuery = q.Encode()

		gothic.BeginAuthHandler(ctx.Writer, ctx.Request)
	}
}

func (ah *AuthHandler) GoogleOauthCallback() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := gothic.CompleteUserAuth(ctx.Writer, ctx.Request)
		if err != nil {
			log.Println("OAuth error:", err)
			ctx.JSON(http.StatusTemporaryRedirect, "Login fail")
			return
		}

		fmt.Println(user)

		// Example: Print user info or store in session/db
		// log.Println("User:", user.Name, user.Email, user.UserID)
		//call check user func

		provider := ctx.Param("provider")
		filter := map[string]any{
			"email":    user.Email,
			"provider": provider,
		}

		// fmt.Println(user.Email, " - ", provider)

		FoundUser, err := ah.UserCase.CheckUserExist(ctx, filter)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				//signup
				ctx.JSON(http.StatusBadRequest, "Signup")
				return

			} else {
				ctx.JSON(http.StatusBadRequest, "Something went wrong")
				return
			}
		}

		fmt.Print(FoundUser)

		// Payload := map[string]string{
		// 	"id":    FoundUser.ID.Hex(),
		// 	"email": FoundUser.Email,
		// }

		// token, err := utils.GenToken(Payload, ah.cfg.JWTAccessTokenSecret)
		// if err != nil {
		// 	ctx.JSON(http.StatusBadRequest, "Token creation error")
		// 	return
		// }
		// ctx.SetCookie("jwt", token, 86400, "/", "", false, true)

		// Redirect to frontend after success
		// ctx.JSON(http.StatusOK, "Login Success")
		ctx.Redirect(http.StatusTemporaryRedirect, ah.cfg.ClientUrl)
	}
}
