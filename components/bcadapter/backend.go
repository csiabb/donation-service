/*
 * Copyright ArxanChain Ltd. 2020 All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package bcadapter

import (
	"github.com/csiabb/donation-service/structs"
)

//go:generate mockgen -destination=mock_bcadapter/mock_bcadapter.go -package=mock_bcadapter github.com/csiabb/donation-service/components/bcadapter IBCAdapter

// IBCAdapter defines the interface to request block chain adapter
type IBCAdapter interface {
	Register(accountID string) (*structs.RegisterResp, error)
	Pubs(bcID string, pubs []*string) ([]*structs.PubResp, error)
}
