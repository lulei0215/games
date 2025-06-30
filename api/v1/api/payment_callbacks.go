package api

import (
	"encoding/json"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/api"
	apiReq "github.com/flipped-aurora/gin-vue-admin/server/model/api/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/payment"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PaymentCallbacksApi struct{}

// CreatePaymentCallbacks paymentCallbacks
// @Tags PaymentCallbacks
// @Summary paymentCallbacks
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body api.PaymentCallbacks true "paymentCallbacks"
// @Success 200 {object} response.Response{msg=string} ""
// @Router /paymentCallbacks/createPaymentCallbacks [post]
func (paymentCallbacksApi *PaymentCallbacksApi) CreatePaymentCallbacks(c *gin.Context) {
	// Context
	ctx := c.Request.Context()

	var paymentCallbacks api.PaymentCallbacks
	err := c.ShouldBindJSON(&paymentCallbacks)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = paymentCallbacksService.CreatePaymentCallbacks(ctx, &paymentCallbacks)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage(":"+err.Error(), c)
		return
	}
	response.OkWithMessage("", c)
}

// DeletePaymentCallbacks paymentCallbacks
// @Tags PaymentCallbacks
// @Summary paymentCallbacks
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body api.PaymentCallbacks true "paymentCallbacks"
// @Success 200 {object} response.Response{msg=string} ""
// @Router /paymentCallbacks/deletePaymentCallbacks [delete]
func (paymentCallbacksApi *PaymentCallbacksApi) DeletePaymentCallbacks(c *gin.Context) {
	// Context
	ctx := c.Request.Context()

	id := c.Query("id")
	err := paymentCallbacksService.DeletePaymentCallbacks(ctx, id)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage(":"+err.Error(), c)
		return
	}
	response.OkWithMessage("", c)
}

// DeletePaymentCallbacksByIds paymentCallbacks
// @Tags PaymentCallbacks
// @Summary paymentCallbacks
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} ""
// @Router /paymentCallbacks/deletePaymentCallbacksByIds [delete]
func (paymentCallbacksApi *PaymentCallbacksApi) DeletePaymentCallbacksByIds(c *gin.Context) {
	// Context
	ctx := c.Request.Context()

	ids := c.QueryArray("ids[]")
	err := paymentCallbacksService.DeletePaymentCallbacksByIds(ctx, ids)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage(":"+err.Error(), c)
		return
	}
	response.OkWithMessage("", c)
}

// UpdatePaymentCallbacks paymentCallbacks
// @Tags PaymentCallbacks
// @Summary paymentCallbacks
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body api.PaymentCallbacks true "paymentCallbacks"
// @Success 200 {object} response.Response{msg=string} ""
// @Router /paymentCallbacks/updatePaymentCallbacks [put]
func (paymentCallbacksApi *PaymentCallbacksApi) UpdatePaymentCallbacks(c *gin.Context) {
	// ctxcontext
	ctx := c.Request.Context()

	var paymentCallbacks api.PaymentCallbacks
	err := c.ShouldBindJSON(&paymentCallbacks)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = paymentCallbacksService.UpdatePaymentCallbacks(ctx, paymentCallbacks)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage(":"+err.Error(), c)
		return
	}
	response.OkWithMessage("", c)
}

// FindPaymentCallbacks idpaymentCallbacks
// @Tags PaymentCallbacks
// @Summary idpaymentCallbacks
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id query int true "idpaymentCallbacks"
// @Success 200 {object} response.Response{data=api.PaymentCallbacks,msg=string} ""
// @Router /paymentCallbacks/findPaymentCallbacks [get]
func (paymentCallbacksApi *PaymentCallbacksApi) FindPaymentCallbacks(c *gin.Context) {
	// Context
	ctx := c.Request.Context()

	id := c.Query("id")
	repaymentCallbacks, err := paymentCallbacksService.GetPaymentCallbacks(ctx, id)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage(":"+err.Error(), c)
		return
	}
	response.OkWithData(repaymentCallbacks, c)
}

// GetPaymentCallbacksList paymentCallbacks
// @Tags PaymentCallbacks
// @Summary paymentCallbacks
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query apiReq.PaymentCallbacksSearch true "paymentCallbacks"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} ""
// @Router /paymentCallbacks/getPaymentCallbacksList [get]
func (paymentCallbacksApi *PaymentCallbacksApi) GetPaymentCallbacksList(c *gin.Context) {
	// Context
	ctx := c.Request.Context()

	var pageInfo apiReq.PaymentCallbacksSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := paymentCallbacksService.GetPaymentCallbacksInfoList(ctx, pageInfo)
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

// GetPaymentCallbacksPublic paymentCallbacks
// @Tags PaymentCallbacks
// @Summary paymentCallbacks
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} ""
// @Router /paymentCallbacks/getPaymentCallbacksPublic [get]
func (paymentCallbacksApi *PaymentCallbacksApi) GetPaymentCallbacksPublic(c *gin.Context) {
	// Context
	ctx := c.Request.Context()

	//
	// ，C，
	paymentCallbacksService.GetPaymentCallbacksPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "paymentCallbacks",
	}, "", c)
}

// TradeCallback

