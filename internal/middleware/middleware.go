package middleware

import (
	"bytes"
	"encoding/json"
	"example/internal/common/helper/commonhelper"
	"example/internal/common/helper/jwthelper"
	"example/internal/common/helper/loghelper"
	"example/internal/common/helper/redishelper"
	"example/internal/common/helper/responsehelper"
	"example/internal/dto"
	"example/internal/repository"
	"github.com/google/uuid"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type (
	Middleware interface {
		GetBodyDataInMiddleware(appC responsehelper.Gin, req any) error
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

func (m *middleware) GetBodyDataInMiddleware(appC responsehelper.Gin, req any) error {
	var bodyBuffer bytes.Buffer
	reader := io.TeeReader(appC.C.Request.Body, &bodyBuffer)

	// Read the body in the middleware
	bodyBytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}

	// Parse JSON into req
	err = json.Unmarshal(bodyBytes, &req)
	if err != nil {
		loghelper.Logger.Errorf("Got error while parsing body request, err: %v", err)
	}
	// Reset the request body so it can be read again in the controller
	appC.C.Request.Body = ioutil.NopCloser(&bodyBuffer)

	return nil
}

func (m *middleware) SetTraceIdMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		appC := responsehelper.Gin{
			C: c,
		}
		req := &dto.BaseRequestDTO{}

		err := m.GetBodyDataInMiddleware(appC, req)
		if err != nil {
			loghelper.Logger.Errorf("Got error while getting body request, err: %v", err)
			appC.Response(http.StatusBadRequest, responsehelper.INVALID_PARAMS, nil)
			appC.C.Abort()
			return
		}

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
