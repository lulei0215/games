package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
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
	"gorm.io/gorm"
)

// CancelPaymentData 取消提现请求结构体
type CancelPaymentData struct {
	OrderId string `json:"orderId" binding:"required"`
}

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
func (paymentTransactionsApi *PaymentTransactionsApi) CreateTrade2(c *gin.Context) {

	uid := utils.GetRedisUserID(c)
	if uid == 0 {
		response.Result(401, nil, "user fail", c)
		return
	}

	pc := payment.InitPayment2()

	var r apiReq.CreateTradeData2
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

	fmt.Println(1)
	OutTradeNo := fmt.Sprintf("%d%d", time.Now().Unix(), time.Now().UnixNano()%100000)
	status, msg, data := pc.CreatePayin(r, OutTradeNo)
	fmt.Println("status", status)
	fmt.Println("msg", msg)
	fmt.Println("data", data)
	if status != 1 {
		response.FailWithMessage(msg, c)
		return
	}
	totalAmount, err := strconv.ParseFloat(data.TotalAmount, 64)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	fmt.Println(2)
	paymentTransactions := api.PaymentTransactions{
		UserId:          uint(uid),
		MerchantOrderNo: OutTradeNo,
		OrderNo:         data.PayOrderNo,
		TransactionType: 1,
		Amount:          int(totalAmount * 100),
		Currency:        "BRL",
		Status:          "PAYING",
		PayType:         "PIX",
		AccountType:     "PIX",
		AccountNo:       r.PayCardNo,
		AccountName:     r.PayName,
		Content:         "CreateTrade2",
		ClientIp:        c.ClientIP(),
		CallbackUrl:     "",
		RedirectUrl:     "",
		PayUrl:          data.PayURL,
		PayRaw:          data.PayParams,
		ErrorMsg:        "",
		RefCpf:          "",
		RefName:         "",
		PayEmail:        r.PayEmail,
		PayPhone:        r.PayPhone,
		PayBankCode:     r.PayBankCode,
		Type:            3,
	}

	fmt.Println("paymentTransactions::::", paymentTransactions)
	err = paymentTransactionsService.Create(c, paymentTransactions)
	if err != nil {
		utils.FailWithMessageI18n(i18n.MsgCreateRecordFailed, c)
		return
	}
	response.OkWithData(data.PayURL, c)
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
	if user.Balance < float64(r.Amount) {
		utils.FailWithMessageI18n(i18n.MsgInsufficientFunds, c)
		return
	}

	global.GVA_LOG.Info("AdminCreatePayment - Attempting to get user withdrawal accounts",
		zap.Uint("userId", uid),
		zap.Int64("accountId", r.AccountId),
		zap.Any("requestData", r))

	userWithdrawalAccounts, err := userWithdrawalAccountsService.GetUserWithdrawalAccounts(ctx, strconv.FormatInt(r.AccountId, 10))

	if err != nil {
		global.GVA_LOG.Error("AdminCreatePayment - Failed to get user withdrawal accounts",
			zap.Error(err),
			zap.Uint("userId", uid),
			zap.Int64("accountId", r.AccountId))
		utils.FailWithMessageI18n(i18n.MsgAccountNotFound, c)
		return
	}

	paymentTransactions := api.PaymentTransactions{
		UserId:          uint(uid),
		MerchantOrderNo: fmt.Sprintf("ORDER_%d", time.Now().Unix()),
		OrderNo:         "",
		TransactionType: 2,
		Amount:          int(r.Amount) * 100,
		Currency:        "BRL",
		Status:          "WAITING_PAY",
		PayType:         "PIX",
		AccountType:     userWithdrawalAccounts.AccountType,
		AccountNo:       userWithdrawalAccounts.AccountNumber,
		AccountName:     userWithdrawalAccounts.AccountName,
		Content:         "CreatePayment",
		ClientIp:        c.ClientIP(),
		CallbackUrl:     "http://115.227.31.245:8889/callback/payment",
		RedirectUrl:     "",
		PayUrl:          "",
		PayRaw:          "",
		ErrorMsg:        "",
		RefCpf:          userWithdrawalAccounts.CpfNumber,
		RefName:         userWithdrawalAccounts.AccountName,
	}

	// 使用事务函数同时创建支付记录并扣减用户余额
	transactionCode := fmt.Sprintf("CREATE_PAYMENT_%d_%d", user.ID, time.Now().Unix())
	err = utils.CreatePaymentWithBalanceDeduction(c, user.ID, float64(r.Amount), "CreatePayment", func(tx *gorm.DB) error {
		return paymentTransactionsService.CreateWithTx(tx, paymentTransactions)
	}, transactionCode)
	if err != nil {
		global.GVA_LOG.Error("Failed to create payment with balance deduction",
			zap.Error(err),
			zap.Uint("userId", user.ID),
			zap.Float64("amount", float64(r.Amount)))

		if err.Error() == "insufficient balance" {
			utils.FailWithMessageI18n(i18n.MsgInsufficientFunds, c)
		} else {
			utils.FailWithMessageI18n(i18n.MsgSystemError, c)
		}
		return
	}

	global.GVA_LOG.Info("AdminCreatePayment - Successfully processed payment",
		zap.Uint("userId", uint(uid)),
		zap.String("merchantOrderNo", paymentTransactions.MerchantOrderNo),
		zap.String("orderNo", paymentTransactions.OrderNo),
		zap.Int("amount", paymentTransactions.Amount),
		zap.String("status", paymentTransactions.Status))

	utils.OkWithMessageI18n(i18n.MsgWithdrawalPending, c)
}

