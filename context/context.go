/*
Copyright ArxanChain Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package context

import (
	"fmt"
	"github.com/csiabb/donation-service/components/image"

	"github.com/csiabb/donation-service/common/log"
	"github.com/csiabb/donation-service/components/aliyun"
	"github.com/csiabb/donation-service/components/wx"
	"github.com/csiabb/donation-service/config"
	"github.com/csiabb/donation-service/models"
	"github.com/csiabb/donation-service/models/impl"
)

var (
	serverContext *Context
	logger        = log.MustGetLogger("context")
)

// Context the context of service
type Context struct {
	WXClient      wx.IWXClient
	Config        *config.SrvcCfg
	DBStorage     models.IDBBackend
	ALiYunBackend aliyun.IALiYunBackend
	ImageBackend  image.IImageBackend
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
		logger.Errorf("Initialize failed, configure is nil")
		return fmt.Errorf("configure is nil")
	}
	fmt.Println("init config:", c.Config)
	logger.Debugf("Initialization configure: %v", c.Config)

	err := c.initStorage()
	if nil != err {
		logger.Errorf("Initialize database storage failed, %v", err)
		return err
	}

	err = c.initALiYunServices()
	if nil != err {
		logger.Errorf("Initialize aliyun services failed, %v", err)
		return err
	}

	err = c.initWXBackend()
	if nil != err {
		logger.Errorf("Initialize wechat backend failed, %v", err)
		return err
	}

	err = c.initImageBackend()
	if nil != err {
		logger.Errorf("Initialize image backend failed, %v", err)
		return err
	}

	logger.Infof("Initialize context success.")

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

func (c *Context) initWXBackend() error {
	var err error
	c.WXClient, err = wx.NewWXBackend(&c.Config.WXCfg)
	if err != nil {
		logger.Errorf("Failed new wx client: %v", err)
		return err
	}

	return nil
}

func (c *Context) initALiYunServices() error {
	var err error
	c.ALiYunBackend, err = aliyun.NewALiYunBackend(&c.Config.ALiYun)
	if nil != err {
		logger.Errorf("New aliyun services error, %v", err)
		return err
	}

	return nil
}
func (c *Context) initImageBackend() error {
	var err error
	c.ImageBackend, err = image.NewImageBackend(&c.Config.Image, c.WXClient)
	if nil != err {
		logger.Errorf("New aliyun services error, %v", err)
		return err
	}

	return nil
}
