package handlers

import "github.com/gin-gonic/gin"

type UploadHandler interface {
	VideoUpload(c *gin.Context)
	AvatarUpload(c *gin.Context)
}
