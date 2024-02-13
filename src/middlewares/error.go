//SPDX-FileCopyrightText: © 2023 3nets, Inc. <it@3nets.io>

package middlewares

import (
	http "net/http"

	gin "github.com/gin-gonic/gin"

	customErrors "hoplite/utils/errors"
)

func ErrorMiddleware(context *gin.Context) {
	context.Next()

	if err := context.Errors.Last(); err != nil {
		if serviceErr, ok := err.Err.(customErrors.ServiceError); ok {
			context.JSON(serviceErr.Status, serviceErr)
			return
		}

		context.JSON(http.StatusInternalServerError, gin.H{
			"code":    "internal-server-error",
			"message": "Something went wrong.",
		})
		return
	}
}