func (paymentCallbacksApi *PaymentCallbacksApi) TradeCallback(c *gin.Context) {
	pc := payment.InitPayment()

	var callbackData apiReq.TradeCallbackRequest
	err := c.ShouldBindJSON(&callbackData)
	if err != nil {
		fmt.Println("2", callbackData)
		response.FailWithMessage(err.Error(), c)
		return
	}
	//Sign
	dataMap := map[string]interface{}{
		"merchantId":      callbackData.Data.MerchantId,
		"merchantOrderNo": callbackData.Data.MerchantOrderNo,
		"orderNo":         callbackData.Data.OrderNo,
		"amount":          callbackData.Data.Amount,
		"status":          callbackData.Data.Status,
		"currency":        callbackData.Data.Currency,
		"payType":         callbackData.Data.PayType,
	}
	if callbackData.Data.RefCpf != "" {
		dataMap["ref_cpf"] = callbackData.Data.RefCpf
	}
	if callbackData.Data.RefName != "" {
		dataMap["ref_name"] = callbackData.Data.RefName
	}

	if !pc.VerifyCallbackSign(dataMap, callbackData.Data.Sign) {
		return
	}

	callbackDataJson, _ := json.Marshal(callbackData)

	paymentCallbacks := api.PaymentCallbacks{
		MerchantId:      callbackData.Data.MerchantId.(string),
		MerchantOrderNo: callbackData.Data.MerchantOrderNo,
		OrderNo:         callbackData.Data.OrderNo,
		CallbackType:    1,
		Amount:          callbackData.Data.Amount,
		Currency:        callbackData.Data.Currency,
		Status:          callbackData.Data.Status,
		PayType:         callbackData.Data.PayType,
		RefCpf:          callbackData.Data.RefCpf,
		RefName:         callbackData.Data.RefName,
		ErrorMsg:        "",
		CallbackData:    string(callbackDataJson),
		Sign:            callbackData.Data.Sign,
		SignVerified:    true,
		IpAddress:       c.ClientIP(),
		UserAgent:       c.GetHeader("User-Agent"),
		Processed:       false,
		ProcessedTime:   nil,
		RetryCount:      0,
		LastRetryTime:   nil,
		ErrorReason:     "",
		Remark:          "TradeCallback",
	}

	ctx := c.Request.Context()
	err = paymentCallbacksService.Create(ctx, paymentCallbacks)
	if err != nil {
		return
	}
	if callbackData.Data.Status == "PAID" {
		err = paymentTransactionsService.TradeOk(ctx, callbackData.Data.MerchantOrderNo, callbackData.Data.OrderNo, paymentCallbacks)
		if err != nil {
			return
		}
		c.JSON(200, gin.H{
			"success": "ok",
		})
	}
	return
}

// PaymentCallback
func (paymentCallbacksApi *PaymentCallbacksApi) PaymentCallback(c *gin.Context) {

	pc := payment.InitPayment()

	var callbackData apiReq.PaymentCallbackRequest
	err := c.ShouldBindJSON(&callbackData)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// Sign
	dataMap := map[string]interface{}{
		"merchantId":      callbackData.Data.MerchantId,
		"merchantOrderNo": callbackData.Data.MerchantOrderNo,
		"orderNo":         callbackData.Data.OrderNo,
		"amount":          callbackData.Data.Amount,
		"status":          callbackData.Data.Status,
		"currency":        callbackData.Data.Currency,
	}
	fmt.Println("3")
	if callbackData.Data.ErrorMsg != "" {
		dataMap["errorMsg"] = callbackData.Data.ErrorMsg
	}

	if !pc.VerifyCallbackSign(dataMap, callbackData.Data.Sign) {
		return
	}

	callbackDataJson, _ := json.Marshal(callbackData)

	paymentCallbacks := api.PaymentCallbacks{
		MerchantId:      callbackData.Data.MerchantId.(string),
		MerchantOrderNo: callbackData.Data.MerchantOrderNo,
		OrderNo:         callbackData.Data.OrderNo,
		CallbackType:    2,
		Amount:          callbackData.Data.Amount,
		Currency:        callbackData.Data.Currency,
		Status:          callbackData.Data.Status,
		PayType:         "",
		RefCpf:          "",
		RefName:         "",
		ErrorMsg:        "",
		CallbackData:    string(callbackDataJson),
		Sign:            callbackData.Data.Sign,
		SignVerified:    true,
		IpAddress:       c.ClientIP(),
		UserAgent:       c.GetHeader("User-Agent"),
		Processed:       false,
		ProcessedTime:   nil,
		RetryCount:      0,
		LastRetryTime:   nil,
		ErrorReason:     "",
		Remark:          "PaymentCallback",
	}

	ctx := c.Request.Context()
	err = paymentCallbacksService.Create(ctx, paymentCallbacks)
	fmt.Println("5")
	if err != nil {
		return
	}
	fmt.Println("6")
	if callbackData.Data.Status == "SUCCESS" {
		err = paymentTransactionsService.PaymentOk(ctx, callbackData.Data.MerchantOrderNo, callbackData.Data.OrderNo)
		if err != nil {
			return
		}
		c.JSON(200, gin.H{
			"success": "ok",
		})
	}
	return
}
