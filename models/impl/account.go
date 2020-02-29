/*
Copyright ArxanChain Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package impl

import (
	"fmt"

	"github.com/csiabb/donation-service/models/storage"
)

// CreateAccount implement create account interface
func (b *DbBackendImpl) CreateAccount(data *storage.Account) error {
	if nil == data {
		return fmt.Errorf("param is nil")
	}

	return b.GetConn().Create(data).Error
}
