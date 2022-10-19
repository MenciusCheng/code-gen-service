package http

import (
	"github.com/MenciusCheng/code-gen-service/model"
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

func TextGen(c *gin.Context) {
	var req model.TextGenReq

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := svc.TextGen(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": res})
}
