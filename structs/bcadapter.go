/*
 * Copyright ArxanChain Ltd. 2020 All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package structs

// RegisterReq defines the request of user register
type RegisterReq struct {
	AccountID string `json:"account_id"` // user account id
}

// RegisterResp defines the response of user register
type RegisterResp struct {
	Code int              `json:"code"`
	Msg  string           `json:"msg"`
	Data RegisterRespData `json:"data"`
}

// RegisterRespData defines the response data of block chain register interface
type RegisterRespData struct {
	ID string `json:"id"` // block chain id
}

// PubReq defines the request of publicity
type PubReq struct {
	UID       string `json:"uid"` // user id
	Publicity string `json:"publicity"`
}

// PubResp defines the response of publicity
type PubResp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data PubRespData `json:"data"`
}

// PubRespData ...
type PubRespData struct {
	ID string `json:"id"` // block chain id
}

// FundsDonation defines the funds of donation
type FundsDonation struct {
	ID                string           `json:"id"`                   // funds id
	UID               string           `json:"uid"`                  // user id
	DonorName         string           `json:"donor_name"`           // user name of the one who donate
	Time              int64            `json:"time"`                 // donate time
	Amount            string           `json:"amount"`               // the amount of publicity funds
	TargetName        string           `json:"target_name"`          // user name of the one who receive donation
	TargetBankCardNum string           `json:"target_bank_card_num"` // bank card number of charity
	DonationImages    []*DonationImage `json:"donation_images"`      // donation proof images
}

// SuppliesDonation defines the supplies of donation
type SuppliesDonation struct {
	ID              string           `json:"id"`              // supplies id
	UID             string           `json:"uid"`             // user id
	DonorName       string           `json:"donor_name"`      // user name of the one who donate
	BillingAddress  string           `json:"billing_addr"`    // billing address
	Time            int64            `json:"time"`            // donate time
	Name            string           `json:"name"`            // name
	Number          int64            `json:"number"`          // number
	Unit            string           `json:"-"`               // unit
	TargetName      string           `json:"target_name"`     // user name of the one who receive donation
	ShippingAddress string           `json:"shipping_addr"`   // donation shipping address
	WayBillNum      string           `json:"way_bill_num"`    // supplies way bill number
	DonationImages  []*DonationImage `json:"donation_images"` // donation proof images
}

// FundsReceived defines the received funds
type FundsReceived struct {
	ID                string           `json:"id"`                   // funds id
	TargetUID         string           `json:"target_uid"`           // charity user id
	TargetName        string           `json:"target_name"`          // user name of the one who receive donation
	DonorName         string           `json:"donor_name"`           // user name of the one who donate
	Time              int64            `json:"time"`                 // received time
	Amount            string           `json:"amount"`               // the amount of publicity funds
	TargetBankCardNum string           `json:"target_bank_card_num"` // bank card number of charity
	DonationImages    []*DonationImage `json:"donation_images"`      // donation proof images
}

// SuppliesReceived defines the received supplies
type SuppliesReceived struct {
	ID              string           `json:"id"`              // supplies id
	TargetUID       string           `json:"target_uid"`      // charity user id
	DonorName       string           `json:"donor_name"`      // user name of the one who donate
	BillingAddress  string           `json:"billing_addr"`    // billing address
	Time            int64            `json:"time"`            // received time
	Name            string           `json:"name"`            // name
	Number          int64            `json:"number"`          // number
	Unit            string           `json:"-"`               // unit
	TargetName      string           `json:"target_name"`     // user name of the one who receive donation
	ShippingAddress string           `json:"shipping_addr"`   // donation shipping address
	WayBillNum      string           `json:"way_bill_num"`    // supplies way bill number
	DonationImages  []*DonationImage `json:"donation_images"` // donation proof images
}

// FundsDistributed defines the distributed funds
type FundsDistributed struct {
	ID                string           `json:"id"`                   // funds id
	TargetUID         string           `json:"target_uid"`           // charity user id
	TargetName        string           `json:"target_name"`          // user name of the one who receive donation
	TargetBankCardNum string           `json:"target_bank_card_num"` // bank card number of charity
	AidName           string           `json:"aid_name"`             // user name of the one who aided
	Time              int64            `json:"time"`                 // distribute time
	Amount            string           `json:"amount"`               // the amount of publicity funds
	DonationImages    []*DonationImage `json:"donation_images"`      // donation proof images
}

// SuppliesDistributed defines the distributed supplies
type SuppliesDistributed struct {
	ID              string           `json:"id"`              // supplies id
	TargetUID       string           `json:"target_uid"`      // charity user id
	TargetName      string           `json:"target_name"`     // user name of the one who receive donation
	BillingAddress  string           `json:"billing_addr"`    // billing address
	Name            string           `json:"name"`            // name
	Number          int64            `json:"number"`          // number
	Unit            string           `json:"-"`               // unit
	Time            int64            `json:"time"`            // received time
	AidName         string           `json:"aid_name"`        // user name of the one who aided
	ShippingAddress string           `json:"shipping_addr"`   // donation shipping address
	WayBillNum      string           `json:"way_bill_num"`    // supplies way bill number
	DonationImages  []*DonationImage `json:"donation_images"` // donation proof images
}

// DonationImage defines the donation proof image
type DonationImage struct {
	URL  string `json:"url"`  // image url
	Hash string `json:"hash"` // image hash
}
