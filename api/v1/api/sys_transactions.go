package api

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/api"
	apiReq "github.com/flipped-aurora/gin-vue-admin/server/model/api/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SysTransactionsApi struct{}

// CreateSysTransactions 创建sysTransactions表
// @Tags SysTransactions
// @Summary 创建sysTransactions表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body api.SysTransactions true "创建sysTransactions表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /sysTransactions/createSysTransactions [post]
func (sysTransactionsApi *SysTransactionsApi) CreateSysTransactions(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var sysTransactions api.SysTransactions
	err := c.ShouldBindJSON(&sysTransactions)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = sysTransactionsService.CreateSysTransactions(ctx, &sysTransactions)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteSysTransactions 删除sysTransactions表
// @Tags SysTransactions
// @Summary 删除sysTransactions表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body api.SysTransactions true "删除sysTransactions表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /sysTransactions/deleteSysTransactions [delete]
func (sysTransactionsApi *SysTransactionsApi) DeleteSysTransactions(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	id := c.Query("id")
	err := sysTransactionsService.DeleteSysTransactions(ctx, id)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteSysTransactionsByIds 批量删除sysTransactions表
// @Tags SysTransactions
// @Summary 批量删除sysTransactions表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /sysTransactions/deleteSysTransactionsByIds [delete]
func (sysTransactionsApi *SysTransactionsApi) DeleteSysTransactionsByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ids := c.QueryArray("ids[]")
	err := sysTransactionsService.DeleteSysTransactionsByIds(ctx, ids)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateSysTransactions 更新sysTransactions表
// @Tags SysTransactions
// @Summary 更新sysTransactions表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body api.SysTransactions true "更新sysTransactions表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /sysTransactions/updateSysTransactions [put]
func (sysTransactionsApi *SysTransactionsApi) UpdateSysTransactions(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var sysTransactions api.SysTransactions
	err := c.ShouldBindJSON(&sysTransactions)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = sysTransactionsService.UpdateSysTransactions(ctx, sysTransactions)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindSysTransactions 用id查询sysTransactions表
// @Tags SysTransactions
// @Summary 用id查询sysTransactions表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id query int true "用id查询sysTransactions表"
// @Success 200 {object} response.Response{data=api.SysTransactions,msg=string} "查询成功"
// @Router /sysTransactions/findSysTransactions [get]
func (sysTransactionsApi *SysTransactionsApi) FindSysTransactions(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	id := c.Query("id")
	resysTransactions, err := sysTransactionsService.GetSysTransactions(ctx, id)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(resysTransactions, c)
}

// GetSysTransactionsList 分页获取sysTransactions表列表
// @Tags SysTransactions
// @Summary 分页获取sysTransactions表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query apiReq.SysTransactionsSearch true "分页获取sysTransactions表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /sysTransactions/getSysTransactionsList [get]
func (sysTransactionsApi *SysTransactionsApi) GetSysTransactionsList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo apiReq.SysTransactionsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := sysTransactionsService.GetSysTransactionsInfoList(ctx, pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetSysTransactionsPublic 不需要鉴权的sysTransactions表接口
// @Tags SysTransactions
// @Summary 不需要鉴权的sysTransactions表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /sysTransactions/getSysTransactionsPublic [get]
func (sysTransactionsApi *SysTransactionsApi) GetSysTransactionsPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	sysTransactionsService.GetSysTransactionsPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的sysTransactions表接口信息",
	}, "获取成功", c)
}

func (sysTransactionsApi *SysTransactionsApi) Get(c *gin.Context) {
	uid := utils.GetRedisUserID(c)
	if uid == 0 {
		response.FailWithMessage("用户ID不存在", c)
		return
	}
	fmt.Println("uid", uid)
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	sysTransactionsService.GetSysTransactionsPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的sysTransactions表接口信息",
	}, "获取成功", c)
}
