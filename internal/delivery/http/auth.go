package http

import (
	"fmt"
	"log"
	"net/http"
	"project/config"
	"project/internal/handlers"
	"project/internal/models"
	"project/internal/presenter"
	"project/internal/repositories"
	"project/internal/usecase"
	"project/utils"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (ah *AuthHandler) Login(c *gin.Context) {
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

func (ah *AuthHandler) SignUp(c *gin.Context) {
}

func (ah *AuthHandler) GoogleOauth(c *gin.Context) {
	provider := c.Param("provider")
	q := c.Request.URL.Query()
	q.Set("provider", provider)
	c.Request.URL.RawQuery = q.Encode()

	gothic.BeginAuthHandler(c.Writer, c.Request)
}

func (ah *AuthHandler) GoogleOauthCallback(c *gin.Context) {
	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		log.Println("OAuth error:", err)
		c.JSON(http.StatusTemporaryRedirect, "Login fail")
		return
	}

	provider := c.Param("provider")
	filter := map[string]any{
		"email":    user.Email,
		"provider": provider,
	}

	FoundUser, err := ah.UserCase.GetUserExist(c.Request.Context(), filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			//signup handler
			payload := models.User{
				Id:       primitive.NewObjectID(),
				Email:    user.Email,
				FullName: user.Name,
				NickName: user.NickName,
				Provider: user.Provider,
				Avatar:   user.AvatarURL,
			}
			FoundUser, err = ah.UserCase.CreateUser(c.Request.Context(), &payload)
			if err != nil {
				c.JSON(http.StatusInternalServerError, "Signup")
				return
			}

		} else {
			c.JSON(http.StatusBadRequest, "Something went wrong")
			return
		}
	}

	sessionId, err := ah.MongoStore.Save(c.Request.Context(), FoundUser.Id.Hex())
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Something went wrong")
		return
	}

	c.SetCookie("cookie", sessionId, 86400, "/", "", false, true)
	gothic.Logout(c.Writer, c.Request)

	c.Redirect(http.StatusTemporaryRedirect, ah.cfg.ClientUrl)
}

func (ah *AuthHandler) Logout(c *gin.Context) {
	fmt.Print("test")
	if _, err := c.Cookie("cookie"); err == nil {
		c.SetCookie("cookie", "", -1, "/", "", false, true)
	}
	c.JSON(http.StatusOK, "Logout success")
}

func (ah *AuthHandler) GetMe(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusForbidden, "User not authenticate")
		return
	}

	foundUser := user.(models.User)
	c.JSON(http.StatusOK, map[string]any{"data": map[string]any{"user": foundUser}})
}

func (ah *AuthHandler) GetUserInfo(c *gin.Context) {
	userIdString := c.Param("userId")

	FoundUser, err := ah.UserCase.GetUserById(c.Request.Context(), userIdString)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Something went wrong when get userinfo")
	}

	c.JSON(http.StatusOK, map[string]any{"data": FoundUser})
}
