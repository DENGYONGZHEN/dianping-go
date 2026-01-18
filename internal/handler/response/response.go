package response

import "github.com/gin-gonic/gin"

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(200, Response{
		Code: 0,
		Data: data,
	})
}

func Fail(c *gin.Context, code int, msg string) {
	c.JSON(200, Response{
		Code: code,
		Msg:  msg,
	})
}
