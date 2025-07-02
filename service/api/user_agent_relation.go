package api

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/api"
	apiReq "github.com/flipped-aurora/gin-vue-admin/server/model/api/request"
)

type UserAgentRelationService struct{}

// CreateUserAgentRelation userAgentRelation表
// Author [yourname](https://github.com/yourname)
func (userAgentRelationService *UserAgentRelationService) CreateUserAgentRelation(ctx context.Context, userAgentRelation *api.UserAgentRelation) (err error) {
	err = global.GVA_DB.Create(userAgentRelation).Error
	return err
}

// DeleteUserAgentRelation userAgentRelation表
// Author [yourname](https://github.com/yourname)
func (userAgentRelationService *UserAgentRelationService) DeleteUserAgentRelation(ctx context.Context, userId string) (err error) {
	err = global.GVA_DB.Delete(&api.UserAgentRelation{}, "user_id = ?", userId).Error
	return err
}

// DeleteUserAgentRelationByIds userAgentRelation表
// Author [yourname](https://github.com/yourname)
func (userAgentRelationService *UserAgentRelationService) DeleteUserAgentRelationByIds(ctx context.Context, userIds []string) (err error) {
	err = global.GVA_DB.Delete(&[]api.UserAgentRelation{}, "user_id in ?", userIds).Error
	return err
}

// UpdateUserAgentRelation userAgentRelation表
// Author [yourname](https://github.com/yourname)
func (userAgentRelationService *UserAgentRelationService) UpdateUserAgentRelation(ctx context.Context, userAgentRelation api.UserAgentRelation) (err error) {
	err = global.GVA_DB.Model(&api.UserAgentRelation{}).Where("user_id = ?", userAgentRelation.UserId).Updates(&userAgentRelation).Error
	return err
}

// GetUserAgentRelation userIduserAgentRelation表
// Author [yourname](https://github.com/yourname)
func (userAgentRelationService *UserAgentRelationService) GetUserAgentRelation(ctx context.Context, userId string) (userAgentRelation api.UserAgentRelation, err error) {
	err = global.GVA_DB.Where("user_id = ?", userId).First(&userAgentRelation).Error
	return
}

// GetUserAgentRelationInfoList userAgentRelation表
// Author [yourname](https://github.com/yourname)
func (userAgentRelationService *UserAgentRelationService) GetUserAgentRelationInfoList(ctx context.Context, info apiReq.UserAgentRelationSearch) (list []api.UserAgentRelation, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// db
	db := global.GVA_DB.Model(&api.UserAgentRelation{})
	var userAgentRelations []api.UserAgentRelation
	//

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&userAgentRelations).Error
	return userAgentRelations, total, err
}
func (userAgentRelationService *UserAgentRelationService) GetUserAgentRelationPublic(ctx context.Context) {
	//
	//
}

func (userAgentRelationService *UserAgentRelationService) GetList(ctx context.Context, info apiReq.UserAgentRelationSearch, uid uint) (list []api.UserAgentRelationWithUser, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// db
	db := global.GVA_DB.Model(&api.UserAgentRelationWithUser{})
	var userAgentRelations []api.UserAgentRelationWithUser
	//

	// 联表查询，关联 sys_users 表
	db = db.Select("user_agent_relation.*, sys_users.username, sys_users.nick_name, sys_users.phone, sys_users.email, sys_users.balance, sys_users.vip_level, sys_users.enable").
		Joins("LEFT JOIN sys_users ON user_agent_relation.user_id = sys_users.id")

	err = db.Where("user_agent_relation.parent_id_1 = ?", uid).Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Where("user_agent_relation.parent_id_1 = ?", uid).Find(&userAgentRelations).Error
	return userAgentRelations, total, err
}
