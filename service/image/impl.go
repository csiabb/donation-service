/*
Copyright Lingzhu Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package image

import (
	"github.com/csiabb/donation-service/common/log"
	"github.com/csiabb/donation-service/context"
)

var (
	logger = log.MustGetLogger("images-service")
)

// ImagesImpl ...
type ImagesImpl struct {
	context *context.Context
}

// NewImages ...
func NewImages(c *context.Context) (*ImagesImpl, error) {
	return &ImagesImpl{context: c}, nil
}
