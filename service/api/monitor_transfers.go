package api

import (
	"context"
	"errors"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/api"
	apiReq "github.com/flipped-aurora/gin-vue-admin/server/model/api/request"
	"gorm.io/gorm"
)

type MonitorTransfersService struct{}

// CreateMonitorTransfers monitorTransfers表
// Author [yourname](https://github.com/yourname)
func (monitorTransfersService *MonitorTransfersService) CreateMonitorTransfers(ctx context.Context, monitorTransfers *api.MonitorTransfers) (err error) {
	err = global.GVA_DB.Create(monitorTransfers).Error
	return err
}

// DeleteMonitorTransfers monitorTransfers表
// Author [yourname](https://github.com/yourname)
func (monitorTransfersService *MonitorTransfersService) DeleteMonitorTransfers(ctx context.Context, id string) (err error) {
	err = global.GVA_DB.Delete(&api.MonitorTransfers{}, "id = ?", id).Error
	return err
}

// DeleteMonitorTransfersByIds monitorTransfers表
// Author [yourname](https://github.com/yourname)
func (monitorTransfersService *MonitorTransfersService) DeleteMonitorTransfersByIds(ctx context.Context, ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]api.MonitorTransfers{}, "id in ?", ids).Error
	return err
}

// UpdateMonitorTransfers monitorTransfers表
// Author [yourname](https://github.com/yourname)
func (monitorTransfersService *MonitorTransfersService) UpdateMonitorTransfers(ctx context.Context, monitorTransfers api.MonitorTransfers) (err error) {
	err = global.GVA_DB.Model(&api.MonitorTransfers{}).Where("id = ?", monitorTransfers.Id).Updates(&monitorTransfers).Error
	return err
}

// GetMonitorTransfers idmonitorTransfers表
// Author [yourname](https://github.com/yourname)
func (monitorTransfersService *MonitorTransfersService) GetMonitorTransfers(ctx context.Context, id string) (monitorTransfers api.MonitorTransfers, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&monitorTransfers).Error
	return
}

// GetMonitorTransfersInfoList monitorTransfers表
// Author [yourname](https://github.com/yourname)
func (monitorTransfersService *MonitorTransfersService) GetMonitorTransfersInfoList(ctx context.Context, info apiReq.MonitorTransfersSearch) (list []api.MonitorTransfers, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// db
	db := global.GVA_DB.Model(&api.MonitorTransfers{})
	var monitorTransferss []api.MonitorTransfers
	//

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["amount"] = true
	if orderMap[info.Sort] {
		OrderStr = info.Sort
		if info.Order == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&monitorTransferss).Error
	return monitorTransferss, total, err
}
func (monitorTransfersService *MonitorTransfersService) GetMonitorTransfersPublic(ctx context.Context) {
	//
	//
}
func (monitorTransfersService *MonitorTransfersService) CreateAndCheckMonitorTransfers(monitorTransfers api.MonitorTransfers) (trans api.MonitorTransfers, err error) {

	var transfer api.MonitorTransfers
	if !errors.Is(global.GVA_DB.Where("tx_hash = ?", monitorTransfers.TxHash).First(&transfer).Error, gorm.ErrRecordNotFound) { //
		return monitorTransfers, errors.New("")
	}

	err = global.GVA_DB.Create(monitorTransfers).Error
	return monitorTransfers, err
}
