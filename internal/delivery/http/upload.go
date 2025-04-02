package http

import (
	"fmt"
	"net/http"
	"project/internal/handlers"
	"project/internal/models"
	"project/internal/presenter"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
)

type UploadHandler struct {
	cld cloudinary.Cloudinary
}

func NewUploadHandler(cld cloudinary.Cloudinary) handlers.UploadHandler {
	return &UploadHandler{cld: cld}
}

func (uh *UploadHandler) VideoUpload(c *gin.Context) {
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, "Something went wrong when get data from file")
	}

	userData, _ := c.Get("user")
	user, ok := userData.(models.User)
	if !ok {
		c.JSON(http.StatusBadRequest, "Something went wrong when retrieve user data")
	}

	uploadResult, err := uh.cld.Upload.Upload(c.Request.Context(), file, uploader.UploadParams{ResourceType: "video", Folder: "video/" + user.Id.Hex()})
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Upload video fail")
	}

	resData := presenter.VideoInfoRes{Name: uploadResult.OriginalFilename, Url: uploadResult.SecureURL, Size: uploadResult.Bytes, Width: uploadResult.Width, Height: uploadResult.Height}

	fmt.Println("resData: ", resData)

	c.JSON(http.StatusOK, map[string]any{"message": "Success", "data": resData})
}

func (uh *UploadHandler) AvatarUpload(c *gin.Context) {
	file, _, err := c.Request.FormFile("avatar")
	if err != nil {
		c.JSON(http.StatusBadRequest, "Something went wrong when get data from file")
	}

	userData, _ := c.Get("user")
	user, ok := userData.(models.User)
	if !ok {
		c.JSON(http.StatusBadRequest, "Something went wrong when retrieve user data")
	}

	uploadResult, err := uh.cld.Upload.Upload(c.Request.Context(), file, uploader.UploadParams{ResourceType: "image", Folder: "image/" + user.Id.Hex()})
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Upload video fail")
	}

	c.JSON(http.StatusOK, map[string]any{"message": "Success", "data": uploadResult.SecureURL})
}
