package api

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/api"
	apiReq "github.com/flipped-aurora/gin-vue-admin/server/model/api/request"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
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
		response.Result(401, nil, "", c)
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

func (sysTransactionsApi *SysTransactionsApi) Betting(c *gin.Context) {
	fmt.Println("Request Headers:")
	for k, v := range c.Request.Header {
		fmt.Printf("%s: %v\n", k, v)
	}

	var r apiReq.Betting

	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(r, utils.ApiBettingVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	uid := utils.GetRedisUserID(c)
	if uid == 0 {
		response.Result(401, nil, "", c)
		return
	}
	var user system.SysUser
	redisuser, _ := global.GVA_REDIS.Get(c, fmt.Sprintf("user_%d", uid)).Result()
	if redisuser == "" {
		response.Result(401, nil, "", c)
		return
	}

	err = json.Unmarshal([]byte(redisuser), &user)
	if err != nil {
		global.GVA_LOG.Error("Failed to unmarshal user data", zap.Error(err))
		response.Result(401, nil, "Failed to get user data", c)
		return
	}
	if float64(r.Coin) > user.Balance {
		response.Result(400, nil, "Insufficient balance", c)
		return
	}
	newBalance := user.Balance - float64(r.Coin)
	user.Balance = newBalance
	updatedUserJson, err := json.Marshal(user)
	if err != nil {
		global.GVA_LOG.Error("Failed to marshal updated user data", zap.Error(err))
		response.Result(500, nil, "Internal server error", c)
		return
	}
	pipe := global.GVA_REDIS.Pipeline()
	pipe.Set(c, fmt.Sprintf("user_%d", uid), string(updatedUserJson), 24*time.Hour)

	balanceUpdate := map[string]interface{}{
		"user_id":    uid,
		"balance":    newBalance,
		"type":       "betting",
		"coin":       r.Coin,
		"room":       r.Room,
		"created_at": time.Now().Format("2006-01-02 15:04:05"),
	}
	balanceUpdateJson, _ := json.Marshal(balanceUpdate)
	pipe.LPush(c, "balance_updates", string(balanceUpdateJson))
	_, err = pipe.Exec(c)
	if err != nil {
		global.GVA_LOG.Error("Failed to update balance in Redis", zap.Error(err))
		response.Result(500, nil, "Failed to update balance", c)
		return
	}
	// 返回成功响应
	response.OkWithMessage("Betting successful", c)
	return
}

// jiesuan
func (sysTransactionsApi *SysTransactionsApi) Win(c *gin.Context) {

	var r apiReq.DecryptRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		response.FailWithMessage("Invalid request format: "+err.Error(), c)
		return
	}

	jsonData, err := json.MarshalIndent(r, "", "    ")
	if err != nil {
		response.FailWithMessage("Failed to process request data", c)
		return
	}

	decrypted, err := utils.CBCDecrypt(string(jsonData))
	if err != nil {
		global.GVA_LOG.Error("Decryption failed", zap.Error(err))
		response.FailWithMessage("Decryption failed: "+err.Error(), c)
		return
	}
	decryptedStr, ok := decrypted.(string)
	if !ok {
		response.FailWithMessage("Decryption result is not a string", c)
		return
	}
	var der apiReq.Settle
	if err := json.Unmarshal([]byte(decryptedStr), &der); err != nil {
		response.FailWithMessage("Failed to unmarshal decrypted data: "+err.Error(), c)
		return
	}

	var lotteryinput utils.LotteryInput
	currentTimestamp := time.Now().Unix()
	lotteryinput.TimeStamp = currentTimestamp

	start := int64(0)
	stop := int64(0)
	results, _ := global.GVA_REDIS.LRange(c, "lottery_results", start, stop).Result()
	if len(results) == 0 {
		lotteryinput.PreviousSeedHash = ""
		global.GVA_REDIS.Expire(c, "lottery_results", 30*24*time.Hour)
	} else {
		lastResult := results[0]
		var record utils.LotteryResult
		if err := json.Unmarshal([]byte(lastResult), &record); err != nil {
			return
		}
		lotteryinput.PreviousSeedHash = record.CurrentSeedHash
	}
	var currentInput utils.LotteryInput
	currentInput = utils.LotteryInput{
		PreviousSeedHash: lotteryinput.PreviousSeedHash,
		TimeStamp:        time.Now().Unix(),
	}
	result, err := utils.GenerateLuckyNumber(currentInput)
	if err != nil {
		fmt.Printf("fail: %v\n", err)
		return
	}
	result.SessionId = der.SessionId
	result.Gid = der.Gid
	recordJson, _ := json.Marshal(result)
	global.GVA_REDIS.LPush(c, "lottery_results", string(recordJson)).Err()

	res, err := utils.CBCEncrypt(result.LuckyNumber)
	if err != nil {
		global.GVA_LOG.Error("CBCEncrypt failed", zap.Error(err))
		response.FailWithMessage("CBCEncrypt failed: "+err.Error(), c)
		return
	}
	response.OkWithDetailed(res, "", c)
}

// jiaoyan
func (sysTransactionsApi *SysTransactionsApi) CheckWin(c *gin.Context) {
	var verifyInput utils.VerifyInput
	err := c.ShouldBindJSON(&verifyInput)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	fmt.Println("verifyInput", verifyInput)
	isValid, _, err := utils.VerifyLottery(verifyInput)
	if err != nil {
		response.FailWithMessage("fail", c)
		return
	}
	if !isValid {
		response.FailWithMessage("fail", c)
		return
	}
	response.OkWithMessage("ok", c)
}
