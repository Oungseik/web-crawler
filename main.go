package main

import (
	"github.com/gin-gonic/gin"
	"web-crawler/api"
)

func main() {
	router := gin.Default()

	api.RegisterRoutes(router)

	router.Run(":8080")
}
