/*
Copyright ArxanChain Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package database

// DBConnectCfg database connect configure
type DBConnectCfg struct {
	Enabled      bool
	Driver       string // postgres, mysql
	Address      string
	DBname       string
	User         string
	Password     string
	SSLMode      string
	MaxIdleConns int
	MaxConns     int
}
