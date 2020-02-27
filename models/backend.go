/*
Copyright ArxanChain Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package models

import (
	"github.com/csiabb/donation-service/structs"

	"github.com/jinzhu/gorm"
)

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
	CreateSupplies(supplies *PubSupplies) error
	QuerySupplies(uid, userType, pubType string, params *structs.QueryParams) ([]*PubSupplies, error)
	QueryPubByUserType(userType string, params *structs.QueryParams) ([]*structs.PubUserItem, error)

	// org
	CreateOrganizations(*DonationStat) error
	QueryOrganizations(params *structs.QueryParams) ([]*structs.OrganizationsItems, error)
	QueryOrganizationInformation(uid string) (*structs.OrganizationInformationItem, error)
}
