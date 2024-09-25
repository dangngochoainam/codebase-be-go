package middleware

import (
	"example/internal/common/helper/commonhelper"
	"example/internal/dto"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func SetTraceIdMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := &dto.BaseRequestDTO{}

		err := c.ShouldBind(req)
		if err != nil {
			fmt.Errorf("Error: %s\n", err)
		}
		if req.TraceId == "" {
			req.TraceId = uuid.New().String()
		}

		fmt.Printf("The request with uuid %s is started \n", req.TraceId)

		c.Set(string(commonhelper.HeaderKeyType_TraceId), req.TraceId)

		c.Next()

		fmt.Printf("The request with uuid %s is served \n", req.TraceId)
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
