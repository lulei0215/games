package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

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

	// Print all raw POST data
	body, err := c.GetRawData()
	if err != nil {
		fmt.Printf("Error reading raw data: %v\n", err)
	} else {
		fmt.Printf("=== TradeCallback Raw POST Data ===\n")
		fmt.Printf("Raw Body: %s\n", string(body))
		fmt.Printf("Content-Type: %s\n", c.GetHeader("Content-Type"))
		fmt.Printf("User-Agent: %s\n", c.GetHeader("User-Agent"))
		fmt.Printf("Client IP: %s\n", c.ClientIP())
		fmt.Printf("Request URL: %s\n", c.Request.URL.String())
		fmt.Printf("Request Method: %s\n", c.Request.Method)
		fmt.Printf("All Headers:\n")
		for key, values := range c.Request.Header {
			fmt.Printf("  %s: %v\n", key, values)
		}
		fmt.Printf("=== End Raw POST Data ===\n")
	}

	// Re-set the request body for form binding
	if len(body) > 0 {
		c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
	}

	// Handle form-encoded data instead of JSON
	var callbackData apiReq.TradeCallbackFormRequest
	err = c.ShouldBind(&callbackData)
	if err != nil {
		fmt.Printf("Form binding error: %v\n", err)
		fmt.Printf("Attempted to bind data: %+v\n", callbackData)
		response.FailWithMessage(err.Error(), c)
		return
	}

	//Sign
	dataMap := map[string]interface{}{
		"merchantId":      callbackData.MerchantId,
		"merchantOrderNo": callbackData.MerchantOrderNo,
		"orderNo":         callbackData.OrderNo,
		"amount":          callbackData.Amount,
		"status":          callbackData.Status,
		"currency":        callbackData.Currency,
		"payType":         callbackData.PayType,
	}
	if callbackData.RefCpf != "" {
		dataMap["ref_cpf"] = callbackData.RefCpf
	}
	if callbackData.RefName != "" {
		dataMap["ref_name"] = callbackData.RefName
	}

	if !pc.VerifyCallbackSign(dataMap, callbackData.Sign) {
		return
	}

	callbackDataJson, _ := json.Marshal(callbackData)

	paymentCallbacks := api.PaymentCallbacks{
		MerchantId:      callbackData.MerchantId,
		MerchantOrderNo: callbackData.MerchantOrderNo,
		OrderNo:         callbackData.OrderNo,
		CallbackType:    1,
		Amount:          callbackData.Amount,
		Currency:        callbackData.Currency,
		Status:          callbackData.Status,
		PayType:         callbackData.PayType,
		RefCpf:          callbackData.RefCpf,
		RefName:         callbackData.RefName,
		ErrorMsg:        "",
		CallbackData:    string(callbackDataJson),
		Sign:            callbackData.Sign,
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
	if callbackData.Status == "PAID" {
		err = paymentTransactionsService.TradeOk(ctx, callbackData.MerchantOrderNo, callbackData.OrderNo, paymentCallbacks)
		if err != nil {
			return
		}
		c.JSON(200, gin.H{
			"success": "ok",
		})
	}
}

// PaymentCallback
func (paymentCallbacksApi *PaymentCallbacksApi) PaymentCallback(c *gin.Context) {

	pc := payment.InitPayment()

	// Print all raw POST data
	body, err := c.GetRawData()
	if err != nil {
		fmt.Printf("Error reading raw data: %v\n", err)
	} else {
		fmt.Printf("=== PaymentCallback Raw POST Data ===\n")
		fmt.Printf("Raw Body: %s\n", string(body))
		fmt.Printf("Content-Type: %s\n", c.GetHeader("Content-Type"))
		fmt.Printf("User-Agent: %s\n", c.GetHeader("User-Agent"))
		fmt.Printf("Client IP: %s\n", c.ClientIP())
		fmt.Printf("Request URL: %s\n", c.Request.URL.String())
		fmt.Printf("Request Method: %s\n", c.Request.Method)
		fmt.Printf("All Headers:\n")
		for key, values := range c.Request.Header {
			fmt.Printf("  %s: %v\n", key, values)
		}
		fmt.Printf("=== End Raw POST Data ===\n")
	}

	// Re-set the request body for JSON binding
	if len(body) > 0 {
		c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
	}

	var callbackData apiReq.PaymentCallbackFormRequest
	err = c.ShouldBind(&callbackData)
	if err != nil {
		fmt.Printf("JSON binding error: %v\n", err)
		fmt.Printf("Attempted to bind data: %+v\n", callbackData)
		response.FailWithMessage(err.Error(), c)
		return
	}

	// Sign

	dataMap := map[string]interface{}{
		"merchantId":      callbackData.MerchantId,
		"merchantOrderNo": callbackData.MerchantOrderNo,
		"orderNo":         callbackData.OrderNo,
		"amount":          callbackData.Amount,
		"status":          callbackData.Status,
		"currency":        callbackData.Currency,
	}
	fmt.Println("3")
	if callbackData.ErrorMsg != "" {
		dataMap["errorMsg"] = callbackData.ErrorMsg
	}

	if !pc.VerifyCallbackSign(dataMap, callbackData.Sign) {
		return
	}

	callbackDataJson, _ := json.Marshal(callbackData)

	paymentCallbacks := api.PaymentCallbacks{
		MerchantId:      callbackData.MerchantId,
		MerchantOrderNo: callbackData.MerchantOrderNo,
		OrderNo:         callbackData.OrderNo,
		CallbackType:    2,
		Amount:          callbackData.Amount,
		Currency:        callbackData.Currency,
		Status:          callbackData.Status,
		PayType:         "",
		RefCpf:          "",
		RefName:         "",
		ErrorMsg:        "",
		CallbackData:    string(callbackDataJson),
		Sign:            callbackData.Sign,
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
	if callbackData.Status == "SUCCESS" {
		err = paymentTransactionsService.PaymentOk(ctx, callbackData.MerchantOrderNo, callbackData.OrderNo)
		if err != nil {
			return
		}
		c.JSON(200, gin.H{
			"success": "ok",
		})
	}
}
