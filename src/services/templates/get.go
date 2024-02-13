//SPDX-FileCopyrightText: © 2023 3nets, Inc. <it@3nets.io>

package templates

import (
	gin "github.com/gin-gonic/gin"

	customErrors "hoplite/utils/errors"
)

func GetTemplate(context *gin.Context) ([]byte, error) {
	// Logic

	// Simulating an error
	return nil, customErrors.ServiceError{
		Status:  404,
		Code:    "get-template-failed",
		Message: "Get template failed",
		Details: "Some details here",
		TraceId: "traceid0x",
	}
}
