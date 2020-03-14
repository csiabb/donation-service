/*
Copyright Lingzhu Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package image

import (
	"bytes"
	"encoding/json"
	"github.com/csiabb/donation-service/common/rest"
	"github.com/csiabb/donation-service/config"
	"github.com/csiabb/donation-service/context"
	"github.com/csiabb/donation-service/models/mock_backend"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func Init(t *testing.T) (*gomock.Controller, *RestHandler, *mock_backend.MockIDBBackend, *httptest.ResponseRecorder, *gin.Context) {
	mockCtl := gomock.NewController(t)
	mockBackend := mock_backend.NewMockIDBBackend(mockCtl)

	// init mock handler
	handler := RestHandler{}
	handler.srvcContext = &context.Context{}
	handler.srvcContext.DBStorage = mockBackend
	handler.srvcContext.Config = &config.SrvcCfg{ALiYun: config.ALiYunCfg{
		Endpoint:        "https://oss-cn-beijing.aliyuncs.com",
		AccessKeyID:     "test",
		AccessKeySecret: "test",
		BucketName:      "donation-oss",
	}}

	// init test mode gin
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	return mockCtl, &handler, mockBackend, w, c
}

// TestRestHandler_Upload test the upload of image
func TestRestHandler_Upload(t *testing.T) {
	mockCtl, handler, _, w, c := Init(t)
	defer mockCtl.Finish()

	imgFile, _ := os.Open("../../build/bin/bg.png")
	defer imgFile.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, _ := writer.CreateFormFile("image_file", filepath.Base("../../build/bin/bg.png"))
	io.Copy(part, imgFile)
	writer.Close()

	c.Request, _ = http.NewRequest(http.MethodPost, "/api/v1/image/upload", body)
	c.Request.Header.Add("Content-Type", writer.FormDataContentType())

	handler.Upload(c)
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
