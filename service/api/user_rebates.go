
package api

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/api"
    apiReq "github.com/flipped-aurora/gin-vue-admin/server/model/api/request"
)

type UserRebatesService struct {}
// CreateUserRebates userRebates表
// Author [yourname](https://github.com/yourname)
func (userRebatesService *UserRebatesService) CreateUserRebates(ctx context.Context, userRebates *api.UserRebates) (err error) {
	err = global.GVA_DB.Create(userRebates).Error
	return err
}

// DeleteUserRebates userRebates表
// Author [yourname](https://github.com/yourname)
func (userRebatesService *UserRebatesService)DeleteUserRebates(ctx context.Context, id string) (err error) {
	err = global.GVA_DB.Delete(&api.UserRebates{},"id = ?",id).Error
	return err
}

// DeleteUserRebatesByIds userRebates表
// Author [yourname](https://github.com/yourname)
func (userRebatesService *UserRebatesService)DeleteUserRebatesByIds(ctx context.Context, ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]api.UserRebates{},"id in ?",ids).Error
	return err
}

// UpdateUserRebates userRebates表
// Author [yourname](https://github.com/yourname)
func (userRebatesService *UserRebatesService)UpdateUserRebates(ctx context.Context, userRebates api.UserRebates) (err error) {
	err = global.GVA_DB.Model(&api.UserRebates{}).Where("id = ?",userRebates.Id).Updates(&userRebates).Error
	return err
}

// GetUserRebates iduserRebates表
// Author [yourname](https://github.com/yourname)
func (userRebatesService *UserRebatesService)GetUserRebates(ctx context.Context, id string) (userRebates api.UserRebates, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&userRebates).Error
	return
}
// GetUserRebatesInfoList userRebates表
// Author [yourname](https://github.com/yourname)
func (userRebatesService *UserRebatesService)GetUserRebatesInfoList(ctx context.Context, info apiReq.UserRebatesSearch) (list []api.UserRebates, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // db
	db := global.GVA_DB.Model(&api.UserRebates{})
    var userRebatess []api.UserRebates
    //
    
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&userRebatess).Error
	return  userRebatess, total, err
}
func (userRebatesService *UserRebatesService)GetUserRebatesPublic(ctx context.Context) {
    // 
    // 
}
