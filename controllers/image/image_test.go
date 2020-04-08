/*
Copyright Lingzhu Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package image

import (
	"bytes"
	"encoding/json"
	"image"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/csiabb/donation-service/common/rest"
	"github.com/csiabb/donation-service/components/aliyun/mock_backend"
	image_mock "github.com/csiabb/donation-service/components/image/mock_backend"
	wx_mock "github.com/csiabb/donation-service/components/wx/mock_wx"
	"github.com/csiabb/donation-service/config"
	"github.com/csiabb/donation-service/context"
	"github.com/csiabb/donation-service/models"
	storage "github.com/csiabb/donation-service/models/mock_backend"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
)

func Init(t *testing.T) (*gomock.Controller, *RestHandler, *storage.MockIDBBackend, *mock_backend.MockIALiYunBackend, *image_mock.MockIImageBackend,
	*wx_mock.MockIWXClient, *httptest.ResponseRecorder, *gin.Context) {
	mockCtl := gomock.NewController(t)
	mockBackend := storage.NewMockIDBBackend(mockCtl)
	aliyunMockBackend := mock_backend.NewMockIALiYunBackend(mockCtl)
	imageMockBackend := image_mock.NewMockIImageBackend(mockCtl)
	accMockBackend := wx_mock.NewMockIWXClient(mockCtl)
	// init mock handler
	handler := RestHandler{}
	handler.srvcContext = &context.Context{}
	handler.srvcContext.DBStorage = mockBackend
	handler.srvcContext.ALiYunBackend = aliyunMockBackend
	handler.srvcContext.ImageBackend = imageMockBackend
	handler.srvcContext.WXClient = accMockBackend
	handler.srvcContext.DBStorage = mockBackend
	handler.srvcContext.Config = &config.SrvcCfg{
		LocalFileSystem: "",
	}

	// init test mode gin
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	return mockCtl, &handler, mockBackend, aliyunMockBackend, imageMockBackend, accMockBackend, w, c
}

// TestRestHandler_Upload test the upload of image
func TestRestHandler_Upload(t *testing.T) {
	mockCtl, handler, _, aliyunMockBackend, _, _, w, c := Init(t)
	defer mockCtl.Finish()

	imgFile, _ := os.Open("../../build/bin/bg.png")
	defer imgFile.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, _ := writer.CreateFormFile("image_file", filepath.Base("../../build/bin/bg.png"))
	io.Copy(part, imgFile)
	writer.Close()

	aliyunMockBackend.EXPECT().UploadObject(gomock.Any(), gomock.Any()).Return(nil)
	c.Request, _ = http.NewRequest(http.MethodPost, "/api/v1/image/upload", body)
	c.Request.Header.Add(rest.HeaderContentType, writer.FormDataContentType())

	handler.Upload(c)
	CommRespCheck(t, w)
}

// TestRestHandler_Share test the share of image
func TestRestHandler_Share(t *testing.T) {
	mockCtl, handler, mockBackend, aliyunMockBackend, imageMockBackend, _, w, c := Init(t)
	defer mockCtl.Finish()

	dst := image.NewNRGBA(image.Rect(0, 0, 120, 120*150/150))
	url := "/api/v1/image/draw?draw_type=prove&donation_type=supplies&donation_id=supplies_id&scene=1012&is_share=true"

	aliyunMockBackend.EXPECT().UploadObject(gomock.Any(), gomock.Any()).Return(nil)
	aliyunMockBackend.EXPECT().IsExist(gomock.Any()).Return(false, nil)
	imageMockBackend.EXPECT().CreateDonationImage(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Return(dst, nil)
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

	c.Request, _ = http.NewRequest(http.MethodGet, url, nil)
	c.Request.Header.Add("Accept", "application/json")

	handler.Draw(c)
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
