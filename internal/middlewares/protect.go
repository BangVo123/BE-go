package middlewares

import (
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

func (m *MiddlewareManager) Protect(c *gin.Context) {
	cookie, err := c.Cookie("cookie")
	if err != nil {
		if err != http.ErrNoCookie {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error:": "Access denied"})
			return
		} else {
			tokenString := c.GetHeader("Authorization")
			if tokenString == "" {
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error:": "Access denied"})
				return
			}

			token := strings.Split(tokenString, " ")[1]

			tokenClaims, err := utils.VerifyToken(token, m.cfg.JWTAccessTokenSecret)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Authentication fail"})
				return
			}

			Expired := time.Unix(tokenClaims.Expired, 0)
			if time.Now().After(Expired) {
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Authentication fail"})
				return
			}

			FoundUser, err := m.UserCase.GetUserExist(c.Request.Context(), map[string]any{"_id": tokenClaims.Id, "provider": nil})
			if err != nil {
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Authentication fail"})
				return
			}
			c.Set("user", FoundUser)
		}
	} else {
		// fmt.Print(cookie)
		userId, err := m.MongoStore.Load(cookie)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Access Denied"})
			return
		}

		FoundUser, err := m.UserCase.GetUserById(c.Request.Context(), userId)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Access denied"})
			return
		}
		c.Set("user", *FoundUser)
	}

	c.Next()
}
