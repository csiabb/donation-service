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
	logger = log.MustGetLogger("donations-service")
)

// DonationsImpl ...
type DonationsImpl struct {
	context *context.Context
}

// NewDonations ...
func NewDonations(c *context.Context) (*DonationsImpl, error) {
	return &DonationsImpl{context: c}, nil
}
