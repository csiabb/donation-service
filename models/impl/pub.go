/*
Copyright ArxanChain Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package impl

import (
	"fmt"
	"time"

	"github.com/csiabb/donation-service/common/rest"
	"github.com/csiabb/donation-service/common/utils"
	"github.com/csiabb/donation-service/models/storage"
	"github.com/csiabb/donation-service/structs"

	"github.com/jinzhu/gorm"
)

const (
	sqlQueryPublicityByUserType = "select * from (select id, uid, user_type, aid_uid, target_uid, pub_type, pay_type, amount, null, null, null, tx_id, remark, block_type, block_height, block_time, created_at as time from pub_funds where user_type = ? and created_at >= ? and created_at <= ? union all select id, uid, user_type, aid_uid, target_uid, pub_type, null, null, name, number, unit, tx_id, remark, block_type, block_height, block_time, created_at as time from pub_supplies where user_type = ? and created_at >= ? and created_at <= ?) as temp order by temp.time limit ? offset ?"
)

// CreateFunds implement receive funds interface
func (b *DbBackendImpl) CreateFunds(data *storage.PubFunds) error {
	if nil == data {
		return fmt.Errorf("param is nil")
	}

	data.ID = utils.GenerateUUID()
	return b.GetConn().Create(data).Error
}

// QueryFunds implement query funds interface
func (b *DbBackendImpl) QueryFunds(uid, userType, pubType string, params *structs.QueryParams) ([]*storage.PubFunds, error) {
	if params.PageNum < 1 {
		params.PageNum = rest.PageNum
	}

	if params.PageLimit < 1 {
		params.PageLimit = rest.PageLimit
	}

	where := b.GetConn().Model(&storage.PubFunds{})

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

	var out []*storage.PubFunds
	offset := (params.PageNum - 1) * params.PageLimit

	if err := where.Offset(offset).Limit(params.PageLimit).Find(&out).Count(&params.Total).Order("created_at desc").Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			e := fmt.Errorf("records not found")
			logger.Error(e)
			return nil, e
		}

		logger.Errorf("query funds records error: %v", err)
		return nil, err
	}

	return out, nil
}

// CreateSupplies defines the created supplies
func (b *DbBackendImpl) CreateSupplies(data *storage.PubSupplies) error {
	if nil == data {
		return fmt.Errorf("param is nil")
	}

	data.ID = utils.GenerateUUID()
	return b.GetConn().Create(data).Error
}

// QuerySupplies defines the query supplies
func (b *DbBackendImpl) QuerySupplies(uid, userType, pubType string, params *structs.QueryParams) ([]*storage.PubSupplies, error) {
	if params.PageNum < 1 {
		params.PageNum = rest.PageNum
	}

	if params.PageLimit < 1 {
		params.PageLimit = rest.PageLimit
	}

	where := b.GetConn().Model(&storage.PubFunds{})

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

	var out []*storage.PubSupplies
	offset := (params.PageNum - 1) * params.PageLimit

	if err := where.Offset(offset).Limit(params.PageLimit).Find(&out).Count(&params.Total).Order("created_at desc").Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			e := fmt.Errorf("records not found")
			logger.Error(e)
			return nil, e
		}

		logger.Errorf("query funds record error: %v", err)
		return nil, err
	}

	return out, nil
}

// QueryPubByUserType defines the query of publicity by user type
func (b *DbBackendImpl) QueryPubByUserType(userType string, params *structs.QueryParams) ([]*structs.PubUserItem, error) {
	if params.StartTime > 0 && params.EndTime > 0 {
		if params.EndTime < params.StartTime {
			return nil, fmt.Errorf("end time can not less than start time")
		}
	} else {
		now := time.Now()
		params.EndTime = now.Unix()
		params.StartTime = params.EndTime - rest.TenDaysBySecond
	}

	offset := (params.PageNum - 1) * params.PageLimit
	var out []*structs.PubUserItem
	err := b.GetConn().Raw(sqlQueryPublicityByUserType, userType, time.Unix(params.StartTime, 0), time.Unix(params.EndTime, 0), userType, time.Unix(params.StartTime, 0), time.Unix(params.EndTime, 0), params.PageLimit, offset).Scan(&out).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			e := fmt.Errorf("records not found")
			logger.Error(e)
			return nil, e
		}

		logger.Errorf("query records error: %v", err)
		return nil, err
	}

	params.Total = int64(len(out))
	return out, nil
}

// QueryFundsDetail defines query publicity funds detail
func (b *DbBackendImpl) QueryFundsDetail(id string) (*storage.FundsDetail, error) {
	if id == "" {
		e := fmt.Errorf("id can not be \\'\\'")
		logger.Error(e)
		return nil, e
	}

	detail := storage.FundsDetail{}
	if err := b.GetConn().Where(&storage.PubFunds{ID: id}).First(&detail.Funds).Error; err != nil {
		e := fmt.Errorf("query funds error, %v", err)
		logger.Error(e)
		return nil, e
	}

	if err := b.GetConn().Where(&storage.Address{UID: detail.Funds.UID, Type: rest.AddrTypeShipping}).First(&detail.BillingAddr).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			e := fmt.Errorf("billing address records not found")
			logger.Debug(e)
		} else {
			e := fmt.Errorf("query billing address error, %v", err)
			logger.Error(e)
			return nil, e
		}
	}

	if err := b.GetConn().Where(&storage.Address{UID: detail.Funds.TargetUID, Type: rest.AddrTypeShipping}).First(&detail.ShippingAddr).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			e := fmt.Errorf("shipping address records not found")
			logger.Debug(e)
		} else {
			e := fmt.Errorf("query shipping address error, %v", err)
			logger.Error(e)
			return nil, e
		}
	}

	if err := b.GetConn().Where(&storage.Image{RelatedID: id, Type: rest.ImageProof}).Find(&detail.ProofImages).Error; err != nil {
		e := fmt.Errorf("query images error, %v", err)
		logger.Error(e)
		return nil, e
	}

	return &detail, nil
}

// QuerySuppliesDetail defines query publicity funds detail
func (b *DbBackendImpl) QuerySuppliesDetail(id string) (*storage.SuppliesDetail, error) {
	if id == "" {
		e := fmt.Errorf("id can not be \\'\\'")
		logger.Error(e)
		return nil, e
	}

	detail := storage.SuppliesDetail{}
	if err := b.GetConn().Where(&storage.PubSupplies{ID: id}).First(&detail.Supplies).Error; err != nil {
		e := fmt.Errorf("query supplies error, %v", err)
		logger.Error(e)
		return nil, e
	}

	if err := b.GetConn().Where(&storage.Address{UID: detail.Supplies.UID, Type: rest.AddrTypeShipping}).First(&detail.BillingAddr).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			e := fmt.Errorf("billing address records not found")
			logger.Debug(e)
		} else {
			e := fmt.Errorf("query billing address error, %v", err)
			logger.Error(e)
			return nil, e
		}
	}

	if err := b.GetConn().Where(&storage.Address{UID: detail.Supplies.TargetUID, Type: rest.AddrTypeShipping}).First(&detail.ShippingAddr).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			e := fmt.Errorf("shipping address records not found")
			logger.Debug(e)
		} else {
			e := fmt.Errorf("query shipping address error, %v", err)
			logger.Error(e)
			return nil, e
		}
	}

	if err := b.GetConn().Where(&storage.Image{RelatedID: id, Type: rest.ImageProof}).Find(&detail.ProofImages).Error; err != nil {
		e := fmt.Errorf("query images error, %v", err)
		logger.Error(e)
		return nil, e
	}

	return &detail, nil
}
