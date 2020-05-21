/*
 * Copyright ArxanChain Ltd. 2020 All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package bcadapter

import "github.com/csiabb/donation-service/common/log"

var (
	logger = log.MustGetLogger("bcadapter")
)

// BackendImpl ...
type BackendImpl struct {
	Config *Config
}

// NewBCAdapterBackend ...
func NewBCAdapterBackend(c *Config) (*BackendImpl, error) {
	logger.Infof("creating bc adapter service ...")
	d := &BackendImpl{Config: c}
	return d, nil
}
