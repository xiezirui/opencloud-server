package controller

import (
	"github.com/gin-gonic/gin"
	"opencloud-server/response"
	"opencloud-server/utils"
	"os"
)

func UploadFile(context *gin.Context) {

	FID := utils.GetFID()

	file, _ := context.FormFile("file")

	os.Mkdir("./static/upload/"+FID, 0666)

	context.SaveUploadedFile(file, "./static/upload/"+FID+"/"+file.Filename)

	response.Success(context, gin.H{"fid": FID, "fname": file.Filename}, "ok")
}
