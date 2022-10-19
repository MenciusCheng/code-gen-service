package http

import (
	"github.com/MenciusCheng/code-gen-service/conf"
	"github.com/MenciusCheng/code-gen-service/service"
	"github.com/MenciusCheng/code-gen-service/util/log"
	"github.com/gin-gonic/gin"
)

var (
	svc *service.Service

	httpServer *gin.Engine
)

// Init create a server and run it
func Init(s *service.Service, conf *conf.Config) {
	svc = s

	// new http server
	httpServer = gin.Default()

	// register handler with http route
	initRoute(httpServer)

	// start a http server
	go func() {
		if err := httpServer.Run(conf.Address); err != nil {
			log.Error("http server start failed, err %v", err)
		}
	}()
}

func Shutdown() {
	if svc != nil {
		svc.Close()
	}
}
