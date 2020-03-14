/*
Copyright Lingzhu Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package aliyun

import (
	"github.com/csiabb/donation-service/common/log"
	a "github.com/csiabb/donation-service/components/aliyun"
)

var (
	logger = log.MustGetLogger("models/aliyun")
)

// ALiYunBackendImpl ...
type ALiYunBackendImpl struct {
	a.ALiYunClient
}

// NewALiYunBackend ...
func NewALiYunBackend(cfg *a.ALiYunCfg) (*ALiYunBackendImpl, error) {
	logger.Infof("creating aliyun service ...")
	d := &ALiYunBackendImpl{ALiYunClient: a.ALiYunClient{ALiYunConfig: cfg}}
	return d, nil
}
