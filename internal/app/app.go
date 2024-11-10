package app

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/sluucke/golang-boilerplate/config"
	v1 "github.com/sluucke/golang-boilerplate/internal/controller/http/v1"
	"github.com/sluucke/golang-boilerplate/pkg/httpserver"
)

func Run(cfg *config.Env) {

	handler := gin.New()
	v1.NewRouter(handler)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.Port))

	err := <-httpServer.Notify()
	//  Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Fatal("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		log.Fatal(fmt.Errorf("app - Run - httpServer.Notify: %w", err))

	}

	err = httpServer.Shutdown()

	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
