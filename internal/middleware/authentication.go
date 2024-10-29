package middleware

import (
	"context"
	"errors"
	"example/internal/common/helper/commonhelper"
	"example/internal/common/helper/loghelper"
	"example/internal/common/helper/redishelper"
	"example/internal/common/helper/responsehelper"
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
)

func (m *middleware) AuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		appC := responsehelper.Gin{
			C: c,
		}

		if slices.Contains(m.anonymousAccessURLs, appC.C.Request.URL.Path) {
			appC.C.Next()
			return
		}

		accessToken := appC.C.GetHeader("Authorization")
		if accessToken == "" {
			loghelper.Logger.Errorf("Got error while getting token, err: %v", errors.New("Authorization header not found"))
			appC.Response(http.StatusInternalServerError, responsehelper.INVALID_PARAMS, nil)
			appC.C.Abort()
			return
		}

		accessToken = accessToken[len("Bearer "):]
		tokenPayloadPublic, err := m.jwtHelper.VerifyToken(accessToken)
		if err != nil {
			loghelper.Logger.Errorf("Got error while verifing token, err: %v", errors.New("Token invalid"))
			appC.Response(http.StatusInternalServerError, responsehelper.INVALID_PARAMS, nil)
			appC.C.Abort()
			return
		}

		var userId string
		err = m.redisSession.Get(context.Background(), redishelper.GenerateRedisSessionKey(redishelper.ACCESS_TOKEN, tokenPayloadPublic.Key), &userId)
		if err != nil {
			loghelper.Logger.Errorf("Got error while getting userId, err: %v", err)
			appC.Response(http.StatusInternalServerError, responsehelper.INVALID_PARAMS, nil)
			appC.C.Abort()
			return
		}

		userEntity, err := m.userRepository.FindUserById(userId)
		if err != nil {
			loghelper.Logger.Errorf("Got error while finding user, err: %v", err)
			appC.Response(http.StatusInternalServerError, responsehelper.INVALID_PARAMS, nil)
			appC.C.Abort()
			return
		}

		appC.C.Set(string(commonhelper.ContextKey_Key), tokenPayloadPublic.Key)
		appC.C.Set(string(commonhelper.ContextKey_User), userEntity)
		appC.C.Next()
	}
}
