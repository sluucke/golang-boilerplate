package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sluucke/go-boilerplate/bootstrap"
	"github.com/sluucke/go-boilerplate/domain"
)

type IndexController struct {
	Env *bootstrap.Env
}

func (lc *IndexController) Index(c *gin.Context) {
	indexReponse := domain.IndexResponse{
		Message: "Hello NIGGA",
	}

	c.JSON(http.StatusOK, indexReponse)
}