// CancelPayment 取消提现
func (paymentTransactionsApi *PaymentTransactionsApi) CancelPayment(c *gin.Context) {
	uid := utils.GetRedisUserID(c)
	global.GVA_LOG.Info("CancelPayment Step 1: GetRedisUserID", zap.Uint("userId", uid))
	if uid == 0 {
		utils.UnauthorizedI18n(c)
		return
	}

	ctx := c.Request.Context()

	var r CancelPaymentData
	err := c.ShouldBindJSON(&r)
	global.GVA_LOG.Info("CancelPayment Step 2: ShouldBindJSON", zap.Any("requestData", r), zap.Error(err))
	if err != nil {
		utils.FailWithMessageI18n(i18n.MsgInvalidRequest, c)
		return
	}

	// 验证订单号不能为空
	if r.OrderId == "" {
		utils.FailWithMessageI18n(i18n.MsgInvalidRequest, c)
		return
	}

	global.GVA_LOG.Info("CancelPayment Step 3: Get payment transaction by orderId", zap.String("orderId", r.OrderId))

	// 根据订单号查询提现记录
	paymentTransaction, err := paymentTransactionsService.GetByOrderNo(ctx, r.OrderId)
	global.GVA_LOG.Info("CancelPayment Step 3.1: GetByOrderNo result", zap.Any("paymentTransaction", paymentTransaction), zap.Error(err))
	if err != nil {
		global.GVA_LOG.Error("Failed to get payment transaction", zap.Error(err))
		utils.FailWithMessageI18n(i18n.MsgRecordNotFound, c)
		return
	}

	// 检查记录是否存在
	if paymentTransaction.Id == 0 {
		utils.FailWithMessageI18n(i18n.MsgRecordNotFound, c)
		return
	}

	// 检查用户权限 - 只能取消自己的提现
	if paymentTransaction.UserId != uint(uid) {
		global.GVA_LOG.Error("User trying to cancel another user's withdrawal",
			zap.Uint("requestUserId", uid),
			zap.Uint("recordUserId", paymentTransaction.UserId))
		utils.FailWithMessageI18n(i18n.MsgUnauthorized, c)
		return
	}

	// 检查状态 - 只能取消待处理的提现
	if paymentTransaction.Status != "WAITING_PAY" {
		global.GVA_LOG.Error("Cannot cancel withdrawal with status", zap.String("status", paymentTransaction.Status))
		utils.FailWithMessageI18n(i18n.MsgCannotCancelWithdrawal, c)
		return
	}

	global.GVA_LOG.Info("CancelPayment Step 4: Update payment transaction status and add balance back to user")

	// 使用事务函数同时更新支付状态并加回用户余额
	updateData := api.PaymentTransactions{
		Status:  "CANCELLED",
		Content: "CancelPayment",
	}

	transactionCode := fmt.Sprintf("CANCEL_PAYMENT_%d_%s_%d", paymentTransaction.UserId, r.OrderId, time.Now().Unix())
	err = utils.CancelPaymentWithBalanceAddition(c, paymentTransaction.UserId, float64(paymentTransaction.Amount)/100, "CancelPayment", func(tx *gorm.DB) error {
		return paymentTransactionsService.UpdateWithTx(tx, r.OrderId, updateData)
	}, transactionCode)
	if err != nil {
		global.GVA_LOG.Error("Failed to cancel payment with balance addition",
			zap.Error(err),
			zap.Uint("userId", paymentTransaction.UserId),
			zap.Float64("amount", float64(paymentTransaction.Amount)/100))
		utils.FailWithMessageI18n(i18n.MsgSystemError, c)
		return
	}

	global.GVA_LOG.Info("CancelPayment Step 6: Success", zap.String("orderId", r.OrderId))
	utils.OkWithMessageI18n(i18n.MsgWithdrawalCancelled, c)
}

