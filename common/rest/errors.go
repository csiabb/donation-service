/*
Copyright ArxanChain Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package rest

// common error code
const (
	SuccCode                  = 0    // succeed
	InvalidParamsErrCode      = 1000 // params invalid
	MissingParamsErrCode      = 1001 // missing params
	ParseRequestParamsError   = 1002 // parse request params error
	DatabaseOperationFailed   = 1003 // database operate failed
	SerializeDataFail         = 1004 // serialize data fail
	DeserializeDataFail       = 1005 // deserialize data fail
	DatabaseUnavailable       = 1007 // database not available
	DatabaseDisabled          = 1008 // database disabled
	PermissionDenied          = 1009 // permission denied
	InternalServerFailure     = 1012 // internal server failure
	RepeatRegistration        = 1013 // repeat registration
	PubToBlockChainFailure    = 1014 // publicity to block chain failure
	BlockChainCallBackTimeout = 1015 // block chain call back timeout
	TokenMissingCode          = 1014 // 缺少token
	TokenExpiredCode          = 1015 // 授权已过期
	TokenNotValidYetCode      = 1016 // 授权未生效
	TokenNotMalformedCode     = 1017 // Auth-Token格式错误
	TokenInvalidCode          = 1018 // Auth-Token无效
)

// wechat error code
const (
	WXLoginFailed     = 2100 // server login failed
	WXUnboundDID      = 2101 // wechat account not bind, not auth
	WXAlreadyboundDID = 2102 // wechat account not bind, auth fine
	WhitelistNotExist = 2103 // user not in white list
)

const (
	TokenMsgMissing     = "token missing"
	TokenMsgExpired     = "token expired"
	TokenMsgNotValidYet = "token not valid yet"
	TokenMsgInvalid     = "token invalid"
	TokenGenerateFailed = "generate token failed"
)
