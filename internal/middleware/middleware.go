package middleware

import (
	"example/internal/common/helper/commonhelper"
	"example/internal/common/helper/jwthelper"
	"example/internal/common/helper/loghelper"
	"example/internal/common/helper/redishelper"
	"example/internal/dto"
	"example/internal/repository"
	"time"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

type (
	Middleware interface {
		SetTimeMsMiddleware() gin.HandlerFunc
		SetTraceIdMiddleware() gin.HandlerFunc
		ValidateRequestMiddleware(obj any) gin.HandlerFunc
		AuthenticationMiddleware() gin.HandlerFunc
	}
	middleware struct {
		jwtHelper           jwthelper.JwtHelper
		redisSession        redishelper.RedisSessionHelper
		anonymousAccessURLs []string
		userRepository      repository.UserRepository
	}
)

func NewMiddleware(jwtHelper jwthelper.JwtHelper, redisSession redishelper.RedisSessionHelper, anonymousAccessURLs []string, userRepository repository.UserRepository) Middleware {
	return &middleware{
		jwtHelper:           jwtHelper,
		redisSession:        redisSession,
		anonymousAccessURLs: anonymousAccessURLs,
		userRepository:      userRepository,
	}
}

func (m *middleware) SetTraceIdMiddleware() gin.HandlerFunc {
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

func (m *middleware) SetTimeMsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		traceId := c.GetString(string(commonhelper.HeaderKeyType_TraceId))
		executeTimeDuration := time.Since(start)
		loghelper.Logger.Debugf("Execute time duration of traceId %s is %s", traceId, executeTimeDuration)
	}
}
