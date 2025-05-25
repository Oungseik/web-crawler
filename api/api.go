package api

import (
	"github.com/gin-gonic/gin"
	"web-crawler/api/middlewares"
	"web-crawler/api/routers"
)

func RegisterRoutes(r *gin.Engine) {
	apiGroup := r.Group("/api")

	apiGroup.Use(middlewares.AuthMw)
	routers.RegisterHealthRoutes(apiGroup)
}
