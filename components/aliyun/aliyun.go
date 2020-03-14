/*
Copyright Lingzhu Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package aliyun

// GormClient gorm database client
type ALiYunClient struct {
	ALiYunConfig *ALiYunCfg
}

// GetConn get database connect
func (d *ALiYunClient) GetConn() *ALiYunCfg {
	return d.ALiYunConfig
}
