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

// CreateSysTransactions sysTransactions
// @Tags SysTransactions
// @Summary sysTransactions
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body api.SysTransactions true "sysTransactions"
// @Success 200 {object} response.Response{msg=string} ""
// @Router /sysTransactions/createSysTransactions [post]
func (sysTransactionsApi *SysTransactionsApi) CreateSysTransactions(c *gin.Context) {
	// Context
	ctx := c.Request.Context()

	var sysTransactions api.SysTransactions
	err := c.ShouldBindJSON(&sysTransactions)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = sysTransactionsService.CreateSysTransactions(ctx, &sysTransactions)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage(":"+err.Error(), c)
		return
	}
	response.OkWithMessage("", c)
}

// DeleteSysTransactions sysTransactions
// @Tags SysTransactions
// @Summary sysTransactions
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body api.SysTransactions true "sysTransactions"
// @Success 200 {object} response.Response{msg=string} ""
// @Router /sysTransactions/deleteSysTransactions [delete]
func (sysTransactionsApi *SysTransactionsApi) DeleteSysTransactions(c *gin.Context) {
	// Context
	ctx := c.Request.Context()

	id := c.Query("id")
	err := sysTransactionsService.DeleteSysTransactions(ctx, id)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage(":"+err.Error(), c)
		return
	}
	response.OkWithMessage("", c)
}

// DeleteSysTransactionsByIds sysTransactions
// @Tags SysTransactions
// @Summary sysTransactions
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} ""
// @Router /sysTransactions/deleteSysTransactionsByIds [delete]
func (sysTransactionsApi *SysTransactionsApi) DeleteSysTransactionsByIds(c *gin.Context) {
	// Context
	ctx := c.Request.Context()

	ids := c.QueryArray("ids[]")
	err := sysTransactionsService.DeleteSysTransactionsByIds(ctx, ids)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage(":"+err.Error(), c)
		return
	}
	response.OkWithMessage("", c)
}

// UpdateSysTransactions sysTransactions
// @Tags SysTransactions
// @Summary sysTransactions
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body api.SysTransactions true "sysTransactions"
// @Success 200 {object} response.Response{msg=string} ""
// @Router /sysTransactions/updateSysTransactions [put]
func (sysTransactionsApi *SysTransactionsApi) UpdateSysTransactions(c *gin.Context) {
	// ctxcontext
	ctx := c.Request.Context()

	var sysTransactions api.SysTransactions
	err := c.ShouldBindJSON(&sysTransactions)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = sysTransactionsService.UpdateSysTransactions(ctx, sysTransactions)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage(":"+err.Error(), c)
		return
	}
	response.OkWithMessage("", c)
}

// FindSysTransactions idsysTransactions
// @Tags SysTransactions
// @Summary idsysTransactions
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id query int true "idsysTransactions"
// @Success 200 {object} response.Response{data=api.SysTransactions,msg=string} ""
// @Router /sysTransactions/findSysTransactions [get]
func (sysTransactionsApi *SysTransactionsApi) FindSysTransactions(c *gin.Context) {
	// Context
	ctx := c.Request.Context()

	id := c.Query("id")
	resysTransactions, err := sysTransactionsService.GetSysTransactions(ctx, id)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage(":"+err.Error(), c)
		return
	}
	response.OkWithData(resysTransactions, c)
}

// GetSysTransactionsList sysTransactions
// @Tags SysTransactions
// @Summary sysTransactions
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query apiReq.SysTransactionsSearch true "sysTransactions"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} ""
// @Router /sysTransactions/getSysTransactionsList [get]
func (sysTransactionsApi *SysTransactionsApi) GetSysTransactionsList(c *gin.Context) {
	// Context
	ctx := c.Request.Context()

	var pageInfo apiReq.SysTransactionsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := sysTransactionsService.GetSysTransactionsInfoList(ctx, pageInfo)
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

// GetSysTransactionsPublic sysTransactions
// @Tags SysTransactions
// @Summary sysTransactions
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} ""
// @Router /sysTransactions/getSysTransactionsPublic [get]
func (sysTransactionsApi *SysTransactionsApi) GetSysTransactionsPublic(c *gin.Context) {
	// Context
	ctx := c.Request.Context()

	//
	// ，C，
	sysTransactionsService.GetSysTransactionsPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "sysTransactions",
	}, "", c)
}

func (sysTransactionsApi *SysTransactionsApi) Get(c *gin.Context) {
	uid := utils.GetRedisUserID(c)
	if uid == 0 {
		response.FailWithMessage("ID", c)
		return
	}
	fmt.Println("uid", uid)
	// Context
	ctx := c.Request.Context()

	//
	// ，C，
	sysTransactionsService.GetSysTransactionsPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "sysTransactions",
	}, "", c)
}
