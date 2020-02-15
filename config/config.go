/*
Copyright ArxanChain Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package config

import (
	"fmt"
	
	"github.com/op/go-logging"
)

var (
	logger = logging.MustGetLogger("config")

	// LeaveOnInt quit server on int signal
	LeaveOnInt = true
	// LeaveOnTerm quit server on terminate signal
	LeaveOnTerm = true
)

// SrvcCfg  service configure
type SrvcCfg struct {
	ServerGeneral ServerGeneralCfg
	Log           LogCfg
}

// ServerGeneralCfg general configure of service
type ServerGeneralCfg struct {
	Host string
	Port int
}

// LogCfg config with rolling backend
// MaxSize is the maximum size in megabytes
// MaxBackups is the maximum number of old log files to retain
// MaxAge is the maximum number of days to retain old log files
type LogCfg struct {
	LogFile    string
	LogLevel   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
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
