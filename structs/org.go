/*
Copyright Lingzhu Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package structs

import (
	"time"

	"github.com/shopspring/decimal"
)

// QueryOrgCharitiesRequest defines the request of query organizations
type QueryOrgCharitiesRequest struct {
	PageNum   int   `form:"page_num"`   // page num
	PageLimit int   `form:"page_limit"` // page limit
	StartTime int64 `form:"start_time"` // start time
	EndTime   int64 `form:"end_time"`   // end time
}

// OrgCharitiesItems defines the struct of organization item
type OrgCharitiesItems struct {
	ID                  string          `json:"id"`                   // donation stat id
	UID                 string          `json:"uid"`                  // user id of the one who donate
	URL                 string          `json:"url"`                  // organization logo
	NickName            string          `json:"nick_name"`            // nick name
	ReceivedFunds       decimal.Decimal `json:"received_funds"`       // receiving funds
	ReceivedSupplies    int64           `json:"received_supplies"`    // receiving supplies
	DistributedFunds    decimal.Decimal `json:"distributed_funds"`    // distributing  funds
	DistributedSupplies int64           `json:"distributed_supplies"` // distributing supplies
	CreatedAt           int64           `json:"created_at"`           // create time
	Time                time.Time       `json:"-"`                    // time
}

// ConvertTime defines the covert of created_at
func (ois *OrgCharitiesItems) ConvertTime() {
	ois.CreatedAt = ois.Time.UTC().Unix()
}

// QueryOrgCharitiesResp defines the response of organizations
type QueryOrgCharitiesResp struct {
	PageNum   int                  `json:"page_num"`   // page num
	PageLimit int                  `json:"page_limit"` // page limit
	StartTime int64                `json:"start_time"` // start time
	EndTime   int64                `json:"end_time"`   // end time
	Total     int64                `json:"total"`      // total number of query result
	Results   []*OrgCharitiesItems `json:"results"`    // orgs items
}

// OrgCharitiesDetailRequest defines the request of query charities detail
type OrgCharitiesDetailRequest struct {
	UID string `form:"uid"` // user id of the one who donate
}

// OrgCharitiesDetailItem defines the struct of charities detail item
type OrgCharitiesDetailItem struct {
	UID         string `json:"uid"`           // user id of the one who donate
	URL         string `json:"url"`           // image url
	NickName    string `json:"nick_name"`     // nick name
	Country     string `json:"country"`       // country
	Province    string `json:"province"`      // province
	City        string `json:"city"`          // city
	District    string `json:"district"`      // district
	Address     string `json:"address"`       // detail address
	Phone       string `json:"phone"`         // phone num
	BankCardNum string `json:"bank_card_num"` // bank card num
	Remark      string `json:"remark"`        // remark
}
