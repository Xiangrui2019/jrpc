package routers

import (
	"jrpc/api"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/jrpc/api/v1")
	{
		v1.POST("/ping", api.Ping)
	}

	return router
}
