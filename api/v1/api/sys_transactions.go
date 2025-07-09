package api

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/api"
	apiReq "github.com/flipped-aurora/gin-vue-admin/server/model/api/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"

	"reflect"
	"sort"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	signUtils "github.com/flipped-aurora/gin-vue-admin/server/utils/sign"
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

// Settlement
func (sysTransactionsApi *SysTransactionsApi) Settle(c *gin.Context) {

	var settleList apiReq.SettleList
	err := c.ShouldBindJSON(&settleList)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 验签逻辑
	if settleList.Sign == "" {
		response.FailWithMessage("签名不能为空", c)
		return
	}

	// 构建验签参数（与TypeScript保持一致）
	verifyParams := make(map[string]interface{})

	// 添加timestamp
	if settleList.Timestamp != "" {
		verifyParams["timestamp"] = settleList.Timestamp
	}

	// 添加list（复杂对象会被转换为[object Object]）
	if len(settleList.List) > 0 {
		verifyParams["list"] = settleList.List
	}

	// 调试信息：输出构建的参数
	global.GVA_LOG.Info("验签参数构建完成",
		zap.Any("verifyParams", verifyParams),
		zap.String("receivedSign", settleList.Sign))

	// 使用验签工具类验证签名
	isValid := signUtils.VerifySign(verifyParams, settleList.Sign)

	if !isValid {
		// 生成正确的签名用于调试
		correctSign := signUtils.GenerateSign(verifyParams)

		global.GVA_LOG.Error("签名验证失败",
			zap.String("receivedSign", settleList.Sign),
			zap.String("correctSign", correctSign),
			zap.Any("verifyParams", verifyParams))

		// 返回详细的错误信息，包括签名字符串
		response.FailWithDetailed(gin.H{
			"error":        "签名验证失败",
			"receivedSign": settleList.Sign,
			"correctSign":  correctSign,
			"verifyParams": verifyParams,
			"signString":   getSignString(verifyParams), // 添加签名字符串用于调试
		}, "签名验证失败", c)
		return
	}

	// 处理结算逻辑
	for _, record := range settleList.List {
		fmt.Println("der:", record.Coin)

		redisuser, _ := global.GVA_REDIS.Get(c, fmt.Sprintf("user_%s", record.UserCode)).Result()
		if redisuser == "" {
			response.Result(401, nil, "", c)
			return
		}
		var userJson system.ApiSysUser
		err = json.Unmarshal([]byte(redisuser), &userJson)
		if err != nil {
			global.GVA_LOG.Error("Failed to unmarshal user data", zap.Error(err))
		} else {
			newBalance := userJson.Balance + record.Win
			if newBalance < 0 {
				global.GVA_LOG.Error("<0:",
					zap.String("UserCode", record.UserCode),
					zap.Float64("OriginalBalance", userJson.Balance),
					zap.Float64("WinAmount", record.Win),
					zap.Float64("CalculatedBalance", newBalance),
					zap.String("Username", userJson.Username),
				)
				newBalance = 0
			}
			userJson.Balance = math.Round(newBalance*100) / 100
			updatedUserJson, err := json.Marshal(userJson)
			if err != nil {
				global.GVA_LOG.Error("Failed to marshal updated user data", zap.Error(err))
			} else {
				err = global.GVA_REDIS.Set(c, fmt.Sprintf("user_%s", record.UserCode), string(updatedUserJson), 0).Err()
				if err != nil {
					global.GVA_LOG.Error("Failed to save user data to Redis", zap.Error(err))
				}
			}
		}
	}

	// Save the entire SettleRecords structure to Redis Hash using timestamp as key
	timestamp := time.Now().Unix()
	settleKey := fmt.Sprintf("Settle_%d", timestamp)
	derJson, err := json.Marshal(settleList.List)
	if err != nil {
		global.GVA_LOG.Error("Failed to marshal SettleRecords", zap.Error(err))
	} else {
		// Save data to Redis Hash
		err = global.GVA_REDIS.HSet(c, "Settle_Records", settleKey, string(derJson)).Err()
		if err != nil {
			global.GVA_LOG.Error("Failed to save SettleRecords to Redis", zap.Error(err))
		}
	}

	response.OkWithMessage("ok", c)
}

