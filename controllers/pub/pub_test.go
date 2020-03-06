/*
 * Copyright ArxanChain Ltd. 2020 All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package pub

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/csiabb/donation-service/common/rest"
	"github.com/csiabb/donation-service/context"
	"github.com/csiabb/donation-service/models"
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

func TestReceiveFunds(t *testing.T) {
	mockCtl, handler, mockBackend, w, c := Init(t)
	defer mockCtl.Finish()

	// post body
	body := bytes.NewBufferString("{" +
		"\"uid\":\"uid_test\", " +
		"\"aid_uid\":\"\", " +
		"\"user_type\":\"normal\", " +
		"\"target_uid\":\"target_uid_test\", " +
		"\"pub_type\":\"donate\", " +
		"\"pay_type\":\"wechat\", " +
		"\"amount\":1000.000, " +
		"\"remark\":\"this is a remark message\"" +
		"}")

	mockBackend.EXPECT().CreateFunds(gomock.Any()).Return(nil)

	// mock request
	c.Request, _ = http.NewRequest(http.MethodPost, "/api/v1/pub/funds/receive", body)
	c.Request.Header.Add("Content-Type", "application/json")
	handler.ReceiveFunds(c)
	CommRespCheck(t, w)
}

func TestReceiveSupplies(t *testing.T) {
	mockCtl, handler, mockBackend, w, c := Init(t)
	defer mockCtl.Finish()

	// post body
	body := bytes.NewBufferString("{" +
		"\"uid\":\"uid_test\", " +
		"\"aid_uid\":\"\", " +
		"\"user_type\":\"normal\", " +
		"\"target_uid\":\"target_uid_test\", " +
		"\"pub_type\":\"donate\", " +
		"\"name\":\"3M 一次性口罩\", " +
		"\"number\":1000, " +
		"\"unit\":\"个\", " +
		"\"remark\":\"this is a remark message\"" +
		"}")

	mockBackend.EXPECT().CreateSupplies(gomock.Any()).Return(nil)

	// mock request
	c.Request, _ = http.NewRequest(http.MethodPost, "/api/v1/pub/supplies", body)
	c.Request.Header.Add("Content-Type", "application/json")
	handler.ReceiveSupplies(c)
	CommRespCheck(t, w)
}

func TestQueryFunds(t *testing.T) {
	mockCtl, handler, mockBackend, w, c := Init(t)
	defer mockCtl.Finish()

	mockBackend.EXPECT().QueryFunds(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return([]*models.PubFunds{
		{
			ID:          "id",
			UID:         "uid_test",
			UserType:    "normal",
			AidUID:      "aid_uid",
			TargetUID:   "target_uid",
			PubType:     "donate",
			PayType:     "wechat",
			Amount:      decimal.NewFromInt(20),
			TxID:        "",
			Remark:      "this is a remark",
			BlockType:   "",
			BlockHeight: 0,
			BlockTime:   0,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			DeletedAt:   nil,
		},
	}, nil)

	url := "/api/v1/pub/funds?uid=&user_type=normal&start_time=0&end_time=0&page_num=1&page_limit=10"

	// mock request
	c.Request, _ = http.NewRequest(http.MethodGet, url, nil)
	c.Request.Header.Add("Accept", "application/json")
	handler.QueryFunds(c)
	CommRespCheck(t, w)
}

func TestQuerySupplies(t *testing.T) {
	mockCtl, handler, mockBackend, w, c := Init(t)
	defer mockCtl.Finish()

	mockBackend.EXPECT().QuerySupplies(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return([]*models.PubSupplies{
		{
			ID:          "id",
			UID:         "uid_test",
			UserType:    "normal",
			AidUID:      "aid_uid",
			TargetUID:   "target_uid",
			PubType:     "donate",
			Name:        "3M 一次性口罩",
			Number:      100,
			Unit:        "个",
			TxID:        "",
			Remark:      "this is a remark",
			BlockType:   "",
			BlockHeight: 0,
			BlockTime:   0,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			DeletedAt:   nil,
		},
	}, nil)

	url := "/api/v1/pub/supplies?uid=&target_uid=&user_type=normal&pub_type=&start_time=0&end_time=0&page_num=1&page_limit=10"

	// mock request
	c.Request, _ = http.NewRequest(http.MethodGet, url, nil)
	c.Request.Header.Add("Accept", "application/json")
	handler.QuerySupplies(c)
	CommRespCheck(t, w)
}

func TestQueryFundsDetail(t *testing.T) {
	mockCtl, handler, mockBackend, w, c := Init(t)
	defer mockCtl.Finish()

	// mock db
	mockBackend.EXPECT().QueryFundsDetail(gomock.Any()).Return(&models.FundsDetail{
		Funds: models.PubFunds{
			ID:                "funds_id",
			UID:               "uid_test",
			DonorName:         "donor_name_test",
			UserType:          "normal",
			AidUID:            "aid_uid",
			AidName:           "aid_name_test",
			AidBankCardNum:    "2233-9933-2232-2323",
			TargetUID:         "target_uid",
			TargetName:        "target_name_test",
			TargetBankCardNum: "2233-9933-2232-9233",
			PubType:           "donate",
			PayType:           "wechat",
			Amount:            decimal.NewFromInt(20),
			TxID:              "",
			Remark:            "remark test",
			BlockType:         "",
			BlockHeight:       0,
			BlockTime:         0,
			CreatedAt:         time.Now(),
			UpdatedAt:         time.Time{},
			DeletedAt:         nil,
		},
		BillingAddr: models.Address{
			ID:        "address_billing_id",
			UID:       "uid_test",
			Type:      "billing",
			Country:   "cn",
			Province:  "jiangsu",
			City:      "xuzhou",
			District:  "huabei",
			Address:   "xihuanlu50",
			ZipCode:   "221411",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: nil,
		},
		ShippingAddr: models.Address{
			ID:        "address_shipping_id",
			UID:       "uid_test",
			Type:      "shipping",
			Country:   "cn",
			Province:  "beijing",
			City:      "beijing",
			District:  "huabei",
			Address:   "tiananmen",
			ZipCode:   "100000",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: nil,
		},
		ProofImages: []*models.Image{
			{
				ID:        "image_id",
				RelatedID: "funds_id",
				Type:      "proof",
				URL:       "www.baidu.com",
				Hash:      "aabbcc",
				Format:    "png",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: nil,
			},
		},
	}, nil)

	// mock request
	c.Request, _ = http.NewRequest(http.MethodGet, "/api/v1/pub/funds/detail?funds_id=uid_test", nil)
	c.Request.Header.Add("Accept", "application/json")
	handler.QueryFundsDetail(c)
	CommRespCheck(t, w)
}

func TestQuerySuppliesDetail(t *testing.T) {
	mockCtl, handler, mockBackend, w, c := Init(t)
	defer mockCtl.Finish()

	// mock db
	mockBackend.EXPECT().QuerySuppliesDetail(gomock.Any()).Return(&models.SuppliesDetail{
		Supplies: models.PubSupplies{
			ID:          "funds_id",
			UID:         "uid_test",
			DonorName:   "donor_name_test",
			UserType:    "normal",
			AidUID:      "aid_uid",
			AidName:     "aid_name_test",
			TargetUID:   "target_uid",
			TargetName:  "target_name_test",
			PubType:     "donate",
			Name:        "3M 一次性口罩",
			Number:      200,
			Unit:        "箱",
			TxID:        "",
			Remark:      "remark test",
			BlockType:   "",
			BlockHeight: 0,
			BlockTime:   0,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Time{},
			DeletedAt:   nil,
		},
		BillingAddr: models.Address{
			ID:        "address_billing_id",
			UID:       "uid_test",
			Type:      "billing",
			Country:   "cn",
			Province:  "jiangsu",
			City:      "xuzhou",
			District:  "huabei",
			Address:   "xihuanlu50",
			ZipCode:   "221411",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: nil,
		},
		ShippingAddr: models.Address{
			ID:        "address_shipping_id",
			UID:       "uid_test",
			Type:      "shipping",
			Country:   "cn",
			Province:  "beijing",
			City:      "beijing",
			District:  "huabei",
			Address:   "tiananmen",
			ZipCode:   "100000",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: nil,
		},
		ProofImages: []*models.Image{
			{
				ID:        "image_id",
				RelatedID: "funds_id",
				Type:      "proof",
				URL:       "www.baidu.com",
				Hash:      "aabbcc",
				Format:    "png",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: nil,
			},
		},
	}, nil)

	// mock request
	c.Request, _ = http.NewRequest(http.MethodGet, "/api/v1/pub/supplies/detail?supplies_id=uid_test", nil)
	c.Request.Header.Add("Accept", "application/json")
	handler.QuerySuppliesDetail(c)
	CommRespCheck(t, w)
}

func TestPubUserListList(t *testing.T) {
	mockCtl, handler, mockBackend, w, c := Init(t)
	defer mockCtl.Finish()

	mockBackend.EXPECT().QueryPubByUserType(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return([]*structs.PubUserItem{
		{
			ID:          "id_test",
			Type:        "funds",
			UID:         "donor_uid",
			DonorName:   "donor_name",
			UserType:    "normal",
			AidUID:      "aid_uid_test",
			AidName:     "aid_name",
			TargetUID:   "target_uid_test",
			TargetName:  "target_name",
			PubType:     "donate",
			PayType:     "wechat",
			Amount:      "1000.00",
			Name:        "",
			Number:      0,
			Unit:        "",
			TxID:        "",
			Remark:      "",
			BlockType:   "",
			BlockHeight: 0,
			BlockTime:   0,
			CreatedAt:   0,
			Time:        time.Now(),
		},
		{
			ID:          "id_test_2",
			Type:        "supplies",
			UID:         "donor_uid_2",
			DonorName:   "donor_name_2",
			UserType:    "normal",
			AidUID:      "aid_uid_test_2",
			AidName:     "aid_name_2",
			TargetUID:   "target_uid_test",
			TargetName:  "target_name_2",
			PubType:     "donate",
			PayType:     "",
			Amount:      "",
			Name:        "3M 一次性口罩",
			Number:      3000,
			Unit:        "个",
			TxID:        "",
			Remark:      "",
			BlockType:   "",
			BlockHeight: 0,
			BlockTime:   0,
			CreatedAt:   0,
			Time:        time.Now(),
		},
	}, nil)

	// mock request
	c.Request, _ = http.NewRequest(http.MethodGet, "/api/v1/pub/list?user_type=normal&target_uid=target_uid_test&pub_type=donate&page_num=1&page_limit=10", nil)
	c.Request.Header.Add("Accept", "application/json")
	handler.PubUserList(c)
	CommRespCheck(t, w)
}

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
