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
	"github.com/csiabb/donation-service/models"
	"github.com/csiabb/donation-service/structs"

	"github.com/jinzhu/gorm"
)

const (
	sqlQueryPublicityByUserType = "select * from (select id, 'funds' as type, uid, donor_name, user_type, aid_uid, aid_name, target_uid, target_name, pub_type, pay_type, amount, null as name, null as number, null as unit, tx_id, remark, block_type, block_height, block_time, created_at as time from pub_funds where user_type = ? and pub_type = ? and created_at >= ? and created_at <= ? union all select id, 'supplies' as type, uid, donor_name, user_type, aid_uid, aid_name, target_uid, target_name, pub_type, null as pay_type, null as amount, name, number, unit, tx_id, remark, block_type, block_height, block_time, created_at as time from pub_supplies where user_type = ? and pub_type = ? and created_at >= ? and created_at <= ?) as temp order by temp.time limit ? offset ?"
	sqlQueryPublicityByCharity  = "select * from (select id, 'funds' as type, uid, donor_name, user_type, aid_uid, aid_name, target_uid, target_name, pub_type, pay_type, amount, null as name, null as number, null as unit, tx_id, remark, block_type, block_height, block_time, created_at as time from pub_funds where target_uid = ? and pub_type = ? and created_at >= ? and created_at <= ? union all select id, 'supplies' as type, uid, donor_name, user_type, aid_uid, aid_name, target_uid, target_name, pub_type, null as pay_type, null as amount, name, number, unit, tx_id, remark, block_type, block_height, block_time, created_at as time from pub_supplies where target_uid = ? and pub_type = ? and created_at >= ? and created_at <= ?) as temp order by temp.time limit ? offset ?"
)

// CreateFunds implement receive funds interface
func (b *DbBackendImpl) CreateFunds(tx *gorm.DB, data *models.PubFunds) error {
	if nil == data {
		return fmt.Errorf("param is nil")
	}

	data.ID = utils.GenerateUUID()
	err := tx.Create(data).Error
	return err
}

// CreateImages implement create images interface
func (b *DbBackendImpl) CreateImages(tx *gorm.DB, data []*models.Image) error {
	if nil == data {
		return fmt.Errorf("param is nil")
	}

	for _, v := range data {
		err := tx.Create(v).Error
		if err != nil {
			return err
		}
	}

	return nil
}

