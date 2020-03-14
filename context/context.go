/*
Copyright ArxanChain Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package context

import (
	"fmt"
	"github.com/csiabb/donation-service/common/log"
	"github.com/csiabb/donation-service/config"
	"github.com/csiabb/donation-service/models"
	"github.com/csiabb/donation-service/models/aliyun"
	"github.com/csiabb/donation-service/models/impl"
)

var (
	serverContext *Context
	logger        = log.MustGetLogger("context")
)

// Context the context of service
type Context struct {
	Config         *config.SrvcCfg
	DBStorage      models.IDBBackend
	ALiYunServices models.IALiYunBackend
}

// GetServerContext ...
func GetServerContext() *Context {
	if serverContext == nil {
		serverContext = &Context{}
	}
	return serverContext
}

// Init init service context
func (c *Context) Init() error {
	if nil == c.Config {
		logger.Errorf("Initalize faild, configure is nil")
		return fmt.Errorf("configure is nil")
	}
	fmt.Println("init config:", c.Config)
	logger.Debugf("Initalization configure: %v", c.Config)

	err := c.initStorage()
	if nil != err {
		logger.Errorf("Initalize database storage faild, %v", err)
		return err
	}

	err = c.initALiYunServices()
	if nil != err {
		logger.Errorf("Initalize aliyun services faild, %v", err)
		return err
	}

	logger.Infof("initalize context success.")

	return nil
}

func (c *Context) initStorage() error {
	if !c.Config.Database.Enabled {
		logger.Infof("database is disabled")
		return nil
	}

	var err error
	c.DBStorage, err = impl.NewDBBackend(&c.Config.Database)
	if nil != err {
		logger.Errorf("New database backend error, %v", err)
		return err
	}

	return nil
}

func (c *Context) initALiYunServices() error {
	var err error
	c.ALiYunServices, err = aliyun.NewALiYunBackend(&c.Config.ALiYun)
	if nil != err {
		logger.Errorf("New aliyun services error, %v", err)
		return err
	}

	return nil
}
