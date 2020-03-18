/*
 * Copyright ArxanChain Ltd. 2020 All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package models

import (
	"github.com/csiabb/donation-service/structs"

	"github.com/jinzhu/gorm"
)

//go:generate mockgen -destination=mock_backend/mock_backend.go -package=mock_backend github.com/csiabb/donation-service/models IDBBackend

// IDBBackend database operate interface
type IDBBackend interface {
	// get database transaction
	GetDBTransaction() *gorm.DB
	DBTransactionCommit(*gorm.DB)
	DBTransactionRollback(*gorm.DB)

	// account
	QueryAccount(openID, uid string) (*Account, error)
	CreateAccount(*Account) error

	// publicity
	CreateFunds(*gorm.DB, *PubFunds) error
	UpdateFunds(tx *gorm.DB, fundsID, blockID string) error
	UpdateFundsBC(tx *gorm.DB, blockID string, funds *PubFunds) error
	QueryFunds(uid, targetUID, userType, pubType string, params *structs.QueryParams) ([]*PubFunds, error)
	QueryFundsDetail(id string) (*FundsDetail, error)
	CreateSupplies(*gorm.DB, []*PubSupplies) error
	UpdateSuppliesList(*gorm.DB, []*PubSupplies, []*structs.PubResp) error
	UpdateSuppliesBC(tx *gorm.DB, blockID string, supplies *PubSupplies) error
	QuerySupplies(uid, targetUID, userType, pubType string, params *structs.QueryParams) ([]*PubSupplies, error)
	QuerySuppliesDetail(id string) (*SuppliesDetail, error)
	QueryPubByUserType(userType, targetUID, pubType string, params *structs.QueryParams) ([]*structs.PubUserItem, error)
	CreateImages(tx *gorm.DB, data []*Image) error
	CreateAddresses(tx *gorm.DB, data []*Address) error

	// org
	CreateOrganization(*DonationStat) error
	QueryOrgCharities(params *structs.QueryParams) ([]*structs.OrgCharitiesItems, error)
	QueryOrgCharitiesDetail(uid string) (*structs.OrgCharitiesDetailItem, error)
}
