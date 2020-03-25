/*
 * Copyright ArxanChain Ltd. 2020 All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package pub

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/csiabb/donation-service/common/rest"
	"github.com/csiabb/donation-service/components/bcadapter/mock_bcadapter"
	"github.com/csiabb/donation-service/context"
	"github.com/csiabb/donation-service/models"
	"github.com/csiabb/donation-service/models/mock_backend"
	"github.com/csiabb/donation-service/structs"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/jinzhu/gorm"
	"github.com/rafaeljusto/redigomock"
	"github.com/shopspring/decimal"
)

const (
	urlPubFunds          = "/api/v1/pub/funds"
	urlPubFundsDetail    = "/api/v1/pub/funds/detail"
	urlPubSupplies       = "/api/v1/pub/supplies"
	urlPubSuppliesDetail = "/api/v1/pub/supplies/detail"
	urlPubList           = "/api/v1/pub/list"
)

const (
	fundsBodyJSON = `{
  "uid": "uid_test",
  "donor_uid": "donor_uid_test",
  "donor_name": "donor_name",
  "user_type": "normal",
  "target_uid": "target_uid_test",
  "target_name": "target_name",
  "target_bank_card_num": "1111-2222-3333-4444",
  "pub_type": "donate",
  "pay_type": "wechat",
  "amount": 100,
  "remark": "remark message",
  "proof_images": [
    {
      "type": "proof",
      "url": "www.baidu.com/aaa.png",
      "hash": "laedjakahshsh",
      "format": "png"
    }
  ]
}`

	suppliesBodyJSON = `{
  "uid": "uid_test",
  "donor_uid": "donor_uid_test",
  "donor_name": "donor_name",
  "user_type": "normal",
  "target_uid": "target_uid_test",
  "target_name": "target_name",
  "pub_type": "donate",
  "supplies_item": [
    {
      "name": "3M 一次性口罩",
      "number": 2000,
      "unit": "箱"
    },
    {
      "name": "75%医用酒精",
      "number": 300,
      "unit": "瓶"
    }
  ],
  "remark": "remark message",
  "way_bill_num": "0202-1728-9393",
  "billing_addr": {
    "province": "江苏省",
    "city": "新沂市",
    "address": "轻工路西208号"
  },
  "shipping_addr": {
    "city": "北京市",
    "address": "昌平区西环里小区22号楼"
  },
  "proof_images": [
    {
      "type": "proof",
      "url": "www.google.com/aaa.png",
      "index": "laedjakahshsh",
      "format": "png"
    },
    {
      "type": "proof",
      "url": "www.google.com/bbb.png",
      "index": "shsh",
      "format": "png"
    }
  ]
}`
)

func Init(t *testing.T) (*gomock.Controller, *RestHandler, *mock_backend.MockIDBBackend, *mock_bcadapter.MockIBCAdapter, *httptest.ResponseRecorder, *gin.Context) {
	mockCtl := gomock.NewController(t)
	mockBackend := mock_backend.NewMockIDBBackend(mockCtl)
	mockBCAdapter := mock_bcadapter.NewMockIBCAdapter(mockCtl)

	// init test mode gin
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// init redigo mock connection
	redisCli := redigomock.NewConn()
	redisCli.Command(rest.RedisGet, "block_id_1").Expect("tx_id_test")

	// init mock handler
	handler := RestHandler{}
	handler.srvcContext = &context.Context{}
	handler.srvcContext.DBStorage = mockBackend
	handler.srvcContext.IBCAdapter = mockBCAdapter
	handler.srvcContext.RedisCli = redisCli

	return mockCtl, &handler, mockBackend, mockBCAdapter, w, c
}

func TestReceiveFundsSucceed(t *testing.T) {
	mockCtl, handler, mockBackend, mockBCAdapter, w, c := Init(t)
	defer mockCtl.Finish()

	db := &gorm.DB{}
	mockBackend.EXPECT().GetDBTransaction().Return(db)
	mockBackend.EXPECT().CreateFunds(gomock.Any(), gomock.Any()).Return(nil)
	mockBackend.EXPECT().CreateImages(gomock.Any(), gomock.Any()).Return(nil)
	mockBackend.EXPECT().QueryAccount(gomock.Any(), gomock.Any()).Return(&models.Account{
		ID:             "account_id",
		Access:         "access",
		Password:       "Aa111111",
		NickName:       "nick_name",
		Type:           "normal",
		Phone:          "18518265711",
		Email:          "kellan@qq.com",
		KycStatus:      "verified",
		Bank:           "",
		BankCardNum:    "",
		TaxID:          "",
		ShippingAddrID: "",
		DID:            "did:axn:da-322e9abb-841e-4778-be61-93741d8f4621",
		Remark:         "",
		OpenID:         "",
		UnionID:        "",
		AppID:          "",
		CreatedAt:      time.Now(),
	}, nil)
	mockBCAdapter.EXPECT().Pubs(gomock.Any(), gomock.Any()).Return([]*structs.PubResp{
		{
			Code: 0,
			Msg:  "",
			Data: structs.PubRespData{ID: "block_id_1"},
		},
	}, nil)
	mockBackend.EXPECT().UpdateFunds(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
	mockBackend.EXPECT().DBTransactionCommit(gomock.Any())

	// mock request
	c.Request, _ = http.NewRequest(http.MethodPost, urlPubFunds, bytes.NewBufferString(fundsBodyJSON))
	c.Request.Header.Add(rest.HeaderContentType, rest.HeaderApplicationJSON)
	handler.ReceiveFunds(c)
	CommRespCheck(t, w)
}

func TestReceiveFundsParams(t *testing.T) {
	mockCtl, handler, _, _, w, c := Init(t)
	defer mockCtl.Finish()

	// post body
	body := bytes.NewBufferString(`{}`)

	// mock request
	c.Request, _ = http.NewRequest(http.MethodPost, urlPubFunds, body)
	c.Request.Header.Add(rest.HeaderContentType, rest.HeaderApplicationJSON)
	handler.ReceiveFunds(c)

	_, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Errorf("io read err, %v", err)
	}

	if w.Code != http.StatusBadRequest {
		t.Error("params check failed")
	}
}

func TestReceiveFundsDB(t *testing.T) {
	mockCtl, handler, mockBackend, _, w, c := Init(t)
	defer mockCtl.Finish()

	// post body
	body := bytes.NewBufferString(fundsBodyJSON)

	db := &gorm.DB{}
	mockBackend.EXPECT().GetDBTransaction().Return(db)
	mockBackend.EXPECT().CreateFunds(gomock.Any(), gomock.Any()).Return(nil)
	mockBackend.EXPECT().CreateImages(gomock.Any(), gomock.Any()).Return(errors.New("create funds failed"))
	mockBackend.EXPECT().DBTransactionRollback(db)

	// mock request
	c.Request, _ = http.NewRequest(http.MethodPost, urlPubFunds, body)
	c.Request.Header.Add(rest.HeaderContentType, rest.HeaderApplicationJSON)
	handler.ReceiveFunds(c)
	_, err := ioutil.ReadAll(w.Body)

	if err != nil {
		t.Errorf("io read err, %v", err)
	}

	if w.Code != http.StatusInternalServerError {
		t.Error("create funds check failed")
	}
}

func TestQueryFundsSucceed(t *testing.T) {
	mockCtl, handler, mockBackend, _, w, c := Init(t)
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

	url := urlPubFunds + "?uid=&user_type=normal&start_time=0&end_time=0&page_num=1&page_limit=10"

	// mock request
	c.Request, _ = http.NewRequest(http.MethodGet, url, nil)
	c.Request.Header.Add(rest.HeaderAccept, rest.HeaderApplicationJSON)
	handler.QueryFunds(c)
	CommRespCheck(t, w)
}

func TestQueryFundsDB(t *testing.T) {
	mockCtl, handler, mockBackend, _, w, c := Init(t)
	defer mockCtl.Finish()

	mockBackend.EXPECT().QueryFunds(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, errors.New("query funds failed"))

	url := urlPubFunds + "?uid=&user_type=normal&start_time=0&end_time=0&page_num=1&page_limit=10"

	// mock request
	c.Request, _ = http.NewRequest(http.MethodGet, url, nil)
	c.Request.Header.Add(rest.HeaderAccept, rest.HeaderApplicationJSON)
	handler.QueryFunds(c)
	_, err := ioutil.ReadAll(w.Body)

	if err != nil {
		t.Errorf("io read err, %v", err)
	}

	if w.Code != http.StatusInternalServerError {
		t.Error("query funds check failed")
	}
}

func TestQueryFundsDetailSucceed(t *testing.T) {
	mockCtl, handler, mockBackend, _, w, c := Init(t)
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
	c.Request, _ = http.NewRequest(http.MethodGet, urlPubFundsDetail+"?uid=uid_test", nil)
	c.Request.Header.Add(rest.HeaderAccept, rest.HeaderApplicationJSON)
	handler.QueryFundsDetail(c)
	CommRespCheck(t, w)
}

func TestQueryFundsDetailDB(t *testing.T) {
	mockCtl, handler, mockBackend, _, w, c := Init(t)
	defer mockCtl.Finish()

	// mock db
	mockBackend.EXPECT().QueryFundsDetail(gomock.Any()).Return(nil, errors.New("query funds detail failed"))

	// mock request
	c.Request, _ = http.NewRequest(http.MethodGet, urlPubFundsDetail+"?uid=uid_test", nil)
	c.Request.Header.Add(rest.HeaderAccept, rest.HeaderApplicationJSON)
	handler.QueryFundsDetail(c)
	_, err := ioutil.ReadAll(w.Body)

	if err != nil {
		t.Errorf("io read err, %v", err)
	}

	if w.Code != http.StatusInternalServerError {
		t.Error("query funds detail check failed")
	}
}

func TestReceiveSuppliesSucceed(t *testing.T) {
	mockCtl, handler, mockBackend, mockBCAdapter, w, c := Init(t)
	defer mockCtl.Finish()

	// mock db
	mockBackend.EXPECT().QueryAccount(gomock.Any(), gomock.Any()).Return(&models.Account{
		ID:             "account_id",
		Access:         "access",
		Password:       "Aa111111",
		NickName:       "nick_name",
		Type:           "normal",
		Phone:          "18518265711",
		Email:          "kellan@qq.com",
		KycStatus:      "verified",
		Bank:           "",
		BankCardNum:    "",
		TaxID:          "",
		ShippingAddrID: "",
		DID:            "did:axn:da-eecf83b7-2bc9-49f2-8005-e0e39f606450",
		Remark:         "",
		OpenID:         "",
		UnionID:        "",
		AppID:          "",
		CreatedAt:      time.Now(),
	}, nil)
	db := &gorm.DB{}
	mockBackend.EXPECT().GetDBTransaction().Return(db)
	mockBackend.EXPECT().CreateSupplies(gomock.Any(), gomock.Any()).Return(nil)
	mockBackend.EXPECT().CreateAddresses(gomock.Any(), gomock.Any()).Return(nil)
	mockBackend.EXPECT().CreateImages(gomock.Any(), gomock.Any()).Return(nil)
	mockBCAdapter.EXPECT().Pubs(gomock.Any(), gomock.Any()).Return([]*structs.PubResp{
		{
			Code: 0,
			Msg:  "",
			Data: structs.PubRespData{ID: "block_id_1"},
		},
	}, nil)
	mockBackend.EXPECT().UpdateSuppliesList(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
	mockBackend.EXPECT().DBTransactionCommit(gomock.Any())

	// mock request
	c.Request, _ = http.NewRequest(http.MethodPost, urlPubSupplies, bytes.NewBufferString(suppliesBodyJSON))
	c.Request.Header.Add(rest.HeaderContentType, rest.HeaderApplicationJSON)
	handler.ReceiveSupplies(c)
	CommRespCheck(t, w)
}

func TestReceiveSuppliesParams(t *testing.T) {
	mockCtl, handler, _, _, w, c := Init(t)
	defer mockCtl.Finish()

	// mock request
	c.Request, _ = http.NewRequest(http.MethodPost, urlPubSupplies, bytes.NewBufferString(`{
  "uid": "uid_test",
  "donor_name": "donor_name",
  "user_type": "normal",
  "target_uid": "target_uid_test",
  "target_name": "target_name",
  "pub_type": "donate"}`))

	c.Request.Header.Add(rest.HeaderContentType, rest.HeaderApplicationJSON)
	handler.ReceiveSupplies(c)
	_, err := ioutil.ReadAll(w.Body)

	if err != nil {
		t.Errorf("io read err, %v", err)
	}

	if w.Code != http.StatusBadRequest {
		t.Error("receive supplies params check failed")
	}
}

func TestReceiveSuppliesDB(t *testing.T) {
	mockCtl, handler, mockBackend, _, w, c := Init(t)
	defer mockCtl.Finish()

	// mock db
	mockBackend.EXPECT().QueryAccount(gomock.Any(), gomock.Any()).Return(&models.Account{
		ID:             "account_id",
		Access:         "access",
		Password:       "Aa111111",
		NickName:       "nick_name",
		Type:           "normal",
		Phone:          "18518265711",
		Email:          "kellan@qq.com",
		KycStatus:      "verified",
		Bank:           "",
		BankCardNum:    "",
		TaxID:          "",
		ShippingAddrID: "",
		DID:            "did:axn:da-eecf83b7-2bc9-49f2-8005-e0e39f606450",
		Remark:         "",
		OpenID:         "",
		UnionID:        "",
		AppID:          "",
		CreatedAt:      time.Now(),
	}, nil)
	db := &gorm.DB{}
	mockBackend.EXPECT().GetDBTransaction().Return(db)
	mockBackend.EXPECT().CreateSupplies(gomock.Any(), gomock.Any()).Return(nil)
	mockBackend.EXPECT().CreateAddresses(gomock.Any(), gomock.Any()).Return(nil)
	mockBackend.EXPECT().CreateImages(gomock.Any(), gomock.Any()).Return(errors.New("create images failed"))
	mockBackend.EXPECT().DBTransactionRollback(db)

	// mock request
	c.Request, _ = http.NewRequest(http.MethodPost, urlPubSupplies, bytes.NewBufferString(suppliesBodyJSON))
	c.Request.Header.Add(rest.HeaderContentType, rest.HeaderApplicationJSON)
	handler.ReceiveSupplies(c)
	_, err := ioutil.ReadAll(w.Body)

	if err != nil {
		t.Errorf("io read err, %v", err)
	}

	if w.Code != http.StatusInternalServerError {
		t.Error("create supplies check failed")
	}
}

func TestQuerySuppliesSucceed(t *testing.T) {
	mockCtl, handler, mockBackend, _, w, c := Init(t)
	defer mockCtl.Finish()

	mockBackend.EXPECT().QuerySupplies(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return([]*models.PubSupplies{
		{
			ID:          "supplies_id",
			WayBillNum:  "320045006492",
			UID:         "uid_test",
			DonorName:   "donor_name_test",
			UserType:    "normal",
			AidUID:      "",
			AidName:     "",
			TargetUID:   "target_uid_test",
			TargetName:  "target_name_test",
			PubType:     "donate",
			Name:        "3M 一次性口罩",
			Number:      200,
			Unit:        "个",
			TxID:        "",
			Remark:      "",
			BlockType:   "",
			BlockHeight: 0,
			BlockTime:   0,
			CreatedAt:   time.Now(),
		},
	}, nil)

	url := urlPubSupplies + "uid=&target_uid=&pub_type=distribute&user_type=&start_time=0&end_time=0&page_num=1&page_limit=10"
	c.Request, _ = http.NewRequest(http.MethodGet, url, nil)
	c.Request.Header.Add(rest.HeaderAccept, rest.HeaderApplicationJSON)
	handler.QuerySupplies(c)
	CommRespCheck(t, w)
}

func TestQuerySuppliesDetailSucceed(t *testing.T) {
	mockCtl, handler, mockBackend, _, w, c := Init(t)
	defer mockCtl.Finish()

	mockBackend.EXPECT().QuerySuppliesDetail(gomock.Any()).Return(&models.SuppliesDetail{
		Supplies: models.PubSupplies{
			ID:          "supplies_id_test",
			WayBillNum:  "8292-2323-3232",
			UID:         "uid_test",
			DonorName:   "donor_name_test",
			UserType:    "normal",
			AidUID:      "",
			AidName:     "",
			TargetUID:   "target_uid_test",
			TargetName:  "target_name_test",
			PubType:     "donate",
			Name:        "3M 一次性口罩",
			Number:      2320,
			Unit:        "个",
			TxID:        "",
			Remark:      "",
			BlockType:   "",
			BlockHeight: 0,
			BlockTime:   0,
			CreatedAt:   time.Now(),
		},
		BillingAddr: models.Address{
			ID:        "addr_id",
			UID:       "uid_test",
			RelatedID: "supplies_id_test",
			Type:      "billing",
			Country:   "中国",
			Province:  "江苏",
			City:      "徐州",
			District:  "云龙",
			Address:   "云龙湖小区32号",
			ZipCode:   "221400",
			CreatedAt: time.Now(),
		},
		ShippingAddr: models.Address{
			ID:        "addr_id",
			UID:       "uid_test",
			RelatedID: "supplies_id_test",
			Type:      "shipping",
			Country:   "中国",
			Province:  "",
			City:      "北京",
			District:  "昌平区",
			Address:   "龙泽苑小区32号",
			ZipCode:   "100000",
			CreatedAt: time.Now(),
		},
		ProofImages: []*models.Image{
			{
				ID:        "image_id",
				RelatedID: "supplies_id_test",
				Type:      "proof",
				URL:       "www.baidu.com/aaa.png",
				Hash:      "aadkaaka",
				Index:     "aabbcc",
				Format:    "png",
				CreatedAt: time.Now(),
			},
		},
	}, nil)

	c.Request, _ = http.NewRequest(http.MethodGet, urlPubSuppliesDetail+"?supplies_id=aaa", nil)
	c.Request.Header.Add(rest.HeaderAccept, rest.HeaderApplicationJSON)
	handler.QuerySuppliesDetail(c)
	CommRespCheck(t, w)
}

func TestQuerySuppliesDetailParam(t *testing.T) {
	mockCtl, handler, _, _, w, c := Init(t)
	defer mockCtl.Finish()

	c.Request, _ = http.NewRequest(http.MethodGet, urlPubSuppliesDetail, nil)
	c.Request.Header.Add(rest.HeaderAccept, rest.HeaderApplicationJSON)
	handler.QuerySuppliesDetail(c)
	_, err := ioutil.ReadAll(w.Body)

	if err != nil {
		t.Errorf("io read err, %v", err)
	}

	if w.Code != http.StatusBadRequest {
		t.Error("query supplies detail param check failed")
	}
}

func TestPubUserListSucceed(t *testing.T) {
	mockCtl, handler, mockBackend, _, w, c := Init(t)
	defer mockCtl.Finish()

	mockBackend.EXPECT().QueryPubByUserType(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return([]*structs.PubUserItem{
		{
			ID:          "funds_test",
			Type:        "funds",
			UID:         "uit_test",
			DonorName:   "donor_name_test",
			UserType:    "normal",
			AidUID:      "",
			AidName:     "",
			TargetUID:   "target_uid_test",
			TargetName:  "target_name_test",
			PubType:     "donate",
			PayType:     "wechat",
			Amount:      "2200",
			Name:        "",
			Number:      0,
			Unit:        "",
			TxID:        "",
			Remark:      "remark test",
			BlockType:   "",
			BlockHeight: 0,
			BlockTime:   0,
			CreatedAt:   time.Now().Unix(),
		},
	}, nil)

	url := urlPubList + "?user_type=&target_uid=uid_charity_2&pub_type=distribute&start_time=0&end_time=0&page_num=1&page_limit=50"
	c.Request, _ = http.NewRequest(http.MethodGet, url, nil)
	c.Request.Header.Add(rest.HeaderAccept, rest.HeaderApplicationJSON)
	handler.PubUserList(c)
	CommRespCheck(t, w)
}

func TestPubUserListParams(t *testing.T) {
	mockCtl, handler, _, _, w, c := Init(t)
	defer mockCtl.Finish()

	c.Request, _ = http.NewRequest(http.MethodGet, urlPubList, nil)
	c.Request.Header.Add(rest.HeaderAccept, rest.HeaderApplicationJSON)
	handler.PubUserList(c)
	_, err := ioutil.ReadAll(w.Body)

	if err != nil {
		t.Errorf("io read err, %v", err)
	}

	if w.Code != http.StatusBadRequest {
		t.Error("pub list param check failed")
	}
}

func TestPubUserListDB(t *testing.T) {
	mockCtl, handler, mockBackend, _, w, c := Init(t)
	defer mockCtl.Finish()

	mockBackend.EXPECT().QueryPubByUserType(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, errors.New("records not exist"))

	url := urlPubList + "?user_type=&target_uid=uid_charity_2&pub_type=distribute&start_time=0&end_time=0&page_num=1&page_limit=50"
	c.Request, _ = http.NewRequest(http.MethodGet, url, nil)
	c.Request.Header.Add(rest.HeaderAccept, rest.HeaderApplicationJSON)
	handler.PubUserList(c)
	_, err := ioutil.ReadAll(w.Body)

	if err != nil {
		t.Errorf("io read err, %v", err)
	}

	if w.Code != http.StatusInternalServerError {
		t.Error("pub list db check failed")
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
