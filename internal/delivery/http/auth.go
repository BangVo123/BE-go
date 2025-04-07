package http

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"project/config"
	"project/internal/handlers"
	"project/internal/models"
	"project/internal/presenter"
	"project/internal/repositories"
	"project/internal/usecase"
	"project/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	mail "gopkg.in/gomail.v2"
)

type AuthHandler struct {
	UserCase     usecase.UserCase
	CodeUserCase usecase.CodeUseCase
	cfg          *config.Configuration
	MongoStore   *repositories.MongoSessionStore
	dialer       mail.Dialer
}

func NewAuthHandler(userCase usecase.UserCase, config *config.Configuration, mongoStore *repositories.MongoSessionStore, CodeUseCase usecase.CodeUseCase, dialer mail.Dialer) handlers.AuthHandler {
	return &AuthHandler{UserCase: userCase, cfg: config, MongoStore: mongoStore, CodeUserCase: CodeUseCase, dialer: dialer}
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
	var creds presenter.RegisterReq
	err := c.ShouldBind(&creds)
	if err != nil {
		c.JSON(http.StatusNoContent, "Some thing went wrong when get creds")
		return
	}

	Code, err := ah.CodeUserCase.GetCode(c.Request.Context(), creds.Email, "auth")
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]any{"error1: ": err})
		return
	}

	codeNumber, err := strconv.Atoi(creds.Code)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]any{"error2: ": err})
		return
	}

	if Code.Code != codeNumber {
		c.JSON(http.StatusBadRequest, "Credentials invalid")
		return
	}

	//add user
	userId, err := ah.UserCase.SignUp(c.Request.Context(), creds)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]any{"error3: ": err})
		return
	}

	//get user and gen token here

	Payload := map[string]string{
		"id":    userId,
		"email": creds.Email,
	}

	token, err := utils.GenToken(Payload, ah.cfg.JWTAccessTokenSecret)
	if err != nil {
		c.JSON(404, gin.H{"error4": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Signup success", "metadata": map[string]any{"token": token}})
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

	var userId primitive.ObjectID

	FoundUser, err := ah.UserCase.GetUserExist(c.Request.Context(), filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			userId = primitive.NewObjectID()
			//signup handler
			payload := models.User{
				Id:       userId,
				Email:    user.Email,
				FullName: user.Name,
				NickName: user.NickName,
				Provider: user.Provider,
				Avatar:   user.AvatarURL,
			}
			err = ah.UserCase.CreateUser(c.Request.Context(), &payload)
			if err != nil {
				c.JSON(http.StatusInternalServerError, "Something went wrong")
				return
			}

		} else {
			c.JSON(http.StatusBadRequest, "Something went wrong")
			return
		}
	}

	if userId.IsZero() {
		userId = FoundUser.Id
	}

	sessionId, err := ah.MongoStore.Save(c.Request.Context(), userId.Hex())
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

// forgot and reset
func (ah *AuthHandler) Forgot(c *gin.Context) {
	var creds presenter.ForgotPasswordReq
	err := c.ShouldBind(&creds)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error when get body of request")
		return
	}

	err = ah.CodeUserCase.DeleteCode(c.Request.Context(), map[string]any{"email": creds.Email})
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]any{"error:": err})
		return
	}

	Code := 100000 + rand.Intn(999999-100000)
	Payload := models.DigitCode{
		Id:        primitive.NewObjectID(),
		Email:     creds.Email,
		Code:      Code,
		Type:      creds.Type,
		ExpiredAt: time.Now().Add(5 * time.Minute),
	}
	err = ah.CodeUserCase.CreateCode(c.Request.Context(), &Payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]any{"error:": err})
		return
	}

	err = utils.SendMail(&ah.dialer, ah.cfg.Email, creds.Email, strconv.Itoa(Code))
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]any{"error: ": err})
		return
	}

	c.JSON(http.StatusOK, map[string]any{"message": "Please check mail to get code"})
}

func (ah *AuthHandler) Reset(c *gin.Context) {
	var creds presenter.ResetPasswordReq
	err := c.ShouldBind(&creds)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]any{"error1:": err})
		return
	}

	Code, err := ah.CodeUserCase.GetCode(c.Request.Context(), creds.Email, "reset")
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]any{"error2:": err})
		return
	}

	codeNumber, err := strconv.Atoi(creds.DigitCode)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]any{"error: ": err})
		return
	}

	if Code.Code != codeNumber {
		c.JSON(http.StatusBadRequest, "Credentials invalid")
		return
	}

	if Code.ExpiredAt.Before(time.Now()) {
		c.JSON(http.StatusInternalServerError, map[string]any{"message:": "Code is expired"})
		return
	}

	err = ah.UserCase.Reset(c.Request.Context(), Code.Email, creds.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]any{"error:": err})
		return
	}

	c.JSON(http.StatusOK, map[string]any{"message:": "Success"})
}

// func (ah *AuthHandler) GetMe(c *gin.Context) {
// 	user, exists := c.Get("user")
// 	if !exists {
// 		c.JSON(http.StatusForbidden, "User not authenticate")
// 		return
// 	}

// 	foundUser := user.(models.User)
// 	c.JSON(http.StatusOK, map[string]any{"data": map[string]any{"user": foundUser}})
// }

// func (ah *AuthHandler) GetUserInfo(c *gin.Context) {
// 	userIdString := c.Param("userId")

// 	FoundUser, err := ah.UserCase.GetUserById(c.Request.Context(), userIdString)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, "Something went wrong when get userinfo")
// 	}

// 	c.JSON(http.StatusOK, map[string]any{"data": FoundUser})
// }
