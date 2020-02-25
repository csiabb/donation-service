/*
Copyright ArxanChain Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package rest

// common error code
const (
	SuccCode                = 0    // 正确
	InvalidParamsErrCode    = 1000 // 参数无效
	MissingParamsErrCode    = 1001 // 缺少参数
	ParseRequestParamsError = 1002 // 解析请求体失败
	DatabaseOperationFailed = 1003 // 数据库操作失败
	SerializeDataFail       = 1004 // 序列化数据失败
	DeserializeDataFail     = 1005 // 反序列化(解析)数据失败
	DatabaseUnavailable     = 1007 // 数据库不可用
	DatabaseDisabled        = 1008 // 数据库已禁用
	PermissionDenied        = 1009 // 没有权限
	InternalServerFailure   = 1012 // 服务内部错误
)
