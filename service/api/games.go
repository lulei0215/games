package api

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/api"
	apiReq "github.com/flipped-aurora/gin-vue-admin/server/model/api/request"
)

type GamesService struct{}

// CreateGames games表
// Author [yourname](https://github.com/yourname)
func (gamesService *GamesService) CreateGames(ctx context.Context, games *api.Games) (err error) {
	err = global.GVA_DB.Create(games).Error
	return err
}

// DeleteGames games表
// Author [yourname](https://github.com/yourname)
func (gamesService *GamesService) DeleteGames(ctx context.Context, id string) (err error) {
	err = global.GVA_DB.Delete(&api.Games{}, "id = ?", id).Error
	return err
}

// DeleteGamesByIds games表
// Author [yourname](https://github.com/yourname)
func (gamesService *GamesService) DeleteGamesByIds(ctx context.Context, ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]api.Games{}, "id in ?", ids).Error
	return err
}

// UpdateGames games表
// Author [yourname](https://github.com/yourname)
func (gamesService *GamesService) UpdateGames(ctx context.Context, games api.Games) (err error) {
	err = global.GVA_DB.Model(&api.Games{}).Where("id = ?", games.Id).Updates(&games).Error
	return err
}

// GetGames idgames表
// Author [yourname](https://github.com/yourname)
func (gamesService *GamesService) GetGames(ctx context.Context, id string) (games api.Games, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&games).Error
	return
}

// GetGamesInfoList games表
// Author [yourname](https://github.com/yourname)
func (gamesService *GamesService) GetGamesInfoList(ctx context.Context, info apiReq.GamesSearch) (list []api.Games, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// db
	db := global.GVA_DB.Model(&api.Games{})
	var gamess []api.Games
	//

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&gamess).Error
	return gamess, total, err
}
func (gamesService *GamesService) GetGamesPublic(ctx context.Context) {
	//
	//
}
