package api

import (
	"go-web-starter/internal/dao"
	"go-web-starter/pkg/jwt"
	"go-web-starter/pkg/response"

	"github.com/gin-gonic/gin"
)

type LoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var req LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, err.Error())
		return
	}
	user, err := dao.GetByUsername(req.Username)
	if err != nil || user.Password != req.Password {
		response.Fail(c, "login failed")
		return
	}
	token, _ := jwt.Generate(user.ID)
	response.OK(c, gin.H{"token": token})
}

func ListUsers(c *gin.Context) {
	users := dao.List()
	response.OK(c, users)
}

func CreateUser(c *gin.Context) {
	var u dao.UserCreateReq
	if err := c.ShouldBindJSON(&u); err != nil {
		response.Fail(c, err.Error())
		return
	}
	dao.Create(u)
	response.OK(c, nil)
}
