package http

import (
	"fmt"
	"net/http"
	"project/internal/handlers"
	"project/internal/models"
	"project/internal/usecase"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type VideoHandler struct {
	VideoUseCase usecase.VideoUseCase
	cld          cloudinary.Cloudinary
}

func NewVideoHandler(VideoUseCase usecase.VideoUseCase, cld cloudinary.Cloudinary) handlers.VideoHandler {
	return &VideoHandler{VideoUseCase: VideoUseCase, cld: cld}
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
	c.JSON(http.StatusOK, map[string]any{"data": videos})
}

func (vh *VideoHandler) AddVideo(c *gin.Context) {
	var AddVideo models.Video
	if err := c.BindJSON(&AddVideo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, _ := c.Get("user")
	userInfo, ok := user.(models.User)
	if !ok {
		c.JSON(http.StatusBadRequest, "Something went wrong when extract userinfo")
		return
	}

	AddVideo.Id = primitive.NewObjectID()
	AddVideo.BelongTo = userInfo.Id

	err := vh.VideoUseCase.AddVideo(c.Request.Context(), AddVideo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Add video success")
}

type VideoWithOwnerInfoHandler struct {
	VideoWithOwnerInfoCase usecase.VideoWithOwnerInfoUseCase
}

func NewVideoWithOwnerInfoHandler(VideoWithOwnerInfoCase usecase.VideoWithOwnerInfoUseCase) handlers.VideoWithOwnerInfoHandler {
	return &VideoWithOwnerInfoHandler{VideoWithOwnerInfoCase: VideoWithOwnerInfoCase}
}

func (vh *VideoWithOwnerInfoHandler) GetVideos(c *gin.Context) {
	pageString := c.Query("page")
	limitString := c.Query("limit")

	var videos *[]models.VideoWithOwnerInfo
	videos, err := vh.VideoWithOwnerInfoCase.GetVideosWithOwnerInfo(c.Request.Context(), pageString, limitString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Something went wrong")
		fmt.Print("error:", err)
		return
	}
	c.JSON(http.StatusOK, map[string]any{"data": videos})
}
