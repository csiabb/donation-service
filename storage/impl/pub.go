package impl

import (
	"fmt"
	"github.com/csiabb/donation-service/common/rest"
	"github.com/csiabb/donation-service/common/utils"
	"github.com/csiabb/donation-service/storage"
	"github.com/csiabb/donation-service/structs"
	"github.com/jinzhu/gorm"
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
func (b *DbBackendImpl) QueryFunds(uid string, params *structs.QueryParams) ([]*storage.PubFunds, error) {
	if uid == "" {
		return nil, fmt.Errorf("uid can not be \\'\\'")
	}

	if params.PageNum < 1 {
		params.PageNum = rest.PageNum
	}

	if params.PageLimit < 1 {
		params.PageLimit = rest.PageLimit
	}

	where := b.GetConn().Model(&storage.PubFunds{}).Where("uid = ?", uid)

	if params.StartTime > 0 && params.EndTime > 0 {
		if params.EndTime < params.StartTime {
			return nil, fmt.Errorf("end time can not less than start time")
		}

		where = where.Where("created_at >= ?", params.StartTime)
		where = where.Where("created_at <= ?", params.EndTime)
	}

	var out []*storage.PubFunds
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
