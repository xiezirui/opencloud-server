package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CORSMiddleWare() gin.HandlerFunc {
	return func(context *gin.Context) {
		origin := context.Request.Header.Get("Origin")
		context.Header("Access-Control-Allow-Origin", origin)
		context.Header("Access-Control-Max-Age", "86400")
		context.Header("Access-Control-Allow-Methods", "*")
		context.Header("Access-Control-Allow-Headers", "*")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
		context.Header("Access-Control-Allow-Credentials", "true")

		fmt.Println("-------")

		if context.Request.Method == http.MethodOptions {
			context.AbortWithStatus(200)
		}
		context.Next()

	}
}
