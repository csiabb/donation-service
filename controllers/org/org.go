/*
Copyright Lingzhu Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package org

import (
	"fmt"
	"net/http"

	"github.com/csiabb/donation-service/common/rest"
	"github.com/csiabb/donation-service/structs"

	"github.com/gin-gonic/gin"
)

// QueryOrganizations defines the request of query organizations list
func (h *RestHandler) QueryOrganizations(c *gin.Context) {
	logger.Infof("Got query organizations request")

	req := &structs.QueryOrganizationsRequest{}
	if err := c.BindQuery(req); err != nil {
		e := fmt.Errorf("invalid parameters: %s", err.Error())
		logger.Error(e)
		c.JSON(http.StatusBadRequest, rest.ErrorResponse(rest.ParseRequestParamsError, e.Error()))
		return
	}
	logger.Debugf("request params, %v", req)

	params := &structs.QueryParams{
		PageNum:   req.PageNum,
		PageLimit: req.PageLimit,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
	}

	donationStats, err := h.srvcContext.DBStorage.QueryOrganizations(params)
	if err != nil {
		e := fmt.Errorf("query organizations error : %s", err.Error())
		logger.Error(e)
		c.JSON(http.StatusInternalServerError, rest.ErrorResponse(rest.DatabaseOperationFailed, e.Error()))
		return
	}

	c.JSON(http.StatusOK, rest.SuccessResponse(&structs.QueryOrganizationsResp{
		Total:     params.Total,
		PageNum:   params.PageNum,
		PageLimit: params.PageLimit,
		StartTime: params.StartTime,
		EndTime:   params.EndTime,
		Results:   donationStats,
	}))
	logger.Info("response query organizations success.")
	return
}
