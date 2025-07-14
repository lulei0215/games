
package api

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/api"
    apiReq "github.com/flipped-aurora/gin-vue-admin/server/model/api/request"
)

type UserBetRecordService struct {}
// CreateUserBetRecord userBetRecord表
// Author [yourname](https://github.com/yourname)
func (userBetRecordService *UserBetRecordService) CreateUserBetRecord(ctx context.Context, userBetRecord *api.UserBetRecord) (err error) {
	err = global.GVA_DB.Create(userBetRecord).Error
	return err
}

// DeleteUserBetRecord userBetRecord表
// Author [yourname](https://github.com/yourname)
func (userBetRecordService *UserBetRecordService)DeleteUserBetRecord(ctx context.Context, id string) (err error) {
	err = global.GVA_DB.Delete(&api.UserBetRecord{},"id = ?",id).Error
	return err
}

// DeleteUserBetRecordByIds userBetRecord表
// Author [yourname](https://github.com/yourname)
func (userBetRecordService *UserBetRecordService)DeleteUserBetRecordByIds(ctx context.Context, ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]api.UserBetRecord{},"id in ?",ids).Error
	return err
}

// UpdateUserBetRecord userBetRecord表
// Author [yourname](https://github.com/yourname)
func (userBetRecordService *UserBetRecordService)UpdateUserBetRecord(ctx context.Context, userBetRecord api.UserBetRecord) (err error) {
	err = global.GVA_DB.Model(&api.UserBetRecord{}).Where("id = ?",userBetRecord.Id).Updates(&userBetRecord).Error
	return err
}

// GetUserBetRecord iduserBetRecord表
// Author [yourname](https://github.com/yourname)
func (userBetRecordService *UserBetRecordService)GetUserBetRecord(ctx context.Context, id string) (userBetRecord api.UserBetRecord, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&userBetRecord).Error
	return
}
// GetUserBetRecordInfoList userBetRecord表
// Author [yourname](https://github.com/yourname)
func (userBetRecordService *UserBetRecordService)GetUserBetRecordInfoList(ctx context.Context, info apiReq.UserBetRecordSearch) (list []api.UserBetRecord, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // db
	db := global.GVA_DB.Model(&api.UserBetRecord{})
    var userBetRecords []api.UserBetRecord
    //
    
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&userBetRecords).Error
	return  userBetRecords, total, err
}
func (userBetRecordService *UserBetRecordService)GetUserBetRecordPublic(ctx context.Context) {
    // 
    // 
}
