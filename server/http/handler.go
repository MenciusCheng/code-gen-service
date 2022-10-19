package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ping(c *gin.Context) {
	if err := svc.Ping(c.Request.Context()); err != nil {
		errMsg := map[string]string{"errorMsg": err.Error()}
		c.JSON(http.StatusInternalServerError, errMsg)
		return
	}
	okMsg := map[string]string{"result": "ok"}
	c.JSON(http.StatusOK, okMsg)
}
