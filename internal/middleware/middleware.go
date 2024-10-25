package middleware

import (
	"example/internal/common/helper/commonhelper"
	"example/internal/common/helper/loghelper"
	"example/internal/dto"
	"fmt"
	"github.com/google/uuid"
	"time"

	"github.com/gin-gonic/gin"
)

func SetTraceIdMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := &dto.BaseRequestDTO{}

		_ = c.ShouldBindBodyWithJSON(req)

		if req.TraceId == "" {
			req.TraceId = uuid.New().String()
		}

		loghelper.Logger.Debug("The request with uuid %s is started \n", req.TraceId)

		c.Set(string(commonhelper.HeaderKeyType_TraceId), req.TraceId)

		c.Next()

		loghelper.Logger.Debugf("The request with uuid %s is served \n", req.TraceId)
	}
}

func SetTimeMsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		executeTimeDuration := time.Since(start)
		fmt.Println("executeTimeDuration", executeTimeDuration)
	}
}
