/*
Copyright Arxan Chain Ltd. 2020 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

                 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package service

import (
	"donation-service/common/metadata"
	"donation-service/config"
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
