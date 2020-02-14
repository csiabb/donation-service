/*
Copyright Arxan Chain Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package config

import (
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
}

// GetServiceCfg returns the configurations for the service
func GetServiceCfg(progName string) *SrvcCfg {
	rcfg := SrvcCfg{}
	logger.Debugf("starting client with configuration: %+v", rcfg)

	return &rcfg
}
