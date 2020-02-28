/*
Copyright Lingzhu Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package impl

import (
	"fmt"
	"time"

	"github.com/csiabb/donation-service/common/rest"
	"github.com/csiabb/donation-service/common/utils"
	"github.com/csiabb/donation-service/models"
	"github.com/csiabb/donation-service/structs"

	"github.com/jinzhu/gorm"
)

const (
	sqlQueryDonationStatAndAccountInfo = "select donation_stat.id, donation_stat.uid, donation_stat.received_funds, donation_stat.received_supplies, donation_stat.distributed_funds, donation_stat.distributed_supplies, donation_stat.created_at, account.nick_name, image.url from donation_stat full join account on donation_stat.uid = account.id full join image on image.related_id = donation_stat.uid where image.type= ? and  donation_stat.created_at >= ? and donation_stat.created_at <= ? order by  donation_stat.created_at limit ? offset ?"
)

// CreateOrganizations implement create the donation statistics of organization interface
func (b *DbBackendImpl) CreateOrganizations(data *models.DonationStat) error {
	if nil == data {
		return fmt.Errorf("param is nil")
	}

	data.ID = utils.GenerateUUID()
	return b.GetConn().Create(data).Error
}

// QueryOrganizations implement query the donation statistics of organization interface
func (b *DbBackendImpl) QueryOrganizations(params *structs.QueryParams) ([]*structs.OrganizationsItems, error) {
	if params.PageNum < 1 {
		params.PageNum = rest.PageNum
	}

	if params.PageLimit < 1 {
		params.PageLimit = rest.PageLimit
	}

	if params.StartTime < 0 || params.EndTime < 0 {
		return nil, fmt.Errorf("the time value is less than 0")
	}

	if params.StartTime > 0 && params.EndTime > 0 {
		if params.EndTime < params.StartTime {
			return nil, fmt.Errorf("end time can not less than start time")
		}
	}

	var out []*structs.OrganizationsItems
	offset := (params.PageNum - 1) * params.PageLimit

	where := b.GetConn().Model(&structs.OrganizationsItems{})
	if err := where.Raw(sqlQueryDonationStatAndAccountInfo, "org", time.Unix(params.StartTime, 0),
		time.Unix(params.EndTime, 0), params.PageLimit, offset).Scan(&out).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			e := fmt.Errorf("records not found")
			logger.Error(e)
			return nil, e
		}

		logger.Errorf("query organizations records error: %v", err)
		return nil, err
	}

	return out, nil
}
