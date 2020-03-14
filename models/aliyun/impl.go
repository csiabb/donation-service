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

// BackendImpl ...
type BackendImpl struct {
	a.Client
}

// NewALiYunBackend ...
func NewALiYunBackend(cfg *a.Config) (*BackendImpl, error) {
	logger.Infof("creating aliyun service ...")
	d := &BackendImpl{Client: a.Client{ALiYunConfig: cfg}}
	return d, nil
}
