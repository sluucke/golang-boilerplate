package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/sluucke/golang-boilerplate/docs"
)

// NewRouter -.
// Swagger spec:
// @title       Golang Boilerplate
// @description Just a simple golang boilerplate with clean arch
// @version     1.0
// @host        localhost:3000
// @BasePath    /v1
func NewRouter(handler *gin.Engine) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	handler.GET("/swagger/*any", swaggerHandler)

	handler.GET("/health", func(c *gin.Context) { c.Status(http.StatusOK) })

	handler.GET("/metrics", gin.WrapH(promhttp.Handler()))

	h := handler.Group("/v1")
	{
		newHelloRoutes(h)
	}
}
