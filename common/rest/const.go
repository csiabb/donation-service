/*
Copyright ArxanChain Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package rest

// user type
const (
	UserTypeNormal = "normal" // normal user
	UserTypeOrg    = "org"    // organization
	UserTypeAdmin  = "admin"  // admin
)

// publicity type
const (
	PubTypeDonate     = "donate"     // donation by aid user
	PubTypeReceive    = "receive"    // receive by aided user
	PubTypeDistribute = "distribute" // distribute by third part user
)

// pay type
const (
	PayTypeOffline    = "offline"    // offline payment
	PayTypeWeChat     = "wechat"     // wechat payment
	PayTypeAliPay     = "alipay"     // alipay
	PayTypeUnionPay   = "unionpay"   // unionpay
	PayTypeCreditCard = "creditcard" // creditcard
)

// query page default value
const (
	PageLimit      = 10                // default page limit
	PageNum        = 1                 // default page number
	TenDayBySecond = 10 * 24 * 60 * 60 // seconds of ten days
)

// the type of addresses
const (
	AddrTypeReg      = "reg"      // user register address
	AddrTypeShipping = "shipping" // user shipping address
)

// the type of image
const (
	ImageAvatar     = "avatar" // avatar image of user
	ImageProof      = "proof"  // proof image of donation
	ImageIDCardHead = "head"   // front image of id card
	ImageIDCardBack = "back"   // back image of id card
)
