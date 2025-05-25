package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterHealthRoutes(rg *gin.RouterGroup) {
	health := rg.Group("/health")
	health.GET("", getHealth)
}

func getHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "server is up and running"})
}
