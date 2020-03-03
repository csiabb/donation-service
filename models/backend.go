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

	CreateAccount(*Account) error

	// publicity
	CreateFunds(*PubFunds) error
	QueryFunds(uid, userType, pubType string, params *structs.QueryParams) ([]*PubFunds, error)
	QueryFundsDetail(id string) (*FundsDetail, error)
	CreateSupplies(supplies *PubSupplies) error
	QuerySupplies(uid, userType, pubType string, params *structs.QueryParams) ([]*PubSupplies, error)
	QuerySuppliesDetail(id string) (*SuppliesDetail, error)
	QueryPubByUserType(userType string, params *structs.QueryParams) ([]*structs.PubUserItem, error)

	// org
	CreateOrganizations(*DonationStat) error
	QueryOrganizations(params *structs.QueryParams) ([]*structs.OrganizationsItems, error)
	QueryOrganizationDetail(uid string) (*structs.OrganizationDetailItem, error)
}
