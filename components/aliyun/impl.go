/*
Copyright Lingzhu Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package aliyun

import (
	"github.com/csiabb/donation-service/common/log"
)

var (
	logger = log.MustGetLogger("models/aliyun")
)

// BackendImpl ...
type BackendImpl struct {
	Client
}

// NewALiYunBackend ...
func NewALiYunBackend(cfg *Config) (*BackendImpl, error) {
	logger.Infof("creating aliyun service ...")
	d := &BackendImpl{Client: Client{ALiYunConfig: cfg}}
	return d, nil
}
