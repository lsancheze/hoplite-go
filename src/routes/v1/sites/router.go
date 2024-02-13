//SPDX-FileCopyrightText: © 2023 3nets, Inc. <it@3nets.io>

package sites

import (
	siteEntities "hoplite/entities/sites"
	middlewares "hoplite/middlewares"
	services "hoplite/services/sites"

	gin "github.com/gin-gonic/gin"
)

func Router(routerGroup *gin.RouterGroup, uri string) {
	group := routerGroup.Group(uri)

	// Entity validator schema
	sitesValidator := &siteEntities.Sites{}
	// Entity routes
	group.POST("/", middlewares.ValidationMiddleware(sitesValidator), middlewares.RouteMiddleware(services.CreateSite))
}