func String(i int64) string {
	return strconv.FormatInt(i, 10)
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

	global.GVA_LOG.Info("AdminCreatePayment - Attempting to get user withdrawal accounts",
		zap.Uint("userId", uid),
		zap.Int64("accountId", r.AccountId),
		zap.Any("requestData", r))

	userWithdrawalAccounts, err := userWithdrawalAccountsService.GetUserWithdrawalAccounts(ctx, strconv.FormatInt(r.AccountId, 10))

	if err != nil {
		global.GVA_LOG.Error("AdminCreatePayment - Failed to get user withdrawal accounts",
			zap.Error(err),
			zap.Uint("userId", uid),
			zap.Int64("accountId", r.AccountId))
		utils.FailWithMessageI18n(i18n.MsgAccountNotFound, c)
		return
	}

	global.GVA_LOG.Info("AdminCreatePayment - Successfully retrieved user withdrawal accounts",
		zap.Uint("userId", uid),
		zap.Int64("accountId", r.AccountId),
		zap.Any("userWithdrawalAccounts", userWithdrawalAccounts))

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

	// 使用事务函数同时创建支付记录并扣减用户余额
	transactionCode := fmt.Sprintf("ADMIN_CREATE_PAYMENT_%d_%d", user.ID, time.Now().Unix())
	err = utils.CreatePaymentWithBalanceDeduction(c, user.ID, float64(r.Amount), "AdminCreatePayment", func(tx *gorm.DB) error {
		return paymentTransactionsService.CreateWithTx(tx, paymentTransactions)
	}, transactionCode)
	if err != nil {
		global.GVA_LOG.Error("Failed to create payment with balance deduction",
			zap.Error(err),
			zap.Uint("userId", user.ID),
			zap.Float64("amount", float64(r.Amount)))

		if err.Error() == "insufficient balance" {
			utils.FailWithMessageI18n(i18n.MsgInsufficientFunds, c)
		} else {
			utils.FailWithMessageI18n(i18n.MsgSystemError, c)
		}
		return
	}

	global.GVA_LOG.Info("AdminCreatePayment - Successfully processed payment",
		zap.Uint("userId", uint(uid)),
		zap.String("merchantOrderNo", paymentResponse.Data.MerchantOrderNo),
		zap.String("orderNo", paymentResponse.Data.OrderNo),
		zap.Int64("amount", paymentResponse.Data.Amount),
		zap.String("status", paymentResponse.Data.Status))

	utils.OkWithMessageI18n(i18n.MsgWithdrawalPending, c)
}

