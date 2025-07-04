package api

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/api"
    apiReq "github.com/flipped-aurora/gin-vue-admin/server/model/api/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type UserRebatesApi struct {}



// CreateUserRebates userRebates表
// @Tags UserRebates
// @Summary userRebates表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body api.UserRebates true "userRebates表"
// @Success 200 {object} response.Response{msg=string} ""
// @Router /userRebates/createUserRebates [post]
func (userRebatesApi *UserRebatesApi) CreateUserRebates(c *gin.Context) {
    // Context
    ctx := c.Request.Context()

	var userRebates api.UserRebates
	err := c.ShouldBindJSON(&userRebates)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = userRebatesService.CreateUserRebates(ctx,&userRebates)
	if err != nil {
        global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage(":" + err.Error(), c)
		return
	}
    response.OkWithMessage("", c)
}

// DeleteUserRebates userRebates表
// @Tags UserRebates
// @Summary userRebates表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body api.UserRebates true "userRebates表"
// @Success 200 {object} response.Response{msg=string} ""
// @Router /userRebates/deleteUserRebates [delete]
func (userRebatesApi *UserRebatesApi) DeleteUserRebates(c *gin.Context) {
    // Context
    ctx := c.Request.Context()

	id := c.Query("id")
	err := userRebatesService.DeleteUserRebates(ctx,id)
	if err != nil {
        global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage(":" + err.Error(), c)
		return
	}
	response.OkWithMessage("", c)
}

// DeleteUserRebatesByIds userRebates表
// @Tags UserRebates
// @Summary userRebates表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} ""
// @Router /userRebates/deleteUserRebatesByIds [delete]
func (userRebatesApi *UserRebatesApi) DeleteUserRebatesByIds(c *gin.Context) {
    // Context
    ctx := c.Request.Context()

	ids := c.QueryArray("ids[]")
	err := userRebatesService.DeleteUserRebatesByIds(ctx,ids)
	if err != nil {
        global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage(":" + err.Error(), c)
		return
	}
	response.OkWithMessage("", c)
}

// UpdateUserRebates userRebates表
// @Tags UserRebates
// @Summary userRebates表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body api.UserRebates true "userRebates表"
// @Success 200 {object} response.Response{msg=string} ""
// @Router /userRebates/updateUserRebates [put]
func (userRebatesApi *UserRebatesApi) UpdateUserRebates(c *gin.Context) {
    // ctxcontext
    ctx := c.Request.Context()

	var userRebates api.UserRebates
	err := c.ShouldBindJSON(&userRebates)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = userRebatesService.UpdateUserRebates(ctx,userRebates)
	if err != nil {
        global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage(":" + err.Error(), c)
		return
	}
	response.OkWithMessage("", c)
}

// FindUserRebates iduserRebates表
// @Tags UserRebates
// @Summary iduserRebates表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id query int true "iduserRebates表"
// @Success 200 {object} response.Response{data=api.UserRebates,msg=string} ""
// @Router /userRebates/findUserRebates [get]
func (userRebatesApi *UserRebatesApi) FindUserRebates(c *gin.Context) {
    // Context
    ctx := c.Request.Context()

	id := c.Query("id")
	reuserRebates, err := userRebatesService.GetUserRebates(ctx,id)
	if err != nil {
        global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage(":" + err.Error(), c)
		return
	}
	response.OkWithData(reuserRebates, c)
}
// GetUserRebatesList userRebates表
// @Tags UserRebates
// @Summary userRebates表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query apiReq.UserRebatesSearch true "userRebates表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} ""
// @Router /userRebates/getUserRebatesList [get]
func (userRebatesApi *UserRebatesApi) GetUserRebatesList(c *gin.Context) {
    // Context
    ctx := c.Request.Context()

	var pageInfo apiReq.UserRebatesSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := userRebatesService.GetUserRebatesInfoList(ctx,pageInfo)
	if err != nil {
	    global.GVA_LOG.Error("!", zap.Error(err))
        response.FailWithMessage(":" + err.Error(), c)
        return
    }
    response.OkWithDetailed(response.PageResult{
        List:     list,
        Total:    total,
        Page:     pageInfo.Page,
        PageSize: pageInfo.PageSize,
    }, "", c)
}

// GetUserRebatesPublic userRebates表
// @Tags UserRebates
// @Summary userRebates表
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} ""
// @Router /userRebates/getUserRebatesPublic [get]
func (userRebatesApi *UserRebatesApi) GetUserRebatesPublic(c *gin.Context) {
    // Context
    ctx := c.Request.Context()

    // 
    // ，C，
    userRebatesService.GetUserRebatesPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "userRebates表",
    }, "", c)
}