// Lottery draw
func (sysTransactionsApi *SysTransactionsApi) Lottery(c *gin.Context) {

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

// Verify win
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

// SOL exchange coins, coins exchange SOL
func (sysTransactionsApi *SysTransactionsApi) Exchange(c *gin.Context) {

	var verifyInput apiReq.MonitorTransfer
	err := c.ShouldBindJSON(&verifyInput)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	fmt.Println("verifyInput", verifyInput)

	response.OkWithMessage("ok", c)
}
func (sysTransactionsApi *SysTransactionsApi) Config(c *gin.Context) {

	// Get config field from POST request
	var requestData struct {
		Config string `json:"config"`
	}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		response.FailWithMessage("Request parameter error", c)
		return
	}

	if requestData.Config == "" {
		response.FailWithMessage("Config field cannot be empty", c)
		return
	}

	storedCode, err := global.GVA_REDIS.Get(c, requestData.Config).Result()
	if err != nil {
		response.FailWithMessage("Code error", c)
		return
	}

	// Parse JSON string to map
	var configData map[string]interface{}
	if err := json.Unmarshal([]byte(storedCode), &configData); err != nil {
		response.FailWithMessage("JSON parsing failed", c)
		return
	}

	response.OkWithData(configData, c)
}

// GetSettleListFromRedis Get settlement list data from Redis and process
func (sysTransactionsApi *SysTransactionsApi) GetSettleListFromRedis(c *gin.Context) {
	// Get all data from Settle_Records Hash in Redis
	settleRecords, err := global.GVA_REDIS.HGetAll(c, "Settle_Records").Result()
	if err != nil {
		global.GVA_LOG.Error("Failed to get Settle_Records from Redis", zap.Error(err))
		response.FailWithMessage("Failed to get settlement list: "+err.Error(), c)
		return
	}

	if len(settleRecords) == 0 {
		response.OkWithDetailed(gin.H{
			"list":  []interface{}{},
			"total": 0,
		}, "No settlement data", c)
		return
	}

	// Process each settlement record
	var processedList []map[string]interface{}
	var processedKeys []string
	var totalRecordsProcessed int
	var totalRebateProcessed int

	for key, settleData := range settleRecords {
		var settleRecords []apiReq.SettleRecord
		err := json.Unmarshal([]byte(settleData), &settleRecords)
		if err != nil {
			global.GVA_LOG.Error("Failed to unmarshal settle record",
				zap.Error(err),
				zap.String("data", settleData),
				zap.String("key", key))
			continue
		}

		// Process each settlement record
		for _, record := range settleRecords {
			// Get user invitation relationship and process rebate
			processUserRebate(c, record)
			totalRebateProcessed++

			processedRecord := map[string]interface{}{
				"UserCode": record.UserCode,
				"Win":      record.Win,
				"Coin":     record.Coin,
				"Key":      key,
			}
			processedList = append(processedList, processedRecord)
			totalRecordsProcessed++
		}

		// Record processed keys
		processedKeys = append(processedKeys, key)
	}

	// Delete processed data
	if len(processedKeys) > 0 {
		err = global.GVA_REDIS.HDel(c, "Settle_Records", processedKeys...).Err()
		if err != nil {
			global.GVA_LOG.Error("Failed to delete processed Settle_Records from Redis",
				zap.Error(err),
				zap.Strings("keysToDelete", processedKeys))
		}
	}

	response.OkWithDetailed(gin.H{
		"list":           processedList,
		"total":          len(processedList),
		"processed_keys": processedKeys,
	}, "Get settlement list successfully, processed data cleaned", c)
}

// GetUserInvitationRelation Get user invitation relationship API
func (sysTransactionsApi *SysTransactionsApi) GetUserInvitationRelation(c *gin.Context) {
	// Get user ID from request parameters
	userIdStr := c.Query("userId")
	if userIdStr == "" {
		response.FailWithMessage("User ID cannot be empty", c)
		return
	}

	userId, err := strconv.ParseUint(userIdStr, 10, 32)
	if err != nil {
		response.FailWithMessage("User ID format error", c)
		return
	}

	// Get invitation relationship
	relation, err := getUserInvitationRelation(c, uint(userId))
	if err != nil {
		global.GVA_LOG.Error("Failed to get user invitation relation",
			zap.Error(err),
			zap.Uint64("userId", userId))
		response.FailWithMessage("Failed to get invitation relationship", c)
		return
	}

	// Get superior user detailed information
	var result map[string]interface{}
	if relation != nil {
		result = make(map[string]interface{})

		// Get level 1 superior information
		if level1Id, ok := relation["level1"].(float64); ok && level1Id > 0 {
			level1User, err := getUserFromRedis(c, int(level1Id))
			if err == nil {
				result["level1_user"] = level1User
			}
		}

		// Get level 2 superior information
		if level2Id, ok := relation["level2"].(float64); ok && level2Id > 0 {
			level2User, err := getUserFromRedis(c, int(level2Id))
			if err == nil {
				result["level2_user"] = level2User
			}
		}

		result["relation"] = relation
	}

	response.OkWithDetailed(result, "Get invitation relationship successfully", c)
}

