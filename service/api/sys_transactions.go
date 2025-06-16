package api

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/api"
	apiReq "github.com/flipped-aurora/gin-vue-admin/server/model/api/request"
)

type SysTransactionsService struct{}

// CreateSysTransactions sysTransactions
// Author [yourname](https://github.com/yourname)
func (sysTransactionsService *SysTransactionsService) CreateSysTransactions(ctx context.Context, sysTransactions *api.SysTransactions) (err error) {
	err = global.GVA_DB.Create(sysTransactions).Error
	return err
}

// DeleteSysTransactions sysTransactions
// Author [yourname](https://github.com/yourname)
func (sysTransactionsService *SysTransactionsService) DeleteSysTransactions(ctx context.Context, id string) (err error) {
	err = global.GVA_DB.Delete(&api.SysTransactions{}, "id = ?", id).Error
	return err
}

// DeleteSysTransactionsByIds sysTransactions
// Author [yourname](https://github.com/yourname)
func (sysTransactionsService *SysTransactionsService) DeleteSysTransactionsByIds(ctx context.Context, ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]api.SysTransactions{}, "id in ?", ids).Error
	return err
}

// UpdateSysTransactions sysTransactions
// Author [yourname](https://github.com/yourname)
func (sysTransactionsService *SysTransactionsService) UpdateSysTransactions(ctx context.Context, sysTransactions api.SysTransactions) (err error) {
	err = global.GVA_DB.Model(&api.SysTransactions{}).Where("id = ?", sysTransactions.Id).Updates(&sysTransactions).Error
	return err
}

// GetSysTransactions idsysTransactions
// Author [yourname](https://github.com/yourname)
func (sysTransactionsService *SysTransactionsService) GetSysTransactions(ctx context.Context, id string) (sysTransactions api.SysTransactions, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&sysTransactions).Error
	return
}

// GetSysTransactionsInfoList sysTransactions
// Author [yourname](https://github.com/yourname)
func (sysTransactionsService *SysTransactionsService) GetSysTransactionsInfoList(ctx context.Context, info apiReq.SysTransactionsSearch) (list []api.SysTransactions, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// db
	db := global.GVA_DB.Model(&api.SysTransactions{})
	var sysTransactionss []api.SysTransactions
	//

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&sysTransactionss).Error
	return sysTransactionss, total, err
}
func (sysTransactionsService *SysTransactionsService) GetSysTransactionsPublic(ctx context.Context) {
	//
	//
}
