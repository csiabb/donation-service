/*
Copyright Lingzhu Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package image

import (
	"fmt"
	"net/http"

	"github.com/csiabb/donation-service/common/rest"
	"github.com/csiabb/donation-service/common/utils"
	"github.com/csiabb/donation-service/structs"

	"github.com/gin-gonic/gin"
)

// Upload defines the upload of image
func (h *RestHandler) Upload(c *gin.Context) {
	logger.Info("got image upload request")

	c.Request.ParseMultipartForm(32 << 20)
	fileRec, _, err := c.Request.FormFile("image_file")
	if err != nil {
		e := fmt.Errorf("invalid parameters: %s", err.Error())
		logger.Error(e)
		c.JSON(http.StatusBadRequest, rest.ErrorResponse(rest.InvalidParamsErrCode, e.Error()))
		return
	}
	logger.Debugf("request params %v", fileRec)

	imageFileTag := utils.GenerateUUID()

	if err = h.srvcContext.ALiYunServices.UploadObject(imageFileTag+".png", fileRec); err != nil {
		e := fmt.Errorf("image upload error : %s", err.Error())
		logger.Error(e)
		c.JSON(http.StatusBadRequest, rest.ErrorResponse(rest.InternalServerFailure, e.Error()))
		return
	}

	c.JSON(http.StatusOK, rest.SuccessResponse(&structs.ImageUploadResp{
		ID: imageFileTag,
	}))

	logger.Info("response image upload success.")
	return
}
