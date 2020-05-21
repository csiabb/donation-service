/*
 * Copyright ArxanChain Ltd. 2020 All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package bc

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/csiabb/donation-service/common/rest"
	"github.com/csiabb/donation-service/context"
	"github.com/csiabb/donation-service/models/mock_backend"
	"github.com/csiabb/donation-service/structs"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/jinzhu/gorm"
	"github.com/rafaeljusto/redigomock"
)

const (
	urlBCCallBack = "bc/cb"
)

const (
	bccbBodyJSON = `{
  "blockchain": "cornerstone-chain",
  "id": "did:axn:da-eecf83b7-2bc9-49f2-8005-e0e39f606450",
  "block_num": 3322,
  "tx_id": "kandkalakna9ejdlalajahbabzgzfaftqub",
  "time": 1584932344
}`
)

func Init(t *testing.T) (*gomock.Controller, *RestHandler, *mock_backend.MockIDBBackend, *httptest.ResponseRecorder, *gin.Context) {
	mockCtl := gomock.NewController(t)
	mockBackend := mock_backend.NewMockIDBBackend(mockCtl)

	// init test mode gin
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// init redigo mock connection
	redisCli := redigomock.NewConn()

	// init mock handler
	handler := RestHandler{}
	handler.srvcContext = &context.Context{}
	handler.srvcContext.DBStorage = mockBackend
	handler.srvcContext.RedisCli = redisCli

	return mockCtl, &handler, mockBackend, w, c
}

func TestBlockChainCallBackSucceed(t *testing.T) {
	mockCtl, handler, mockBackend, w, c := Init(t)
	defer mockCtl.Finish()

	db := &gorm.DB{}
	mockBackend.EXPECT().GetDBTransaction().Return(db)
	mockBackend.EXPECT().UpdateFundsBC(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
	mockBackend.EXPECT().UpdateSuppliesBC(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
	mockBackend.EXPECT().DBTransactionCommit(gomock.Any())

	// mock request
	c.Request, _ = http.NewRequest(http.MethodPost, urlBCCallBack, bytes.NewBufferString(bccbBodyJSON))
	c.Request.Header.Add(rest.HeaderContentType, rest.HeaderApplicationJSON)
	handler.BlockChainCallBack(c)
	CommRespCheck(t, w)
}

func TestBlockChainCallBackParams(t *testing.T) {
	mockCtl, handler, _, w, c := Init(t)
	defer mockCtl.Finish()

	// mock request
	c.Request, _ = http.NewRequest(http.MethodPost, urlBCCallBack, bytes.NewBufferString(`{}`))
	c.Request.Header.Add(rest.HeaderContentType, rest.HeaderApplicationJSON)
	handler.BlockChainCallBack(c)
	_, err := ioutil.ReadAll(w.Body)

	if err != nil {
		t.Errorf("io read err, %v", err)
	}

	if w.Code != http.StatusBadRequest {
		t.Error("params check failed")
	}
}

func TestBlockChainCallBackDB(t *testing.T) {
	mockCtl, handler, mockBackend, w, c := Init(t)
	defer mockCtl.Finish()

	db := &gorm.DB{}
	mockBackend.EXPECT().GetDBTransaction().Return(db)
	mockBackend.EXPECT().UpdateFundsBC(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
	mockBackend.EXPECT().UpdateSuppliesBC(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("update supplies error"))
	mockBackend.EXPECT().DBTransactionRollback(gomock.Any())

	// mock request
	c.Request, _ = http.NewRequest(http.MethodPost, urlBCCallBack, bytes.NewBufferString(bccbBodyJSON))
	c.Request.Header.Add(rest.HeaderContentType, rest.HeaderApplicationJSON)
	handler.BlockChainCallBack(c)
	_, err := ioutil.ReadAll(w.Body)

	if err != nil {
		t.Errorf("io read err, %v", err)
	}

	if w.Code != http.StatusInternalServerError {
		t.Error("update supplies check failed")
	}
}

func CommRespCheck(t *testing.T, w *httptest.ResponseRecorder) {
	b, err := ioutil.ReadAll(w.Body)

	if err != nil {
		t.Errorf("io read err, %v", err)
	}

	if w.Code == 200 {
		resp := &structs.BCCBResp{}
		err := json.Unmarshal(b, resp)

		if err != nil {
			t.Errorf("unmarshal error, %v", err)
		}

		if resp.Code != "success" {
			t.Error(resp.Code, resp.Msg)
		}
	} else {
		t.Error(w.Code, string(b))
	}
}
