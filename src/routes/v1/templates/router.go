//SPDX-FileCopyrightText: © 2023 3nets, Inc. <it@3nets.io>

package templates

import (
	middlewares "hoplite/middlewares"
	services "hoplite/services/templates"

	gin "github.com/gin-gonic/gin"
)

func Router(routerGroup *gin.RouterGroup, uri string) {
	group := routerGroup.Group(uri)

	// Entity routes
	group.GET("/", middlewares.RouteMiddleware(services.GetTemplate))
}
