package pub

import (
	"fmt"
	"github.com/csiabb/donation-service/common/rest"
	"github.com/csiabb/donation-service/storage"
	"github.com/csiabb/donation-service/structs"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"net/http"
)

func (h *RestHandler) CreateFunds(c *gin.Context) {
	logger.Info("got create funds request")

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

	funds := &storage.PubFunds{
		UID:       req.UID,
		AidUID:    req.AidUID,
		TargetUID: req.TargetUID,
		PubType:   req.PubType,
		PayType:   req.PayType,
		Amount:    req.Amount,
		Remark:    req.Remark,
	}

	err := h.srvcContext.DBStorage.CreateFunds(funds)
	if err != nil {
		e := fmt.Errorf("create funds error : %s", err.Error())
		logger.Error(e)
		c.JSON(http.StatusInternalServerError, rest.ErrorResponse(rest.DatabaseOperationFailed, e.Error()))
	}

	c.JSON(http.StatusOK, rest.SuccessResponse(nil))
	logger.Infof("response receive funds success.")
	return
}

func (h *RestHandler) DistributeFunds(c *gin.Context) {
}

func (h *RestHandler) QueryFunds(c *gin.Context) {
	logger.Info("got query funds request")

	req := &structs.QueryFundsRequest{}
	var err error
	if err = c.BindQuery(req); err != nil {
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

	result, err := h.srvcContext.DBStorage.QueryFunds(req.UID, params)
	if err != nil {
		e := fmt.Errorf("query funds error : %s", err.Error())
		logger.Error(e)
		c.JSON(http.StatusInternalServerError, rest.ErrorResponse(rest.DatabaseOperationFailed, e.Error()))
		return
	}

	payload := make([]*structs.QueryFundsItems, 0)
	for _, v := range result {
		payload = append(payload, &structs.QueryFundsItems{
			ID:          v.ID,
			UID:         v.UID,
			AidUID:      v.AidUID,
			TargetUID:   v.TargetUID,
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
	logger.Infof("response query funds success.")
	return
}

func (h *RestHandler) ReceiveSupplies(c *gin.Context) {
}

func (h *RestHandler) DistributeSupplies(c *gin.Context) {
}

func (h *RestHandler) QuerySupplies(c *gin.Context) {
	return
}
