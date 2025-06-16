package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/announcement/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/announcement/model/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var Info = new(info)

type info struct{}

// CreateInfo
// @Tags Info
// @Summary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Info true ""
// @Success 200 {object} response.Response{msg=string} ""
// @Router /info/createInfo [post]
func (a *info) CreateInfo(c *gin.Context) {
	var info model.Info
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceInfo.CreateInfo(&info)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage("", c)
		return
	}
	response.OkWithMessage("", c)
}

// DeleteInfo
// @Tags Info
// @Summary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Info true ""
// @Success 200 {object} response.Response{msg=string} ""
// @Router /info/deleteInfo [delete]
func (a *info) DeleteInfo(c *gin.Context) {
	ID := c.Query("ID")
	err := serviceInfo.DeleteInfo(ID)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage("", c)
		return
	}
	response.OkWithMessage("", c)
}

// DeleteInfoByIds
// @Tags Info
// @Summary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} ""
// @Router /info/deleteInfoByIds [delete]
func (a *info) DeleteInfoByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := serviceInfo.DeleteInfoByIds(IDs); err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage("", c)
		return
	}
	response.OkWithMessage("", c)
}

// UpdateInfo
// @Tags Info
// @Summary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Info true ""
// @Success 200 {object} response.Response{msg=string} ""
// @Router /info/updateInfo [put]
func (a *info) UpdateInfo(c *gin.Context) {
	var info model.Info
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceInfo.UpdateInfo(info)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage("", c)
		return
	}
	response.OkWithMessage("", c)
}

// FindInfo id
// @Tags Info
// @Summary id
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Info true "id"
// @Success 200 {object} response.Response{data=model.Info,msg=string} ""
// @Router /info/findInfo [get]
func (a *info) FindInfo(c *gin.Context) {
	ID := c.Query("ID")
	reinfo, err := serviceInfo.GetInfo(ID)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage("", c)
		return
	}
	response.OkWithData(reinfo, c)
}

// GetInfoList
// @Tags Info
// @Summary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.InfoSearch true ""
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} ""
// @Router /info/getInfoList [get]
func (a *info) GetInfoList(c *gin.Context) {
	var pageInfo request.InfoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := serviceInfo.GetInfoInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage("", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "", c)
}

// GetInfoDataSource Info
// @Tags Info
// @Summary Info
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} ""
// @Router /info/getInfoDataSource [get]
func (a *info) GetInfoDataSource(c *gin.Context) {
	//
	dataSource, err := serviceInfo.GetInfoDataSource()
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage("", c)
		return
	}
	response.OkWithData(dataSource, c)
}

// GetInfoPublic
// @Tags Info
// @Summary
// @accept application/json
// @Produce application/json
// @Param data query request.InfoSearch true ""
// @Success 200 {object} response.Response{data=object,msg=string} ""
// @Router /info/getInfoPublic [get]
func (a *info) GetInfoPublic(c *gin.Context) {
	//  ，C，
	response.OkWithDetailed(gin.H{"info": ""}, "", c)
}
