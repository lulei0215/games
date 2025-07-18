package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/api"
	apiReq "github.com/flipped-aurora/gin-vue-admin/server/model/api/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserBetRecordApi struct{}

// CreateUserBetRecord userBetRecord表
// @Tags UserBetRecord
// @Summary userBetRecord表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body api.UserBetRecord true "userBetRecord表"
// @Success 200 {object} response.Response{msg=string} ""
// @Router /userBetRecord/createUserBetRecord [post]
func (userBetRecordApi *UserBetRecordApi) CreateUserBetRecord(c *gin.Context) {
	// Context
	ctx := c.Request.Context()

	var userBetRecord api.UserBetRecord
	err := c.ShouldBindJSON(&userBetRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = userBetRecordService.CreateUserBetRecord(ctx, &userBetRecord)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage(":"+err.Error(), c)
		return
	}
	response.OkWithMessage("", c)
}

// DeleteUserBetRecord userBetRecord表
// @Tags UserBetRecord
// @Summary userBetRecord表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body api.UserBetRecord true "userBetRecord表"
// @Success 200 {object} response.Response{msg=string} ""
// @Router /userBetRecord/deleteUserBetRecord [delete]
func (userBetRecordApi *UserBetRecordApi) DeleteUserBetRecord(c *gin.Context) {
	// Context
	ctx := c.Request.Context()

	id := c.Query("id")
	err := userBetRecordService.DeleteUserBetRecord(ctx, id)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage(":"+err.Error(), c)
		return
	}
	response.OkWithMessage("", c)
}

// DeleteUserBetRecordByIds userBetRecord表
// @Tags UserBetRecord
// @Summary userBetRecord表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} ""
// @Router /userBetRecord/deleteUserBetRecordByIds [delete]
func (userBetRecordApi *UserBetRecordApi) DeleteUserBetRecordByIds(c *gin.Context) {
	// Context
	ctx := c.Request.Context()

	ids := c.QueryArray("ids[]")
	err := userBetRecordService.DeleteUserBetRecordByIds(ctx, ids)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage(":"+err.Error(), c)
		return
	}
	response.OkWithMessage("", c)
}

// UpdateUserBetRecord userBetRecord表
// @Tags UserBetRecord
// @Summary userBetRecord表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body api.UserBetRecord true "userBetRecord表"
// @Success 200 {object} response.Response{msg=string} ""
// @Router /userBetRecord/updateUserBetRecord [put]
func (userBetRecordApi *UserBetRecordApi) UpdateUserBetRecord(c *gin.Context) {
	// ctxcontext
	ctx := c.Request.Context()

	var userBetRecord api.UserBetRecord
	err := c.ShouldBindJSON(&userBetRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = userBetRecordService.UpdateUserBetRecord(ctx, userBetRecord)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage(":"+err.Error(), c)
		return
	}
	response.OkWithMessage("", c)
}

// FindUserBetRecord iduserBetRecord表
// @Tags UserBetRecord
// @Summary iduserBetRecord表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id query int true "iduserBetRecord表"
// @Success 200 {object} response.Response{data=api.UserBetRecord,msg=string} ""
// @Router /userBetRecord/findUserBetRecord [get]
func (userBetRecordApi *UserBetRecordApi) FindUserBetRecord(c *gin.Context) {
	// Context
	ctx := c.Request.Context()

	id := c.Query("id")
	reuserBetRecord, err := userBetRecordService.GetUserBetRecord(ctx, id)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage(":"+err.Error(), c)
		return
	}
	response.OkWithData(reuserBetRecord, c)
}

// GetUserBetRecordList userBetRecord表
// @Tags UserBetRecord
// @Summary userBetRecord表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query apiReq.UserBetRecordSearch true "userBetRecord表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} ""
// @Router /userBetRecord/getUserBetRecordList [get]
func (userBetRecordApi *UserBetRecordApi) GetUserBetRecordList(c *gin.Context) {
	// Context
	ctx := c.Request.Context()

	var pageInfo apiReq.UserBetRecordSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := userBetRecordService.GetUserBetRecordInfoList(ctx, pageInfo)
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
func (userBetRecordApi *UserBetRecordApi) GetMyBetRecordList(c *gin.Context) {

	uid := utils.GetRedisUserID(c)
	if uid == 0 {
		response.Result(401, nil, "", c)
		return
	}
	// Context
	ctx := c.Request.Context()

	var pageInfo apiReq.UserBetRecordSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := userBetRecordService.GetMyUserBetRecordInfoList(ctx, pageInfo, uid)
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

// GetUserBetRecordPublic userBetRecord表
// @Tags UserBetRecord
// @Summary userBetRecord表
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} ""
// @Router /userBetRecord/getUserBetRecordPublic [get]
func (userBetRecordApi *UserBetRecordApi) GetUserBetRecordPublic(c *gin.Context) {
	// Context
	ctx := c.Request.Context()

	//
	// ，C，
	userBetRecordService.GetUserBetRecordPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "userBetRecord表",
	}, "", c)
}
