/*
Copyright ArxanChain Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package impl

import (
	"github.com/csiabb/donation-service/common/log"
	"github.com/csiabb/donation-service/components/database"
	"github.com/csiabb/donation-service/models"
)

var (
	logger = log.MustGetLogger("storage")
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
}
