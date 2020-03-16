/*
Copyright Lingzhu Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package image

// BackendImpl ...
type BackendImpl struct {
	Client
}

// NewImageBackend ...
func NewImageBackend(cfg *Config) (*BackendImpl, error) {
	logger.Infof("creating db backend ...")
	d := &BackendImpl{Client: Client{ImageConfig: cfg}}

	return d, nil
}
