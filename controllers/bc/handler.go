/*
Copyright Lingzhu Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package bc

import (
	"github.com/csiabb/donation-service/common/log"
	"github.com/csiabb/donation-service/context"
)

var (
	logger = log.MustGetLogger("bc")
)

// RestHandler block chain handler
type RestHandler struct {
	srvcContext *context.Context
}

// NewRestHandler ...
func NewRestHandler(c *context.Context) (*RestHandler, error) {
	return &RestHandler{srvcContext: c}, nil
}
