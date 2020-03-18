/*
 * Copyright ArxanChain Ltd. 2020 All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package structs

// BCCBReq defines the request of block chain call back information
type BCCBReq struct {
	BlockChain string `json:"blockchain" binding:"required"` // block chain name
	ID         string `json:"id" binding:"required"`         // block chain id
	BlockNum   int64  `json:"block_num" binding:"required"`  // block number in block chain
	TxID       string `json:"tx_id" binding:"required"`      // publicity data tx id of block chain tx
	Time       int64  `json:"time" binding:"required"`       // time
}

// BCCBResp defines the response of block chain call back information
type BCCBResp struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}
