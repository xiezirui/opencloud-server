package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"opencloud-server/router"
)

func main() {
	Server := gin.Default()

	Server.StaticFS("/static", http.Dir("./static"))

	router.RunRouter(Server)

	Server.Run(":1016")
}
