
package api

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/api"
    apiReq "github.com/flipped-aurora/gin-vue-admin/server/model/api/request"
)

type SysTransactionsService struct {}
// CreateSysTransactions 创建sysTransactions表记录
// Author [yourname](https://github.com/yourname)
func (sysTransactionsService *SysTransactionsService) CreateSysTransactions(ctx context.Context, sysTransactions *api.SysTransactions) (err error) {
	err = global.GVA_DB.Create(sysTransactions).Error
	return err
}

// DeleteSysTransactions 删除sysTransactions表记录
// Author [yourname](https://github.com/yourname)
func (sysTransactionsService *SysTransactionsService)DeleteSysTransactions(ctx context.Context, id string) (err error) {
	err = global.GVA_DB.Delete(&api.SysTransactions{},"id = ?",id).Error
	return err
}

// DeleteSysTransactionsByIds 批量删除sysTransactions表记录
// Author [yourname](https://github.com/yourname)
func (sysTransactionsService *SysTransactionsService)DeleteSysTransactionsByIds(ctx context.Context, ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]api.SysTransactions{},"id in ?",ids).Error
	return err
}

// UpdateSysTransactions 更新sysTransactions表记录
// Author [yourname](https://github.com/yourname)
func (sysTransactionsService *SysTransactionsService)UpdateSysTransactions(ctx context.Context, sysTransactions api.SysTransactions) (err error) {
	err = global.GVA_DB.Model(&api.SysTransactions{}).Where("id = ?",sysTransactions.Id).Updates(&sysTransactions).Error
	return err
}

// GetSysTransactions 根据id获取sysTransactions表记录
// Author [yourname](https://github.com/yourname)
func (sysTransactionsService *SysTransactionsService)GetSysTransactions(ctx context.Context, id string) (sysTransactions api.SysTransactions, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&sysTransactions).Error
	return
}
// GetSysTransactionsInfoList 分页获取sysTransactions表记录
// Author [yourname](https://github.com/yourname)
func (sysTransactionsService *SysTransactionsService)GetSysTransactionsInfoList(ctx context.Context, info apiReq.SysTransactionsSearch) (list []api.SysTransactions, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&api.SysTransactions{})
    var sysTransactionss []api.SysTransactions
    // 如果有条件搜索 下方会自动创建搜索语句
    
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&sysTransactionss).Error
	return  sysTransactionss, total, err
}
func (sysTransactionsService *SysTransactionsService)GetSysTransactionsPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
