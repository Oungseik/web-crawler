package health

import "github.com/gin-gonic/gin"

type GetCheckHealthReq struct{}

type GetCheckHealthRes struct {
	Message string `json:"message"`
}

func CheckHealth(c *gin.Context, in *GetCheckHealthReq) (*GetCheckHealthRes, error) {
	return &GetCheckHealthRes{Message: "server is up and running"}, nil
}
