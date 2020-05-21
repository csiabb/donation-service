/*
Copyright ArxanChain Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package rest

// CommonResponse common rest response
type CommonResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

//SuccessResponse generate success response
func SuccessResponse(data interface{}) *CommonResponse {
	var res = CommonResponse{}
	res.Code = SuccCode
	res.Msg = ""
	res.Data = data
	return &res
}

//ErrorResponse generate error response
//Param code: error code
//Param msg: additional messages for the error, can be nil
func ErrorResponse(errCode int, msg string) *CommonResponse {
	var res = CommonResponse{}
	res.Code = errCode
	res.Msg = msg
	return &res
}
