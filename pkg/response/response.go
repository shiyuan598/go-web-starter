package response

import "github.com/gin-gonic/gin"

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func OK(c *gin.Context, data interface{}) {
	c.JSON(200, Result{Code: 0, Msg: "ok", Data: data})
}

func Fail(c *gin.Context, msg string) {
	c.JSON(200, Result{Code: -1, Msg: msg})
}
