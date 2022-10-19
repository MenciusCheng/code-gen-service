package http

import "github.com/gin-gonic/gin"

func initRoute(r *gin.Engine) {

	r.GET("/ping", ping)

	r.POST("/api/text/gen", TextGen)
}
