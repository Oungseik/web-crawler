package health

import "github.com/gin-gonic/gin"

type GetCheckHealthReq struct{}

type GetCheckHealthRes struct {
	message string
}

func CheckHealth(c *gin.Context, in *GetCheckHealthReq) (*GetCheckHealthRes, error) {
	return &GetCheckHealthRes{message: "server is up and running"}, nil
}
