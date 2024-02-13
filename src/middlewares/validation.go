//SPDX-FileCopyrightText: © 2023 3nets, Inc. <it@3nets.io>

package middlewares

import (
	regex "hoplite/utils"
	customErrors "hoplite/utils/errors"

	gin "github.com/gin-gonic/gin"
	binding "github.com/gin-gonic/gin/binding"
	validator "github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
	validate.RegisterValidation("resourceId", regex.ResourceId)
	validate.RegisterValidation("name", regex.Name)
	validate.RegisterValidation("description", regex.Description)
}

func ValidationMiddleware(validator interface{}) gin.HandlerFunc {
	return func(context *gin.Context) {
		err := context.ShouldBindBodyWith(validator, binding.JSON)
		if err != nil {
			context.Error(&gin.Error{
				Err: customErrors.ServiceError{
					Status:  422,
					Code:    "bind-error",
					Message: err.Error(),
				},
				Type: gin.ErrorTypePrivate,
			})
			context.Abort()
			return
		}

		if err := validate.Struct(validator); err != nil {
			context.Error(&gin.Error{
				Err: customErrors.ServiceError{
					Status:  422,
					Code:    "validation-failed",
					Message: err.Error(),
				},
				Type: gin.ErrorTypePrivate,
			})
			context.Abort()
			return
		}

		context.Next()
	}
}
