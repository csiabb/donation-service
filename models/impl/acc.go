/*
 * Copyright ArxanChain Ltd. 2020 All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package impl

import (
	"fmt"

	"github.com/csiabb/donation-service/models"
)

// CheckAccount implement check user account exist or not
func (b *DbBackendImpl) CheckAccount(openID string) (*models.Account, error) {
	where := b.GetConn().Model(&models.Account{})

	if openID != "" {
		where = where.Where("open_id = ?", openID)
	}

	acc := &models.Account{}
	err := where.First(acc).Error
	return acc, err
}

// CreateAccount implement create user account
func (b *DbBackendImpl) CreateAccount(data *models.Account) error {
	if nil == data {
		return fmt.Errorf("param is nil")
	}

	return b.GetConn().Create(data).Error
}
