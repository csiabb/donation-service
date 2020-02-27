/*
Copyright Lingzhu Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package structs

import (
	"github.com/shopspring/decimal"
)

// QueryFundsRequest defines the request of query organizations
type QueryOrganizationsRequest struct {
	PageNum   int   `form:"page_num"`   // page num
	PageLimit int   `form:"page_limit"` // page limit
	StartTime int64 `form:"start_time"` // start time
	EndTime   int64 `form:"end_time"`   // end time
}

// OrganizationInformationItem defines the struct of organization item
type OrganizationsItems struct {
	ID                  string          `json:"id"`                   // funds id
	UID                 string          `json:"uid"`                  // user id of the one who donate
	URL                 string          `json:"url"`                  // organization logo
	NickName            string          `json:"nick_name"`            // nick name
	ReceivedFunds       decimal.Decimal `json:"received_funds"`       // receiving funds
	ReceivedSupplies    int64           `json:"received_supplies"`    // receiving supplies
	DistributedFunds    decimal.Decimal `json:"distributed_funds"`    // distributing  funds
	DistributedSupplies int64           `json:"distributed_supplies"` // distributing supplies
	CreatedAt           int64           `json:"created_at"`           // create time
}

// QueryOrganizationsResp defines the response of organizations
type QueryOrganizationsResp struct {
	PageNum   int                   `json:"page_num"`   // page num
	PageLimit int                   `json:"page_limit"` // page limit
	StartTime int64                 `json:"start_time"` // start time
	EndTime   int64                 `json:"end_time"`   // end time
	Total     int64                 `json:"total"`      // total number of query result
	Results   []*OrganizationsItems `json:"results"`    // funds items
}
