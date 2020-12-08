package api

import (
	"github.com/gin-gonic/gin"
	"github.com/starbuling-l/StarBlog/pkg/e"
	"github.com/starbuling-l/StarBlog/pkg/upload"
	"log"
	"net/http"
)

func UploadImage(ctx *gin.Context) {
	code := e.SUCCESS
	data := make(map[string]string)
	file, image, err := ctx.Request.FormFile("image")
	if err != nil {
		log.Printf("ctx.Request.FormFile err:%v", err)
		code = e.ERROR
	}
	if image == nil {
		code = e.INVALID_PARAMS
	} else {
		imageName := upload.GetImageName(image.Filename)
		fullPath := upload.GetImageFullPath()
		savePath := upload.GetImagePath()
		src := fullPath + savePath
		if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
			code = e.ERROR_UPLOAD_CHECK_IMAGE_FORMAT
		} else {
			err := upload.CheckImage(fullPath)
			if err != nil {
				code = e.ERROR_UPLOAD_CHECK_IMAGE_FAIL
			} else if err := ctx.SaveUploadedFile(image, src); err != nil {
				code = e.ERROR_UPLOAD_SAVE_IMAGE_FAIL
			} else {
				data["image_url"] = upload.GetImageFullUrl(imageName)
				data["image_save_url"] = savePath + imageName
			}
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
