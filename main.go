package main

import (
	"github.com/gin-gonic/gin"
	swagger "github.com/num30/gin-swagger-ui"
	"github.com/wI2L/fizz"
	"github.com/wI2L/fizz/openapi"
	"web-crawler/api"
)

func main() {
	router := gin.Default()
	router.Use(gin.Logger())

	f := fizz.NewFromEngine(router)
	info := &openapi.Info{
		Title:       "Service API",
		Description: "API to perform web crawling",
		Version:     "0.1",
	}

	f.GET("/openapi.json", nil, f.OpenAPI(info, "json"))
	swagger.AddOpenApiUIHandler(router, "/api-docs", "/openapi.json")

	api.RegisterRoutes(f)

	f.Engine().Run(":8080")
}
