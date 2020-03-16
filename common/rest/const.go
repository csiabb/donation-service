/*
Copyright ArxanChain Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package rest

import "image/color"

// user type
const (
	UserTypeNormal     = "normal"  // normal user
	UserTypeOrg        = "org"     // organization
	UserTypeAdmin      = "admin"   // admin
	UserTypeOrgCharity = "charity" // charity
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
	AddrReg      = "reg"      // user register address
	AddrShipping = "shipping" // user shipping address
	AddrBilling  = "billing"  // user billing address
)

// the type of image
const (
	ImageAvatar     = "avatar" // avatar image of user
	ImageProof      = "proof"  // proof image of donation
	ImageIDCardHead = "head"   // front image of id card
	ImageIDCardBack = "back"   // back image of id card
)

// the type of items donated
const (
	DonatedTypeFunds    = "funds"
	DonatedTypeSupplies = "supplies"
)

//
const (
	ArxanChainHttp = "https://boxdev.arxanchain.com"
)

// the type of share
const (
	DonationProve = "donation_prove"
	Home          = "home"
)

// define share default info
const (
	ShareTitle    = "众行公益链"
	ShareIcon     = ""
	ShareImageURL = ""
)

// image width and height
const (
	Width     = 750
	SubWidth  = 655
	SubHeight = 597
)

// define qr code size
const (
	QrCodeSize = 120
)

// define type of color
var (
	Color1 = color.RGBA{R: 227, G: 184, B: 121, A: 255}
	Color2 = color.RGBA{R: 201, G: 112, B: 3, A: 255}
	Color3 = color.RGBA{R: 252, G: 235, B: 198, A: 255}
)

// image title
var (
	Title = []string{"捐赠者：", "接收者：", "捐助物资：", "区块链高度：", "存证唯一表示：", "上链时间："}
)
