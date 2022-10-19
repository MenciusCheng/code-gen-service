package main

import (
	"github.com/MenciusCheng/code-gen-service/conf"
	"github.com/MenciusCheng/code-gen-service/server/http"
	"github.com/MenciusCheng/code-gen-service/service"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// init local config
	cfg, err := conf.Init()
	if err != nil {
		panic(err)
	}

	// create a service instance
	srv := service.New(cfg)

	// init and start http server
	http.Init(srv, cfg)

	defer http.Shutdown()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-sigChan
		log.Printf("get a signal %s\n", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Println("public.wechat.service server exit now...")
			return
		case syscall.SIGHUP:
		default:
		}
	}
}
