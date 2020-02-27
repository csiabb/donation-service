/*
Copyright ArxanChain Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package rest

// publicity type
const (
	PubTypeDonate     = "donate"
	PubTypeReceive    = "receive"
	PubTypeDistribute = "distribute"
)

// pay type
const (
	PayTypeOffline    = "offline"
	PayTypeWeChat     = "wechat"
	PayTypeAliPay     = "alipay"
	PayTypeUnionPay   = "unionpay"
	PayTypeCreditCard = "creditcard"
)

// query page default value
const (
	PageLimit      = 10                // default page limit
	PageNum        = 1                 // default page number
	TenDayBySecond = 10 * 24 * 60 * 60 // seconds of ten days
)
