package http

import (
	"log"
	"net/http"
	"project/config"
	"project/internal/handlers"
	"project/internal/presenter"
	"project/internal/repositories"
	"project/internal/usecase"
	"project/utils"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthHandler struct {
	UserCase   usecase.UserCase
	cfg        *config.Configuration
	MongoStore *repositories.MongoSessionStore
}

func NewAuthHandler(userCase usecase.UserCase, config *config.Configuration, mongoStore *repositories.MongoSessionStore) handlers.AuthHandler {
	return &AuthHandler{UserCase: userCase, cfg: config, MongoStore: mongoStore}
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
			"id":    user.Id.Hex(),
			"email": user.Email,
		}

		token, err := utils.GenToken(Payload, ah.cfg.JWTAccessTokenSecret)
		if err != nil {
			c.JSON(404, gin.H{"error2": err.Error()})
		}

		c.JSON(http.StatusOK, gin.H{"message": "Login success", "metadata": map[string]any{"token": token}})
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

		provider := ctx.Param("provider")
		filter := map[string]any{
			"email":    user.Email,
			"provider": provider,
		}

		FoundUser, err := ah.UserCase.CheckUserExist(ctx, filter)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				//signup handler here

				ctx.JSON(http.StatusBadRequest, "Signup")
				return

			} else {
				ctx.JSON(http.StatusBadRequest, "Something went wrong")
				return
			}
		}

		sessionId, err := ah.MongoStore.Save(FoundUser.Id.Hex())
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, "Something went wrong")
			return
		}

		ctx.SetCookie("cookie", sessionId, 86400, "/", "", false, true)

		ctx.Redirect(http.StatusTemporaryRedirect, ah.cfg.ClientUrl)
	}
}
