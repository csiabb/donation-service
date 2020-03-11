/*
 * Copyright ArxanChain Ltd. 2020 All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package structs

// LoginRequest defines the request of user registration
type LoginRequest struct {
	Access      string `json:"access"`                       // user name
	Phone       string `json:"phone"`                        // user phone
	Email       string `json:"email"`                        // user email
	Password    string `json:"password"`                     // user password
	Source      string `json:"source" binding:"required"`    // register source
	Nickname    string `json:"nickname"`                     // nickname
	SmsCode     string `json:"sms_code"`                     // sms code
	UserType    string `json:"user_type" binding:"required"` // user type
	CertCode    string `json:"cert_code"`                    // temporary login credentials of wechat
	ComponentID string `json:"component_id"`                 // component id of wechat
	ID          int    `json:"id"`                           // id of wechat app user
	AppID       string `json:"app_id"`                       // app id
	Remark      string `json:"remark"`                       // user description
}

// LoginResp defines the response of user registration
type LoginResp struct {
	UID string `json:"uid"` // user id
}

//CheckFingerPrintRequest ...
type CheckFingerPrintRequest struct {
	ID            int    `json:"id" binding:"required"`
	JSONString    string `json:"json_string"`
	JSONSignature string `json:"json_signature"`
}

// FingerRequest ...
type FingerRequest struct {
	OpenID        string `json:"openid"`
	JSONString    string `json:"json_string"`
	JSONSignature string `json:"json_signature"`
}

// TokenResponse ...
type TokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

// FingerResponse ...
type FingerResponse struct {
	IsOK    bool   `json:"is_ok"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}
