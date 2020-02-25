/*
Copyright ArxanChain Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package structs

import (
	"github.com/shopspring/decimal"
)

// ReceiveFundsRequest defines the request of receiving funds
type ReceiveFundsRequest struct {
	UID       string          `json:"uid" binding:"required"`        // user id of the one who donate
	UserType  string          `json:"user_type"`                     // user type
	AidUID    string          `json:"aid_uid"`                       // user id of the one who aided
	TargetUID string          `json:"target_uid" binding:"required"` // user id of charity
	PubType   string          `json:"pub_type" binding:"required"`   // publicity type
	PayType   string          `json:"pay_type"`                      // pay type
	Amount    decimal.Decimal `json:"amount" binding:"required"`     // pay amount
	Remark    string          `json:"remark"`                        // remark text
}

// QueryFundsRequest defines the request of query funds
type QueryFundsRequest struct {
	UID       string `form:"uid"`        // user id of the one who donate
	UserType  string `form:"user_type"`  // user type
	PubType   string `form:"pub_type"`   // publicity type
	PageNum   int    `form:"page_num"`   // page num
	PageLimit int    `form:"page_limit"` // page limit
	StartTime int64  `form:"start_time"` // start time
	EndTime   int64  `form:"end_time"`   // end time
}

// QueryFundsResp defines the response of funds
type QueryFundsResp struct {
	PageNum   int                `json:"page_num"`   // page num
	PageLimit int                `json:"page_limit"` // page limit
	StartTime int64              `json:"start_time"` // start time
	EndTime   int64              `json:"end_time"`   // end time
	Total     int64              `json:"total"`      // total number of query result
	Results   []*QueryFundsItems `json:"results"`    // funds items
}

// QueryFundsItems defines the struct of funds item
type QueryFundsItems struct {
	ID          string `json:"id"`           // funds id
	UID         string `json:"uid"`          // user id
	UserType    string `json:"user_type"`    // user type
	AidUID      string `json:"aid_uid"`      // aid user id
	TargetUID   string `json:"target_uid"`   // user id of charity
	PubType     string `json:"pub_type"`     // the type of publicity
	PayType     string `json:"pay_type"`     // pay type
	Amount      string `json:"amount"`       // the amount of publicity funds
	TxID        string `json:"tx_id"`        // block chain tx id
	Remark      string `json:"remark"`       // remark
	BlockType   string `json:"block_type"`   // block type
	BlockHeight int64  `json:"block_height"` // block height
	BlockTime   int64  `json:"block_time"`   // block time
	CreatedAt   int64  `json:"create_at"`    // create time
}

// ReceiveSuppliesRequest defines the struct of received supplies
type ReceiveSuppliesRequest struct {
	UID       string `json:"uid"`                           // user id
	UserType  string `json:"user_type"`                     // user type
	AidUID    string `json:"aid_uid"`                       // aid user id
	TargetUID string `json:"target_uid" binding:"required"` // user id of charity
	PubType   string `json:"pub_type" binding:"required"`   // the type of publicity
	Name      string `json:"name" binding:"required"`       // name
	Number    int64  `json:"number" binding:"required"`     // number
	Unit      string `json:"unit" binding:"required"`       // unit
	Remark    string `json:"remark"`                        // remark
}

// QuerySuppliesRequest defines the request of supplies
type QuerySuppliesRequest struct {
	UID       string `form:"uid" binding:"required"` // user id of the one who donate
	UserType  string `form:"user_type"`              // user type
	PubType   string `form:"pub_type"`               // publicity type
	PageNum   int    `form:"page_num"`               // page num
	PageLimit int    `form:"page_limit"`             // page limit
	StartTime int64  `form:"start_time"`             // start time
	EndTime   int64  `form:"end_time"`               // end time
}

// QuerySuppliesResp defines the response of supplies
type QuerySuppliesResp struct {
	PageNum   int                   `json:"page_num"`   // page num
	PageLimit int                   `json:"page_limit"` // page limit
	StartTime int64                 `json:"start_time"` // start time
	EndTime   int64                 `json:"end_time"`   // end time
	Total     int64                 `json:"total"`      // total number of query result
	Results   []*QuerySuppliesItems `json:"results"`    // funds items
}

// QuerySuppliesItems defines the struct of supplies item
type QuerySuppliesItems struct {
	ID          string `json:"id"`           // funds id
	UID         string `json:"uid"`          // user id
	UserType    string `json:"user_type"`    // user type
	AidUID      string `json:"aid_uid"`      // aid user id
	TargetUID   string `json:"target_uid"`   // user id of charity
	PubType     string `json:"pub_type"`     // the type of publicity
	Name        string `json:"name"`         // name
	Number      int64  `json:"number"`       // number
	Unit        string `json:"unit"`         // unit
	TxID        string `json:"tx_id"`        // block chain tx id
	Remark      string `json:"remark"`       // remark
	BlockType   string `json:"block_type"`   // block type
	BlockHeight int64  `json:"block_height"` // block height
	BlockTime   int64  `json:"block_time"`   // block time
	CreatedAt   int64  `json:"create_at"`    // create time
}
