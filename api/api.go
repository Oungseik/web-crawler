package api

import (
	"github.com/gin-gonic/gin"
	routers "web-crawler/api/routers"
)

func RegisterRoutes(router *gin.Engine) {
	apiGroup := router.Group("/api")

	routers.RegisterHealthRoutes(apiGroup)
}
