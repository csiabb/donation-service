/*
Copyright ArxanChain Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package models

import (
	"time"
)

// Account account model
type Account struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
