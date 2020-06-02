/*
 * Copyright ArxanChain Ltd. 2020 All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package bc

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/csiabb/donation-service/common/rest"
	"github.com/csiabb/donation-service/models"
	"github.com/csiabb/donation-service/structs"

	"github.com/gin-gonic/gin"
)

var (
	lock sync.Mutex
)

// BlockChainCallBack defines call back of block chain
func (h *RestHandler) BlockChainCallBack(c *gin.Context) {
	logger.Info("got bc call back request")

	req := &structs.BCCBReq{}
	if err := c.BindJSON(req); err != nil {
		e := fmt.Errorf("invalid parameters, %s", err.Error())
		logger.Error(e)
		c.JSON(http.StatusBadRequest, rest.ErrorResponse(rest.ParseRequestParamsError, e.Error()))
		return
	}
	logger.Debugf("request params, %v", req)

	funds := &models.PubFunds{
		BlockType:   req.BlockChain,
		TxID:        req.TxID,
		BlockHeight: req.BlockNum,
		BlockTime:   req.Time,
	}

	supplies := &models.PubSupplies{
		BlockType:   req.BlockChain,
		TxID:        req.TxID,
		BlockHeight: req.BlockNum,
		BlockTime:   req.Time,
	}

	tx := h.srvcContext.DBStorage.GetDBTransaction()
	err := h.srvcContext.DBStorage.UpdateFundsBC(tx, req.ID, funds)
	if err != nil {
		h.srvcContext.DBStorage.DBTransactionRollback(tx)
		e := fmt.Errorf("update bc info to funds error, %s", err.Error())
		logger.Error(e)
		c.JSON(http.StatusInternalServerError, rest.ErrorResponse(rest.InternalServerFailure, e.Error()))
		return
	}

	err = h.srvcContext.DBStorage.UpdateSuppliesBC(tx, req.ID, supplies)
	if err != nil {
		h.srvcContext.DBStorage.DBTransactionRollback(tx)
		e := fmt.Errorf("update bc info to supplies error, %s", err.Error())
		logger.Error(e)
		c.JSON(http.StatusInternalServerError, rest.ErrorResponse(rest.InternalServerFailure, e.Error()))
		return
	}
	h.srvcContext.DBStorage.DBTransactionCommit(tx)

	key := req.ID
	err = h.srvcContext.Redis.Set(context.Background(), key, req.TxID, time.Duration(60*60)*time.Second).Err()
	if err != nil {
		e := fmt.Errorf("update bc info error when set redis, %s", err.Error())
		logger.Error(e)
		c.JSON(http.StatusInternalServerError, rest.ErrorResponse(rest.InternalServerFailure, e.Error()))
		return
	}

	c.JSON(http.StatusOK, &structs.BCCBResp{Code: "success", Msg: ""})
	logger.Info("response bc call back success.")
	return
}
