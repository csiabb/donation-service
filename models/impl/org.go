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
	sqlQueryDonationStatAndAccountInfo       = "select temp.id, temp.uid, temp.received_funds, temp.received_supplies, temp.distributed_funds, temp.distributed_supplies, temp.time, temp.nick_name, image.url from (select donation_stat.id, donation_stat.uid, donation_stat.received_funds, donation_stat.received_supplies, donation_stat.distributed_funds, donation_stat.distributed_supplies, donation_stat.created_at as time, account.nick_name from donation_stat full join account on donation_stat.uid = account.id and account.type = ? where donation_stat.created_at >= ? and donation_stat.created_at <= ? order by distributed_funds desc, distributed_supplies desc ) as temp left join image on image.related_id = temp.uid and image.type = ? order by temp.time limit ? offset ? "
	sqlQueryDetailDonationStatAndAccountInfo = "select temp.nick_name, temp.remark, temp.uid, temp.phone, temp.bank_card_num, temp.country, temp.district, temp.province, temp.city, temp.address,image.url from (select account.nick_name, account.remark, account.id as uid, account.phone, account.bank_card_num, address.country, address.district, address.province, address.city, address.address from account left join address on address.uid = account.id  and address.type = ? where account.type = ?) as temp left join image on image.related_id = temp.uid and image.type = ? where temp.uid = ?"
)

// CreateOrganization implement create the donation statistics of organization interface
func (b *DbBackendImpl) CreateOrganization(data *models.DonationStat) error {
	if nil == data {
		return fmt.Errorf("param is nil")
	}

	data.ID = utils.GenerateUUID()
	return b.GetConn().Create(data).Error
}

// QueryOrgCharities implement query the donation statistics of organization charity interface
func (b *DbBackendImpl) QueryOrgCharities(params *structs.QueryParams) ([]*structs.OrgCharitiesItems, error) {
	if params.PageNum < 1 {
		params.PageNum = rest.PageNum
	}

	if params.PageLimit < 1 {
		params.PageLimit = rest.PageLimit
	}

	if params.StartTime > 0 && params.EndTime > 0 {
		if params.EndTime < params.StartTime {
			return nil, fmt.Errorf("end time can not less than start time")
		}
	} else {
		now := time.Now()
		params.EndTime = now.Unix()
		params.StartTime = params.EndTime - rest.TenDayBySecond
	}

	var out []*structs.OrgCharitiesItems
	offset := (params.PageNum - 1) * params.PageLimit
	where := b.GetConn().Model(&structs.OrgCharitiesItems{})
	if err := where.Raw(sqlQueryDonationStatAndAccountInfo, rest.UserTypeOrgCharity, time.Unix(params.StartTime, 0),
		time.Unix(params.EndTime, 0), rest.ImageAvatar, params.PageLimit, offset).Scan(&out).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			e := fmt.Errorf("records not found")
			logger.Error(e)
			return nil, e
		}

		logger.Errorf("query organizations records error , %v", err)
		return nil, err
	}

	params.Total = int64(len(out))
	return out, nil
}

// QueryOrgCharitiesDetail implement query detail the donation statistics of charities interface
func (b *DbBackendImpl) QueryOrgCharitiesDetail(uid string) (*structs.OrgCharitiesDetailItem, error) {
	if 0 == len(uid) {
		return nil, fmt.Errorf("param is nil")
	}

	var out structs.OrgCharitiesDetailItem
	where := b.GetConn().Model(&structs.OrgCharitiesDetailItem{})
	if err := where.Raw(sqlQueryDetailDonationStatAndAccountInfo, rest.AddrReg, rest.UserTypeOrgCharity, rest.ImageAvatar, uid).Scan(&out).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			e := fmt.Errorf("record not found")
			logger.Error(e)
			return nil, e
		}

		logger.Errorf("query organizations detail record error , %v", err)
		return nil, err
	}

	return &out, nil
}
