/*
Copyright ArxanChain Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package config

import (
	"fmt"

	"github.com/csiabb/donation-service/common/log"
	"github.com/csiabb/donation-service/components/database"
)

var (
	logger = log.MustGetLogger("config")

	// LeaveOnInt quit server on int signal
	LeaveOnInt = true
	// LeaveOnTerm quit server on terminate signal
	LeaveOnTerm = true
)

// SrvcCfg  service configure
type SrvcCfg struct {
	ServerGeneral ServerGeneralCfg
	Log           log.Config
	Database      database.DBConnectCfg
}

// ServerGeneralCfg general configure of service
type ServerGeneralCfg struct {
	Host string
	Port int
}

// GetServiceCfg returns the configurations for the service
func GetServiceCfg(progName string) *SrvcCfg {
	rcfg := SrvcCfg{}
	parser := initConfig(progName)
	err := parser.Unmarshal(&rcfg)
	if err != nil {
		logger.Panic("Error loading configuration: ", err)
	}
	logger.Debugf("starting client with configuration: %+v", rcfg)
	fmt.Println(rcfg)

	return &rcfg
}
