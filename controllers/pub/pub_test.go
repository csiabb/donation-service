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

	"github.com/csiabb/donation-service/common/rest"
	"github.com/csiabb/donation-service/context"
	"github.com/csiabb/donation-service/models/storage"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
)

func TestReceiveFunds(t *testing.T) {
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

	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()
	mockBackend := storage.NewMockIDBBackend(mockCtl)
	mockBackend.EXPECT().CreateFunds(gomock.Any()).Return(nil)

	// init handler
	handler := RestHandler{}
	handler.srvcContext = &context.Context{}
	handler.srvcContext.DBStorage = mockBackend

	// init test mode gin
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// mock request
	c.Request, _ = http.NewRequest(http.MethodPost, "/api/v1/pub/funds/receive", body)
	c.Request.Header.Add("Content-Type", "application/json")

	handler.ReceiveFunds(c)
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

		t.Log("receive funds succeed")
	} else {
		t.Error(w.Code, string(b))
	}
}
