/*
Copyright Lingzhu Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package donation

import (
	"github.com/csiabb/donation-service/common/log"
	"github.com/csiabb/donation-service/context"
)

var (
	logger = log.MustGetLogger("donation-handler")
)

// RestHandler donation handler
type RestHandler struct {
	srvcContext *context.Context
}

// NewRestHandler ...
func NewRestHandler(c *context.Context) (*RestHandler, error) {
	return &RestHandler{srvcContext: c}, nil
}
