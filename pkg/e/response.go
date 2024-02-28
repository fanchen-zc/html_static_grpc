package e

import (
	"github.com/gin-gonic/gin"
	"time"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Ret       int         `json:"ret"`
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	Data      interface{} `json:"ret_data"`
	Timestamp int64       `json:"timestamp"`
}

func (g *Gin) Res(httpCode, ret, errCode int, msg string, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code:      errCode,
		Msg:       msg,
		Ret:       ret,
		Data:      data,
		Timestamp: time.Now().Unix(),
	})
}

func (g *Gin) Success(msg string, data interface{}) {
	g.Res(SUCCESS, 0, 1, msg, data)
}

func (g *Gin) Fail(code int, msg string, data interface{}) {
	g.Res(SUCCESS, 1, code, msg, data)
}
