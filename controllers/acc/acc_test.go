/*
 * Copyright ArxanChain Ltd. 2020 All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package acc

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/csiabb/donation-service/common/rest"
	"github.com/csiabb/donation-service/components/bcadapter/mock_bcadapter"
	"github.com/csiabb/donation-service/components/wx"
	"github.com/csiabb/donation-service/components/wx/mock_wx"
	"github.com/csiabb/donation-service/config"
	"github.com/csiabb/donation-service/context"
	"github.com/csiabb/donation-service/models"
	"github.com/csiabb/donation-service/models/mock_backend"
	"github.com/csiabb/donation-service/structs"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/jinzhu/gorm"
)

const (
	urlAccLoginWXApp = "api/v1/acc/login/wxapp"
)

const (
	wxLoginBodyJSON = `{
  "access": "access",
  "phone": "18518265711",
  "email": "aaa@icloud.com",
  "password": "Aa111111",
  "source": "aa",
  "nickname": "kellan",
  "sms_code": "112233",
  "user_type": "normal",
  "cert_code": "aaa",
  "component_id": "aabbcc",
  "id": 20,
  "app_id": "xxyyzz",
  "remark": "desc"
}`
)

func Init(t *testing.T) (*gomock.Controller, *RestHandler, *mock_backend.MockIDBBackend, *mock_wx.MockIWXClient, *mock_bcadapter.MockIBCAdapter, *httptest.ResponseRecorder, *gin.Context) {
	mockCtl := gomock.NewController(t)
	mockDB := mock_backend.NewMockIDBBackend(mockCtl)
	mockWX := mock_wx.NewMockIWXClient(mockCtl)
	mockBCAdapter := mock_bcadapter.NewMockIBCAdapter(mockCtl)

	// init mock handler
	handler := RestHandler{}
	handler.srvcContext = &context.Context{}
	handler.srvcContext.DBStorage = mockDB
	handler.srvcContext.WXClient = mockWX
	handler.srvcContext.Config = &config.SrvcCfg{}
	handler.srvcContext.Config.WXCfg = wx.ClientCfg{}
	handler.srvcContext.IBCAdapter = mockBCAdapter

	// init test mode gin
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	return mockCtl, &handler, mockDB, mockWX, mockBCAdapter, w, c
}

func TestLoginWXAppUserExist(t *testing.T) {
	mockCtl, handler, mockDB, mockWX, _, w, c := Init(t)
	defer mockCtl.Finish()

	lr := wx.LoginResponse{
		OpenID:     "qpsjkayuenzvdsgdflf",
		SessionKey: "adkakdakdkad",
		UnionID:    "dapelemajla",
	}

	acc := &models.Account{
		ID:        "uid",
		Access:    "access",
		Password:  "Aa111111",
		NickName:  "nick name",
		Type:      "normal",
		Phone:     "18518265711",
		Email:     "aaa@icloud.com",
		CreatedAt: time.Now(),
	}

	mockWX.EXPECT().WXLogin(gomock.Any(), gomock.Any(), gomock.Any()).Return(lr, nil)
	mockDB.EXPECT().QueryAccount(gomock.Any(), gomock.Any()).Return(acc, nil)

	// mock request
	c.Request, _ = http.NewRequest(http.MethodPost, urlAccLoginWXApp, bytes.NewBufferString(wxLoginBodyJSON))
	c.Request.Header.Add(rest.HeaderContentType, rest.HeaderApplicationJSON)
	handler.LoginWXApp(c)
	CommRespCheck(t, w)
}

func TestLoginWXAppWXResp(t *testing.T) {
	mockCtl, handler, _, mockWX, _, w, c := Init(t)
	defer mockCtl.Finish()

	lr := wx.LoginResponse{
		OpenID:     "",
		SessionKey: "",
		UnionID:    "",
	}

	mockWX.EXPECT().WXLogin(gomock.Any(), gomock.Any(), gomock.Any()).Return(lr, nil)

	// mock request
	c.Request, _ = http.NewRequest(http.MethodPost, urlAccLoginWXApp, bytes.NewBufferString(wxLoginBodyJSON))
	c.Request.Header.Add(rest.HeaderContentType, rest.HeaderApplicationJSON)
	handler.LoginWXApp(c)
	_, err := ioutil.ReadAll(w.Body)

	if err != nil {
		t.Errorf("io read err, %v", err)
	}

	if w.Code != http.StatusBadRequest {
		t.Error("wx server login check failed")
	}
}

func TestLoginWXAppCreateAccount(t *testing.T) {
	mockCtl, handler, mockDB, mockWX, mockBCAdapter, w, c := Init(t)
	defer mockCtl.Finish()

	lr := wx.LoginResponse{
		OpenID:     "qpsjkayuenzvdsgdflf",
		SessionKey: "adkakdakdkad",
		UnionID:    "dapelemajla",
	}

	mockWX.EXPECT().WXLogin(gomock.Any(), gomock.Any(), gomock.Any()).Return(lr, nil)
	mockDB.EXPECT().QueryAccount(gomock.Any(), gomock.Any()).Return(nil, gorm.ErrRecordNotFound)
	mockBCAdapter.EXPECT().Register(gomock.Any()).Return(&structs.RegisterResp{
		Code: 0,
		Msg:  "",
		Data: structs.RegisterRespData{ID: "aabbcc"},
	}, nil)
	mockDB.EXPECT().CreateAccount(gomock.Any()).Return(nil)

	// mock request
	c.Request, _ = http.NewRequest(http.MethodPost, urlAccLoginWXApp, bytes.NewBufferString(wxLoginBodyJSON))
	c.Request.Header.Add(rest.HeaderContentType, rest.HeaderApplicationJSON)
	handler.LoginWXApp(c)
	_, err := ioutil.ReadAll(w.Body)

	if err != nil {
		t.Errorf("io read err, %v", err)
	}

	if w.Code != http.StatusOK {
		t.Error("create account check failed")
	}
}

func CommRespCheck(t *testing.T, w *httptest.ResponseRecorder) {
	b, err := ioutil.ReadAll(w.Body)

	if err != nil {
		t.Errorf("io read err, %v", err)
	}

	if w.Code == 200 {
		resp := &rest.CommonResponse{}
		err := json.Unmarshal(b, resp)

		if err != nil {
			t.Errorf("unmarshal error, %v", err)
		}

		if resp.Code != 0 {
			t.Error(resp.Code, resp.Msg)
		}
	} else {
		t.Error(w.Code, string(b))
	}
}
