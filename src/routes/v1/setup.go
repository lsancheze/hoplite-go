//SPDX-FileCopyrightText: © 2023 3nets, Inc. <it@3nets.io>

package v1

import (
	api "hoplite/api"
	middlewares "hoplite/middlewares"
	sites "hoplite/routes/v1/sites"
	templates "hoplite/routes/v1/templates"
)

func Setup(restApi *api.RestApi, uri string) {
	// v1 Router
	v1router := restApi.Router.Group(uri)
	v1router.Use(middlewares.ErrorMiddleware)

	// Entity routers
	sites.Router(v1router, "/sites")
	templates.Router(v1router, "/templates")
}
