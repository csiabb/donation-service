/*
Copyright ArxanChain Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package version

import (
	"net/http"

	"github.com/csiabb/donation-service/common/metadata"
	"github.com/csiabb/donation-service/common/rest"

	"github.com/gin-gonic/gin"
)

// Version query service version
func (h *RestHandler) Version(c *gin.Context) {
	logger.Infof("Got query service version request")

	c.JSON(http.StatusOK, rest.SuccessResponse(metadata.ProgramVersion.FullVersion()))

	logger.Infof("Response query service version success.")
	return
}
