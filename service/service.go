/*
Copyright ArxanChain Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package service

import (
	"github.com/csiabb/donation-service/common/metadata"
	"github.com/csiabb/donation-service/config"
)

// Server interface ...
type Server interface {
	Start() (err error)
	Shutdown()
}

// NewUninitializedServer creates a Server instance without initializing it
func NewUninitializedServer(c *config.SrvcCfg, sVer *metadata.Version) (Server, error) {
	if DonationServer == nil {
		DonationServer = &ServerImpl{
			config:  c,
			version: sVer,
		}
	}
	return DonationServer, nil
}

// NewServer create unite-did server
func NewServer(c *config.SrvcCfg, sVer *metadata.Version) (Server, error) {
	if DonationServer == nil {
		DonationServer = &ServerImpl{
			config:  c,
			version: sVer,
		}
		err := DonationServer.init()
		if err != nil {
			logger.Errorf("Failed to initialize unite did server, %+v", err)
			return nil, err
		}
	}
	return DonationServer, nil
}
