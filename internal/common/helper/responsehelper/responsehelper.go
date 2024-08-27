package responsehelper

import "github.com/gin-gonic/gin"

type Gin struct {
	C *gin.Context
}

type Response struct {
	SystemCode SystemCode  `json:"systemCode"`
	Msg        string      `json:"msg"`
	Data       interface{} `json:"data"`
}

func (c *Gin) Response(httpCode int, systemCode SystemCode, data interface{}) {
	c.C.JSON(httpCode, Response{
		SystemCode: systemCode,
		Msg:        GetMsg(systemCode),
		Data:       data,
	})
	return
}