// processUserRebate Process user rebate
func processUserRebate(c *gin.Context, record apiReq.SettleRecord) {
	// Convert UserCode to user ID
	userId, err := strconv.ParseUint(record.UserCode, 10, 32)
	if err != nil {
		global.GVA_LOG.Error("Failed to convert user ID",
			zap.Error(err),
			zap.String("UserCode", record.UserCode))
		return
	}

	// Get user invitation relationship
	relation, err := getUserInvitationRelation(c, uint(userId))
	if err != nil {
		global.GVA_LOG.Error("Failed to get user invitation relationship",
			zap.Error(err),
			zap.Uint64("userId", userId),
			zap.String("userCode", record.UserCode))
		return
	}

	if relation == nil {
		return
	}

	// Process level 1 superior rebate
	if level1Id, ok := relation["level1"].(float64); ok && level1Id > 0 {
		// Get level 1 superior rebate rate
		rebateRate1, err := getUserRebateRate(c, int(level1Id))
		if err != nil {
			global.GVA_LOG.Error("Failed to get level 1 superior rebate rate",
				zap.Error(err),
				zap.Int("level1UserId", int(level1Id)))
			return
		}

		// Check if rebate rate is greater than 0
		if rebateRate1 > 0 {
			// Calculate level 1 rebate amount
			rawRebateAmount1 := record.Coin * rebateRate1
			// Round to 2 decimal places
			rebateAmount1 := math.Round(rawRebateAmount1*100) / 100

			addRebateToUser(c, int(level1Id), rebateAmount1, "Level 1 Rebate", record, rebateRate1, 1)
		}
	}

	// Process level 2 superior rebate
	if level2Id, ok := relation["level2"].(float64); ok && level2Id > 0 {
		// Get level 2 superior rebate rate
		rebateRate2, err := getUserRebateRate(c, int(level2Id))
		if err != nil {
			global.GVA_LOG.Error("Failed to get level 2 superior rebate rate",
				zap.Error(err),
				zap.Int("level2UserId", int(level2Id)))
			return
		}

		// Check if rebate rate is greater than 0
		if rebateRate2 > 0 {
			// Calculate level 2 rebate amount
			rawRebateAmount2 := record.Coin * rebateRate2
			// Round to 2 decimal places
			rebateAmount2 := math.Round(rawRebateAmount2*100) / 100

			addRebateToUser(c, int(level2Id), rebateAmount2, "Level 2 Rebate", record, rebateRate2, 2)
		}
	}
}

// getUserInvitationRelation Get user invitation relationship
func getUserInvitationRelation(c *gin.Context, userId uint) (map[string]interface{}, error) {
	key := fmt.Sprintf("invitation_relation_%d", userId)

	result, err := global.GVA_REDIS.Get(c, key).Result()
	if err != nil {
		global.GVA_LOG.Error("Failed to get invitation relationship from Redis",
			zap.Error(err),
			zap.Uint("userId", userId),
			zap.String("redisKey", key))
		return nil, err
	}

	var invitationData map[string]interface{}
	err = json.Unmarshal([]byte(result), &invitationData)
	if err != nil {
		global.GVA_LOG.Error("Failed to parse invitation relationship JSON data",
			zap.Error(err),
			zap.Uint("userId", userId),
			zap.String("rawData", result))
		return nil, err
	}

	return invitationData, nil
}

