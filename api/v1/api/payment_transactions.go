package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/api"
	apiReq "github.com/flipped-aurora/gin-vue-admin/server/model/api/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/i18n"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/payment"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PaymentTransactionsApi struct{}

// CreatePaymentTransactions paymentTransactions
// @Tags PaymentTransactions
// @Summary paymentTransactions
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body api.PaymentTransactions true "paymentTransactions"
// @Success 200 {object} response.Response{msg=string} ""
// @Router /paymentTransactions/createPaymentTransactions [post]
func (paymentTransactionsApi *PaymentTransactionsApi) CreatePaymentTransactions(c *gin.Context) {
	// Context
	ctx := c.Request.Context()

	var paymentTransactions api.PaymentTransactions
	err := c.ShouldBindJSON(&paymentTransactions)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = paymentTransactionsService.CreatePaymentTransactions(ctx, &paymentTransactions)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage(":"+err.Error(), c)
		return
	}
	response.OkWithMessage("", c)
}

// DeletePaymentTransactions paymentTransactions
// @Tags PaymentTransactions
// @Summary paymentTransactions
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body api.PaymentTransactions true "paymentTransactions"
// @Success 200 {object} response.Response{msg=string} ""
// @Router /paymentTransactions/deletePaymentTransactions [delete]
func (paymentTransactionsApi *PaymentTransactionsApi) DeletePaymentTransactions(c *gin.Context) {
	// Context
	ctx := c.Request.Context()

	id := c.Query("id")
	err := paymentTransactionsService.DeletePaymentTransactions(ctx, id)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage(":"+err.Error(), c)
		return
	}
	response.OkWithMessage("", c)
}

// DeletePaymentTransactionsByIds paymentTransactions
// @Tags PaymentTransactions
// @Summary paymentTransactions
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} ""
// @Router /paymentTransactions/deletePaymentTransactionsByIds [delete]
func (paymentTransactionsApi *PaymentTransactionsApi) DeletePaymentTransactionsByIds(c *gin.Context) {
	// Context
	ctx := c.Request.Context()

	ids := c.QueryArray("ids[]")
	err := paymentTransactionsService.DeletePaymentTransactionsByIds(ctx, ids)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage(":"+err.Error(), c)
		return
	}
	response.OkWithMessage("", c)
}

// UpdatePaymentTransactions paymentTransactions
// @Tags PaymentTransactions
// @Summary paymentTransactions
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body api.PaymentTransactions true "paymentTransactions"
// @Success 200 {object} response.Response{msg=string} ""
// @Router /paymentTransactions/updatePaymentTransactions [put]
func (paymentTransactionsApi *PaymentTransactionsApi) UpdatePaymentTransactions(c *gin.Context) {
	// ctxcontext
	ctx := c.Request.Context()

	var paymentTransactions api.PaymentTransactions
	err := c.ShouldBindJSON(&paymentTransactions)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = paymentTransactionsService.UpdatePaymentTransactions(ctx, paymentTransactions)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage(":"+err.Error(), c)
		return
	}
	response.OkWithMessage("", c)
}

// FindPaymentTransactions idpaymentTransactions
// @Tags PaymentTransactions
// @Summary idpaymentTransactions
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id query int true "idpaymentTransactions"
// @Success 200 {object} response.Response{data=api.PaymentTransactions,msg=string} ""
// @Router /paymentTransactions/findPaymentTransactions [get]
func (paymentTransactionsApi *PaymentTransactionsApi) FindPaymentTransactions(c *gin.Context) {
	// Context
	ctx := c.Request.Context()

	id := c.Query("id")
	repaymentTransactions, err := paymentTransactionsService.GetPaymentTransactions(ctx, id)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage(":"+err.Error(), c)
		return
	}
	response.OkWithData(repaymentTransactions, c)
}

// GetPaymentTransactionsList paymentTransactions
// @Tags PaymentTransactions
// @Summary paymentTransactions
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query apiReq.PaymentTransactionsSearch true "paymentTransactions"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} ""
// @Router /paymentTransactions/getPaymentTransactionsList [get]
func (paymentTransactionsApi *PaymentTransactionsApi) GetPaymentTransactionsList(c *gin.Context) {
	// Context
	ctx := c.Request.Context()

	var pageInfo apiReq.PaymentTransactionsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := paymentTransactionsService.GetPaymentTransactionsInfoList(ctx, pageInfo)
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

// GetPaymentTransactionsPublic paymentTransactions
// @Tags PaymentTransactions
// @Summary paymentTransactions
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} ""
// @Router /paymentTransactions/getPaymentTransactionsPublic [get]
func (paymentTransactionsApi *PaymentTransactionsApi) GetPaymentTransactionsPublic(c *gin.Context) {
	// Context
	ctx := c.Request.Context()

	//
	// ，C，
	paymentTransactionsService.GetPaymentTransactionsPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "paymentTransactions",
	}, "", c)
}

