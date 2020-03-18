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

// QueryOrgCharities defines the request of query organization charities list
func (h *RestHandler) QueryOrgCharities(c *gin.Context) {
	logger.Infof("Got query charities request")

	req := &structs.QueryOrgCharitiesRequest{}
	if err := c.Bind(req); err != nil {
		e := fmt.Errorf("invalid parameters, %s", err.Error())
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

	donationStats, err := h.srvcContext.DBStorage.QueryOrgCharities(params)
	if err != nil {
		e := fmt.Errorf("query charities error , %s", err.Error())
		logger.Error(e)
		c.JSON(http.StatusInternalServerError, rest.ErrorResponse(rest.DatabaseOperationFailed, e.Error()))
		return
	}

	for i := 0; i < len(donationStats); i++ {
		donationStats[i].ConvertTime()
	}

	c.JSON(http.StatusOK, rest.SuccessResponse(&structs.QueryOrgCharitiesResp{
		Total:     params.Total,
		PageNum:   params.PageNum,
		PageLimit: params.PageLimit,
		StartTime: params.StartTime,
		EndTime:   params.EndTime,
		Results:   donationStats,
	}))
	logger.Info("response query charities success.")
	return
}

// QueryOrgCharitiesDetail defines the request of query charities detail
func (h *RestHandler) QueryOrgCharitiesDetail(c *gin.Context) {
	logger.Infof("Got query charities detail request")

	req := &structs.OrgCharitiesDetailRequest{}
	if err := c.Bind(req); err != nil {
		e := fmt.Errorf("invalid parameters, %s", err.Error())
		logger.Error(e)
		c.JSON(http.StatusBadRequest, rest.ErrorResponse(rest.ParseRequestParamsError, e.Error()))
		return
	}
	logger.Debugf("request params, %v", req)

	item, err := h.srvcContext.DBStorage.QueryOrgCharitiesDetail(req.UID)
	if err != nil {
		e := fmt.Errorf("query charities detail error , %s", err.Error())
		logger.Error(e)
		c.JSON(http.StatusInternalServerError, rest.ErrorResponse(rest.DatabaseOperationFailed, e.Error()))
		return
	}

	c.JSON(http.StatusOK, rest.SuccessResponse(&structs.OrgCharitiesDetailResp{
		UID:         item.UID,
		URL:         item.URL,
		NickName:    item.NickName,
		Address:     item.Province + item.City + item.District + item.Address,
		Phone:       item.Phone,
		BankCardNum: item.BankCardNum,
		Remark:      item.Remark,
	}))

	logger.Info("response query charities detail success.")
	return
}
