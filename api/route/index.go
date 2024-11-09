package route

import (
	"github.com/gin-gonic/gin"
	"github.com/sluucke/go-boilerplate/api/controller"
	"github.com/sluucke/go-boilerplate/bootstrap"
)

func NewIndexRoute(env *bootstrap.Env, group gin.RouterGroup) {

	indexController := controller.IndexController{
		Env: env,
	}

	group.GET("/", indexController.Index)
}
