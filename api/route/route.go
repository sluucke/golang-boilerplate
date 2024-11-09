package route

import (
	"github.com/gin-gonic/gin"
	"github.com/sluucke/go-boilerplate/bootstrap"
)

func Setup(env *bootstrap.Env, gin *gin.Engine) {
	publicRouter := gin.Group("")

	NewIndexRoute(env, *publicRouter)
}
