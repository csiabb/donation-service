/*
 * Copyright ArxanChain Ltd. 2020 All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package wx

import (
	"net/http"

	"github.com/csiabb/donation-service/structs"
)

// IWXClient defines the wx client interface
type IWXClient interface {
	WXLogin(appID string, secret string, code string) (lres LoginResponse, err error)
	DecryptUserInfo(rawData, encryptedData, signature, iv, ssk string) (ui UserInfo, err error)
	DecryptPhoneNumber(ssk, data, iv string) (phone PhoneNumber, err error)
	CheckFinger(finger structs.FingerRequest, accessToken string) (*structs.FingerResponse, error)
}

// ClientCfg ...
type ClientCfg struct {
	Enabled bool
	AppID   string
	Secret  string
	Name    string
	Env     string
}

// Client wx Client
type Client struct {
	c          *ClientCfg
	HTTPClient *http.Client
}

// Response defines the struct of wx base data
type Response struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// UserInfo defines the user information
type UserInfo struct {
	OpenID    string    `json:"openId"`
	Nickname  string    `json:"nickName"`
	Gender    int       `json:"gender"`
	Province  string    `json:"province"`
	Language  string    `json:"language"`
	Country   string    `json:"country"`
	City      string    `json:"city"`
	Avatar    string    `json:"avatarUrl"`
	UnionID   string    `json:"unionId"`
	Watermark watermark `json:"watermark"`
}

type watermark struct {
	AppID     string `json:"appid"`
	Timestamp int64  `json:"timestamp"`
}

type loginResponse struct {
	Response
	LoginResponse
}

// LoginResponse defines the login response
type LoginResponse struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
}

// PhoneNumber defines the phone number
type PhoneNumber struct {
	PhoneNumber     string    `json:"phoneNumber"`
	PurePhoneNumber string    `json:"purePhoneNumber"`
	CountryCode     string    `json:"countryCode"`
	Watermark       watermark `json:"watermark"`
}

// GetWXACodeRequest ...
type GetWXACodeRequest struct {
	Scene     string      `json:"scene"`
	Width     int64       `json:"width"`
	AutoColor bool        `json:"auto_color"`
	LineColor interface{} `json:"line_color"`
	IsHyaLine bool        `json:"is_hyaline"`
}
