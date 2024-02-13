//SPDX-FileCopyrightText: © 2023 3nets, Inc. <it@3nets.io>

package sites

import (
	"fmt"

	siteEntities "hoplite/entities/sites"

	gin "github.com/gin-gonic/gin"
	binding "github.com/gin-gonic/gin/binding"
)

func CreateSite(context *gin.Context) ([]byte, error) {
	// Logic
	var body siteEntities.Sites
	context.ShouldBindBodyWith(&body, binding.JSON)

	fmt.Println(body.ProviderId)

	// Simulating successful response
	return []byte(`{"_id": "12baf23", "name": "siteA"}`), nil
}
