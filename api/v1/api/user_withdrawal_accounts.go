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

type UserWithdrawalAccountsApi struct{}

// CreateUserWithdrawalAccounts userWithdrawalAccounts
// @Tags UserWithdrawalAccounts
// @Summary userWithdrawalAccounts
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body api.UserWithdrawalAccounts true "userWithdrawalAccounts"
// @Success 200 {object} response.Response{msg=string} ""
// @Router /userWithdrawalAccounts/createUserWithdrawalAccounts [post]
func (userWithdrawalAccountsApi *UserWithdrawalAccountsApi) CreateUserWithdrawalAccounts(c *gin.Context) {
	// Context
	ctx := c.Request.Context()

	var userWithdrawalAccounts api.UserWithdrawalAccounts
	err := c.ShouldBindJSON(&userWithdrawalAccounts)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = userWithdrawalAccountsService.CreateUserWithdrawalAccounts(ctx, &userWithdrawalAccounts)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage(":"+err.Error(), c)
		return
	}
	response.OkWithMessage("", c)
}

// DeleteUserWithdrawalAccounts userWithdrawalAccounts
// @Tags UserWithdrawalAccounts
// @Summary userWithdrawalAccounts
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body api.UserWithdrawalAccounts true "userWithdrawalAccounts"
// @Success 200 {object} response.Response{msg=string} ""
// @Router /userWithdrawalAccounts/deleteUserWithdrawalAccounts [delete]
func (userWithdrawalAccountsApi *UserWithdrawalAccountsApi) DeleteUserWithdrawalAccounts(c *gin.Context) {
	// Context
	ctx := c.Request.Context()

	id := c.Query("id")
	err := userWithdrawalAccountsService.DeleteUserWithdrawalAccounts(ctx, id)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage(":"+err.Error(), c)
		return
	}
	response.OkWithMessage("", c)
}

// DeleteUserWithdrawalAccountsByIds userWithdrawalAccounts
// @Tags UserWithdrawalAccounts
// @Summary userWithdrawalAccounts
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} ""
// @Router /userWithdrawalAccounts/deleteUserWithdrawalAccountsByIds [delete]
func (userWithdrawalAccountsApi *UserWithdrawalAccountsApi) DeleteUserWithdrawalAccountsByIds(c *gin.Context) {
	// Context
	ctx := c.Request.Context()

	ids := c.QueryArray("ids[]")
	err := userWithdrawalAccountsService.DeleteUserWithdrawalAccountsByIds(ctx, ids)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage(":"+err.Error(), c)
		return
	}
	response.OkWithMessage("", c)
}

// UpdateUserWithdrawalAccounts userWithdrawalAccounts
// @Tags UserWithdrawalAccounts
// @Summary userWithdrawalAccounts
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body api.UserWithdrawalAccounts true "userWithdrawalAccounts"
// @Success 200 {object} response.Response{msg=string} ""
// @Router /userWithdrawalAccounts/updateUserWithdrawalAccounts [put]
func (userWithdrawalAccountsApi *UserWithdrawalAccountsApi) UpdateUserWithdrawalAccounts(c *gin.Context) {
	// ctxcontext
	ctx := c.Request.Context()

	var userWithdrawalAccounts api.UserWithdrawalAccounts
	err := c.ShouldBindJSON(&userWithdrawalAccounts)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = userWithdrawalAccountsService.UpdateUserWithdrawalAccounts(ctx, userWithdrawalAccounts)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage(":"+err.Error(), c)
		return
	}
	response.OkWithMessage("", c)
}

// FindUserWithdrawalAccounts iduserWithdrawalAccounts
// @Tags UserWithdrawalAccounts
// @Summary iduserWithdrawalAccounts
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id query int true "iduserWithdrawalAccounts"
// @Success 200 {object} response.Response{data=api.UserWithdrawalAccounts,msg=string} ""
// @Router /userWithdrawalAccounts/findUserWithdrawalAccounts [get]
func (userWithdrawalAccountsApi *UserWithdrawalAccountsApi) FindUserWithdrawalAccounts(c *gin.Context) {
	// Context
	ctx := c.Request.Context()

	id := c.Query("id")
	reuserWithdrawalAccounts, err := userWithdrawalAccountsService.GetUserWithdrawalAccounts(ctx, id)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage(":"+err.Error(), c)
		return
	}
	response.OkWithData(reuserWithdrawalAccounts, c)
}

// GetUserWithdrawalAccountsList userWithdrawalAccounts
// @Tags UserWithdrawalAccounts
// @Summary userWithdrawalAccounts
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query apiReq.UserWithdrawalAccountsSearch true "userWithdrawalAccounts"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} ""
// @Router /userWithdrawalAccounts/getUserWithdrawalAccountsList [get]
func (userWithdrawalAccountsApi *UserWithdrawalAccountsApi) GetUserWithdrawalAccountsList(c *gin.Context) {
	// Context
	ctx := c.Request.Context()

	var pageInfo apiReq.UserWithdrawalAccountsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := userWithdrawalAccountsService.GetUserWithdrawalAccountsInfoList(ctx, pageInfo)
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

func (userWithdrawalAccountsApi *UserWithdrawalAccountsApi) Add(c *gin.Context) {
	// Context
	uid := utils.GetRedisUserID(c)
	if uid == 0 {
		response.Result(401, nil, "user fail", c)
		return
	}

	var userWithdrawalAccounts api.UserWithdrawalAccounts
	err := c.ShouldBindJSON(&userWithdrawalAccounts)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userWithdrawalAccounts.UserId = uid
	userWithdrawalAccounts.IsDefault = 1
	userWithdrawalAccounts.Status = 1
	err = utils.Verify(userWithdrawalAccounts, utils.AddUserWithdrawalAccountsVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = userWithdrawalAccountsService.CreateUserWithdrawalAccounts(c, &userWithdrawalAccounts)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage(":"+err.Error(), c)
		return
	}
	response.OkWithMessage("ok", c)

}

func (userWithdrawalAccountsApi *UserWithdrawalAccountsApi) Del(c *gin.Context) {

	uid := utils.GetRedisUserID(c)
	if uid == 0 {
		response.Result(401, nil, "user fail", c)
		return
	}

	var r apiReq.SetDefaultAccountRequest
	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = userWithdrawalAccountsService.Del(c, r.Id, uid)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage(":"+err.Error(), c)
		return
	}
	response.OkWithMessage("ok", c)

}
func (userWithdrawalAccountsApi *UserWithdrawalAccountsApi) Upd(c *gin.Context) {
	uid := utils.GetRedisUserID(c)
	if uid == 0 {
		response.Result(401, nil, "user fail", c)
		return
	}
	ctx := c.Request.Context()

	var userWithdrawalAccounts api.UserWithdrawalAccounts
	err := c.ShouldBindJSON(&userWithdrawalAccounts)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = userWithdrawalAccountsService.Update(ctx, userWithdrawalAccounts, uid)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage(":"+err.Error(), c)
		return
	}
	response.OkWithMessage("ok", c)
}
func (userWithdrawalAccountsApi *UserWithdrawalAccountsApi) List(c *gin.Context) {
	uid := utils.GetRedisUserID(c)
	if uid == 0 {
		response.Result(401, nil, "user fail", c)
		return
	}
	// Context
	ctx := c.Request.Context()

	var pageInfo apiReq.UserWithdrawalAccountsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := userWithdrawalAccountsService.List(ctx, pageInfo, uid)
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
