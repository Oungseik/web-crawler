package users

import (
	"github.com/gin-gonic/gin"
)

type GetUserReq struct {
	Id    string `path:"id"`
	Token string `header:"Authorization"`
}

type GetUserRes struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Token string `json:"token"`
}

func GetUserById(c *gin.Context, in *GetUserReq) (*GetUserRes, error) {
	return &GetUserRes{Id: in.Id, Token: in.Token, Name: "John Doe"}, nil
}
