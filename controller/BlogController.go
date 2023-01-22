package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"opencloud-server/database"
	"opencloud-server/model/dao"
	"opencloud-server/model/dto"
	"opencloud-server/response"
	"opencloud-server/utils"
)

func AddBlog(context *gin.Context) {
	var blog dao.Blog

	context.Bind(&blog)

	blog.BlogID = utils.GetBID()

	database.AddBlog(&blog)

	response.Success(context, nil, "博客上传成功")
}

func GetBlogs(context *gin.Context) {

	var Req dto.Req

	context.Bind(&Req)

	var blogs []dao.Blog

	database.GetBlogs(Req.Uid, &blogs)

	log.Printf("%v", blogs)

	for index := range blogs {
		blogs[index].CoverPath = "http://localhost:1016/static/upload/" + blogs[index].CoverPath + "/" + blogs[index].CoverName
	}

	response.Success(context, gin.H{"blogs": blogs}, "博客列表加载成功")
}