// addRebateToUser Add rebate to user
func addRebateToUser(c *gin.Context, userId int, rebateAmount float64, rebateType string, record apiReq.SettleRecord, rebateRate float64, rebateLevel int) {
	// Get user information
	redisKey := fmt.Sprintf("user_%d", userId)

	redisuser, err := global.GVA_REDIS.Get(c, redisKey).Result()
	if err != nil {
		global.GVA_LOG.Error("Failed to get user from Redis",
			zap.Error(err),
			zap.Int("userId", userId),
			zap.String("redisKey", redisKey))
		return
	}

	var userJson system.ApiSysUser
	err = json.Unmarshal([]byte(redisuser), &userJson)
	if err != nil {
		global.GVA_LOG.Error("Failed to unmarshal user data",
			zap.Error(err),
			zap.Int("userId", userId),
			zap.String("redisData", redisuser))
		return
	}

	// Calculate new balance
	originalBalance := userJson.Balance

	// Round rebate amount to 2 decimal places
	roundedRebateAmount := math.Round(rebateAmount*100) / 100

	// Calculate new balance
	newBalance := userJson.Balance + roundedRebateAmount

	if newBalance < 0 {
		global.GVA_LOG.Error("Rebate would result in negative balance, set to 0",
			zap.Int("userId", userId),
			zap.Float64("originalBalance", originalBalance),
			zap.Float64("rebateAmount", roundedRebateAmount),
			zap.String("rebateType", rebateType))
		newBalance = 0
	}

	// Final balance rounded to 2 decimal places
	finalBalance := math.Round(newBalance*100) / 100
	userJson.Balance = finalBalance

	// Update user information to Redis
	updatedUserJson, err := json.Marshal(userJson)
	if err != nil {
		global.GVA_LOG.Error("Failed to marshal updated user data",
			zap.Error(err),
			zap.Int("userId", userId))
		return
	}

	err = global.GVA_REDIS.Set(c, redisKey, string(updatedUserJson), 0).Err()
	if err != nil {
		global.GVA_LOG.Error("Failed to save user data to Redis",
			zap.Error(err),
			zap.Int("userId", userId),
			zap.String("redisKey", redisKey))
		return
	}

	// 保存返佣记录到数据库
	saveRebateRecordToDB(c, userId, record, rebateType, rebateLevel, rebateRate, rebateAmount, originalBalance, userJson.Balance)
}

// getUserFromRedis Get user information from Redis
func getUserFromRedis(c *gin.Context, userId int) (*system.ApiSysUser, error) {
	redisuser, err := global.GVA_REDIS.Get(c, fmt.Sprintf("user_%d", userId)).Result()
	if err != nil {
		return nil, err
	}

	var userJson system.ApiSysUser
	err = json.Unmarshal([]byte(redisuser), &userJson)
	if err != nil {
		return nil, err
	}

	return &userJson, nil
}

// getUserRebateRate Get rebate rate from user Redis information
func getUserRebateRate(c *gin.Context, userId int) (float64, error) {
	// Get user information
	user, err := getUserFromRedis(c, userId)
	if err != nil {
		global.GVA_LOG.Error("Failed to get user information",
			zap.Error(err),
			zap.Int("userId", userId))
		return 0, err
	}

	// Check if user level is 0, if so no rebate
	if user.Level == 0 {
		return 0, nil
	}

	// Get rebate rate from user's level field
	// level field stores percentage value (e.g., 3 means 3%, 5 means 5%)
	rebateRate := float64(user.Level) / 100.0 // level=3 means 3%, so level/100

	// Boundary check: ensure rebate rate is within reasonable range (0-100%)
	if rebateRate < 0 {
		global.GVA_LOG.Warn("Rebate rate is negative, set to 0",
			zap.Int("userId", userId),
			zap.Int("userLevel", user.Level),
			zap.Float64("originalRate", rebateRate))
		rebateRate = 0
	} else if rebateRate > 1.0 {
		global.GVA_LOG.Warn("Rebate rate exceeds 100%, set to 100%",
			zap.Int("userId", userId),
			zap.Int("userLevel", user.Level),
			zap.Float64("originalRate", rebateRate))
		rebateRate = 1.0
	}

	return rebateRate, nil
}

// saveRebateRecordToDB Save rebate record to database
func saveRebateRecordToDB(c *gin.Context, userId int, record apiReq.SettleRecord, rebateType string, rebateLevel int, rebateRate float64, rebateAmount float64, balanceBefore float64, balanceAfter float64) {
	// Convert UserCode to user ID
	fromUserId, err := strconv.ParseUint(record.UserCode, 10, 32)
	if err != nil {
		global.GVA_LOG.Error("Failed to convert fromUserId",
			zap.Error(err),
			zap.String("userCode", record.UserCode))
		return
	}

	// Convert BetInfo to JSON
	betInfoJSON, err := json.Marshal(record.BetInfo)
	if err != nil {
		global.GVA_LOG.Error("Failed to serialize BetInfo",
			zap.Error(err),
			zap.Any("betInfo", record.BetInfo))
		return
	}

	// Create rebate record
	rebateRecord := api.UserRebates{
		UserId:            uint(userId),
		FromUserId:        uint(fromUserId),
		FromUserCode:      record.UserCode,
		RebateType:        rebateType,
		RebateLevel:       rebateLevel,
		Coin:              record.Coin,
		Win:               record.Win,
		RebateRate:        rebateRate,
		RebateAmount:      rebateAmount,
		UserBalanceBefore: balanceBefore,
		UserBalanceAfter:  balanceAfter,
		SessionId:         record.SessionID,
		GameType:          record.GameType,
		Area:              record.Area,
		BetInfo:           betInfoJSON,
		Status:            1, // 1-Success
		Remark:            fmt.Sprintf("User %s bet %.2f, got %.2f%% rebate", record.UserCode, record.Coin, rebateRate*100),
	}

	// Save to database
	err = global.GVA_DB.Create(&rebateRecord).Error
	if err != nil {
		global.GVA_LOG.Error("Failed to save rebate record to database",
			zap.Error(err),
			zap.Int("userId", userId),
			zap.String("rebateType", rebateType),
			zap.String("errorType", fmt.Sprintf("%T", err)),
			zap.String("errorDetails", err.Error()))
		return
	}
}

