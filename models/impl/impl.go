/*
Copyright ArxanChain Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package impl

import (
	"github.com/csiabb/donation-service/common/log"
	"github.com/csiabb/donation-service/components/database"
	"github.com/csiabb/donation-service/models"
	"github.com/jinzhu/gorm"
)

var (
	logger = log.MustGetLogger("models")
)

// DbBackendImpl ...
type DbBackendImpl struct {
	database.GormClient
}

// NewDBBackend ...
func NewDBBackend(cfg *database.DBConnectCfg) (*DbBackendImpl, error) {
	logger.Infof("creating db backend ...")
	d := &DbBackendImpl{GormClient: database.GormClient{DbConfig: cfg}}
	if err := d.Init(); err != nil {
		return nil, err
	}

	// Migrate the schema
	logger.Info("migrate the donation-service schema")
	migrateDb(d)
	return d, nil
}

// migrateDb
func migrateDb(d *DbBackendImpl) {
	// Migrate the schema
	d.Db.AutoMigrate(models.Account{})
	d.Db.AutoMigrate(models.Address{})
	d.Db.AutoMigrate(models.DonationStat{})
	d.Db.AutoMigrate(models.PersonKyc{})
	d.Db.AutoMigrate(models.OrgKyc{})
	d.Db.AutoMigrate(models.Image{})
	d.Db.AutoMigrate(models.PubFunds{})
	d.Db.AutoMigrate(models.PubSupplies{})
	d.Db.AutoMigrate(models.Cover{})
}

// GetDBTransaction ...
func (db *DbBackendImpl) GetDBTransaction() *gorm.DB {
	return db.GetConn().Begin()
}

// DBTransactionCommit ...
func (db *DbBackendImpl) DBTransactionCommit(tx *gorm.DB) {
	if tx != nil {
		tx.Commit()
	}
}

// DBTransactionRollback ...
func (db *DbBackendImpl) DBTransactionRollback(tx *gorm.DB) {
	if tx != nil {
		tx.Rollback()
	}
}
