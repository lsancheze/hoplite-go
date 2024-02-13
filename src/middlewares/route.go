//SPDX-FileCopyrightText: © 2023 3nets, Inc. <it@3nets.io>

package middlewares

import (
	"errors"
	http "net/http"

	gin "github.com/gin-gonic/gin"

	customErrors "hoplite/utils/errors"
)

func RouteMiddleware(fn func(context *gin.Context) ([]byte, error)) gin.HandlerFunc {
	return func(context *gin.Context) {
		result, err := fn(context)
		if err != nil {
			if serviceErr, ok := err.(customErrors.ServiceError); ok {
				context.Error(&gin.Error{
					Err:  serviceErr,
					Type: gin.ErrorTypePrivate,
				})
			} else {
				// If it's not a custom error, handle it as a generic error
				context.Error(errors.New("Internal Server Error"))
			}

			return
		}
		context.Data(http.StatusOK, "application/json", result)
	}
}
