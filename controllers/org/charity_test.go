/*
Copyright Lingzhu Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package org

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/csiabb/donation-service/common/rest"
	"github.com/csiabb/donation-service/context"
	"github.com/csiabb/donation-service/models/mock_backend"
	"github.com/csiabb/donation-service/structs"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/shopspring/decimal"
)

func Init(t *testing.T) (*gomock.Controller, *RestHandler, *mock_backend.MockIDBBackend, *httptest.ResponseRecorder, *gin.Context) {
	mockCtl := gomock.NewController(t)
	mockBackend := mock_backend.NewMockIDBBackend(mockCtl)

	// init mock handler
	handler := RestHandler{}
	handler.srvcContext = &context.Context{}
	handler.srvcContext.DBStorage = mockBackend

	// init test mode gin
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	return mockCtl, &handler, mockBackend, w, c
}

// TestRestHandler_QueryOrgCharities test the list of charities
func TestRestHandler_QueryOrgCharities(t *testing.T) {
	mockCtl, handler, mockBackend, w, c := Init(t)
	defer mockCtl.Finish()

	mockBackend.EXPECT().QueryOrgCharities(gomock.Any()).Return([]*structs.OrgCharitiesItems{
		{
			ID:                  "id",
			UID:                 "uid_test",
			URL:                 "url",
			NickName:            "nick_name",
			ReceivedSupplies:    0,
			ReceivedFunds:       decimal.NewFromInt(20),
			DistributedFunds:    decimal.NewFromInt(20),
			DistributedSupplies: 0,
			CreatedAt:           time.Now().UTC().Unix(),
		},
	}, nil)

	url := "/api/v1/org/charities?start_time=0&end_time=0&page_num=1&page_limit=10"

	// mock request
	c.Request, _ = http.NewRequest(http.MethodGet, url, nil)
	c.Request.Header.Add(rest.HeaderAccept, rest.HeaderApplicationJSON)
	handler.QueryOrgCharities(c)
	CommRespCheck(t, w)
}

// TestRestHandler_QueryOrgCharitiesDetail test charity details
func TestRestHandler_QueryOrgCharitiesDetail(t *testing.T) {
	mockCtl, handler, mockBackend, w, c := Init(t)
	defer mockCtl.Finish()

	// mock db
	mockBackend.EXPECT().QueryOrgCharitiesDetail(gomock.Any()).Return(&structs.OrgCharitiesDetailItem{
		UID:         "uid",
		URL:         "url",
		NickName:    "nick_name",
		Country:     "country",
		Province:    "province",
		City:        "city",
		District:    "district",
		Address:     "address",
		Phone:       "phone",
		BankCardNum: "bank_card_num",
		Remark:      "remark",
	}, nil)

	// mock request
	c.Request, _ = http.NewRequest(http.MethodGet, "/api/v1/pub/funds/detail?uid=uid_test", nil)
	c.Request.Header.Add(rest.HeaderAccept, rest.HeaderApplicationJSON)
	handler.QueryOrgCharitiesDetail(c)
	CommRespCheck(t, w)
}

// CommRespCheck http response reply data check
func CommRespCheck(t *testing.T, w *httptest.ResponseRecorder) {
	b, err := ioutil.ReadAll(w.Body)

	if err != nil {
		t.Errorf("io read err, %v", err)
		return
	}

	if w.Code == 200 {
		resp := &rest.CommonResponse{}
		err := json.Unmarshal(b, resp)

		if err != nil {
			t.Errorf("unmarshal error, %v", err)
			return
		}

		if resp.Code != 0 {
			t.Error(resp.Code, resp.Msg)
			return
		}
	} else {
		t.Error(w.Code, string(b))
	}
}
