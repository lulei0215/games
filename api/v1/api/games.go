package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/api"
	apiReq "github.com/flipped-aurora/gin-vue-admin/server/model/api/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GamesApi struct{}

// CreateGames games表
// @Tags Games
// @Summary games表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body api.Games true "games表"
// @Success 200 {object} response.Response{msg=string} ""
// @Router /games/createGames [post]
func (gamesApi *GamesApi) CreateGames(c *gin.Context) {
	// Context
	ctx := c.Request.Context()

	var games api.Games
	err := c.ShouldBindJSON(&games)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = gamesService.CreateGames(ctx, &games)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage(":"+err.Error(), c)
		return
	}
	response.OkWithMessage("", c)
}

// DeleteGames games表
// @Tags Games
// @Summary games表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body api.Games true "games表"
// @Success 200 {object} response.Response{msg=string} ""
// @Router /games/deleteGames [delete]
func (gamesApi *GamesApi) DeleteGames(c *gin.Context) {
	// Context
	ctx := c.Request.Context()

	id := c.Query("id")
	err := gamesService.DeleteGames(ctx, id)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage(":"+err.Error(), c)
		return
	}
	response.OkWithMessage("", c)
}

// DeleteGamesByIds games表
// @Tags Games
// @Summary games表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} ""
// @Router /games/deleteGamesByIds [delete]
func (gamesApi *GamesApi) DeleteGamesByIds(c *gin.Context) {
	// Context
	ctx := c.Request.Context()

	ids := c.QueryArray("ids[]")
	err := gamesService.DeleteGamesByIds(ctx, ids)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage(":"+err.Error(), c)
		return
	}
	response.OkWithMessage("", c)
}

// UpdateGames games表
// @Tags Games
// @Summary games表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body api.Games true "games表"
// @Success 200 {object} response.Response{msg=string} ""
// @Router /games/updateGames [put]
func (gamesApi *GamesApi) UpdateGames(c *gin.Context) {
	// ctxcontext
	ctx := c.Request.Context()

	var games api.Games
	err := c.ShouldBindJSON(&games)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = gamesService.UpdateGames(ctx, games)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage(":"+err.Error(), c)
		return
	}
	response.OkWithMessage("", c)
}

// FindGames idgames表
// @Tags Games
// @Summary idgames表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id query int true "idgames表"
// @Success 200 {object} response.Response{data=api.Games,msg=string} ""
// @Router /games/findGames [get]
func (gamesApi *GamesApi) FindGames(c *gin.Context) {
	// Context
	ctx := c.Request.Context()

	id := c.Query("id")
	regames, err := gamesService.GetGames(ctx, id)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage(":"+err.Error(), c)
		return
	}
	response.OkWithData(regames, c)
}

// GetGamesList games表
// @Tags Games
// @Summary games表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query apiReq.GamesSearch true "games表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} ""
// @Router /games/getGamesList [get]
func (gamesApi *GamesApi) GetGamesList(c *gin.Context) {
	// Context
	ctx := c.Request.Context()

	var pageInfo apiReq.GamesSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := gamesService.GetGamesInfoList(ctx, pageInfo)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage(":"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "", c)
}

// GetGamesPublic games表
// @Tags Games
// @Summary games表
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} ""
// @Router /games/getGamesPublic [get]
func (gamesApi *GamesApi) GetGamesPublic(c *gin.Context) {
	// Context
	ctx := c.Request.Context()

	//
	// ，C，
	gamesService.GetGamesPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "games表",
	}, "", c)
}

func (gamesApi *GamesApi) GetList(c *gin.Context) {
	// Context
	ctx := c.Request.Context()

	var pageInfo apiReq.GamesSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := gamesService.GetGamesInfoList(ctx, pageInfo)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage(":"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "", c)
}
