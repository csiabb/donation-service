package structs

import "github.com/shopspring/decimal"

// ReceiveFundsRequest defines the request of receiving funds
type ReceiveFundsRequest struct {
	UID       string          `json:"uid" binding:"required"`        // user id of the one who donate
	AidUID    string          `json:"aid_uid"`                       // user id of the one who aided
	TargetUID string          `json:"target_uid" binding:"required"` // user id of charity
	PubType   string          `json:"pub_type" binding:"required"`   // publicity type
	PayType   string          `json:"pay_type"`                      // pay type
	Amount    decimal.Decimal `json:"amount" binding:"required"`     // pay amount
	Remark    string          `json:"remark"`                        // remark text
}

// QueryFundsRequest defines the request of query funds
type QueryFundsRequest struct {
	UID       string `form:"uid" binding:"required"` // user id of the one who donate
	PageNum   int    `form:"page_num"`               // page num
	PageLimit int    `form:"page_limit"`             // page limit
	StartTime int64  `form:"start_time"`             // start time
	EndTime   int64  `form:"end_time"`               // end time
}

// QueryFundsResp defines the response of funds
type QueryFundsResp struct {
	PageNum   int                // page num
	PageLimit int                // page limit
	StartTime int64              // start time
	EndTime   int64              // end time
	Total     int64              // total number of query result
	Results   []*QueryFundsItems // funds items
}

// QueryFundsItems defines the struct of funds item
type QueryFundsItems struct {
	ID          string `json:"id"`           // funds id
	UID         string `json:"uid"`          // user id
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
