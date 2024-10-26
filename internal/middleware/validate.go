package middleware

import (
	"example/internal/common/helper/loghelper"
	"example/internal/common/helper/responsehelper"
	"example/internal/common/helper/validatehelper"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

func ValidateRequestMiddleware(obj any) gin.HandlerFunc {
	return func(c *gin.Context) {
		appC := responsehelper.Gin{
			C: c,
		}

		objType := reflect.TypeOf(obj).Elem()
		objValue := reflect.New(objType).Interface()

		// Bind JSON input dynamically to the passed struct type
		if err := c.ShouldBindBodyWithJSON(&objValue); err != nil {
			loghelper.Logger.Errorf("Got error while binding body, err: %v", err)
			appC.Response(http.StatusBadRequest, responsehelper.INVALID_PARAMS, nil)
			c.Abort()
			return
		}

		validate := validatehelper.NewValidate()
		// Validate the dynamically bound struct
		if err := validate.ValidateStruct(objValue); err != nil {
			loghelper.Logger.Errorf("Got error while validate input, err: %v", err)
			appC.Response(http.StatusBadRequest, responsehelper.INVALID_PARAMS, nil)
			c.Abort()
			return
		}

		// Set the validated struct in the context for further use in handlers
		c.Next()
	}
}
