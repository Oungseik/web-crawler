package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func AuthMw(c *gin.Context) {

	fmt.Println("handling authentication middleware")

	c.Next()
}
