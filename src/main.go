//SPDX-FileCopyrightText: © 2021 3nets, Inc. <it@3nets.io>

package main

import (
	api "hoplite/api"
	log "hoplite/logger"

	v1router "hoplite/routes/v1"
)

func main() {
	var err error

	log.Init("Hoplite - Golang Gin")
	log.Level(log.INFO)
	log.Write("main", log.INFO, "GIN Test")

	log.Write("main", log.INFO, "Start REST API...")
	restApi, err := api.Init("0.0.0.0:8088")
	if err != nil {
		log.Write("main", log.ERROR, err.Error())
	}

	log.Write("main", log.INFO, "Set up Router...")
	// v1 router
	v1router.Setup(restApi, "/api/v1")

	log.Write("main", log.INFO, "Run REST API Server...")
	err = restApi.Run()
	if err != nil {
		log.Write("main", log.ERROR, err.Error())
	}
}
