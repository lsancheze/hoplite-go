//SPDX-FileCopyrightText: © 2023 3nets, Inc. <it@3nets.io>

package restapi

import (
	log "hoplite/logger"

	gin "github.com/gin-gonic/gin"
)

type RestApi struct {
	Router *gin.Engine
	Addr   string
}

func Init(addr string) (restApi *RestApi, err error) {
	err = nil
	restApi = &RestApi{}
	restApi.Addr = addr
	restApi.Router = gin.Default()
	return restApi, err
}

func (restApi *RestApi) Run() (err error) {
	err = restApi.Router.Run(restApi.Addr)
	if err != nil {
		log.Write("restapi", log.ERROR, err.Error())
	}
	return err
}
