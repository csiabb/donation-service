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

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/jinzhu/gorm"
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
	body := bytes.NewBufferString("{\n  \"uid\": \"uid_test\",\n  \"donor_name\": \"donor_name\",\n  \"user_type\": \"normal\",\n  \"target_uid\": \"target_uid_test\",\n  \"target_name\": \"target_name\",\n  \"target_bank_card_num\": \"1111-2222-3333-4444\",\n  \"pub_type\": \"donate\",\n  \"pay_type\": \"wechat\",\n  \"amount\": 100,\n  \"remark\": \"remark message\",\n  \"proof_images\": [\n    {\n      \"type\": \"proof\",\n      \"url\": \"www.baidu.com/aaa.png\",\n      \"hash\": \"laedjakahshsh\",\n      \"format\": \"png\"\n    }\n  ]\n}")

	mockBackend.EXPECT().GetDBTransaction().Return(&gorm.DB{})
	mockBackend.EXPECT().CreateFunds(gomock.Any(), gomock.Any()).Return(nil)
	mockBackend.EXPECT().CreateImages(gomock.Any(), gomock.Any()).Return(nil)
	mockBackend.EXPECT().DBTransactionCommit(gomock.Any())

	// mock request
	c.Request, _ = http.NewRequest(http.MethodPost, "/api/v1/pub/funds/receive", body)
	c.Request.Header.Add("Content-Type", "application/json")
	handler.ReceiveFunds(c)
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
			PubType:     "pub_type",
			PayType:     "pay_type",
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

	url := "/api/v1/pub/funds/query?uid=&user_type=normal&start_time=0&end_time=0&page_num=1&page_limit=10"

	// mock request
	c.Request, _ = http.NewRequest(http.MethodGet, url, nil)
	c.Request.Header.Add("Accept", "application/json")
	handler.QueryFunds(c)
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
			PubType:           "pub_type",
			PayType:           "pay_type",
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
				Index:     "adkadkadk",
				Format:    "png",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: nil,
			},
		},
	}, nil)

	// mock request
	c.Request, _ = http.NewRequest(http.MethodGet, "/api/v1/pub/funds/detail?uid=uid_test", nil)
	c.Request.Header.Add("Accept", "application/json")
	handler.QueryFundsDetail(c)
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
