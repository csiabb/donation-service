/*
Copyright ArxanChain Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package rest

const (
	// publicity type
	PubTypeDonate     = "donate"
	PubTypeReceive    = "receive"
	PubTypeDistribute = "distribute"

	// pay type
	PayTypeOffline    = "offline"
	PayTypeWeChat     = "wechat"
	PayTypeAliPay     = "alipay"
	PayTypeUnionPay   = "unionpay"
	PayTypeCreditCard = "creditcard"
)

const (
	PageLimit      = 10                // default page limit
	PageNum        = 1                 // default page number
	TenDayBySecond = 10 * 24 * 60 * 60 // seconds of ten days
)