// getMapKeys Get all keys from map
func getMapKeys(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// VerifySettleSign 验证结算签名
// @Tags SysTransactions
// @Summary 验证结算签名
// @Accept application/json
// @Produce application/json
// @Param data body apiReq.SettleList true "结算数据"
// @Success 200 {object} response.Response{msg=string} ""
// @Router /sysTransactions/verifySettleSign [post]
func (sysTransactionsApi *SysTransactionsApi) VerifySettleSign(c *gin.Context) {
	var settleList apiReq.SettleList
	err := c.ShouldBindJSON(&settleList)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 验签逻辑
	if settleList.Sign == "" {
		response.FailWithMessage("签名不能为空", c)
		return
	}

	// 构建验签参数（与TypeScript保持一致）
	verifyParams := make(map[string]interface{})

	// 添加timestamp
	if settleList.Timestamp != "" {
		verifyParams["timestamp"] = settleList.Timestamp
	}

	// 添加list（复杂对象会被转换为[object Object]）
	if len(settleList.List) > 0 {
		verifyParams["list"] = settleList.List
	}

	// 调试信息：输出构建的参数
	global.GVA_LOG.Info("验签参数构建完成",
		zap.Any("verifyParams", verifyParams),
		zap.String("receivedSign", settleList.Sign))

	// 使用验签工具类验证签名
	isValid := signUtils.VerifySign(verifyParams, settleList.Sign)

	if !isValid {
		// 生成正确的签名用于调试
		correctSign := signUtils.GenerateSign(verifyParams)

		global.GVA_LOG.Error("签名验证失败",
			zap.String("receivedSign", settleList.Sign),
			zap.String("correctSign", correctSign),
			zap.Any("verifyParams", verifyParams))

		// 返回详细的错误信息，包括签名字符串
		response.FailWithDetailed(gin.H{
			"error":        "签名验证失败",
			"receivedSign": settleList.Sign,
			"correctSign":  correctSign,
			"verifyParams": verifyParams,
			"signString":   getSignString(verifyParams), // 添加签名字符串用于调试
		}, "签名验证失败", c)
		return
	}

	global.GVA_LOG.Info("签名验证成功",
		zap.String("sign", settleList.Sign),
		zap.Any("params", verifyParams))

	response.OkWithMessage("验签成功", c)
}

// getSignString 获取签名字符串用于调试
func getSignString(params map[string]interface{}) string {
	// 复制验签工具的逻辑来生成签名字符串

	// 1. 按照ASCII码排序参数
	var keys []string
	for key := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	// 2. 拼接参数字符串
	var signStr strings.Builder
	for _, key := range keys {
		value := params[key]
		if value != nil && value != "" {
			// 对于复杂对象，使用[object Object]格式（与TypeScript保持一致）
			var valueStr string
			if isComplexValueForDebug(value) {
				valueStr = "[object Object]"
			} else {
				valueStr = fmt.Sprintf("%v", value)
			}
			signStr.WriteString(fmt.Sprintf("%s=%s&", key, valueStr))
		}
	}

	// 3. 加上签名key
	signStr.WriteString(fmt.Sprintf("key=%s", "GAME_2025_SIGN_KEY_8F7E6D5C4B3A2918_9A8B7C6D5E4F3210"))

	return signStr.String()
}

// isComplexValueForDebug 判断是否为复杂值（用于调试）
func isComplexValueForDebug(value interface{}) bool {
	switch v := value.(type) {
	case []interface{}, map[string]interface{}:
		return true
	default:
		// 使用反射检查是否为切片或结构体
		val := reflect.ValueOf(v)
		if val.Kind() == reflect.Slice || val.Kind() == reflect.Struct {
			return true
		}
		return false
	}
}
