/*
Copyright ArxanChain Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package database

import (
	"fmt"
	"strings"

	"github.com/csiabb/donation-service/common/log"

	"github.com/jinzhu/gorm"

	// import database driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	logger = log.MustGetLogger("components/database")
)

// IGormClient gorm database interface
type IGormClient interface {
	Init() error
	GetConn() *gorm.DB
	Close()
}

// GormClient gorm database client
type GormClient struct {
	DbConfig *DBConnectCfg
	Db       *gorm.DB
}

// NewGormClient create gorm database client instance
func NewGormClient(cfg *DBConnectCfg) (IGormClient, error) {
	logger.Infof("creating db backend ...")
	d := &GormClient{DbConfig: cfg}
	if err := d.Init(); err != nil {
		return nil, err
	}
	return d, nil
}

// Init init gorm database client
func (d *GormClient) Init() error {
	// get database host and port for address configuration
	addr := strings.Split(d.DbConfig.Address, ":")
	dbHost, dbPort := addr[0], addr[1]

	var connURL string
	if "postgres" == d.DbConfig.Driver {
		connURL = fmt.Sprintf("dbname=%s host=%s port=%s sslmode=%s user=%s password=%s",
			d.DbConfig.DBname,
			dbHost,
			dbPort,
			d.DbConfig.SSLMode,
			d.DbConfig.User,
			d.DbConfig.Password)
	} else if "mysql" == d.DbConfig.Driver {
		connURL = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
			d.DbConfig.User,
			d.DbConfig.Password,
			d.DbConfig.DBname)
	}

	db, err := gorm.Open(d.DbConfig.Driver, connURL)
	if err != nil {
		logger.Errorf("failed to connect to database %s: %s", d.DbConfig.Address, err)
		return err
	}
	// do not use the pluralization table  name
	db.SingularTable(true)
	d.Db = db
	if d.DbConfig.MaxConns > 0 {
		db.DB().SetMaxOpenConns(d.DbConfig.MaxConns)
	}
	if d.DbConfig.MaxIdleConns > 0 {
		db.DB().SetMaxIdleConns(d.DbConfig.MaxIdleConns)
	}
	return nil
}

// GetConn get database connect
func (d *GormClient) GetConn() *gorm.DB {
	return d.Db
}

// Close close gorm database client
func (d *GormClient) Close() {
	logger.Infof("stopping db backend ...")
	if err := d.Db.Close(); err != nil {
		logger.Warningf("failed to stop db backend: %s", err)
	}
}
