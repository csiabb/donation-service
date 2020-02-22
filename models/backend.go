/*
Copyright ArxanChain Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package models

// IDBBackend database operate interface
type IDBBackend interface {
	CreateAccount(*Account) error
	ListDonationsByFund(*Request) ([]*PubFunds, error)
	ListDonationsBySupply(*Request) ([]*PubSupply, error)
}
