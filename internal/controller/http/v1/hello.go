package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type helloRoutes struct {
}

func newHelloRoutes(handler *gin.RouterGroup) {
	r := &helloRoutes{}

	h := handler.Group("/hello")
	{
		h.GET("/", r.world)
	}
}

type worldResponse struct {
	Message string `json:"message"`
}

// @Summary     Hello World
// @Description Show hello world
// @Tags  	    hello
// @Success     200 {object} worldResponse
// @Router      /hello [get]
func (r *helloRoutes) world(c *gin.Context) {
	c.JSON(http.StatusOK, worldResponse{
		Message: "Hello World",
	})
}