// CreateTrade
func (paymentTransactionsApi *PaymentTransactionsApi) CreateTrade(c *gin.Context) {

	uid := utils.GetRedisUserID(c)
	if uid == 0 {
		response.Result(401, nil, "user fail", c)
		return
	}

	pc := payment.InitPayment()

	var r apiReq.CreateTradeData
	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(r, utils.CreateTradeVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//
	formData := url.Values{}
	formData.Set("merchantId", pc.MerchantId)
	formData.Set("merchantOrderNo", fmt.Sprintf("ORDER_%d", time.Now().Unix()))
	formData.Set("amount", strconv.FormatInt(r.Amount*100, 10))
	formData.Set("payType", "PIX_QRCODE")
	formData.Set("currency", "BRL")
	formData.Set("content", "CreateTrade")
	formData.Set("clientIp", "115.227.31.245")
	formData.Set("callback", "http://115.227.31.245:8889/callback/trade")
	formData.Set("redirect", "http://115.227.31.245:7072")

	for k, v := range formData {
		fmt.Printf("  %s: %s\n", k, v[0])
	}

	signature := pc.GenerateFormSign(formData)
	formData.Set("sign", signature)

	resp, err := http.PostForm(pc.BaseURL+"/api/open/merchant/trade/create", formData)
	if err != nil {
		fmt.Printf(" post fail: %v\n", err)
		response.FailWithMessage("post fail", c)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf(" read all fail: %v\n", err)
		response.FailWithMessage("read all fail", c)
		return
	}

	if resp.StatusCode != 200 {
		response.FailWithMessage("StatusCode no 200", c)
		return
	}
	var paymentResponse payment.CreateTradeResponse

	if err := json.Unmarshal(body, &paymentResponse); err != nil {
		fmt.Printf(" json fail: %v\n", err)
		response.FailWithMessage("json fail", c)

		return
	}
	paymentTransactions := api.PaymentTransactions{
		UserId:          uint(uid),
		MerchantOrderNo: paymentResponse.Data.MerchantOrderNo,
		OrderNo:         paymentResponse.Data.OrderNo,
		TransactionType: 1,
		Amount:          int(paymentResponse.Data.Amount),
		Currency:        paymentResponse.Data.Currency,
		Status:          paymentResponse.Data.Status,
		PayType:         "PIX",
		AccountType:     "",
		AccountNo:       "",
		AccountName:     "",
		Content:         "CreateTrade",
		ClientIp:        c.ClientIP(),
		CallbackUrl:     "",
		RedirectUrl:     "",
		PayUrl:          "",
		PayRaw:          "",
		ErrorMsg:        "",
		RefCpf:          "",
		RefName:         "",
	}

	err = paymentTransactionsService.Create(c, paymentTransactions)
	if err != nil {

		utils.FailWithMessageI18n(i18n.MsgCreateRecordFailed, c)
		return
	}
	response.OkWithData(paymentResponse.Data.PayUrl, c)
}

// CreatePayment
func (paymentTransactionsApi *PaymentTransactionsApi) CreatePayment(c *gin.Context) {

	uid := utils.GetRedisUserID(c)
	if uid == 0 {
		utils.UnauthorizedI18n(c)
		return
	}

	ctx := c.Request.Context()

	var r apiReq.CreatePaymentData
	err := c.ShouldBindJSON(&r)
	if err != nil {
		utils.FailWithMessageI18n(i18n.MsgInvalidRequest, c)
		return
	}
	err = utils.Verify(r, utils.CreateTradeVerify)
	if err != nil {
		utils.FailWithMessageI18n(i18n.MsgInvalidAmount, c)
		return
	}

	var user system.ApiSysUser
	redisuser, _ := global.GVA_REDIS.Get(c, fmt.Sprintf("user_%d", uid)).Result()
	if redisuser == "" {
		utils.UnauthorizedI18n(c)
		return
	}
	err = json.Unmarshal([]byte(redisuser), &user)
	if err != nil {
		global.GVA_LOG.Error("Failed to unmarshal user data", zap.Error(err))
		utils.UnauthorizedI18n(c)
		return
	}
	fmt.Println("user.Balance", user.Balance)
	fmt.Println("user.ID", user.ID)
	fmt.Println("r.Amount", r.Amount)
	if user.Balance < float64(r.Amount) {
		utils.FailWithMessageI18n(i18n.MsgInsufficientFunds, c)
		return
	}

	userWithdrawalAccounts, err := userWithdrawalAccountsService.GetUserWithdrawalAccounts(ctx, r.Id)
	if err != nil {

		utils.FailWithMessageI18n(i18n.MsgAccountNotFound, c)
		return
	}

	paymentTransactions := api.PaymentTransactions{
		UserId:          uint(uid),
		MerchantOrderNo: "",
		OrderNo:         fmt.Sprintf("ORDER_%d", time.Now().Unix()),
		TransactionType: 2,
		Amount:          int(r.Amount * 100),
		Currency:        "BRL",
		Status:          "WAITING_PAY",
		PayType:         "PIX",
		AccountType:     userWithdrawalAccounts.AccountType,
		AccountNo:       userWithdrawalAccounts.AccountNumber,
		AccountName:     userWithdrawalAccounts.AccountName,
		Content:         "CreatePayment",
		ClientIp:        c.ClientIP(),
		CallbackUrl:     "",
		RedirectUrl:     "",
		PayUrl:          "",
		PayRaw:          "",
		ErrorMsg:        "",
		RefCpf:          userWithdrawalAccounts.CpfNumber,
		RefName:         userWithdrawalAccounts.AccountName,
	}

	err = paymentTransactionsService.Create(ctx, paymentTransactions)
	if err != nil {

		utils.FailWithMessageI18n(i18n.MsgCreateRecordFailed, c)
		return
	}

	user.Balance = user.Balance - float64(r.Amount)
	userJson, err := json.Marshal(user)
	if err != nil {
		global.GVA_LOG.Error("CreatePayment Failed to marshal user data", zap.Error(err))
	} else {
		err = global.GVA_REDIS.Set(c, fmt.Sprintf("user_%d", user.ID), string(userJson), 0).Err()
		if err != nil {
			global.GVA_LOG.Error("CreatePayment Failed to save user data to Redis", zap.Error(err))
		}
	}

	utils.OkWithMessageI18n(i18n.MsgWithdrawalPending, c)
}

func (paymentTransactionsApi *PaymentTransactionsApi) AdminCreatePayment(c *gin.Context) {

	uid := utils.GetRedisUserID(c)
	if uid == 0 {
		utils.UnauthorizedI18n(c)
		return
	}

	pc := payment.InitPayment()

	ctx := c.Request.Context()

	var r apiReq.CreatePaymentData
	err := c.ShouldBindJSON(&r)
	if err != nil {
		utils.FailWithMessageI18n(i18n.MsgInvalidRequest, c)
		return
	}
	err = utils.Verify(r, utils.CreateTradeVerify)
	if err != nil {
		utils.FailWithMessageI18n(i18n.MsgInvalidAmount, c)
		return
	}

	var user system.ApiSysUser
	redisuser, _ := global.GVA_REDIS.Get(c, fmt.Sprintf("user_%d", uid)).Result()
	if redisuser == "" {
		utils.UnauthorizedI18n(c)
		return
	}
	err = json.Unmarshal([]byte(redisuser), &user)
	if err != nil {
		global.GVA_LOG.Error("Failed to unmarshal user data", zap.Error(err))
		utils.UnauthorizedI18n(c)
		return
	}
	if user.Balance < float64(r.Amount) {
		utils.FailWithMessageI18n(i18n.MsgInsufficientFunds, c)
		return
	}

	userWithdrawalAccounts, err := userWithdrawalAccountsService.GetUserWithdrawalAccounts(ctx, r.Id)
	if err != nil {

		utils.FailWithMessageI18n(i18n.MsgAccountNotFound, c)
		return
	}

	formData := url.Values{}
	formData.Set("merchantId", pc.MerchantId)
	formData.Set("merchantOrderNo", fmt.Sprintf("ORDER_%d", time.Now().Unix()))
	formData.Set("amount", strconv.FormatInt(r.Amount*100, 10))
	formData.Set("currency", "BRL")
	formData.Set("accountType", userWithdrawalAccounts.AccountType)
	formData.Set("accountNo", userWithdrawalAccounts.AccountNumber)
	formData.Set("accountName", userWithdrawalAccounts.AccountName)

	// GenerateFormSign
	signature := pc.GenerateFormSign(formData)
	formData.Set("sign", signature)

	for k, v := range formData {
		if k == "sign" {
			fmt.Printf("  %s: %s (GenerateFormSign)\n", k, v[0])
		} else {
			fmt.Printf("  %s: %s\n", k, v[0])
		}
	}

	resp, err := http.PostForm(pc.BaseURL+"/api/open/merchant/payment/create", formData)
	if err != nil {

		utils.FailWithMessageI18n(i18n.MsgWithdrawalFailed, c)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {

		utils.FailWithMessageI18n(i18n.MsgWithdrawalFailed, c)
		return
	}

	if resp.StatusCode != 200 {

		utils.FailWithMessageI18n(i18n.MsgWithdrawalFailed, c)
		return
	}

	var paymentResponse payment.CreatePaymentResponse
	if err := json.Unmarshal(body, &paymentResponse); err != nil {
		global.GVA_LOG.Error("paymentResponse JSON fail",
			zap.Error(err),
			zap.String("responseBody", string(body)),
		)
		utils.FailWithMessageI18n(i18n.MsgWithdrawalFailed, c)
		return
	}

	paymentTransactions := api.PaymentTransactions{
		UserId:          uint(uid),
		MerchantOrderNo: paymentResponse.Data.MerchantOrderNo,
		OrderNo:         paymentResponse.Data.OrderNo,
		TransactionType: 2,
		Amount:          int(paymentResponse.Data.Amount),
		Currency:        paymentResponse.Data.Currency,
		Status:          paymentResponse.Data.Status,
		PayType:         "PIX",
		AccountType:     userWithdrawalAccounts.AccountType,
		AccountNo:       userWithdrawalAccounts.AccountNumber,
		AccountName:     userWithdrawalAccounts.AccountName,
		Content:         "CreatePayment",
		ClientIp:        c.ClientIP(),
		CallbackUrl:     "",
		RedirectUrl:     "",
		PayUrl:          "",
		PayRaw:          "",
		ErrorMsg:        "",
		RefCpf:          userWithdrawalAccounts.CpfNumber,
		RefName:         userWithdrawalAccounts.AccountName,
	}

	err = paymentTransactionsService.Create(ctx, paymentTransactions)
	if err != nil {
		global.GVA_LOG.Error("paymentTransactionsService.Create",
			zap.Error(err),
			zap.Any("paymentTransactions", paymentTransactions),
			zap.Any("userWithdrawalAccounts", userWithdrawalAccounts),
			zap.Any("paymentResponse", paymentResponse),
		)
		utils.FailWithMessageI18n(i18n.MsgCreateRecordFailed, c)
		return
	}

	user.Balance = user.Balance - float64(r.Amount)
	userJson, err := json.Marshal(user)
	if err != nil {
		global.GVA_LOG.Error("CreatePayment Failed to marshal user data", zap.Error(err))
	} else {
		err = global.GVA_REDIS.Set(c, fmt.Sprintf("user_%d", user.ID), string(userJson), 0).Err()
		if err != nil {
			global.GVA_LOG.Error("CreatePayment Failed to save user data to Redis", zap.Error(err))
		}
	}

	global.GVA_LOG.Info("CreatePayment",
		zap.Uint("userId", uint(uid)),
		zap.String("merchantOrderNo", paymentResponse.Data.MerchantOrderNo),
		zap.String("orderNo", paymentResponse.Data.OrderNo),
		zap.Int64("amount", paymentResponse.Data.Amount),
		zap.String("status", paymentResponse.Data.Status),
	)

	utils.OkWithMessageI18n(i18n.MsgWithdrawalPending, c)
}
func (paymentTransactionsApi *PaymentTransactionsApi) GetPaymentList(c *gin.Context) {

	uid := utils.GetRedisUserID(c)
	if uid == 0 {
		utils.UnauthorizedI18n(c)
		return
	}

	ctx := c.Request.Context()

	var pageInfo apiReq.PaymentTransactionsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var transactionType = 2
	list, total, err := paymentTransactionsService.GetPaymentList(ctx, pageInfo, uid, transactionType)
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

func (paymentTransactionsApi *PaymentTransactionsApi) GetTradeList(c *gin.Context) {

	uid := utils.GetRedisUserID(c)
	if uid == 0 {
		utils.UnauthorizedI18n(c)
		return
	}

	ctx := c.Request.Context()

	var pageInfo apiReq.PaymentTransactionsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var transactionType = 1
	list, total, err := paymentTransactionsService.GetPaymentList(ctx, pageInfo, uid, transactionType)
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
