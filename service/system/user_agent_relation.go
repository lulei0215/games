
package system

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
    systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
)

type UserAgentRelationService struct {}
// CreateUserAgentRelation userAgentRelation表
// Author [yourname](https://github.com/yourname)
func (userAgentRelationService *UserAgentRelationService) CreateUserAgentRelation(ctx context.Context, userAgentRelation *system.UserAgentRelation) (err error) {
	err = global.GVA_DB.Create(userAgentRelation).Error
	return err
}

// DeleteUserAgentRelation userAgentRelation表
// Author [yourname](https://github.com/yourname)
func (userAgentRelationService *UserAgentRelationService)DeleteUserAgentRelation(ctx context.Context, userId string) (err error) {
	err = global.GVA_DB.Delete(&system.UserAgentRelation{},"user_id = ?",userId).Error
	return err
}

// DeleteUserAgentRelationByIds userAgentRelation表
// Author [yourname](https://github.com/yourname)
func (userAgentRelationService *UserAgentRelationService)DeleteUserAgentRelationByIds(ctx context.Context, userIds []string) (err error) {
	err = global.GVA_DB.Delete(&[]system.UserAgentRelation{},"user_id in ?",userIds).Error
	return err
}

// UpdateUserAgentRelation userAgentRelation表
// Author [yourname](https://github.com/yourname)
func (userAgentRelationService *UserAgentRelationService)UpdateUserAgentRelation(ctx context.Context, userAgentRelation system.UserAgentRelation) (err error) {
	err = global.GVA_DB.Model(&system.UserAgentRelation{}).Where("user_id = ?",userAgentRelation.UserId).Updates(&userAgentRelation).Error
	return err
}

// GetUserAgentRelation userIduserAgentRelation表
// Author [yourname](https://github.com/yourname)
func (userAgentRelationService *UserAgentRelationService)GetUserAgentRelation(ctx context.Context, userId string) (userAgentRelation system.UserAgentRelation, err error) {
	err = global.GVA_DB.Where("user_id = ?", userId).First(&userAgentRelation).Error
	return
}
// GetUserAgentRelationInfoList userAgentRelation表
// Author [yourname](https://github.com/yourname)
func (userAgentRelationService *UserAgentRelationService)GetUserAgentRelationInfoList(ctx context.Context, info systemReq.UserAgentRelationSearch) (list []system.UserAgentRelation, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // db
	db := global.GVA_DB.Model(&system.UserAgentRelation{})
    var userAgentRelations []system.UserAgentRelation
    //
    
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&userAgentRelations).Error
	return  userAgentRelations, total, err
}
func (userAgentRelationService *UserAgentRelationService)GetUserAgentRelationPublic(ctx context.Context) {
    // 
    // 
}