func (paymentTransactionsApi *PaymentTransactionsApi) CreatePayment2(c *gin.Context) {
	global.GVA_LOG.Info("CreatePayment2 - Starting payment process")

	uid := utils.GetRedisUserID(c)
	global.GVA_LOG.Info("CreatePayment2 - Retrieved user ID from Redis", zap.Uint("uid", uid))

	if uid == 0 {
		global.GVA_LOG.Warn("CreatePayment2 - User ID is 0, unauthorized")
		utils.UnauthorizedI18n(c)
		return
	}

	global.GVA_LOG.Info("CreatePayment2 - Initializing payment client")
	pc := payment.InitPayment2()

	ctx := c.Request.Context()

	global.GVA_LOG.Info("CreatePayment2 - Binding JSON request")
	var r apiReq.CreatePaymentData2
	err := c.ShouldBindJSON(&r)
	if err != nil {
		global.GVA_LOG.Error("CreatePayment2 - Failed to bind JSON request", zap.Error(err))
		utils.FailWithMessageI18n(i18n.MsgInvalidRequest, c)
		return
	}
	global.GVA_LOG.Info("CreatePayment2 - Request data bound successfully", zap.Any("requestData", r))

	global.GVA_LOG.Info("CreatePayment2 - Verifying request data")
	err = utils.Verify(r, utils.CreateTradeVerify)
	if err != nil {
		global.GVA_LOG.Error("CreatePayment2 - Request verification failed", zap.Error(err))
		utils.FailWithMessageI18n(i18n.MsgInvalidAmount, c)
		return
	}
	global.GVA_LOG.Info("CreatePayment2 - Request verification passed")

	global.GVA_LOG.Info("CreatePayment2 - Getting user data from Redis")
	var user system.ApiSysUser
	redisuser, _ := global.GVA_REDIS.Get(c, fmt.Sprintf("user_%d", uid)).Result()
	if redisuser == "" {
		global.GVA_LOG.Warn("CreatePayment2 - User data not found in Redis", zap.Uint("uid", uid))
		utils.UnauthorizedI18n(c)
		return
	}
	global.GVA_LOG.Info("CreatePayment2 - User data retrieved from Redis", zap.String("redisUser", redisuser))

	err = json.Unmarshal([]byte(redisuser), &user)
	if err != nil {
		global.GVA_LOG.Error("CreatePayment2 - Failed to unmarshal user data", zap.Error(err))
		utils.UnauthorizedI18n(c)
		return
	}
	if user.Robot == 1 {
		global.GVA_LOG.Error("Robot == 1")
		utils.FailWithMessageI18n(i18n.MsgWithdrawalFailed, c)
		return
	}
	global.GVA_LOG.Info("CreatePayment2 - User data unmarshaled successfully", zap.Any("user", user))

	global.GVA_LOG.Info("CreatePayment2 - Parsing total amount", zap.String("totalAmount", r.TotalAmount))
	amount, err := strconv.ParseFloat(r.TotalAmount, 64)
	if err != nil {
		global.GVA_LOG.Error("CreatePayment2 - Failed to parse total amount", zap.Error(err), zap.String("totalAmount", r.TotalAmount))
		utils.FailWithMessageI18n(i18n.MsgInvalidAmount, c)
		return
	}
	global.GVA_LOG.Info("CreatePayment2 - Total amount parsed successfully", zap.Float64("amount", amount))

	global.GVA_LOG.Info("CreatePayment2 - Checking user balance", zap.Float64("userBalance", user.Balance), zap.Float64("requestAmount", amount))
	if user.Balance < amount {
		global.GVA_LOG.Warn("CreatePayment2 - Insufficient funds", zap.Float64("userBalance", user.Balance), zap.Float64("requestAmount", amount))
		utils.FailWithMessageI18n(i18n.MsgInsufficientFunds, c)
		return
	}
	global.GVA_LOG.Info("CreatePayment2 - Balance check passed")

	global.GVA_LOG.Info("CreatePayment2 - Attempting to get user withdrawal accounts",
		zap.Uint("userId", uid),
		zap.String("accountId", r.AccountId),
		zap.Any("requestData", r))

	userWithdrawalAccounts, err := userWithdrawalAccountsService.GetUserWithdrawalAccounts(ctx, r.AccountId)

	if err != nil {
		global.GVA_LOG.Error("CreatePayment2 - Failed to get user withdrawal accounts",
			zap.Error(err),
			zap.Uint("userId", uid),
			zap.String("accountId", r.AccountId))
		utils.FailWithMessageI18n(i18n.MsgAccountNotFound, c)
		return
	}

	global.GVA_LOG.Info("CreatePayment2 - Successfully retrieved user withdrawal accounts",
		zap.Uint("userId", uid),
		zap.String("accountId", r.AccountId),
		zap.Any("userWithdrawalAccounts", userWithdrawalAccounts))

	// 判断 BankAcctNo 是否是邮箱格式
	global.GVA_LOG.Info("CreatePayment2 - Checking if account number is email format", zap.String("accountNumber", userWithdrawalAccounts.AccountNumber))
	accEmail := ""
	if strings.Contains(userWithdrawalAccounts.AccountNumber, "@") {
		accEmail = userWithdrawalAccounts.AccountNumber
		global.GVA_LOG.Info("CreatePayment2 - Account number is email format", zap.String("accEmail", accEmail))
	} else {
		global.GVA_LOG.Info("CreatePayment2 - Account number is not email format")
	}

	global.GVA_LOG.Info("CreatePayment2 - Creating cashout request form data")
	formData := payment.CashoutCreateRequest{
		MerNo:         pc.MerchantId,
		CurrencyCode:  "BRL",
		OutTradeNo:    fmt.Sprintf("ORDER_%d%d", time.Now().Unix(), time.Now().UnixNano()%100000),
		TotalAmount:   r.TotalAmount,
		RandomNo:      fmt.Sprintf("%d", time.Now().UnixNano()%100000000000000),
		BankCode:      userWithdrawalAccounts.BankCode,
		BankAcctName:  userWithdrawalAccounts.AccountName,
		BankFirstName: userWithdrawalAccounts.FirstName,
		BankLastName:  userWithdrawalAccounts.LastName,
		BankAcctNo:    userWithdrawalAccounts.AccountNumber,
		AccPhone:      userWithdrawalAccounts.Phone,
		AccEmail:      accEmail,
		NotifyUrl:     "https://api.bzgame777.com/callback/payment2",
		IdentityNo:    userWithdrawalAccounts.CpfNumber,
		IdentityType:  userWithdrawalAccounts.AccountType,
	}
	global.GVA_LOG.Info("CreatePayment2 - Form data created", zap.Any("formData", formData))

	global.GVA_LOG.Info("CreatePayment2 - Calling CreateCashout API")
	code, msg, response := pc.CreateCashout(formData)
	global.GVA_LOG.Info("CreatePayment2 - CreateCashout API response",
		zap.Int("code", code),
		zap.String("msg", msg),
		zap.Any("response", response))

	if code != 0 {
		global.GVA_LOG.Error("CreatePayment2 - CreateCashout API failed", zap.Int("code", code), zap.String("msg", msg))
		utils.FailWithMessageI18n(i18n.MsgWithdrawalFailed, c)
		fmt.Println(msg, response)
		return
	}
	global.GVA_LOG.Info("CreatePayment2 - CreateCashout API succeeded")

	// Convert response.TotalAmount (string) to int safely
	global.GVA_LOG.Info("CreatePayment2 - Parsing response TotalAmount", zap.String("responseTotalAmount", response.TotalAmount))
	amountInt, err := strconv.Atoi(response.TotalAmount)
	if err != nil {
		global.GVA_LOG.Error("CreatePayment2 - Failed to parse TotalAmount", zap.String("TotalAmount", response.TotalAmount), zap.Error(err))
		utils.FailWithMessageI18n(i18n.MsgSystemError, c)
		return
	}
	global.GVA_LOG.Info("CreatePayment2 - TotalAmount parsed successfully", zap.Int("amountInt", amountInt))

	global.GVA_LOG.Info("CreatePayment2 - Creating payment transaction record")
	paymentTransactions := api.PaymentTransactions{
		UserId:          uint(uid),
		MerchantOrderNo: response.OutTradeNo,
		OrderNo:         response.RemitOrderNo,
		TransactionType: 2,
		Amount:          amountInt * 100,
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
	global.GVA_LOG.Info("CreatePayment2 - Payment transaction record created", zap.Any("paymentTransactions", paymentTransactions))

	global.GVA_LOG.Info("CreatePayment2 - Creating payment transaction and deducting user balance in transaction")
	// 使用事务函数同时创建支付记录并扣减用户余额
	transactionCode := fmt.Sprintf("CREATE_PAYMENT2_%d_%d", user.ID, time.Now().Unix())
	err = utils.CreatePaymentWithBalanceDeduction(c, user.ID, amount, "CreatePayment2", func(tx *gorm.DB) error {
		return paymentTransactionsService.CreateWithTx(tx, paymentTransactions)
	}, transactionCode)
	if err != nil {
		global.GVA_LOG.Error("CreatePayment2 - Failed to create payment with balance deduction",
			zap.Error(err),
			zap.Uint("userId", user.ID),
			zap.Float64("amount", amount))

		if err.Error() == "insufficient balance" {
			global.GVA_LOG.Warn("CreatePayment2 - Insufficient balance error")
			utils.FailWithMessageI18n(i18n.MsgInsufficientFunds, c)
		} else {
			global.GVA_LOG.Error("CreatePayment2 - Unknown error during balance deduction")
			utils.FailWithMessageI18n(i18n.MsgSystemError, c)
		}
		return
	}
	global.GVA_LOG.Info("CreatePayment2 - Payment transaction and balance deduction completed successfully")

	global.GVA_LOG.Info("CreatePayment2 - Successfully processed payment",
		zap.Uint("userId", uint(uid)),
		zap.String("merchantOrderNo", response.OutTradeNo),
		zap.String("orderNo", response.RemitOrderNo),
		zap.Int("amount", amountInt),
		zap.String("status", response.ResultCode))

	global.GVA_LOG.Info("CreatePayment2 - Payment process completed successfully")
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
