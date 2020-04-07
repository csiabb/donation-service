/*
Copyright ArxanChain Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package rest

// comm http header
const (
	HeaderAccept          = "Accept"
	HeaderContentType     = "Content-Type"
	HeaderApplicationJSON = "application/json"
)

// user source
const (
	SourceWechat = "wx"
	SourceApp    = "app"
	SourceWeb    = "web"
)

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
	DonatedTypeFunds    = "funds"    // funds of donation
	DonatedTypeSupplies = "supplies" // supplies of donation
)

// the type of share
const (
	Prove = "prove" // donation prove of share
	Home  = "home"  // home of share
)

// define default info of share
const (
	DrawTitle           = "众行公益链"                     // title of draw
	DrawIcon            = "csiabb.png"                // icon url of draw
	DrawImageURL        = "home_share.png"            // home image url of draw
	DrawHomeContent     = "体验区块链技术！众行公益链邀请您一起，见证爱心行动" // home content of draw
	DrawDonationContent = "你的爱心行动被永久登记到区块链啦～你也来试试吧"   // donation content of draw
)

// define length and width of donation image
const (
	Width     = 750 // width of donation prove image
	SubWidth  = 655 // subWith of donation prove image
	SubHeight = 597 // subHeight of donation prove image
)

//  define information of qr
const (
	QrCodeSize = 120       // qr code size
	QRContent  = "长按识别二维码" // qr code content
)

// define redis cmd
const (
	RedisSet      = "SET"
	RedisGet      = "GET"
	RedisExpireAt = "EXPIREAT"
)
