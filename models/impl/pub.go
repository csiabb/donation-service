/*
Copyright ArxanChain Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package impl

import (
	"fmt"

	"github.com/csiabb/donation-service/common/rest"
	"github.com/csiabb/donation-service/common/utils"
	"github.com/csiabb/donation-service/models"
	"github.com/csiabb/donation-service/structs"

	"github.com/jinzhu/gorm"
)

// CreateFunds implement receive funds interface
func (b *DbBackendImpl) CreateFunds(data *models.PubFunds) error {
	if nil == data {
		return fmt.Errorf("param is nil")
	}

	data.ID = utils.GenerateUUID()
	return b.GetConn().Create(data).Error
}

// QueryFunds implement query funds interface
func (b *DbBackendImpl) QueryFunds(uid, userType, pubType string, params *structs.QueryParams) ([]*models.PubFunds, error) {
	if params.PageNum < 1 {
		params.PageNum = rest.PageNum
	}

	if params.PageLimit < 1 {
		params.PageLimit = rest.PageLimit
	}

	where := b.GetConn().Model(&models.PubFunds{})

	if uid != "" {
		where = where.Where("uid = ?", uid)
	}

	if params.StartTime > 0 && params.EndTime > 0 {
		if params.EndTime < params.StartTime {
			return nil, fmt.Errorf("end time can not less than start time")
		}

		where = where.Where("created_at >= ?", params.StartTime)
		where = where.Where("created_at <= ?", params.EndTime)
	}

	if userType != "" {
		where = where.Where("user_type = ?", userType)
	}

	if pubType != "" {
		where = where.Where("pub_type = ?", pubType)
	}

	var out []*models.PubFunds
	offset := (params.PageNum - 1) * params.PageLimit

	if err := where.Offset(offset).Limit(params.PageLimit).Find(&out).Count(&params.Total).Order("created_at desc").Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			e := fmt.Errorf("record not found")
			logger.Error(e)
			return nil, e
		}

		logger.Errorf("query funds record error: %v", err)
		return nil, err
	}

	return out, nil
}

// CreateSupplies defines the created supplies
func (b *DbBackendImpl) CreateSupplies(data *models.PubSupplies) error {
	if nil == data {
		return fmt.Errorf("param is nil")
	}

	data.ID = utils.GenerateUUID()
	return b.GetConn().Create(data).Error
}

// QuerySupplies defines the query supplies
func (b *DbBackendImpl) QuerySupplies(uid, userType, pubType string, params *structs.QueryParams) ([]*models.PubSupplies, error) {
	if params.PageNum < 1 {
		params.PageNum = rest.PageNum
	}

	if params.PageLimit < 1 {
		params.PageLimit = rest.PageLimit
	}

	where := b.GetConn().Model(&models.PubFunds{})

	if uid != "" {
		where = where.Where("uid = ?", uid)
	}

	if params.StartTime > 0 && params.EndTime > 0 {
		if params.EndTime < params.StartTime {
			return nil, fmt.Errorf("end time can not less than start time")
		}

		where = where.Where("created_at >= ?", params.StartTime)
		where = where.Where("created_at <= ?", params.EndTime)
	}

	if userType != "" {
		where = where.Where("user_type = ?", userType)
	}

	if pubType != "" {
		where = where.Where("pub_type = ?", pubType)
	}

	var out []*models.PubSupplies
	offset := (params.PageNum - 1) * params.PageLimit

	if err := where.Offset(offset).Limit(params.PageLimit).Find(&out).Count(&params.Total).Order("created_at desc").Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			e := fmt.Errorf("record not found")
			logger.Error(e)
			return nil, e
		}

		logger.Errorf("query funds record error: %v", err)
		return nil, err
	}

	return out, nil
}
