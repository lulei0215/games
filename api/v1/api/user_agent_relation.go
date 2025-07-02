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

type UserAgentRelationApi struct{}

// CreateUserAgentRelation userAgentRelation表
// @Tags UserAgentRelation
// @Summary userAgentRelation表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body api.UserAgentRelation true "userAgentRelation表"
// @Success 200 {object} response.Response{msg=string} ""
// @Router /userAgentRelation/createUserAgentRelation [post]
func (userAgentRelationApi *UserAgentRelationApi) CreateUserAgentRelation(c *gin.Context) {
	// Context
	ctx := c.Request.Context()

	var userAgentRelation api.UserAgentRelation
	err := c.ShouldBindJSON(&userAgentRelation)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = userAgentRelationService.CreateUserAgentRelation(ctx, &userAgentRelation)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage(":"+err.Error(), c)
		return
	}
	response.OkWithMessage("", c)
}

// DeleteUserAgentRelation userAgentRelation表
// @Tags UserAgentRelation
// @Summary userAgentRelation表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body api.UserAgentRelation true "userAgentRelation表"
// @Success 200 {object} response.Response{msg=string} ""
// @Router /userAgentRelation/deleteUserAgentRelation [delete]
func (userAgentRelationApi *UserAgentRelationApi) DeleteUserAgentRelation(c *gin.Context) {
	// Context
	ctx := c.Request.Context()

	userId := c.Query("userId")
	err := userAgentRelationService.DeleteUserAgentRelation(ctx, userId)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage(":"+err.Error(), c)
		return
	}
	response.OkWithMessage("", c)
}

// DeleteUserAgentRelationByIds userAgentRelation表
// @Tags UserAgentRelation
// @Summary userAgentRelation表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} ""
// @Router /userAgentRelation/deleteUserAgentRelationByIds [delete]
func (userAgentRelationApi *UserAgentRelationApi) DeleteUserAgentRelationByIds(c *gin.Context) {
	// Context
	ctx := c.Request.Context()

	userIds := c.QueryArray("userIds[]")
	err := userAgentRelationService.DeleteUserAgentRelationByIds(ctx, userIds)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage(":"+err.Error(), c)
		return
	}
	response.OkWithMessage("", c)
}

// UpdateUserAgentRelation userAgentRelation表
// @Tags UserAgentRelation
// @Summary userAgentRelation表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body api.UserAgentRelation true "userAgentRelation表"
// @Success 200 {object} response.Response{msg=string} ""
// @Router /userAgentRelation/updateUserAgentRelation [put]
func (userAgentRelationApi *UserAgentRelationApi) UpdateUserAgentRelation(c *gin.Context) {
	// ctxcontext
	ctx := c.Request.Context()

	var userAgentRelation api.UserAgentRelation
	err := c.ShouldBindJSON(&userAgentRelation)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = userAgentRelationService.UpdateUserAgentRelation(ctx, userAgentRelation)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage(":"+err.Error(), c)
		return
	}
	response.OkWithMessage("", c)
}

// FindUserAgentRelation iduserAgentRelation表
// @Tags UserAgentRelation
// @Summary iduserAgentRelation表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param userId query int true "iduserAgentRelation表"
// @Success 200 {object} response.Response{data=api.UserAgentRelation,msg=string} ""
// @Router /userAgentRelation/findUserAgentRelation [get]
func (userAgentRelationApi *UserAgentRelationApi) FindUserAgentRelation(c *gin.Context) {
	// Context
	ctx := c.Request.Context()

	userId := c.Query("userId")
	reuserAgentRelation, err := userAgentRelationService.GetUserAgentRelation(ctx, userId)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage(":"+err.Error(), c)
		return
	}
	response.OkWithData(reuserAgentRelation, c)
}

// GetUserAgentRelationList userAgentRelation表
// @Tags UserAgentRelation
// @Summary userAgentRelation表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query apiReq.UserAgentRelationSearch true "userAgentRelation表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} ""
// @Router /userAgentRelation/getUserAgentRelationList [get]
func (userAgentRelationApi *UserAgentRelationApi) GetUserAgentRelationList(c *gin.Context) {
	// Context
	ctx := c.Request.Context()

	var pageInfo apiReq.UserAgentRelationSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := userAgentRelationService.GetUserAgentRelationInfoList(ctx, pageInfo)
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

// GetUserAgentRelationPublic userAgentRelation表
// @Tags UserAgentRelation
// @Summary userAgentRelation表
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} ""
// @Router /userAgentRelation/getUserAgentRelationPublic [get]
func (userAgentRelationApi *UserAgentRelationApi) GetUserAgentRelationPublic(c *gin.Context) {
	// Context
	ctx := c.Request.Context()

	//
	// ，C，
	userAgentRelationService.GetUserAgentRelationPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "userAgentRelation表",
	}, "", c)
}
func (userAgentRelationApi *UserAgentRelationApi) GetList(c *gin.Context) {
	uid := utils.GetRedisUserID(c)
	if uid == 0 {
		response.Result(401, nil, "", c)
		return
	}
	// Context
	ctx := c.Request.Context()

	var pageInfo apiReq.UserAgentRelationSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := userAgentRelationService.GetList(ctx, pageInfo, uid)

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
