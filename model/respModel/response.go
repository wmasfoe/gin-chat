package respModel

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func Result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

const (
	SuccessCode = http.StatusOK
	ErrorCode   = 0
)

func SuccessResp(data interface{}, msg string, c *gin.Context) {
	Result(SuccessCode, data, msg, c)
}
func SuccessWithData(data interface{}, c *gin.Context) {
	SuccessResp(data, "成功", c)
}
func SuccessWithMsg(msg string, c *gin.Context) {
	SuccessResp(make(map[string]interface{}), msg, c)
}

func FailResp(data interface{}, msg string, c *gin.Context) {
	Result(ErrorCode, data, msg, c)
}
func FailWithMsg(msg string, c *gin.Context) {
	FailResp(make(map[string]interface{}), msg, c)
}
func FailWithCode(code int, c *gin.Context) {
	msg, ok := CodeMsgMapping[code]
	if ok {
		Result(code, make(map[string]interface{}), msg, c)
		return
	}
	Result(ErrorCode, make(map[string]interface{}), "未知错误", c)
}
