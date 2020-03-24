/*
Copyright ArxanChain Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package structs

import (
	"time"

	"github.com/csiabb/donation-service/common/rest"

	"github.com/shopspring/decimal"
)

// ReceiveFundsRequest defines the request of receiving funds
type ReceiveFundsRequest struct {
	UID               string                  `json:"uid" binding:"required"`                  // user id of the one who donate
	DonorUID          string                  `json:"donor_uid" binding:"required"`            // user did
	DonorName         string                  `json:"donor_name" binding:"required"`           // user name of the one who donate
	UserType          string                  `json:"user_type" binding:"required"`            // user type
	TargetUID         string                  `json:"target_uid" binding:"required"`           // user id of charity
	TargetName        string                  `json:"target_name" binding:"required"`          // user name of the one who receive donation
	TargetBankCardNum string                  `json:"target_bank_card_num" binding:"required"` // target bank card number
	PubType           string                  `json:"pub_type" binding:"required"`             // publicity type
	PayType           string                  `json:"pay_type"`                                // pay type
	Amount            decimal.Decimal         `json:"amount" binding:"required"`               // pay amount
	Remark            string                  `json:"remark"`                                  // remark text
	PubProofImage     []*PubProofImageRequest `json:"proof_images" binding:"required"`         // images of proof
}

// GetUIDByFundsReq implement get funds uid
func (rsr *ReceiveFundsRequest) GetUIDByFundsReq() string {
	switch rsr.PubType {
	case rest.PubTypeDonate:
		return rsr.UID
	case rest.PubTypeDistribute:
		return rsr.TargetUID
	case rest.PubTypeReceive:
		return rsr.DonorUID
	default:
		return ""
	}
}

// QueryFundsRequest defines the request of query funds
type QueryFundsRequest struct {
	UID       string `form:"uid"`        // user id of the one who donate
	TargetUID string `form:"target_uid"` // user id of charity
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
	ID                string `json:"id"`                   // funds id
	UID               string `json:"uid"`                  // user id
	DonorName         string `json:"donor_name"`           // user name of the one who donate
	UserType          string `json:"user_type"`            // user type
	AidUID            string `json:"aid_uid"`              // aid user id
	AidName           string `json:"aid_name"`             // user name of the one who accept donation
	AidBankCardNum    string `json:"aid_bank_card_num"`    // aid bank card number
	TargetUID         string `json:"target_uid"`           // user id of charity
	TargetName        string `json:"target_name"`          // user name of the one who receive donation
	TargetBankCardNum string `json:"target_bank_card_num"` // target bank card number
	PubType           string `json:"pub_type"`             // the type of publicity
	PayType           string `json:"pay_type"`             // pay type
	Amount            string `json:"amount"`               // the amount of publicity funds
	TxID              string `json:"tx_id"`                // block chain tx id
	Remark            string `json:"remark"`               // remark
	BlockType         string `json:"block_type"`           // block type
	BlockHeight       int64  `json:"block_height"`         // block height
	BlockTime         int64  `json:"block_time"`           // block time
	CreatedAt         int64  `json:"created_at"`           // created time
}

// FundsDetailRequest defines the request of query detail funds
type FundsDetailRequest struct {
	FundsID string `form:"funds_id"` // id of funds
}

// ReceiveSuppliesRequest defines the struct of received supplies
type ReceiveSuppliesRequest struct {
	UID             string                  `json:"uid"`                              // user id
	DonorUID        string                  `json:"donor_uid"`                        // user id of the one who donate
	DonorName       string                  `json:"donor_name"`                       // user name of the one who donate
	UserType        string                  `json:"user_type"`                        // user type
	TargetUID       string                  `json:"target_uid" binding:"required"`    // user id of charity
	TargetName      string                  `json:"target_name"`                      // user name of the one who receive donation
	PubType         string                  `json:"pub_type" binding:"required"`      // the type of publicity
	SuppliesItem    []*SuppliesItem         `json:"supplies_item" binding:"required"` // the supplies item
	Remark          string                  `json:"remark"`                           // remark
	WayBillNum      string                  `json:"way_bill_num"`                     // supplies way bill number
	BillingAddress  PubAddress              `json:"billing_addr"`                     // billing address
	ShippingAddress PubAddress              `json:"shipping_addr"`                    // donation shipping address
	PubProofImage   []*PubProofImageRequest `json:"proof_images" binding:"required"`  // images of proof
}

// GetUIDBySuppliesReq defines get the uid of who originated
func (rsr *ReceiveSuppliesRequest) GetUIDBySuppliesReq() string {
	switch rsr.PubType {
	case rest.PubTypeDonate:
		return rsr.UID
	case rest.PubTypeDistribute:
		return rsr.TargetUID
	case rest.PubTypeReceive:
		return rsr.DonorUID
	default:
		return ""
	}
}

// SuppliesItem defines the struct item of received supplies
type SuppliesItem struct {
	Name   string `json:"name" binding:"required"`   // name
	Number int64  `json:"number" binding:"required"` // number
	Unit   string `json:"unit" binding:"required"`   // unit
}

// QuerySuppliesRequest defines the request of supplies
type QuerySuppliesRequest struct {
	UID       string `form:"uid"`        // user id of the one who donate
	TargetUID string `form:"target_uid"` // user id of charity
	UserType  string `form:"user_type"`  // user type
	PubType   string `form:"pub_type"`   // publicity type
	PageNum   int    `form:"page_num"`   // page num
	PageLimit int    `form:"page_limit"` // page limit
	StartTime int64  `form:"start_time"` // start time
	EndTime   int64  `form:"end_time"`   // end time
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
	ID          string `json:"id"`           // supplies id
	WayBillNum  string `json:"way_bill_num"` // supplies way bill number
	UID         string `json:"uid"`          // user id
	DonorName   string `json:"donor_name"`   // user name of the one who donate
	UserType    string `json:"user_type"`    // user type
	AidUID      string `json:"aid_uid"`      // aid user id
	AidName     string `json:"aid_name"`     // user name of the one who accept donation
	TargetUID   string `json:"target_uid"`   // user id of charity
	TargetName  string `json:"target_name"`  // user name of the one who receive donation
	PubType     string `json:"pub_type"`     // the type of publicity
	Name        string `json:"name"`         // name
	Number      int64  `json:"number"`       // number
	Unit        string `json:"unit"`         // unit
	TxID        string `json:"tx_id"`        // block chain tx id
	Remark      string `json:"remark"`       // remark
	BlockType   string `json:"block_type"`   // block type
	BlockHeight int64  `json:"block_height"` // block height
	BlockTime   int64  `json:"block_time"`   // block time
	CreatedAt   int64  `json:"created_at"`   // created time
}

// SuppliesDetailRequest defines the request of query detail supplies
type SuppliesDetailRequest struct {
	SuppliesID string `form:"supplies_id" binding:"required"` // id of supplies
}

// PubUserRequest defines the request of publicity information
type PubUserRequest struct {
	UserType  string `form:"user_type"`                   // user type
	TargetUID string `form:"target_uid"`                  // user id of charity
	PubType   string `form:"pub_type" binding:"required"` // publicity type
	PageNum   int    `form:"page_num"`                    // page num
	PageLimit int    `form:"page_limit"`                  // page limit
	StartTime int64  `form:"start_time"`                  // start time
	EndTime   int64  `form:"end_time"`                    // end time
}

// PubUserResp defines the response of publicity information
type PubUserResp struct {
	PageNum     int            `json:"page_num"`     // page num
	PageLimit   int            `json:"page_limit"`   // page limit
	StartTime   int64          `json:"start_time"`   // start time
	EndTime     int64          `json:"end_time"`     // end time
	Total       int64          `json:"total"`        // total number of query result
	SuppliesNum int64          `json:"supplies_num"` // total number of query supplies
	FundsNum    int64          `json:"funds_num"`    // total number of query funds
	Results     []*PubUserItem `json:"results"`      // funds items
}

// PubUserItem defines the item of publicity information
type PubUserItem struct {
	ID          string    `json:"id"`           // funds id
	Type        string    `json:"type"`         // publicity type
	UID         string    `json:"uid"`          // user id
	DonorName   string    `json:"donor_name"`   // user name of the one who donate
	UserType    string    `json:"user_type"`    // user type
	AidUID      string    `json:"aid_uid"`      // aid user id
	AidName     string    `json:"aid_name"`     // user name of the one who accept donation
	TargetUID   string    `json:"target_uid"`   // user id of charity
	TargetName  string    `json:"target_name"`  // user name of the one who receive donation
	PubType     string    `json:"pub_type"`     // the type of publicity
	PayType     string    `json:"pay_type"`     // pay type of funds
	Amount      string    `json:"amount"`       // the amount of funds
	Name        string    `json:"name"`         // name of supplies
	Number      int64     `json:"number"`       // number of supplies
	Unit        string    `json:"unit"`         // unit of supplies
	TxID        string    `json:"tx_id"`        // block chain tx id
	Remark      string    `json:"remark"`       // remark
	BlockType   string    `json:"block_type"`   // block type
	BlockHeight int64     `json:"block_height"` // block height
	BlockTime   int64     `json:"block_time"`   // block time
	CreatedAt   int64     `json:"created_at"`   // created time
	Time        time.Time `json:"-"`            // time
}

// ConvertTime defines the covert of created_at
func (pui *PubUserItem) ConvertTime() {
	pui.CreatedAt = pui.Time.Unix()
}

// Count defines the number of items donated
func (pui *PubUserItem) Count(fundsNum, suppliesNum *int64) {
	if pui.Type == rest.DonatedTypeFunds {
		*fundsNum++
	} else if pui.Type == rest.DonatedTypeSupplies {
		*suppliesNum++
	}
}

// PubFundsDetail defines the detail information of publicity funds
type PubFundsDetail struct {
	PubFunds        QueryFundsItems      `json:"pub_funds"`     // publicity funds
	BillingAddress  PubAddress           `json:"billing_addr"`  // billing address
	ShippingAddress PubAddress           `json:"shipping_addr"` // shipping address
	ProofImages     []*PubProofImageResp `json:"proof_images"`  // the proof of donation
}

// PubSuppliesDetail defines the detail information of publicity supplies
type PubSuppliesDetail struct {
	PubSupplies     QuerySuppliesItems   `json:"pub_supplies"`  // publicity supplies
	BillingAddress  PubAddress           `json:"billing_addr"`  // billing address
	ShippingAddress PubAddress           `json:"shipping_addr"` // shipping address
	ProofImages     []*PubProofImageResp `json:"proof_images"`  // the proof of donation
}

// PubAddress defines the shipping address of publicity funds detail information
type PubAddress struct {
	ID       string `json:"id"`       // address id
	Type     string `json:"type"`     // address type
	Country  string `json:"country"`  // country
	Province string `json:"province"` // province
	City     string `json:"city"`     // city
	District string `json:"district"` // district
	Address  string `json:"address"`  // detail address
	ZipCode  string `json:"zip_code"` // zip code
}

// PubProofImageRequest defines the request proof image of publicity detail information
type PubProofImageRequest struct {
	Type   string `json:"type" binding:"required"`  // image type
	URL    string `json:"url" binding:"required"`   // image url
	Index  string `json:"index" binding:"required"` // image index
	Format string `json:"format"`                   // image file format
}

// PubProofImageResp defines the proof image of publicity detail information
type PubProofImageResp struct {
	ID     string `json:"id"`     // image id
	Type   string `json:"type"`   // user type
	URL    string `json:"url"`    // image url
	Hash   string `json:"hash"`   // image hash
	Index  string `json:"index"`  // image index
	Format string `json:"format"` // image file format
}
