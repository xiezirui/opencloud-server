package main

import (
	"github.com/gin-gonic/gin"
	"opencloud-server/router"
)

func main() {
	Server := gin.Default()

	router.RunRouter(Server)

	Server.Run(":1016")
}
