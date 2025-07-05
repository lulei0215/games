package api

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/api"
	apiReq "github.com/flipped-aurora/gin-vue-admin/server/model/api/request"
)

type UserRebatesService struct{}

// CreateUserRebates userRebates表
// Author [yourname](https://github.com/yourname)
func (userRebatesService *UserRebatesService) CreateUserRebates(ctx context.Context, userRebates *api.UserRebates) (err error) {
	err = global.GVA_DB.Create(userRebates).Error
	return err
}

// DeleteUserRebates userRebates表
// Author [yourname](https://github.com/yourname)
func (userRebatesService *UserRebatesService) DeleteUserRebates(ctx context.Context, id string) (err error) {
	err = global.GVA_DB.Delete(&api.UserRebates{}, "id = ?", id).Error
	return err
}

// DeleteUserRebatesByIds userRebates表
// Author [yourname](https://github.com/yourname)
func (userRebatesService *UserRebatesService) DeleteUserRebatesByIds(ctx context.Context, ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]api.UserRebates{}, "id in ?", ids).Error
	return err
}

// UpdateUserRebates userRebates表
// Author [yourname](https://github.com/yourname)
func (userRebatesService *UserRebatesService) UpdateUserRebates(ctx context.Context, userRebates api.UserRebates) (err error) {
	err = global.GVA_DB.Model(&api.UserRebates{}).Where("id = ?", userRebates.Id).Updates(&userRebates).Error
	return err
}

// GetUserRebates iduserRebates表
// Author [yourname](https://github.com/yourname)
func (userRebatesService *UserRebatesService) GetUserRebates(ctx context.Context, id string) (userRebates api.UserRebates, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&userRebates).Error
	return
}

// GetUserRebatesInfoList userRebates表
// Author [yourname](https://github.com/yourname)
func (userRebatesService *UserRebatesService) GetUserRebatesInfoList(ctx context.Context, info apiReq.UserRebatesSearch) (list []api.UserRebates, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// db
	db := global.GVA_DB.Model(&api.UserRebates{})
	var userRebatess []api.UserRebates

	// Add query conditions
	if info.Id != nil {
		db = db.Where("id = ?", *info.Id)
	}
	if info.UserId != nil {
		db = db.Where("user_id = ?", *info.UserId)
	}
	if info.FromUserId != nil {
		db = db.Where("from_user_id = ?", *info.FromUserId)
	}
	if info.FromUserCode != "" {
		db = db.Where("from_user_code LIKE ?", "%"+info.FromUserCode+"%")
	}
	if info.RebateType != "" {
		db = db.Where("rebate_type LIKE ?", "%"+info.RebateType+"%")
	}
	if info.RebateLevel != nil {
		db = db.Where("rebate_level = ?", *info.RebateLevel)
	}
	if info.Coin != nil {
		db = db.Where("coin = ?", *info.Coin)
	}
	if info.Win != nil {
		db = db.Where("win = ?", *info.Win)
	}
	if info.RebateRate != nil {
		db = db.Where("rebate_rate = ?", *info.RebateRate)
	}
	if info.RebateAmount != nil {
		db = db.Where("rebate_amount = ?", *info.RebateAmount)
	}
	if info.UserBalanceBefore != nil {
		db = db.Where("user_balance_before = ?", *info.UserBalanceBefore)
	}
	if info.UserBalanceAfter != nil {
		db = db.Where("user_balance_after = ?", *info.UserBalanceAfter)
	}
	if info.SessionId != "" {
		db = db.Where("session_id LIKE ?", "%"+info.SessionId+"%")
	}
	if info.GameType != nil {
		db = db.Where("game_type = ?", *info.GameType)
	}
	if info.Area != "" {
		db = db.Where("area LIKE ?", "%"+info.Area+"%")
	}
	if info.BetInfo != "" {
		db = db.Where("bet_info LIKE ?", "%"+info.BetInfo+"%")
	}
	if info.Status != nil {
		db = db.Where("status = ?", *info.Status)
	}
	if info.Remark != "" {
		db = db.Where("remark LIKE ?", "%"+info.Remark+"%")
	}
	if info.CreatedAtStart != nil {
		db = db.Where("created_at >= ?", *info.CreatedAtStart)
	}
	if info.CreatedAtEnd != nil {
		db = db.Where("created_at <= ?", *info.CreatedAtEnd)
	}
	if info.UpdatedAtStart != nil {
		db = db.Where("updated_at >= ?", *info.UpdatedAtStart)
	}
	if info.UpdatedAtEnd != nil {
		db = db.Where("updated_at <= ?", *info.UpdatedAtEnd)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&userRebatess).Error
	return userRebatess, total, err
}
func (userRebatesService *UserRebatesService) GetUserRebatesPublic(ctx context.Context) {
	//
	//
}
