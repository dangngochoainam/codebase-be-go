package middleware

import (
	"example/internal/common/helper/loghelper"
	"example/internal/common/helper/responsehelper"
	"example/internal/common/helper/validatehelper"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

func (m *middleware) ValidateRequestMiddleware(obj any) gin.HandlerFunc {
	return func(c *gin.Context) {
		appC := responsehelper.Gin{
			C: c,
		}

		objType := reflect.TypeOf(obj).Elem()
		objValue := reflect.New(objType).Interface()

		if err := appC.C.ShouldBindBodyWithJSON(&objValue); err != nil {
			loghelper.Logger.Errorf("Got error while binding body, err: %v", err)
			appC.Response(http.StatusBadRequest, responsehelper.INVALID_PARAMS, nil)
			appC.C.Abort()
			return
		}

		validate := validatehelper.NewValidate()
		if err := validate.ValidateStruct(objValue); err != nil {
			loghelper.Logger.Errorf("Got error while validate input, err: %v", err)
			appC.Response(http.StatusBadRequest, responsehelper.INVALID_PARAMS, nil)
			appC.C.Abort()
			return
		}

		appC.C.Next()
	}
}
