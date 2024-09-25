package responsehelper

import (
	"example/internal/common/helper/commonhelper"

	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

type BaseResponseDTO struct {
	TraceId    string     `json:"traceId"`
	SystemCode SystemCode `json:"systemCode"`
	Message    string     `json:"message"`
	Data       any        `json:"data"`
}

func (c *Gin) Response(httpCode int, systemCode SystemCode, data any) {

	traceId, _ := c.C.Get(string(commonhelper.HeaderKeyType_TraceId))

	c.C.JSON(httpCode, BaseResponseDTO{
		TraceId:    traceId.(string),
		SystemCode: systemCode,
		Data:       data,
		Message:    GetMsg(systemCode),
	})
	return
}
