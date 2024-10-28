package middleware

import (
	"errors"
	"example/internal/common/helper/loghelper"
	"example/internal/common/helper/responsehelper"
	"github.com/gin-gonic/gin"
	"net/http"
	"slices"
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
		err := m.jwtHelper.VerifyToken(accessToken)
		if err != nil {
			loghelper.Logger.Errorf("Got error while verifing token, err: %v", errors.New("Token invalid"))
			appC.Response(http.StatusInternalServerError, responsehelper.INVALID_PARAMS, nil)
			appC.C.Abort()
			return
		}

		appC.C.Next()
	}
}
