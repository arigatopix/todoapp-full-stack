package app

import (
	e "server/pkg/e"

	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (g *Gin) Response(httpCode, eCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code:    eCode,
		Message: e.GetMessage(eCode),
		Data:    data,
	})
}
