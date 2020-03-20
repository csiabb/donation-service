/*
 * Copyright ArxanChain Ltd. 2020 All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package wx

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/csiabb/donation-service/components/wx/utils"
	"github.com/csiabb/donation-service/structs"

	"github.com/hunterhug/go_image/graphics"
	"github.com/op/go-logging"
)

var logger = logging.MustGetLogger("wx")

var (
	newDx     = 120
	width     = int64(280)
	lineColor = map[string]string{"r": "201", "g": "112", "b": "3"}
)

const (
	codeAPI        = "/sns/jscode2session"
	grantType      = "authorization_code"
	wxAddress      = "https://api.weixin.qq.com"
	checkFingerURL = "%s/cgi-bin/soter/verify_signature?access_token=%s"
	accessTokenURL = "%s/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
	wxacodeURL     = "%s/wxa/getwxacodeunlimit?access_token=%s"
)

// NewWXBackend returns a handle to the agent endpoints
func NewWXBackend(config *ClientCfg) (IWXClient, error) {
	return &Client{c: config, HTTPClient: &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}}, nil
}

// WXLogin get the login data
func (c *Client) WXLogin(appID string, secret string, code string) (lres LoginResponse, err error) {
	if code == "" {
		err = errors.New("code can not be null")
		return
	}

	api, err := code2url(appID, secret, code, wxAddress)
	if err != nil {
		return
	}

	res, err := http.Get(api)
	if err != nil {
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		err = errors.New("WeChat service failed to login")
		return
	}

	var data loginResponse
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return
	}

	if data.ErrCode != 0 {
		err = errors.New(data.ErrMsg)
		return
	}

	lres = data.LoginResponse
	return
}

// DecryptUserInfo decrypt the user information
func (c *Client) DecryptUserInfo(rawData, encryptedData, signature, iv, ssk string) (ui UserInfo, err error) {
	if signature != "" {
		if ok := utils.Validate(rawData, ssk, signature); !ok {
			err = errors.New("validate data failed")
			return
		}
	}

	bts, err := utils.Decrypt(ssk, encryptedData, iv)
	if err != nil {
		return
	}

	err = json.Unmarshal(bts, &ui)
	return
}

// DecryptPhoneNumber decrypt phone number
func (c *Client) DecryptPhoneNumber(ssk, data, iv string) (phone PhoneNumber, err error) {
	bts, err := utils.CBCDecrypt(ssk, data, iv)
	if err != nil {
		return
	}

	err = json.Unmarshal(bts, &phone)
	return
}

func code2url(appID, secret, code, baseURL string) (string, error) {

	url, err := url.Parse(baseURL + codeAPI)
	if err != nil {
		return "", err
	}

	query := url.Query()

	query.Set("appid", appID)
	query.Set("secret", secret)
	query.Set("js_code", code)
	query.Set("grant_type", grantType)

	url.RawQuery = query.Encode()

	return url.String(), nil
}

// CheckFinger ...
func (c *Client) CheckFinger(finger structs.FingerRequest, accessToken string) (*structs.FingerResponse, error) {
	wxURL := fmt.Sprintf(checkFingerURL, wxAddress, accessToken)
	pushByte, err := json.Marshal(finger)
	body := bytes.NewBuffer(pushByte)
	response, err := c.HTTPClient.Post(wxURL, "application/json;charset=utf-8", body)
	if err != nil {
		logger.Errorf("Failed to push message: %v", err)
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		e := fmt.Errorf("Get response status is not ok")
		logger.Error(e)
		return nil, err
	}
	defer response.Body.Close()
	res, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logger.Errorf("Failed to read response body: %v", err)
		return nil, err
	}
	s := structs.FingerResponse{}
	err = json.Unmarshal(res, &s)
	if err != nil {
		logger.Errorf("Failed to UnMarshal response body: %v", err)
		return nil, err
	}
	return &s, nil
}

// GetAccessToken ...
func (c *Client) GetAccessToken(appID string, secret string) (string, error) {
	wxURL := fmt.Sprintf(accessTokenURL, wxAddress, appID, secret)
	logger.Debugf("Get access token url: %v", wxURL)
	response, err := http.Get(wxURL)
	if err != nil {
		logger.Errorf("Failed to do get access token: %v", err)
		return "", err
	}
	if response.StatusCode != http.StatusOK {
		logger.Errorf("Get response status is not ok")
		return "", err
	}
	defer response.Body.Close()
	res, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logger.Errorf("Failed to read response body: %v", err)
		return "", err
	}
	body := structs.TokenResponse{}
	err = json.Unmarshal(res, &body)
	if err != nil {
		logger.Errorf("Failed to UnMarshal response body: %v", err)
		return "", err
	}
	return body.AccessToken, nil
}

// GetWXQrCode get WeChat qr code
func (c *Client) GetWXQrCode(token string, scene string) (image.Image, error) {
	wxURL := fmt.Sprintf(wxacodeURL, wxAddress, token)
	logger.Debugf("Get wx code url: %v", wxURL)
	var req = &GetWXQRCodeRequest{
		Scene:     scene,
		Width:     width,
		LineColor: lineColor,
		IsHyaLine: true,
	}

	pushByte, err := json.Marshal(req)
	body := bytes.NewBuffer(pushByte)

	response, err := http.Post(wxURL, "application/json", body)
	if err != nil {
		logger.Errorf("Failed to do get wx qr code : %v", err)
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		logger.Errorf("Get response status is not ok")
		return nil, err
	}
	defer response.Body.Close()

	qrCode, err := png.Decode(response.Body)
	if err != nil {
		return nil, err
	}

	// adjust the size
	dst := image.NewRGBA(image.Rect(0, 0, newDx, newDx*qrCode.Bounds().Dy()/qrCode.Bounds().Dx()))
	if err = graphics.Scale(dst, qrCode); err != nil {
		return nil, err
	}

	return dst, nil
}
