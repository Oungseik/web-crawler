package main

import (
	"github.com/gin-gonic/gin"
	swagger "github.com/num30/gin-swagger-ui"
	"github.com/wI2L/fizz"
	"github.com/wI2L/fizz/openapi"
	"web-crawler/api"
)

func main() {
	r := gin.Default()
	r.Use(gin.Logger())

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Next()
	})

	f := fizz.NewFromEngine(r)
	info := &openapi.Info{
		Title:       "Service API",
		Description: "API to perform web crawling",
		Version:     "0.1",
	}

	f.GET("/api-docs/openapi.json", nil, f.OpenAPI(info, "json"))
	swagger.AddOpenApiUIHandler(r, "/api-docs/ui", "/api-docs/openapi.json")

	api.RegisterRoutes(f)

	f.Engine().Run(":8080")
}
