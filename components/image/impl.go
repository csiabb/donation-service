/*
Copyright Lingzhu Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package image

import "github.com/csiabb/donation-service/components/wx"

// BackendImpl ...
type BackendImpl struct {
	Client
}

// NewImageBackend ...
func NewImageBackend(cfg *Config, wxClient wx.IWXClient) (*BackendImpl, error) {
	logger.Infof("creating db backend ...")
	d := &BackendImpl{Client: Client{ImageConfig: cfg, WXClient: wxClient}}

	return d, nil
}
