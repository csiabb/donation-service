/*
 * Copyright ArxanChain Ltd. 2020 All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package acc

import (
	"fmt"
	"net/http"

	"github.com/csiabb/donation-service/common/rest"
	"github.com/csiabb/donation-service/common/utils"
	"github.com/csiabb/donation-service/models"
	"github.com/csiabb/donation-service/structs"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// LoginWXApp defines the user login
func (h *RestHandler) LoginWXApp(c *gin.Context) {
	logger.Info("got login request")

	req := &structs.LoginRequest{}
	if err := c.BindJSON(req); err != nil {
		e := fmt.Errorf("invalid parameters, %s", err.Error())
		logger.Error(e)
		c.JSON(http.StatusBadRequest, rest.ErrorResponse(rest.ParseRequestParamsError, e.Error()))
		return
	}
	logger.Debugf("request params, %v", req)

	if req.CertCode == "" || (req.ID <= 0 && req.AppID == "") {
		logger.Error("invalid params cert code, id or app id")
		c.JSON(http.StatusBadRequest, rest.ErrorResponse(rest.InvalidParamsErrCode, "invalid params cert code, id or app id"))
		return
	}

	wxApp := h.srvcContext.Config.WXCfg
	wxCredentials, err := h.srvcContext.WXClient.WXLogin(wxApp.AppID, wxApp.Secret, req.CertCode)
	if err != nil {
		logger.Errorf("get user info by call wx service return %+v", err)
		c.JSON(http.StatusBadRequest, rest.ErrorResponse(rest.WXLoginFailed, err.Error()))
		return
	}
	logger.Debugf("get user info by call wx service return %+v", wxCredentials)

	if wxCredentials.OpenID == "" {
		logger.Error("request params is invalid")
		c.JSON(http.StatusBadRequest, rest.ErrorResponse(rest.InvalidParamsErrCode, "cert code is invalid"))
		return
	}

	user, err := h.srvcContext.DBStorage.QueryAccount(wxCredentials.OpenID, "")
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			logger.Errorf("query account return err %v", err)
			c.JSON(http.StatusInternalServerError, rest.ErrorResponse(rest.InternalServerFailure, err.Error()))
			return
		}
	} else {
		if user.ID != "" {
			logger.Debug("user already exists")
			c.JSON(http.StatusOK, rest.SuccessResponse(&structs.LoginResp{
				UID: user.ID,
			}))
			return
		}
	}

	id := utils.GenerateUUID()
	bcResp, err := h.srvcContext.IBCAdapter.Register(id)
	if err != nil {
		e := fmt.Errorf("register on block chain failed, %v", err)
		logger.Error(e)
		c.JSON(http.StatusInternalServerError, rest.ErrorResponse(rest.InternalServerFailure, e.Error()))
		return
	}

	acc := &models.Account{
		ID:       id,
		Access:   req.Access,
		Password: req.Password,
		NickName: req.Nickname,
		Phone:    req.Phone,
		Email:    req.Email,
		Remark:   req.Remark,
		OpenID:   wxCredentials.OpenID,
		AppID:    wxApp.AppID,
		UnionID:  wxCredentials.UnionID,
		DID:      bcResp.Data.ID,
	}

	err = h.srvcContext.DBStorage.CreateAccount(acc)
	if err != nil {
		e := fmt.Errorf("create account error, %v", err)
		logger.Error(e.Error())
		c.JSON(http.StatusInternalServerError, rest.ErrorResponse(rest.InternalServerFailure, e.Error()))
		return
	}

	c.JSON(http.StatusOK, rest.SuccessResponse(&structs.LoginResp{
		UID: acc.ID,
	}))
}
