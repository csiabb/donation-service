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
		e := fmt.Errorf("invalid parameters, %s", err.Error())
		logger.Error(e)
		c.JSON(http.StatusBadRequest, rest.ErrorResponse(rest.InvalidParamsErrCode, e.Error()))
		return
	}
	logger.Debugf("request params %v", fileRec)

	imageFileTag := utils.GenerateUUID()

	if err = h.srvcContext.ALiYunBackend.UploadObject(imageFileTag+".png", fileRec); err != nil {
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

// Share define share of image
func (h *RestHandler) Share(c *gin.Context) {
	logger.Infof("Got query share request")

	req := &structs.ShareRequest{}
	if err := c.Bind(req); err != nil {
		e := fmt.Errorf("invalid parameters: %s", err.Error())
		logger.Error(e)
		c.JSON(http.StatusBadRequest, rest.ErrorResponse(rest.ParseRequestParamsError, e.Error()))
		return
	}
	logger.Debugf("request params, %v", req)

	var err error
	var content, imagURL string
	if req.ShareType == rest.Prove {
		content = rest.ShareDonationContent
		if imagURL, err = h.GetImageURL(req); err != nil {
			e := fmt.Errorf("get image url err : %s", err.Error())
			logger.Error(e)
			c.JSON(http.StatusBadRequest, rest.ErrorResponse(rest.ParseRequestParamsError, e.Error()))
			return
		}
	} else {
		content = rest.ShareHomeContent
		imagURL = rest.ShareImageURL
	}

	c.JSON(http.StatusOK, rest.SuccessResponse(&structs.ShareResp{
		Icon:     rest.ShareIcon,
		Title:    rest.ShareTitle,
		Content:  content,
		ImageURL: imagURL,
	}))

	logger.Info("response share success.")
	return

}
