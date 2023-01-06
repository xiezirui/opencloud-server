package router

import (
	"github.com/gin-gonic/gin"
	"opencloud-server/controller"
	"opencloud-server/middleware"
)

func RunRouter(Server *gin.Engine) {
	Server.Use(middleware.CORSMiddleWare())
	Server.POST("/api/register", controller.Register)
	Server.POST("/api/login", controller.Login)
}
