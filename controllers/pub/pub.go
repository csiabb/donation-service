/*
Copyright ArxanChain Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package pub

import (
	"fmt"
	"net/http"

	"github.com/csiabb/donation-service/common/rest"
	"github.com/csiabb/donation-service/models"
	"github.com/csiabb/donation-service/structs"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

// ReceiveFunds defines the request of received funds
func (h *RestHandler) ReceiveFunds(c *gin.Context) {
	logger.Info("got receive funds request")

	req := &structs.ReceiveFundsRequest{}
	if err := c.BindJSON(req); err != nil {
		e := fmt.Errorf("invalid parameters: %s", err.Error())
		logger.Error(e)
		c.JSON(http.StatusBadRequest, rest.ErrorResponse(rest.ParseRequestParamsError, e.Error()))
		return
	}
	logger.Debugf("request params, %v", req)

	if req.PubType != rest.PubTypeDonate && req.PubType != rest.PubTypeDistribute && req.PubType != rest.PubTypeReceive {
		e := fmt.Errorf("pub type invalid")
		logger.Error(e)
		c.JSON(http.StatusBadRequest, rest.ErrorResponse(rest.InvalidParamsErrCode, e.Error()))
		return
	}

	if req.Amount.LessThanOrEqual(decimal.NewFromInt(0)) {
		e := fmt.Errorf("amount can not less than 0")
		logger.Error(e)
		c.JSON(http.StatusBadRequest, rest.ErrorResponse(rest.InvalidParamsErrCode, e.Error()))
		return
	}

	funds := &models.PubFunds{
		UID:       req.UID,
		UserType:  req.UserType,
		AidUID:    req.AidUID,
		TargetUID: req.TargetUID,
		PubType:   req.PubType,
		PayType:   req.PayType,
		Amount:    req.Amount,
		Remark:    req.Remark,
	}

	err := h.srvcContext.DBStorage.CreateFunds(funds)
	if err != nil {
		e := fmt.Errorf("create funds error, %s", err.Error())
		logger.Error(e)
		c.JSON(http.StatusInternalServerError, rest.ErrorResponse(rest.DatabaseOperationFailed, e.Error()))
	}

	c.JSON(http.StatusOK, rest.SuccessResponse(nil))
	logger.Infof("response receive funds success.")
	return
}

// QueryFunds defines the request of query funds
func (h *RestHandler) QueryFunds(c *gin.Context) {
	logger.Info("got query funds request")

	req := &structs.QueryFundsRequest{}
	var err error
	if err = c.Bind(req); err != nil {
		e := fmt.Errorf("invalid parameters: %s", err.Error())
		logger.Error(e)
		c.JSON(http.StatusBadRequest, rest.ErrorResponse(rest.InvalidParamsErrCode, e.Error()))
		return
	}
	logger.Debugf("request params %v", req)

	params := &structs.QueryParams{
		PageNum:   req.PageNum,
		PageLimit: req.PageLimit,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
	}

	result, err := h.srvcContext.DBStorage.QueryFunds(req.UID, req.TargetUID, req.UserType, req.PubType, params)
	if err != nil {
		e := fmt.Errorf("query funds error, %s", err.Error())
		logger.Error(e)
		c.JSON(http.StatusInternalServerError, rest.ErrorResponse(rest.DatabaseOperationFailed, e.Error()))
		return
	}

	payload := make([]*structs.QueryFundsItems, 0)
	for _, v := range result {
		payload = append(payload, &structs.QueryFundsItems{
			ID:          v.ID,
			UID:         v.UID,
			DonorName:   v.DonorName,
			UserType:    v.UserType,
			AidUID:      v.AidUID,
			AidName:     v.AidName,
			TargetUID:   v.TargetUID,
			TargetName:  v.TargetName,
			PubType:     v.PubType,
			PayType:     v.PayType,
			Amount:      v.Amount.String(),
			TxID:        v.TxID,
			Remark:      v.Remark,
			BlockType:   v.BlockType,
			BlockHeight: v.BlockHeight,
			BlockTime:   v.BlockTime,
			CreatedAt:   v.CreatedAt.Unix(),
		})
	}

	c.JSON(http.StatusOK, rest.SuccessResponse(&structs.QueryFundsResp{
		Total:     params.Total,
		PageNum:   params.PageNum,
		PageLimit: params.PageLimit,
		StartTime: params.StartTime,
		EndTime:   params.EndTime,
		Results:   payload,
	}))
	logger.Info("response query funds success.")
	return
}

// QueryFundsDetail defines the detail information of funds
func (h *RestHandler) QueryFundsDetail(c *gin.Context) {
	logger.Info("got query funds detail request")

	req := &structs.FundsDetailRequest{}
	var err error
	if err = c.Bind(req); err != nil {
		e := fmt.Errorf("invalid parameters: %s", err.Error())
		logger.Error(e)
		c.JSON(http.StatusBadRequest, rest.ErrorResponse(rest.InvalidParamsErrCode, e.Error()))
		return
	}
	logger.Debugf("request params %v", req)

	f, err := h.srvcContext.DBStorage.QueryFundsDetail(req.FundsID)
	if err != nil {
		e := fmt.Errorf("query funds detail error, %s", err.Error())
		logger.Error(e)
		c.JSON(http.StatusInternalServerError, rest.ErrorResponse(rest.DatabaseOperationFailed, e.Error()))
		return
	}

	funds := structs.QueryFundsItems{
		ID:                f.Funds.ID,
		UID:               f.Funds.UID,
		DonorName:         f.Funds.DonorName,
		UserType:          f.Funds.UserType,
		AidUID:            f.Funds.AidUID,
		AidName:           f.Funds.AidName,
		AidBankCardNum:    f.Funds.AidBankCardNum,
		TargetUID:         f.Funds.TargetUID,
		TargetName:        f.Funds.TargetName,
		TargetBankCardNum: f.Funds.TargetBankCardNum,
		PubType:           f.Funds.PubType,
		PayType:           f.Funds.PayType,
		Amount:            f.Funds.Amount.String(),
		TxID:              f.Funds.TxID,
		Remark:            f.Funds.Remark,
		BlockType:         f.Funds.BlockType,
		BlockHeight:       f.Funds.BlockHeight,
		BlockTime:         f.Funds.BlockTime,
		CreatedAt:         f.Funds.CreatedAt.Unix(),
	}

	bAddr := structs.PubAddress{
		ID:       f.BillingAddr.ID,
		Type:     f.BillingAddr.Type,
		Country:  f.BillingAddr.Country,
		Province: f.BillingAddr.Province,
		City:     f.BillingAddr.City,
		District: f.BillingAddr.District,
		Address:  f.BillingAddr.Address,
		ZipCode:  f.BillingAddr.ZipCode,
	}

	sAddr := structs.PubAddress{
		ID:       f.ShippingAddr.ID,
		Type:     f.ShippingAddr.Type,
		Country:  f.ShippingAddr.Country,
		Province: f.ShippingAddr.Province,
		City:     f.ShippingAddr.City,
		District: f.ShippingAddr.District,
		Address:  f.ShippingAddr.Address,
		ZipCode:  f.ShippingAddr.ZipCode,
	}

	images := make([]*structs.PubProofImage, 0)
	for _, v := range f.ProofImages {
		images = append(images, &structs.PubProofImage{
			ID:     v.ID,
			Type:   v.Type,
			URL:    v.URL,
			Hash:   v.Hash,
			Format: v.Format,
		})
	}

	result := structs.PubFundsDetail{
		PubFunds:        funds,
		BillingAddress:  bAddr,
		ShippingAddress: sAddr,
		ProofImages:     images,
	}

	c.JSON(http.StatusOK, rest.SuccessResponse(&result))
	logger.Info("response query funds detail success.")
	return
}

// ReceiveSupplies defines the request of received supplies
func (h *RestHandler) ReceiveSupplies(c *gin.Context) {
	logger.Info("got receive supplies request")

	req := &structs.ReceiveSuppliesRequest{}
	if err := c.BindJSON(req); err != nil {
		e := fmt.Errorf("invalid parameters: %s", err.Error())
		logger.Error(e)
		c.JSON(http.StatusBadRequest, rest.ErrorResponse(rest.ParseRequestParamsError, e.Error()))
		return
	}
	logger.Debugf("request params, %v", req)

	supplies := &models.PubSupplies{
		UID:       req.UID,
		UserType:  req.UserType,
		AidUID:    req.AidUID,
		TargetUID: req.TargetUID,
		PubType:   req.PubType,
		Name:      req.Name,
		Number:    req.Number,
		Unit:      req.Unit,
		Remark:    req.Remark,
	}

	err := h.srvcContext.DBStorage.CreateSupplies(supplies)
	if err != nil {
		e := fmt.Errorf("create supplies error, %s", err.Error())
		logger.Error(e)
		c.JSON(http.StatusInternalServerError, rest.ErrorResponse(rest.DatabaseOperationFailed, e.Error()))
		return
	}

	c.JSON(http.StatusOK, rest.SuccessResponse(nil))
	logger.Info("response create supplies success.")
}

// QuerySupplies defines the request of query supplies
func (h *RestHandler) QuerySupplies(c *gin.Context) {
	logger.Info("got query supplies request")

	req := &structs.QuerySuppliesRequest{}
	var err error
	if err = c.Bind(req); err != nil {
		e := fmt.Errorf("invalid parameters: %s", err.Error())
		logger.Error(e)
		c.JSON(http.StatusBadRequest, rest.ErrorResponse(rest.InvalidParamsErrCode, e.Error()))
		return
	}
	logger.Debugf("request params %v", req)

	params := &structs.QueryParams{
		PageNum:   req.PageNum,
		PageLimit: req.PageLimit,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
	}

	result, err := h.srvcContext.DBStorage.QuerySupplies(req.UID, req.TargetUID, req.UserType, req.PubType, params)
	if err != nil {
		e := fmt.Errorf("query funds error, %s", err.Error())
		logger.Error(e)
		c.JSON(http.StatusInternalServerError, rest.ErrorResponse(rest.DatabaseOperationFailed, e.Error()))
		return
	}

	payload := make([]*structs.QuerySuppliesItems, 0)
	for _, v := range result {
		payload = append(payload, &structs.QuerySuppliesItems{
			ID:          v.ID,
			UID:         v.UID,
			DonorName:   v.DonorName,
			UserType:    v.UserType,
			AidUID:      v.AidUID,
			AidName:     v.AidName,
			TargetUID:   v.TargetUID,
			PubType:     v.PubType,
			Name:        v.Name,
			Number:      v.Number,
			Unit:        v.Unit,
			TxID:        v.TxID,
			Remark:      v.Remark,
			BlockType:   v.BlockType,
			BlockHeight: v.BlockHeight,
			BlockTime:   v.BlockTime,
			CreatedAt:   v.CreatedAt.Unix(),
		})
	}

	c.JSON(http.StatusOK, rest.SuccessResponse(&structs.QuerySuppliesResp{
		Total:     params.Total,
		PageNum:   params.PageNum,
		PageLimit: params.PageLimit,
		StartTime: params.StartTime,
		EndTime:   params.EndTime,
		Results:   payload,
	}))
	logger.Info("response query supplies success.")
	return
}

// QuerySuppliesDetail defines the detail information of supplies
func (h *RestHandler) QuerySuppliesDetail(c *gin.Context) {
	logger.Info("got query supplies detail request")

	req := &structs.SuppliesDetailRequest{}
	var err error
	if err = c.Bind(req); err != nil {
		e := fmt.Errorf("invalid parameters: %s", err.Error())
		logger.Error(e)
		c.JSON(http.StatusBadRequest, rest.ErrorResponse(rest.InvalidParamsErrCode, e.Error()))
		return
	}
	logger.Debugf("request params %v", req)

	s, err := h.srvcContext.DBStorage.QuerySuppliesDetail(req.SuppliesID)
	if err != nil {
		e := fmt.Errorf("query supplies detail error, %s", err.Error())
		logger.Error(e)
		c.JSON(http.StatusInternalServerError, rest.ErrorResponse(rest.DatabaseOperationFailed, e.Error()))
		return
	}

	supplies := structs.QuerySuppliesItems{
		ID:          s.Supplies.ID,
		WayBillNum:  s.Supplies.WayBillNum,
		UID:         s.Supplies.UID,
		DonorName:   s.Supplies.DonorName,
		UserType:    s.Supplies.UserType,
		AidUID:      s.Supplies.AidUID,
		AidName:     s.Supplies.AidName,
		TargetUID:   s.Supplies.TargetUID,
		TargetName:  s.Supplies.TargetName,
		PubType:     s.Supplies.PubType,
		Name:        s.Supplies.Name,
		Number:      s.Supplies.Number,
		Unit:        s.Supplies.Unit,
		TxID:        s.Supplies.TxID,
		Remark:      s.Supplies.Remark,
		BlockType:   s.Supplies.BlockType,
		BlockHeight: s.Supplies.BlockHeight,
		BlockTime:   s.Supplies.BlockTime,
		CreatedAt:   s.Supplies.CreatedAt.Unix(),
	}

	bAddr := structs.PubAddress{
		ID:       s.BillingAddr.ID,
		Type:     s.BillingAddr.Type,
		Country:  s.BillingAddr.Country,
		Province: s.BillingAddr.Province,
		City:     s.BillingAddr.City,
		District: s.BillingAddr.District,
		Address:  s.BillingAddr.Address,
		ZipCode:  s.BillingAddr.ZipCode,
	}

	sAddr := structs.PubAddress{
		ID:       s.ShippingAddr.ID,
		Type:     s.ShippingAddr.Type,
		Country:  s.ShippingAddr.Country,
		Province: s.ShippingAddr.Province,
		City:     s.ShippingAddr.City,
		District: s.ShippingAddr.District,
		Address:  s.ShippingAddr.Address,
		ZipCode:  s.ShippingAddr.ZipCode,
	}

	images := make([]*structs.PubProofImage, 0)
	for _, v := range s.ProofImages {
		images = append(images, &structs.PubProofImage{
			ID:     v.ID,
			Type:   v.Type,
			URL:    v.URL,
			Hash:   v.Hash,
			Format: v.Format,
		})
	}

	result := structs.PubSuppliesDetail{
		PubSupplies:     supplies,
		BillingAddress:  bAddr,
		ShippingAddress: sAddr,
		ProofImages:     images,
	}

	c.JSON(http.StatusOK, rest.SuccessResponse(&result))
	logger.Info("response query supplies detail success.")
	return
}

// PubUserList defines the publicity information of user
func (h *RestHandler) PubUserList(c *gin.Context) {
	logger.Info("got publicity person list request")

	req := &structs.PubUserRequest{}
	var err error
	if err = c.Bind(req); err != nil {
		e := fmt.Errorf("invalid parameters: %s", err.Error())
		logger.Error(e)
		c.JSON(http.StatusBadRequest, rest.ErrorResponse(rest.InvalidParamsErrCode, e.Error()))
		return
	}
	logger.Debugf("request params %v", req)

	params := &structs.QueryParams{
		PageNum:   req.PageNum,
		PageLimit: req.PageLimit,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
	}

	result, err := h.srvcContext.DBStorage.QueryPubByUserType(req.UserType, req.TargetUID, req.PubType, params)
	if err != nil {
		e := fmt.Errorf("query funds error, %s", err.Error())
		logger.Error(e)
		c.JSON(http.StatusInternalServerError, rest.ErrorResponse(rest.DatabaseOperationFailed, e.Error()))
		return
	}

	for _, v := range result {
		v.ConvertTime()
	}

	c.JSON(http.StatusOK, rest.SuccessResponse(&structs.PubUserResp{
		Total:     params.Total,
		PageNum:   params.PageNum,
		PageLimit: params.PageLimit,
		StartTime: params.StartTime,
		EndTime:   params.EndTime,
		Results:   result,
	}))
	logger.Info("response query records success.")
	return
}
