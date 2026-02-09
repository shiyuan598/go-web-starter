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

// Login godoc
// @Summary 登录
// @Tags Auth
// @Accept json
// @Produce json
// @Param data body LoginReq true "登录参数"
// @Success 200 {object} response.Result
// @Router /login [post]
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

// ListUsers godoc
// @Summary 用户列表
// @Tags User
// @Security ApiKeyAuth
// @Success 200 {object} response.Result
// @Router /users [get]
func ListUsers(c *gin.Context) {
	users := dao.List()
	response.OK(c, users)
}

// CreateUser godoc
// @Summary 创建用户
// @Tags User
// @Accept json
// @Security ApiKeyAuth
// @Param data body dao.UserCreateReq true "用户信息"
// @Success 200 {object} response.Result
// @Router /users [post]
func CreateUser(c *gin.Context) {
	var u dao.UserCreateReq
	if err := c.ShouldBindJSON(&u); err != nil {
		response.Fail(c, err.Error())
		return
	}
	dao.Create(u)
	response.OK(c, nil)
}
