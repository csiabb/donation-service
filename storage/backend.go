/*
Copyright ArxanChain Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package storage

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

	// pub
	CreateFunds(*PubFunds) error
	QueryFunds(uid string, params *structs.QueryParams) ([]*PubFunds, error)
}
