package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sluucke/go-boilerplate/api/route"
	"github.com/sluucke/go-boilerplate/bootstrap"
)

func main() {
	app := bootstrap.App()

	env := app.Env

	releaseMode := "debug"

	if env.AppEnv == "prod" {
		releaseMode = "release"
	}

	gin.SetMode(releaseMode)
	gin := gin.Default()

	route.Setup(env, gin)

	gin.Run(env.ServerAddr)
}
