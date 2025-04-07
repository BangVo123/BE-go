package http

import (
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
	LikeUseCase  usecase.LikeUseCase
	LoveUseCase  usecase.LoveUseCase
}

func NewVideoHandler(VideoUseCase usecase.VideoUseCase, cld cloudinary.Cloudinary, LikeUseCase usecase.LikeUseCase, LoveUseCase usecase.LoveUseCase) handlers.VideoHandler {
	return &VideoHandler{VideoUseCase: VideoUseCase, cld: cld, LikeUseCase: LikeUseCase, LoveUseCase: LoveUseCase}
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

func (vh *VideoHandler) Like(c *gin.Context) {
	videoId := c.Param("videoId")
	if videoId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Can not retrieve video id"})
		return
	}

	user, _ := c.Get("user")
	userModel, ok := (user).(models.User)
	if !ok {
		c.JSON(http.StatusBadRequest, "Something went wrong when extract userinfo")
		return
	}

	err := vh.LikeUseCase.Like(c.Request.Context(), videoId, userModel.Id.Hex())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error", "err": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success"})
}

func (vh *VideoHandler) Love(c *gin.Context) {
	videoId := c.Param("videoId")
	if videoId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Can not retrieve video id"})
		return
	}

	user, _ := c.Get("user")
	userModel, ok := (user).(models.User)
	if !ok {
		c.JSON(http.StatusBadRequest, "Something went wrong when extract userinfo")
		return
	}

	err := vh.LoveUseCase.Love(c.Request.Context(), videoId, userModel.Id.Hex())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error", "err": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success"})
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
		return
	}
	c.JSON(http.StatusOK, map[string]any{"data": videos})
}
