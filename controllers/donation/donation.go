/*
Copyright Lingzhu Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package donation

import (
	"github.com/csiabb/donation-service/common/rest"
	"github.com/csiabb/donation-service/models"
	"github.com/csiabb/donation-service/service/donation"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Donation query service donation
func (h *RestHandler) ListDonations(c *gin.Context) {
	logger.Infof("Got query service donation request")

	req := &models.Request{}
	if err := c.BindQuery(req); err != nil {
		logger.Errorf("bind query err, request =%+v, err =%+v", req, err)
		c.JSON(http.StatusOK, rest.ErrorResponse(http.StatusBadRequest, ""))
		return
	}

	logger.Infof("init handler ...")
	handler, err := donation.NewDonations(h.srvcContext)
	if err != nil {
		logger.Errorf("init handler is err, request =%+v, err = %+v", req, err)
		c.JSON(http.StatusOK, rest.ErrorResponse(http.StatusBadRequest, ""))
		return
	}

	logger.Infof("Get a list of donations.")
	donations, err := handler.ListDonations(req)
	if err != nil {
		logger.Errorf("Error getting donation list, request =%+v, err = %+v", req, err)
		c.JSON(http.StatusOK, rest.ErrorResponse(http.StatusBadRequest, ""))
		return
	}

	c.JSON(http.StatusOK, rest.SuccessResponse(donations))
	logger.Infof("Response query service donation success.")
	return
}
