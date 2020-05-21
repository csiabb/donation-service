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

// CreateAccount implement create user account
func (b *DbBackendImpl) CreateAccount(data *models.Account) error {
	if nil == data {
		return fmt.Errorf("param is nil")
	}

	return b.GetConn().Create(data).Error
}

// QueryAccount implement check user account exist or not
func (b *DbBackendImpl) QueryAccount(openID, uid string) (*models.Account, error) {
	where := b.GetConn().Model(&models.Account{})

	if openID != "" {
		where = where.Where("open_id = ?", openID)
	}

	if uid != "" {
		where = where.Where("id = ?", uid)
	}

	acc := &models.Account{}
	err := where.First(acc).Error
	return acc, err
}
