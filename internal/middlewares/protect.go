package middlewares

import (
	"context"
	"net/http"
	"project/config"
	"project/internal/repositories"
	"project/internal/usecase"
	"project/utils"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type MiddlewareManager struct {
	UserCase   usecase.UserCase
	cfg        *config.Configuration
	MongoStore *repositories.MongoSessionStore
}

func NewMiddlewareManager(UserCase usecase.UserCase, cfg *config.Configuration, MongoStore *repositories.MongoSessionStore) *MiddlewareManager {
	return &MiddlewareManager{UserCase: UserCase, cfg: cfg, MongoStore: MongoStore}
}

func (m *MiddlewareManager) Protect() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("cookie")
		if err != nil {
			if err != http.ErrNoCookie {
				c.JSON(http.StatusForbidden, "Access denied")
				return
			} else {
				token := strings.Split(c.Request.Header["Authorization"][0], " ")[1]
				if token == "" {
					c.JSON(http.StatusForbidden, "Access denied")
					return
				}

				tokenClaims, err := utils.VerifyToken(token, m.cfg.JWTAccessTokenSecret)
				if err != nil {
					c.JSON(http.StatusForbidden, "Authentication fail")
					return
				}

				Expired := time.Unix(tokenClaims.Expired, 0)
				if time.Now().After(Expired) {
					c.JSON(http.StatusForbidden, "Authentication fail")
					return
				}

				FoundUser, err := m.UserCase.CheckUserExist(context.TODO(), map[string]any{"_id": tokenClaims.Id})
				if err != nil {
					c.JSON(http.StatusForbidden, "Authentication fail")
					return
				}
				c.Set("user", FoundUser)
			}
		} else {
			// fmt.Print(cookie)
			userId, err := m.MongoStore.Load(cookie)
			if err != nil {
				c.JSON(http.StatusForbidden, "Authentication fail")
				return
			}

			FoundUser, err := m.UserCase.CheckUserExist(context.TODO(), map[string]any{"_id": userId})
			if err != nil {
				c.JSON(http.StatusForbidden, "Authentication fail")
				return
			}
			c.Set("user", FoundUser)
		}

		c.Next()
	}
}