// QueryFunds implement query funds interface
func (b *DbBackendImpl) QueryFunds(uid, targetUID, userType, pubType string, params *structs.QueryParams) ([]*models.PubFunds, error) {
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

	where := b.GetConn().Model(&models.PubFunds{})
	where = where.Where("created_at >= ?", time.Unix(params.StartTime, 0))
	where = where.Where("created_at <= ?", time.Unix(params.EndTime, 0))

	if uid != "" {
		where = where.Where("uid = ?", uid)
	}

	if userType != "" {
		where = where.Where("user_type = ?", userType)
	}

	if pubType != "" {
		where = where.Where("pub_type = ?", pubType)
	}

	if targetUID != "" {
		where = where.Where("target_uid = ?", targetUID)
	}

	var out []*models.PubFunds
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
func (b *DbBackendImpl) CreateSupplies(tx *gorm.DB, data []*models.PubSupplies) error {
	if nil == data {
		return fmt.Errorf("param is nil")
	}

	for _, v := range data {
		err := tx.Create(v).Error
		if err != nil {
			return err
		}
	}

	return nil
}

// CreateAddresses implement create addresses
func (b *DbBackendImpl) CreateAddresses(tx *gorm.DB, data []*models.Address) error {
	if nil == data {
		return fmt.Errorf("param is nil")
	}

	for _, v := range data {
		err := tx.Create(v).Error
		if err != nil {
			return err
		}
	}

	return nil
}

// QuerySupplies defines the query supplies
func (b *DbBackendImpl) QuerySupplies(uid, targetUID, userType, pubType string, params *structs.QueryParams) ([]*models.PubSupplies, error) {
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

	where := b.GetConn().Model(&models.PubFunds{})
	where = where.Where("created_at >= ?", time.Unix(params.StartTime, 0))
	where = where.Where("created_at <= ?", time.Unix(params.EndTime, 0))

	if uid != "" {
		where = where.Where("uid = ?", uid)
	}

	if userType != "" {
		where = where.Where("user_type = ?", userType)
	}

	if pubType != "" {
		where = where.Where("pub_type = ?", pubType)
	}

	if targetUID != "" {
		where = where.Where("target_uid = ?", targetUID)
	}

	var out []*models.PubSupplies
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
func (b *DbBackendImpl) QueryPubByUserType(userType, targetUID, pubType string, params *structs.QueryParams) ([]*structs.PubUserItem, error) {
	if params.StartTime > 0 && params.EndTime > 0 {
		if params.EndTime < params.StartTime {
			return nil, fmt.Errorf("end time can not less than start time")
		}
	} else {
		now := time.Now()
		params.EndTime = now.Unix()
		params.StartTime = params.EndTime - rest.TenDayBySecond
	}

	offset := (params.PageNum - 1) * params.PageLimit
	var out []*structs.PubUserItem
	var err error

	if pubType == "" {
		return nil, fmt.Errorf("pub type can not be \\'\\'")
	}

	if userType == "" && targetUID == "" {
		return nil, fmt.Errorf("user type and target id can not be \\'\\' the same time")
	}

	if userType != "" {
		err = b.GetConn().Raw(sqlQueryPublicityByUserType, userType, pubType, time.Unix(params.StartTime, 0), time.Unix(params.EndTime, 0), userType, pubType, time.Unix(params.StartTime, 0), time.Unix(params.EndTime, 0), params.PageLimit, offset).Scan(&out).Error
	} else if targetUID != "" {
		err = b.GetConn().Raw(sqlQueryPublicityByCharity, targetUID, pubType, time.Unix(params.StartTime, 0), time.Unix(params.EndTime, 0), targetUID, pubType, time.Unix(params.StartTime, 0), time.Unix(params.EndTime, 0), params.PageLimit, offset).Scan(&out).Error
	}

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
func (b *DbBackendImpl) QueryFundsDetail(id string) (*models.FundsDetail, error) {
	if id == "" {
		e := fmt.Errorf("id can not be \\'\\'")
		logger.Error(e)
		return nil, e
	}

	detail := models.FundsDetail{}
	if err := b.GetConn().Where(&models.PubFunds{ID: id}).First(&detail.Funds).Error; err != nil {
		e := fmt.Errorf("query funds error, %v", err)
		logger.Error(e)
		return nil, e
	}

	if err := b.GetConn().Where(&models.Address{UID: detail.Funds.UID, Type: rest.AddrShipping}).First(&detail.BillingAddr).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			e := fmt.Errorf("billing address records not found")
			logger.Debug(e)
		} else {
			e := fmt.Errorf("query billing address error, %v", err)
			logger.Error(e)
			return nil, e
		}
	}

	if err := b.GetConn().Where(&models.Address{UID: detail.Funds.TargetUID, Type: rest.AddrShipping}).First(&detail.ShippingAddr).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			e := fmt.Errorf("shipping address records not found")
			logger.Debug(e)
		} else {
			e := fmt.Errorf("query shipping address error, %v", err)
			logger.Error(e)
			return nil, e
		}
	}

	if err := b.GetConn().Where(&models.Image{RelatedID: id, Type: rest.ImageProof}).Find(&detail.ProofImages).Error; err != nil {
		e := fmt.Errorf("query images error, %v", err)
		logger.Error(e)
		return nil, e
	}

	return &detail, nil
}

// QuerySuppliesDetail defines query publicity funds detail
func (b *DbBackendImpl) QuerySuppliesDetail(id string) (*models.SuppliesDetail, error) {
	if id == "" {
		e := fmt.Errorf("id can not be \\'\\'")
		logger.Error(e)
		return nil, e
	}

	detail := models.SuppliesDetail{}
	if err := b.GetConn().Where(&models.PubSupplies{ID: id}).First(&detail.Supplies).Error; err != nil {
		e := fmt.Errorf("query supplies error, %v", err)
		logger.Error(e)
		return nil, e
	}

	if err := b.GetConn().Where(&models.Address{UID: detail.Supplies.UID, Type: rest.AddrShipping}).First(&detail.BillingAddr).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			e := fmt.Errorf("billing address records not found")
			logger.Debug(e)
		} else {
			e := fmt.Errorf("query billing address error, %v", err)
			logger.Error(e)
			return nil, e
		}
	}

	if err := b.GetConn().Where(&models.Address{UID: detail.Supplies.TargetUID, Type: rest.AddrShipping}).First(&detail.ShippingAddr).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			e := fmt.Errorf("shipping address records not found")
			logger.Debug(e)
		} else {
			e := fmt.Errorf("query shipping address error, %v", err)
			logger.Error(e)
			return nil, e
		}
	}

	if err := b.GetConn().Where(&models.Image{RelatedID: id, Type: rest.ImageProof}).Find(&detail.ProofImages).Error; err != nil {
		e := fmt.Errorf("query images error, %v", err)
		logger.Error(e)
		return nil, e
	}

	return &detail, nil
}
