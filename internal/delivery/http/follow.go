package http

import (
	"net/http"
	"project/internal/handlers"
	"project/internal/models"
	"project/internal/presenter"
	"project/internal/usecase"

	"github.com/gin-gonic/gin"
)

type FollowHandler struct {
	followUseCase usecase.FollowUseCase
}

func NewFollowhandler(followUseCase usecase.FollowUseCase) handlers.FollowHandler {
	return &FollowHandler{followUseCase: followUseCase}
}

func (fh *FollowHandler) Follow(c *gin.Context) {
	user, _ := c.Get("user")
	userObj, ok := user.(models.User)
	if !ok {
		c.JSON(http.StatusBadRequest, "Something went wrong when get user info")
		return
	}

	var followReq presenter.FollowReq
	err := c.ShouldBind(followReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error", "err": err})
		return
	}

	err = fh.followUseCase.Follow(c.Request.Context(), userObj.Id.Hex(), followReq.FollowingId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error", "err": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success"})
}
