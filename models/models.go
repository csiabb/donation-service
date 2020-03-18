/*
 * Copyright ArxanChain Ltd. 2020 All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package models

import (
	"time"

	"github.com/shopspring/decimal"
)

// Account defines the common information of user
type Account struct {
	ID             string `gorm:"type:varchar(256);primary_key"` // user id
	Access         string `gorm:"type:varchar(256)"`             // user name
	Password       string `gorm:"type:varchar(256)"`             // password
	NickName       string `gorm:"type:varchar(64)"`              // nick name
	Type           string `gorm:"type:varchar(16)"`              // user type
	Phone          string `gorm:"type:varchar(32)"`              // phone num
	Email          string `gorm:"type:varchar(128)"`             // email
	KycStatus      string `gorm:"type:varchar(16)"`              // kyc status
	Bank           string `gorm:"type:varchar(64)"`              // bank name
	BankCardNum    string `gorm:"type:varchar(64)"`              // bank card num
	TaxID          string `gorm:"type:varchar(128)"`             // tax id
	ShippingAddrID string `gorm:"type:varchar(256)"`             // shipping address id
	DID            string `gorm:"type:varchar(128)"`             // did
	Remark         string `gorm:"type:text"`                     // description
	OpenID         string `gorm:"type:varchar(256)"`             // open id of wechat app
	UnionID        string `gorm:"type:varchar(256)"`             // id of wechat app
	AppID          string `gorm:"type:varchar(256)"`             // app id
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time `sql:"index"`
}

// Address defines user's address
type Address struct {
	ID        string `gorm:"type:varchar(256);primary_key"` // address id
	UID       string `gorm:"type:varchar(256);not null"`    // user id
	RelatedID string `gorm:"type:varchar(256)"`             // related id
	Type      string `gorm:"type:varchar(16)"`              // address type
	Country   string `gorm:"type:varchar(32)"`              // country
	Province  string `gorm:"type:varchar(32)"`              // province
	City      string `gorm:"type:varchar(32)"`              // city
	District  string `gorm:"type:varchar(32)"`              // district
	Address   string `gorm:"type:varchar(256)"`             // detail address
	ZipCode   string `gorm:"type:varchar(256)"`             // zip code
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// DonationStat defines the donation statistics of organization
type DonationStat struct {
	ID                  string          `gorm:"type:varchar(256);primary_key"` // donation statistics id
	UID                 string          `gorm:"type:varchar(256);not null"`    // user id
	ReceivedSupplies    int64           // receiving supply
	DistributedSupplies int64           // distribute supply
	ReceivedFunds       decimal.Decimal `gorm:"type:decimal(30,4)"` // receiving funds
	DistributedFunds    decimal.Decimal `gorm:"type:decimal(30,4)"` // distribute funds
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           *time.Time `sql:"index"`
}

// PersonKyc defines the kyc information of single person
type PersonKyc struct {
	ID          string `gorm:"type:varchar(256);primary_key"` // person kyc id
	UID         string `gorm:"type:varchar(256);not null"`    // user id
	RealName    string `gorm:"type:varchar(128)"`             // real name
	Gender      string `gorm:"type:varchar(8)"`               // gender
	CertType    string `gorm:"type:varchar(32)"`              // the type of certification
	CertNum     string `gorm:"type:varchar(128)"`             // the num of certification
	Status      string `gorm:"type:varchar(32)"`              // the status of certification
	Remark      string `gorm:"size:1024"`                     // remark
	CertExpired int64  // the expired of certification
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time `sql:"index"`
}

// OrgKyc defines the kyc information of organization
type OrgKyc struct {
	ID          string `gorm:"type:varchar(256);primary_key"` // organization kyc id
	UID         string `gorm:"type:varchar(256);not null"`    // user id
	LegalPerson string `gorm:"type:varchar(64)"`              // the legal person name of organization
	CreditCode  string `gorm:"type:varchar(128)"`             // the credit code of organization
	Name        string `gorm:"type:varchar(256)"`             // the name of organization
	Region      string `gorm:"type:varchar(32)"`              // the region of organization
	CertType    string `gorm:"type:varchar(32)"`              // the type of certification
	Type        string `gorm:"type:varchar(32)"`              // the type of organization
	Status      string `gorm:"type:varchar(32)"`              // the status of certification
	Expired     int64  // expired time
	Remark      string `gorm:"size:1024"` // remark
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time `sql:"index"`
}

// Image defines user's image
type Image struct {
	ID        string `gorm:"type:varchar(256);primary_key"` // image id
	RelatedID string `gorm:"type:varchar(256);not null"`    // the related id
	Type      string `gorm:"type:varchar(64)"`              // image type
	URL       string `gorm:"type:varchar(512)"`             // image url
	Hash      string `gorm:"type:varchar(256)"`             // image hash
	Index     string `gorm:"type:varchar(256)"`             // image index
	Format    string `gorm:"type:varchar(64)"`              // image file format
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// PubFunds defines the publicity funds
type PubFunds struct {
	ID                string          `gorm:"type:varchar(256);primary_key"` // funds id
	UID               string          `gorm:"type:varchar(256)"`             // user id
	DonorName         string          `gorm:"type:varchar(256)"`             // user name of the one who donate
	UserType          string          `gorm:"type:varchar(16)"`              // user type
	AidUID            string          `gorm:"type:varchar(256)"`             // aid user id
	AidName           string          `gorm:"type:varchar(256)"`             // user name of the one who accept donation
	AidBankCardNum    string          `gorm:"type:varchar(64)"`              // bank card number of aid user
	TargetUID         string          `gorm:"type:varchar(256)"`             // user id of charity
	TargetName        string          `gorm:"type:varchar(256)"`             // user name of the one who receive donation
	TargetBankCardNum string          `gorm:"type:varchar(64)"`              // bank card number of charity
	PubType           string          `gorm:"type:varchar(16)"`              // the type of publicity
	PayType           string          `gorm:"type:varchar(16)"`              // pay type
	Amount            decimal.Decimal `gorm:"type:decimal(30,4)"`            // the amount of publicity funds
	Remark            string          `gorm:"size:1024"`                     // remark
	BlockType         string          `gorm:"type:varchar(32)"`              // block type
	BlockID           string          `gorm:"type:varchar(256)"`             // block chain id
	TxID              string          `gorm:"type:varchar(256)"`             // block chain tx id
	BlockHeight       int64           // block height
	BlockTime         int64           // block time
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         *time.Time `sql:"index"`
}

// PubSupplies defines the publicity supplies
type PubSupplies struct {
	ID          string `gorm:"type:varchar(256);primary_key"` // supply id
	WayBillNum  string `gorm:"type:varchar(256)"`             // way bill num
	UID         string `gorm:"type:varchar(256)"`             // user id
	DonorName   string `gorm:"type:varchar(256)"`             // user name of the one who donate
	UserType    string `gorm:"type:varchar(16)"`              // user type
	AidUID      string `gorm:"type:varchar(256)"`             // aid user id
	AidName     string `gorm:"type:varchar(256)"`             // user name of the one who accept donation
	TargetUID   string `gorm:"type:varchar(256)"`             // user id of charity
	TargetName  string `gorm:"type:varchar(256)"`             // user name of the one who receive donation
	PubType     string `gorm:"type:varchar(16)"`              // the type of publicity
	Name        string `gorm:"type:varchar(512)"`             // name
	Number      int64  // number
	Unit        string `gorm:"type:varchar(32)"`  // unit
	BlockID     string `gorm:"type:varchar(256)"` // block chain id
	TxID        string `gorm:"type:varchar(256)"` // block chain tx id
	Remark      string `gorm:"size:1024"`         // remark
	BlockType   string `gorm:"type:varchar(32)"`  // block type
	BlockHeight int64  // block height
	BlockTime   int64  // block time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time `sql:"index"`
}

// Cover defines the introduction information
type Cover struct {
	ID          string `gorm:"type:varchar(256);primary_key"` // cover id
	Information string // cover content
	ImageURL    string `gorm:"type:varchar(256)"` // image url
	SkipURL     string `gorm:"type:varchar(256)"` // skip url
	Weight      int    // show weight
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time `sql:"index"`
}

// FundsDetail defines the detail of funds
type FundsDetail struct {
	Funds        PubFunds
	BillingAddr  Address
	ShippingAddr Address
	ProofImages  []*Image
}

// SuppliesDetail defines the detail of supplies
type SuppliesDetail struct {
	Supplies     PubSupplies
	BillingAddr  Address
	ShippingAddr Address
	ProofImages  []*Image
}
