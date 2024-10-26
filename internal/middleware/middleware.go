package middleware

import (
	"example/internal/common/helper/commonhelper"
	"example/internal/common/helper/loghelper"
	"example/internal/dto"
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

		loghelper.Logger.Debugf("The request with traceId %s is started", req.TraceId)
		c.Set(string(commonhelper.HeaderKeyType_TraceId), req.TraceId)

		c.Next()

		loghelper.Logger.Debugf("The request with traceId %s is served", req.TraceId)
	}
}

func SetTimeMsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		traceId := c.GetString(string(commonhelper.HeaderKeyType_TraceId))
		executeTimeDuration := time.Since(start)
		loghelper.Logger.Debugf("Execute time duration of traceId %s is %s", traceId, executeTimeDuration)
	}
}
