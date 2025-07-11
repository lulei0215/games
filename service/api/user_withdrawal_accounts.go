package api

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/api"
	apiReq "github.com/flipped-aurora/gin-vue-admin/server/model/api/request"
	"go.uber.org/zap"
)

type UserWithdrawalAccountsService struct{}

// CreateUserWithdrawalAccounts userWithdrawalAccounts表
// Author [yourname](https://github.com/yourname)
func (userWithdrawalAccountsService *UserWithdrawalAccountsService) CreateUserWithdrawalAccounts(ctx context.Context, userWithdrawalAccounts *api.UserWithdrawalAccounts) (err error) {
	err = global.GVA_DB.Create(userWithdrawalAccounts).Error
	return err
}

// DeleteUserWithdrawalAccounts userWithdrawalAccounts表
// Author [yourname](https://github.com/yourname)
func (userWithdrawalAccountsService *UserWithdrawalAccountsService) DeleteUserWithdrawalAccounts(ctx context.Context, id string) (err error) {
	err = global.GVA_DB.Delete(&api.UserWithdrawalAccounts{}, "id = ?", id).Error
	return err
}

// DeleteUserWithdrawalAccountsByIds userWithdrawalAccounts表
// Author [yourname](https://github.com/yourname)
func (userWithdrawalAccountsService *UserWithdrawalAccountsService) DeleteUserWithdrawalAccountsByIds(ctx context.Context, ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]api.UserWithdrawalAccounts{}, "id in ?", ids).Error
	return err
}

// UpdateUserWithdrawalAccounts userWithdrawalAccounts表
// Author [yourname](https://github.com/yourname)
func (userWithdrawalAccountsService *UserWithdrawalAccountsService) UpdateUserWithdrawalAccounts(ctx context.Context, userWithdrawalAccounts api.UserWithdrawalAccounts) (err error) {
	err = global.GVA_DB.Model(&api.UserWithdrawalAccounts{}).Where("id = ?", userWithdrawalAccounts.Id).Updates(&userWithdrawalAccounts).Error
	return err
}

// GetUserWithdrawalAccounts iduserWithdrawalAccounts表
// Author [yourname](https://github.com/yourname)
func (userWithdrawalAccountsService *UserWithdrawalAccountsService) GetUserWithdrawalAccounts(ctx context.Context, id string) (userWithdrawalAccounts api.UserWithdrawalAccounts, err error) {
	global.GVA_LOG.Info("GetUserWithdrawalAccounts - Starting database query",
		zap.String("accountId", id))

	err = global.GVA_DB.Where("id = ?", id).First(&userWithdrawalAccounts).Error
	if err != nil {
		global.GVA_LOG.Error("GetUserWithdrawalAccounts - Database query failed",
			zap.Error(err),
			zap.String("accountId", id))
		return
	}

	global.GVA_LOG.Info("GetUserWithdrawalAccounts - Successfully retrieved account",
		zap.String("accountId", id),
		zap.Any("userWithdrawalAccounts", userWithdrawalAccounts))

	return
}

// GetUserWithdrawalAccountsInfoList userWithdrawalAccounts表
// Author [yourname](https://github.com/yourname)
func (userWithdrawalAccountsService *UserWithdrawalAccountsService) GetUserWithdrawalAccountsInfoList(ctx context.Context, info apiReq.UserWithdrawalAccountsSearch) (list []api.UserWithdrawalAccounts, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// db
	db := global.GVA_DB.Model(&api.UserWithdrawalAccounts{})
	var userWithdrawalAccountss []api.UserWithdrawalAccounts
	//

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&userWithdrawalAccountss).Error
	return userWithdrawalAccountss, total, err
}
func (userWithdrawalAccountsService *UserWithdrawalAccountsService) List(ctx context.Context, info apiReq.UserWithdrawalAccountsSearch, userId uint) (list []api.UserWithdrawalAccounts, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&api.UserWithdrawalAccounts{}).Where("user_id = ?", userId)
	var userWithdrawalAccountss []api.UserWithdrawalAccounts
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&userWithdrawalAccountss).Error
	return userWithdrawalAccountss, total, err
}

func (userWithdrawalAccountsService *UserWithdrawalAccountsService) Del(ctx context.Context, id int, userId uint) (err error) {
	err = global.GVA_DB.Delete(&[]api.UserWithdrawalAccounts{}, "id = ? and user_id = ?", id, userId).Error
	return err
}

func (userWithdrawalAccountsService *UserWithdrawalAccountsService) Update(ctx context.Context, userWithdrawalAccounts api.UserWithdrawalAccounts, userid uint) (err error) {
	err = global.GVA_DB.Model(&api.UserWithdrawalAccounts{}).Where("id = ? and user_id = ?", userWithdrawalAccounts.Id, userid).Updates(&userWithdrawalAccounts).Error
	return err
}
