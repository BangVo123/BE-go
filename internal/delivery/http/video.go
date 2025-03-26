package http

import (
	"net/http"
	"project/internal/handlers"
	"project/internal/models"
	"project/internal/usecase"

	"github.com/gin-gonic/gin"
)

type VideoHandler struct {
	VideoUseCase usecase.VideoUseCase
}

func NewVideoHandler(VideoUseCase usecase.VideoUseCase) handlers.VideoHandler {
	return &VideoHandler{VideoUseCase: VideoUseCase}
}

// this func is a gin.HandlerFunc - has param *gin.Context
func (vh *VideoHandler) GetVideos(c *gin.Context) {
	pageString := c.Query("page")
	limitString := c.Query("limit")

	var videos *[]models.Video
	videos, err := vh.VideoUseCase.GetVideos(c.Request.Context(), pageString, limitString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Something went wrong")
		return
	}
	c.JSON(http.StatusOK, map[string]any{"data": &videos})

}
