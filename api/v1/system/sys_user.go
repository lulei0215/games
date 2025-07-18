package system

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"strconv"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/api"
	apiReq "github.com/flipped-aurora/gin-vue-admin/server/model/api/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	systemRes "github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/i18n"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gopkg.in/gomail.v2"
)

// Login
// @Tags     Base
// @Summary
// @Produce   application/json
// @Param    data  body      systemReq.Login                                                true  ", , "
// @Success  200   {object}  response.Response{data=systemRes.LoginResponse,msg=string}  ",token,"
// @Router   /base/login [post]
// func (b *BaseApi) Login(c *gin.Context) {
// 	var l systemReq.Login
// 	err := c.ShouldBindJSON(&l)
// 	key := c.ClientIP()

// 	if err != nil {
// 		response.FailWithMessage(err.Error(), c)
// 		return
// 	}
// 	err = utils.Verify(l, utils.LoginVerify)
// 	if err != nil {
// 		response.FailWithMessage(err.Error(), c)
// 		return
// 	}

// 	//
// 	openCaptcha := global.GVA_CONFIG.Captcha.OpenCaptcha               //
// 	openCaptchaTimeOut := global.GVA_CONFIG.Captcha.OpenCaptchaTimeOut //
// 	v, ok := global.BlackCache.Get(key)
// 	if !ok {
// 		global.BlackCache.Set(key, 1, time.Second*time.Duration(openCaptchaTimeOut))
// 	}

// 	var oc bool = false
// 	if openCaptcha > 0 {
// 		oc = true
// 	} else if ok {
// 		if count, ok := v.(int); ok && count > 5 {
// 			oc = true
// 		}
// 	}

// 	if !oc || (l.CaptchaId != "" && l.Captcha != "" && store.Verify(l.CaptchaId, l.Captcha, true)) {
// 		u := &system.SysUser{Username: l.Username, Password: l.Password}
// 		user, err := userService.Login(u)
// 		if err != nil {
// 			global.GVA_LOG.Error("! !", zap.Error(err))
// 			// +1
// 			global.BlackCache.Increment(key, 1)
// 			response.FailWithMessage("", c)
// 			return
// 		}
// 		if user.Enable != 1 {
// 			global.GVA_LOG.Error("! !")
// 			// +1
// 			global.BlackCache.Increment(key, 1)
// 			response.FailWithMessage("", c)
// 			return
// 		}

// 		// 检查Redis中是否有用户数据
// 		redisKey := fmt.Sprintf("user_%d", user.ID)
// 		redisUser, err := global.GVA_REDIS.Get(c, redisKey).Result()
// 		fmt.Println("redisUser", redisUser)
// 		// if err == nil && redisUser != "" {
// 		// 	// Redis中有用户数据，直接返回
// 		// 	var cachedUser system.SysUser
// 		// 	err = json.Unmarshal([]byte(redisUser), &cachedUser)
// 		// 	if err == nil {
// 		// 		global.GVA_LOG.Info("Using cached user data from Redis",
// 		// 			zap.Uint("userId", user.ID),
// 		// 			zap.String("username", user.Username))
// 		// 		b.TokenNext(c, cachedUser)
// 		// 		response.FailWithMessage("", c)
// 		// 	} else {
// 		// 		global.GVA_LOG.Error("Failed to unmarshal cached user data",
// 		// 			zap.Error(err),
// 		// 			zap.Uint("userId", user.ID))
// 		// 	}
// 		// }

//			// Redis中没有数据或解析失败，使用数据库数据
//			global.GVA_LOG.Info("Using database user data",
//				zap.Uint("userId", user.ID),
//				zap.String("username", user.Username))
//			b.TokenNext(c, *user)
//		}
//		// +1
//		global.BlackCache.Increment(key, 1)
//		response.FailWithMessage("", c)
//	}
func (b *BaseApi) Login(c *gin.Context) {
	var l systemReq.Login
	err := c.ShouldBindJSON(&l)
	key := c.ClientIP()

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(l, utils.LoginVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 判断验证码是否开启
	openCaptcha := global.GVA_CONFIG.Captcha.OpenCaptcha               // 是否开启防爆次数
	openCaptchaTimeOut := global.GVA_CONFIG.Captcha.OpenCaptchaTimeOut // 缓存超时时间
	v, ok := global.BlackCache.Get(key)
	if !ok {
		global.BlackCache.Set(key, 1, time.Second*time.Duration(openCaptchaTimeOut))
	}

	var oc bool = false
	if openCaptcha > 0 {
		oc = true
	} else if ok {
		if count, ok := v.(int); ok && count > 5 {
			oc = true
		}
	}

	if !oc || (l.CaptchaId != "" && l.Captcha != "" && store.Verify(l.CaptchaId, l.Captcha, true)) {
		u := &system.SysUser{Username: l.Username, Password: l.Password}
		user, err := userService.Login(u)
		if err != nil {
			global.GVA_LOG.Error("登陆失败! 用户名不存在或者密码错误!", zap.Error(err))
			// 验证码次数+1
			global.BlackCache.Increment(key, 1)
			response.FailWithMessage("用户名不存在或者密码错误", c)
			return
		}
		if user.Enable != 1 {
			global.GVA_LOG.Error("登陆失败! 用户被禁止登录!")
			// 验证码次数+1
			global.BlackCache.Increment(key, 1)
			response.FailWithMessage("用户被禁止登录", c)
			return
		}
		b.TokenNext(c, *user)
		return
	}
	// 验证码次数+1
	global.BlackCache.Increment(key, 1)
	response.FailWithMessage("验证码错误", c)
}
func (b *BaseApi) Dashboard(c *gin.Context) {
	// 获取今日开始和结束时间
	now := time.Now()
	todayStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	todayEnd := todayStart.Add(24 * time.Hour)

	// 统计数据
	var dashboardData = make(map[string]interface{})

	// 1. 今日注册人数
	var todayRegisterCount int64
	err := global.GVA_DB.Model(&system.SysUser{}).
		Where("created_at >= ? AND created_at < ?", todayStart, todayEnd).
		Count(&todayRegisterCount).Error
	if err != nil {
		global.GVA_LOG.Error("Failed to get today register count", zap.Error(err))
		todayRegisterCount = 0
	}

	// 2. 今日充值人数和总额
	var todayRechargeCount int64
	var todayRechargeAmount float64
	err = global.GVA_DB.Model(&api.PaymentTransactions{}).
		Where("created_at >= ? AND created_at < ? AND transaction_type = ? AND status IN (?)",
			todayStart, todayEnd, 1, []string{"PAID", "SUCCESS", "COMPLETED"}).
		Count(&todayRechargeCount).Error
	if err != nil {
		global.GVA_LOG.Error("Failed to get today recharge count", zap.Error(err))
		todayRechargeCount = 0
	}

	err = global.GVA_DB.Model(&api.PaymentTransactions{}).
		Where("created_at >= ? AND created_at < ? AND transaction_type = ? AND status IN (?)",
			todayStart, todayEnd, 1, []string{"PAID", "SUCCESS", "COMPLETED"}).
		Select("COALESCE(SUM(amount), 0)").
		Scan(&todayRechargeAmount).Error
	if err != nil {
		global.GVA_LOG.Error("Failed to get today recharge amount", zap.Error(err))
		todayRechargeAmount = 0
	}
	// 除以100并保留2位小数
	todayRechargeAmount = float64(int64(todayRechargeAmount/100*100)) / 100

	// 3. 今日提现申请数量和金额
	var todayWithdrawCount int64
	var todayWithdrawAmount float64
	err = global.GVA_DB.Model(&api.PaymentTransactions{}).
		Where("created_at >= ? AND created_at < ? AND transaction_type = ?", todayStart, todayEnd, 2).
		Count(&todayWithdrawCount).Error
	if err != nil {
		global.GVA_LOG.Error("Failed to get today withdraw count", zap.Error(err))
		todayWithdrawCount = 0
	}

	err = global.GVA_DB.Model(&api.PaymentTransactions{}).
		Where("created_at >= ? AND created_at < ? AND transaction_type = ?", todayStart, todayEnd, 2).
		Select("COALESCE(SUM(amount), 0)").
		Scan(&todayWithdrawAmount).Error
	if err != nil {
		global.GVA_LOG.Error("Failed to get today withdraw amount", zap.Error(err))
		todayWithdrawAmount = 0
	}
	// 除以100并保留2位小数
	todayWithdrawAmount = float64(int64(todayWithdrawAmount/100*100)) / 100

	// 4. 总用户数
	var totalUserCount int64
	err = global.GVA_DB.Model(&system.SysUser{}).Count(&totalUserCount).Error
	if err != nil {
		global.GVA_LOG.Error("Failed to get total user count", zap.Error(err))
		totalUserCount = 0
	}

	// 5. 总充值金额
	var totalRechargeAmount float64
	err = global.GVA_DB.Model(&api.PaymentTransactions{}).
		Where("transaction_type = ? AND status IN (?)", 1, []string{"PAID", "SUCCESS", "COMPLETED"}).
		Select("COALESCE(SUM(amount), 0)").
		Scan(&totalRechargeAmount).Error
	if err != nil {
		global.GVA_LOG.Error("Failed to get total recharge amount", zap.Error(err))
		totalRechargeAmount = 0
	}
	// 除以100并保留2位小数
	totalRechargeAmount = float64(int64(totalRechargeAmount/100*100)) / 100

	// 6. 总提现金额
	var totalWithdrawAmount float64
	err = global.GVA_DB.Model(&api.PaymentTransactions{}).
		Where("transaction_type = ? AND status IN (?)", 2, []string{"SUCCESS", "COMPLETED"}).
		Select("COALESCE(SUM(amount), 0)").
		Scan(&totalWithdrawAmount).Error
	if err != nil {
		global.GVA_LOG.Error("Failed to get total withdraw amount", zap.Error(err))
		totalWithdrawAmount = 0
	}
	// 除以100并保留2位小数
	totalWithdrawAmount = float64(int64(totalWithdrawAmount/100*100)) / 100

	// 7. 今日返利总额
	var todayRebateAmount float64
	err = global.GVA_DB.Model(&api.UserRebates{}).
		Where("created_at >= ? AND created_at < ?", todayStart, todayEnd).
		Select("COALESCE(SUM(rebate_amount), 0)").
		Scan(&todayRebateAmount).Error
	if err != nil {
		global.GVA_LOG.Error("Failed to get today rebate amount", zap.Error(err))
		todayRebateAmount = 0
	}
	// 除以100并保留2位小数
	todayRebateAmount = float64(int64(todayRebateAmount/100*100)) / 100

	// 8. 总返利金额
	var totalRebateAmount float64
	err = global.GVA_DB.Model(&api.UserRebates{}).
		Select("COALESCE(SUM(rebate_amount), 0)").
		Scan(&totalRebateAmount).Error
	if err != nil {
		global.GVA_LOG.Error("Failed to get total rebate amount", zap.Error(err))
		totalRebateAmount = 0
	}
	// 除以100并保留2位小数
	totalRebateAmount = float64(int64(totalRebateAmount/100*100)) / 100

	// 9. 今日活跃用户数（有登录记录的用户）
	var todayActiveUsers int64
	err = global.GVA_DB.Model(&system.SysUser{}).
		Where("updated_at >= ? AND updated_at < ?", todayStart, todayEnd).
		Count(&todayActiveUsers).Error
	if err != nil {
		global.GVA_LOG.Error("Failed to get today active users", zap.Error(err))
		todayActiveUsers = 0
	}

	// 10. 待审核提现数量
	var pendingWithdrawCount int64
	err = global.GVA_DB.Model(&api.PaymentTransactions{}).
		Where("transaction_type = ? AND status IN (?)", 2, []string{"WAITING_PAY", "PAYING"}).
		Count(&pendingWithdrawCount).Error
	if err != nil {
		global.GVA_LOG.Error("Failed to get pending withdraw count", zap.Error(err))
		pendingWithdrawCount = 0
	}

	// 11. 待审核提现金额
	var pendingWithdrawAmount float64
	err = global.GVA_DB.Model(&api.PaymentTransactions{}).
		Where("transaction_type = ? AND status IN (?)", 2, []string{"WAITING_PAY", "PAYING"}).
		Select("COALESCE(SUM(amount), 0)").
		Scan(&pendingWithdrawAmount).Error
	if err != nil {
		global.GVA_LOG.Error("Failed to get pending withdraw amount", zap.Error(err))
		pendingWithdrawAmount = 0
	}
	// 除以100并保留2位小数
	pendingWithdrawAmount = float64(int64(pendingWithdrawAmount/100*100)) / 100

	// 12. 今日充值失败数量和金额
	var todayRechargeFailedCount int64
	var todayRechargeFailedAmount float64
	err = global.GVA_DB.Model(&api.PaymentTransactions{}).
		Where("created_at >= ? AND created_at < ? AND transaction_type = ? AND status = ?",
			todayStart, todayEnd, 1, "PAY_FAILED").
		Count(&todayRechargeFailedCount).Error
	if err != nil {
		global.GVA_LOG.Error("Failed to get today recharge failed count", zap.Error(err))
		todayRechargeFailedCount = 0
	}

	err = global.GVA_DB.Model(&api.PaymentTransactions{}).
		Where("created_at >= ? AND created_at < ? AND transaction_type = ? AND status = ?",
			todayStart, todayEnd, 1, "PAY_FAILED").
		Select("COALESCE(SUM(amount), 0)").
		Scan(&todayRechargeFailedAmount).Error
	if err != nil {
		global.GVA_LOG.Error("Failed to get today recharge failed amount", zap.Error(err))
		todayRechargeFailedAmount = 0
	}
	// 除以100并保留2位小数
	todayRechargeFailedAmount = float64(int64(todayRechargeFailedAmount/100*100)) / 100

	// 13. 今日提现失败数量和金额
	var todayWithdrawFailedCount int64
	var todayWithdrawFailedAmount float64
	err = global.GVA_DB.Model(&api.PaymentTransactions{}).
		Where("created_at >= ? AND created_at < ? AND transaction_type = ? AND status IN (?)",
			todayStart, todayEnd, 2, []string{"FAILED", "REJECTED"}).
		Count(&todayWithdrawFailedCount).Error
	if err != nil {
		global.GVA_LOG.Error("Failed to get today withdraw failed count", zap.Error(err))
		todayWithdrawFailedCount = 0
	}

	err = global.GVA_DB.Model(&api.PaymentTransactions{}).
		Where("created_at >= ? AND created_at < ? AND transaction_type = ? AND status IN (?)",
			todayStart, todayEnd, 2, []string{"FAILED", "REJECTED"}).
		Select("COALESCE(SUM(amount), 0)").
		Scan(&todayWithdrawFailedAmount).Error
	if err != nil {
		global.GVA_LOG.Error("Failed to get today withdraw failed amount", zap.Error(err))
		todayWithdrawFailedAmount = 0
	}
	// 除以100并保留2位小数
	todayWithdrawFailedAmount = float64(int64(todayWithdrawFailedAmount/100*100)) / 100

	// 组装返回数据
	dashboardData = map[string]interface{}{
		"today_register_count":         todayRegisterCount,        // 今日注册人数
		"today_recharge_count":         todayRechargeCount,        // 今日充值人数
		"today_recharge_amount":        todayRechargeAmount,       // 今日充值总额
		"today_recharge_failed_count":  todayRechargeFailedCount,  // 今日充值失败数量
		"today_recharge_failed_amount": todayRechargeFailedAmount, // 今日充值失败金额
		"today_withdraw_count":         todayWithdrawCount,        // 今日提现申请数量
		"today_withdraw_amount":        todayWithdrawAmount,       // 今日提现申请金额
		"today_withdraw_failed_count":  todayWithdrawFailedCount,  // 今日提现失败数量
		"today_withdraw_failed_amount": todayWithdrawFailedAmount, // 今日提现失败金额
		"total_user_count":             totalUserCount,            // 总用户数
		"total_recharge_amount":        totalRechargeAmount,       // 总充值金额
		"total_withdraw_amount":        totalWithdrawAmount,       // 总提现金额
		"today_rebate_amount":          todayRebateAmount,         // 今日返利总额
		"total_rebate_amount":          totalRebateAmount,         // 总返利金额
		"today_active_users":           todayActiveUsers,          // 今日活跃用户数
		"pending_withdraw_count":       pendingWithdrawCount,      // 待审核提现数量
		"pending_withdraw_amount":      pendingWithdrawAmount,     // 待审核提现金额
		"today_start":                  todayStart.Format("2006-01-02 15:04:05"),
		"today_end":                    todayEnd.Format("2006-01-02 15:04:05"),
	}

	// 记录统计信息
	global.GVA_LOG.Info("Dashboard statistics generated",
		zap.Int64("today_register_count", todayRegisterCount),
		zap.Int64("today_recharge_count", todayRechargeCount),
		zap.Float64("today_recharge_amount", todayRechargeAmount),
		zap.Int64("today_recharge_failed_count", todayRechargeFailedCount),
		zap.Float64("today_recharge_failed_amount", todayRechargeFailedAmount),
		zap.Int64("today_withdraw_count", todayWithdrawCount),
		zap.Float64("today_withdraw_amount", todayWithdrawAmount),
		zap.Int64("today_withdraw_failed_count", todayWithdrawFailedCount),
		zap.Float64("today_withdraw_failed_amount", todayWithdrawFailedAmount),
		zap.Int64("total_user_count", totalUserCount),
		zap.Float64("total_recharge_amount", totalRechargeAmount),
		zap.Float64("total_withdraw_amount", totalWithdrawAmount),
		zap.Float64("today_rebate_amount", todayRebateAmount),
		zap.Float64("total_rebate_amount", totalRebateAmount),
		zap.Int64("today_active_users", todayActiveUsers),
		zap.Int64("pending_withdraw_count", pendingWithdrawCount),
		zap.Float64("pending_withdraw_amount", pendingWithdrawAmount),
	)

	response.OkWithDetailed(dashboardData, "获取统计数据成功", c)
}

// TokenNext jwt
func (b *BaseApi) TokenNext(c *gin.Context, user system.SysUser) {
	j := &utils.JWT{SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey)} //
	claims := j.CreateClaims(systemReq.BaseClaims{
		UUID:        user.UUID,
		ID:          uint(user.ID),
		NickName:    user.NickName,
		Username:    user.Username,
		AuthorityId: user.AuthorityId,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		global.GVA_LOG.Error("token!", zap.Error(err))
		response.FailWithMessage("token", c)
		return
	}
	if global.GVA_CONFIG.System.UseMultipoint {
		if err := utils.SetRedisJWT(token, user.Username); err != nil {
			global.GVA_LOG.Error("!", zap.Error(err))
			response.FailWithMessage("", c)
			return
		}
	}
	// 由于JWT永不过期，设置一个很长的过期时间
	utils.SetToken(c, token, 365*24*60*60) // 1年

	// 修复空指针异常：检查ExpiresAt是否为nil
	var expiresAt int64
	if claims.RegisteredClaims.ExpiresAt != nil {
		expiresAt = claims.RegisteredClaims.ExpiresAt.Unix() * 1000
	} else {
		// 如果ExpiresAt为nil，设置一个默认的过期时间（1年后）
		expiresAt = time.Now().AddDate(1, 0, 0).Unix() * 1000
	}

	response.OkWithDetailed(systemRes.LoginResponse{
		User:      user,
		Token:     token,
		ExpiresAt: expiresAt,
	}, "ok", c)

	users, _ := global.GVA_REDIS.Get(c, fmt.Sprintf("user_%d", user.ID)).Result()
	if users == "" {
		// 将用户数据序列化为JSON
		userJson, err := json.Marshal(user)
		if err != nil {
			global.GVA_LOG.Error("Failed to marshal user data", zap.Error(err))
		} else {
			// 保存到Redis，设置过期时间为24小时
			err = global.GVA_REDIS.Set(c, fmt.Sprintf("user_%d", user.ID), string(userJson), 0).Err()
			if err != nil {
				global.GVA_LOG.Error("Failed to save user data to Redis", zap.Error(err))
			}
		}
	} else {
		// 用户数据已存在，可以在这里添加其他处理逻辑
	}
}

// Register
// @Tags     SysUser
// @Summary
// @Produce   application/json
// @Param    data  body      systemReq.Register                                            true  ", , , ID"
// @Success  200   {object}  response.Response{data=systemRes.SysUserResponse,msg=string}  ","
// @Router   /user/admin_register [post]
func (b *BaseApi) Register(c *gin.Context) {
	var r systemReq.Register
	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(r, utils.RegisterVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var authorities []system.SysAuthority
	for _, v := range r.AuthorityIds {
		authorities = append(authorities, system.SysAuthority{
			AuthorityId: v,
		})
	}
	user := &system.SysUser{Username: r.Username, NickName: r.NickName, Password: r.Password, HeaderImg: r.HeaderImg, AuthorityId: r.AuthorityId, Authorities: authorities, Enable: r.Enable, Phone: r.Phone, Email: r.Email, Level: r.Level, Balance: r.Balance, Robot: r.Robot}
	userReturn, err := userService.Register(*user)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))

		// 获取语言设置
		lang := c.GetHeader("X-Language")
		if lang == "" {
			acceptLang := c.GetHeader("Accept-Language")
			lang = i18n.GetLangFromHeader(acceptLang)
		} else {
			lang = i18n.NormalizeLang(lang)
		}

		// 根据错误类型返回多语言消息
		var errorMsg string
		if err.Error() == "USERNAME_DUPLICATE" {
			errorMsg = i18n.GetMessage(lang, i18n.MsgUsernameDuplicate)
		} else {
			errorMsg = err.Error()
		}

		response.FailWithDetailed(systemRes.SysUserResponse{User: userReturn}, errorMsg, c)
		return
	}
	response.OkWithDetailed(systemRes.SysUserResponse{User: userReturn}, "", c)
}

// ChangePassword
// @Tags      SysUser
// @Summary
// @Security  ApiKeyAuth
// @Produce  application/json
// @Param     data  body      systemReq.ChangePasswordReq    true  ", , "
// @Success   200   {object}  response.Response{msg=string}  ""
// @Router    /user/changePassword [post]
func (b *BaseApi) ChangePassword(c *gin.Context) {
	var req systemReq.ChangePasswordReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(req, utils.ChangePasswordVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	uid := utils.GetUserID(c)
	u := &system.SysUser{GVA_MODEL: global.GVA_MODEL{ID: uid}, Password: req.Password}
	_, err = userService.ChangePassword(u, req.NewPassword)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage("，", c)
		return
	}
	response.OkWithMessage("", c)
}
func (b *BaseApi) ChangeWithdrawPassword(c *gin.Context) {

	var req systemReq.ChangePasswordReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(req, utils.ChangePasswordVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	uid := utils.GetUserID(c)
	if uid == 0 {
		response.FailWithMessage("ID", c)
		return
	}

	u := &system.SysUser{GVA_MODEL: global.GVA_MODEL{ID: uid}, Password: req.Password}
	_, err = userService.ChangeWithdrawPassword(u, req.NewPassword)
	if err != nil {
		global.GVA_LOG.Error("ChangeWithdrawPassword failed", zap.Error(err))
		lang := c.GetHeader("X-Language")
		if lang == "" {
			acceptLang := c.GetHeader("Accept-Language")
			lang = i18n.GetLangFromHeader(acceptLang)
		} else {
			lang = i18n.NormalizeLang(lang)
		}
		var errorMsg string
		if err.Error() == "old password error" {
			errorMsg = i18n.GetMessage(lang, i18n.MsgWithdrawPasswordError)
		} else {
			errorMsg = err.Error()
		}
		response.FailWithMessage(errorMsg, c)
		return
	}
	response.OkWithMessage("ok", c)
}
func (b *BaseApi) VerifyWithdrawPassword(c *gin.Context) {

	var req systemReq.VerifyWithdrawPasswordReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(req, utils.ChangePasswordVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	uid := utils.GetUserID(c)
	if uid == 0 {
		response.FailWithMessage("ID", c)
		return
	}

	u := &system.SysUser{GVA_MODEL: global.GVA_MODEL{ID: uid}, Password: req.Password}
	err = userService.VerifyWithdrawPassword(u, req.Password)

	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		lang := c.GetHeader("X-Language")
		if lang == "" {
			acceptLang := c.GetHeader("Accept-Language")
			lang = i18n.GetLangFromHeader(acceptLang)
		} else {
			lang = i18n.NormalizeLang(lang)
		}
		var errorMsg string
		if err.Error() == "WITHDRAW_PASSWORD_ERROR" {
			errorMsg = i18n.GetMessage(lang, i18n.MsgWithdrawPasswordError)
		} else {
			errorMsg = err.Error()
		}
		response.FailWithMessage(errorMsg, c)
		return
	}
	response.OkWithMessage("ok", c)
}
func (b *BaseApi) SetWithdrawPassword(c *gin.Context) {

	var req systemReq.VerifyWithdrawPasswordReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(req, utils.ChangePasswordVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	uid := utils.GetUserID(c)
	if uid == 0 {
		response.FailWithMessage("ID", c)
		return
	}

	u := &system.SysUser{GVA_MODEL: global.GVA_MODEL{ID: uid}, WithdrawPassword: req.Password}
	_, err = userService.SetWithdrawPassword(u, req.Password, req.LoginPassword)
	if err != nil {
		global.GVA_LOG.Error("SetWithdrawPassword failed", zap.Error(err))
		lang := c.GetHeader("X-Language")
		if lang == "" {
			acceptLang := c.GetHeader("Accept-Language")
			lang = i18n.GetLangFromHeader(acceptLang)
		} else {
			lang = i18n.NormalizeLang(lang)
		}
		var errorMsg string
		if err.Error() == "LOGIN_PASSWORD_ERROR" {
			errorMsg = i18n.GetMessage(lang, i18n.MsgLoginPasswordError)
		} else {
			errorMsg = err.Error()
		}
		response.FailWithMessage(errorMsg, c)
		return
	}

	response.OkWithMessage("ok", c)
}

// GetUserList
// @Tags      SysUser
// @Summary
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      systemReq.GetUserList                                        true  ", "
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  ",,,,"
// @Router    /user/getUserList [post]
func (b *BaseApi) GetUserList(c *gin.Context) {
	var pageInfo systemReq.GetUserList
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(pageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := userService.GetUserInfoList(pageInfo)
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

// SetUserAuthority
// @Tags      SysUser
// @Summary
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      systemReq.SetUserAuth          true  "UUID, ID"
// @Success   200   {object}  response.Response{msg=string}  ""
// @Router    /user/setUserAuthority [post]
func (b *BaseApi) SetUserAuthority(c *gin.Context) {
	var sua systemReq.SetUserAuth
	err := c.ShouldBindJSON(&sua)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if UserVerifyErr := utils.Verify(sua, utils.SetUserAuthorityVerify); UserVerifyErr != nil {
		response.FailWithMessage(UserVerifyErr.Error(), c)
		return
	}
	userID := utils.GetUserID(c)
	err = userService.SetUserAuthority(userID, sua.AuthorityId)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	claims := utils.GetUserInfo(c)
	claims.AuthorityId = sua.AuthorityId
	token, err := utils.NewJWT().CreateToken(*claims)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	c.Header("new-token", token)
	c.Header("new-expires-at", strconv.FormatInt(claims.ExpiresAt.Unix(), 10))
	utils.SetToken(c, token, int((claims.ExpiresAt.Unix()-time.Now().Unix())/60))
	response.OkWithMessage("", c)
}

// SetUserAuthorities
// @Tags      SysUser
// @Summary
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      systemReq.SetUserAuthorities   true  "UUID, ID"
// @Success   200   {object}  response.Response{msg=string}  ""
// @Router    /user/setUserAuthorities [post]
func (b *BaseApi) SetUserAuthorities(c *gin.Context) {
	var sua systemReq.SetUserAuthorities
	err := c.ShouldBindJSON(&sua)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	authorityID := utils.GetUserAuthorityId(c)
	err = userService.SetUserAuthorities(authorityID, sua.ID, sua.AuthorityIds)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage("", c)
		return
	}
	response.OkWithMessage("", c)
}

// DeleteUser
// @Tags      SysUser
// @Summary
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.GetById                true  "ID"
// @Success   200   {object}  response.Response{msg=string}  ""
// @Router    /user/deleteUser [delete]
func (b *BaseApi) DeleteUser(c *gin.Context) {
	var reqId request.GetById
	err := c.ShouldBindJSON(&reqId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(reqId, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	jwtId := utils.GetUserID(c)
	if jwtId == uint(reqId.ID) {
		response.FailWithMessage(", 。", c)
		return
	}
	err = userService.DeleteUser(reqId.ID)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage("", c)
		return
	}
	response.OkWithMessage("", c)
}

// SetUserInfo
// @Tags      SysUser
// @Summary
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysUser                                             true  "ID, , , "
// @Success   200   {object}  response.Response{data=map[string]interface{},msg=string}  ""
// @Router    /user/setUserInfo [put]
func (b *BaseApi) SetUserInfo(c *gin.Context) {
	var user systemReq.ChangeUserInfo
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(user, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if len(user.AuthorityIds) != 0 {
		authorityID := utils.GetUserAuthorityId(c)
		err = userService.SetUserAuthorities(authorityID, user.ID, user.AuthorityIds)
		if err != nil {
			global.GVA_LOG.Error("!", zap.Error(err))
			response.FailWithMessage("", c)
			return
		}
	}
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = userService.SetUserInfo(system.SysUser{
		GVA_MODEL: global.GVA_MODEL{
			ID: user.ID,
		},
		NickName:  user.NickName,
		HeaderImg: user.HeaderImg,
		Phone:     user.Phone,
		Level:     user.Level,
		Email:     user.Email,
		Enable:    user.Enable,
	})
	if err != nil {
		global.GVA_LOG.Error("Failed to update user info", zap.Error(err))
		response.FailWithMessage("更新用户信息失败", c)
		return
	}
	updatedUser, err := userService.FindUserByUId(user.ID)
	if err != nil {
		global.GVA_LOG.Error("Failed to get updated user info", zap.Error(err))
		response.FailWithMessage("获取更新后的用户信息失败", c)
		return
	}

	userJson, err := json.Marshal(updatedUser)
	if err != nil {
		global.GVA_LOG.Error("Failed to marshal updated user data", zap.Error(err))
	} else {
		err = global.GVA_REDIS.Set(c, fmt.Sprintf("user_%d", user.ID), string(userJson), 0).Err()
		if err != nil {
			global.GVA_LOG.Error("Failed to update user data in Redis", zap.Error(err))
		}
	}

	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage("", c)
		return
	}
	response.OkWithMessage("", c)
}

// SetSelfInfo
// @Tags      SysUser
// @Summary
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysUser                                             true  "ID, , , "
// @Success   200   {object}  response.Response{data=map[string]interface{},msg=string}  ""
// @Router    /user/SetSelfInfo [put]
func (b *BaseApi) SetSelfInfo(c *gin.Context) {
	var user systemReq.ChangeUserInfo
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	user.ID = utils.GetUserID(c)
	err = userService.SetSelfInfo(system.SysUser{
		GVA_MODEL: global.GVA_MODEL{
			ID: user.ID,
		},
		NickName:  user.NickName,
		HeaderImg: user.HeaderImg,
		Phone:     user.Phone,
		Email:     user.Email,
		Enable:    user.Enable,
	})
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage("", c)
		return
	}
	response.OkWithMessage("", c)
}

// SetSelfSetting
// @Tags      SysUser
// @Summary
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      map[string]interface{}  true  ""
// @Success   200   {object}  response.Response{data=map[string]interface{},msg=string}  ""
// @Router    /user/SetSelfSetting [put]
func (b *BaseApi) SetSelfSetting(c *gin.Context) {
	var req common.JSONMap
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = userService.SetSelfSetting(req, utils.GetUserID(c))
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage("", c)
		return
	}
	response.OkWithMessage("", c)
}

// GetUserInfo
// @Tags      SysUser
// @Summary
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{data=map[string]interface{},msg=string}  ""
// @Router    /user/getUserInfo [get]
func (b *BaseApi) GetUserInfo(c *gin.Context) {
	uuid := utils.GetUserUuid(c)
	ReqUser, err := userService.GetUserInfo(uuid)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage("", c)
		return
	}
	response.OkWithDetailed(gin.H{"userInfo": ReqUser}, "", c)
}

// ResetPassword
// @Tags      SysUser
// @Summary
// @Security  ApiKeyAuth
// @Produce  application/json
// @Param     data  body      system.SysUser                 true  "ID"
// @Success   200   {object}  response.Response{msg=string}  ""
// @Router    /user/resetPassword [post]
func (b *BaseApi) ResetPassword(c *gin.Context) {
	var rps systemReq.ResetPassword
	err := c.ShouldBindJSON(&rps)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = userService.ResetPassword(rps.ID, rps.Password)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage(""+err.Error(), c)
		return
	}
	response.OkWithMessage("", c)
}
func (b *BaseApi) ResetWithdrawPassword(c *gin.Context) {
	// ID
	uid := utils.GetUserID(c)
	if uid == 0 {
		response.FailWithMessage("ID", c)
		return
	}

	var rps systemReq.WithdrawPassword
	err := c.ShouldBindJSON(&rps)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	rps.ID = uid
	err = userService.ResetWithdrawPassword(rps.ID, rps.Password)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))
		response.FailWithMessage(""+err.Error(), c)
		return
	}
	response.OkWithMessage("", c)
}
func (b *BaseApi) ApiLogin(c *gin.Context) {
	var l systemReq.ApiLogin
	err := c.ShouldBindJSON(&l)
	key := c.ClientIP()

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	u := &system.SysUser{Username: l.Username, Password: l.Password}
	user, err := userService.ApiLogin(u)
	if err != nil {
		global.GVA_LOG.Error("! !", zap.Error(err))
		global.BlackCache.Increment(key, 1)

		// 获取语言设置
		lang := c.GetHeader("X-Language")
		if lang == "" {
			acceptLang := c.GetHeader("Accept-Language")
			lang = i18n.GetLangFromHeader(acceptLang)
		} else {
			lang = i18n.NormalizeLang(lang)
		}

		// 根据错误类型返回多语言消息
		var errorMsg string
		if err.Error() == "password error" {
			errorMsg = i18n.GetMessage(lang, i18n.MsgLoginPasswordError)
		} else {
			errorMsg = i18n.GetMessage(lang, i18n.MsgFailed)
		}

		response.FailWithMessage(errorMsg, c)
		return
	}
	if user.Enable != 1 {
		global.GVA_LOG.Error("! !")
		// +1

		// 获取语言设置
		lang := c.GetHeader("X-Language")
		if lang == "" {
			acceptLang := c.GetHeader("Accept-Language")
			lang = i18n.GetLangFromHeader(acceptLang)
		} else {
			lang = i18n.NormalizeLang(lang)
		}

		response.FailWithMessage(i18n.GetMessage(lang, i18n.MsgFailed), c)
		return
	}

	// 直接从MySQL数据库查询最新的用户数据
	global.GVA_LOG.Info("Querying latest user data from MySQL for API login",
		zap.Uint("userId", user.ID),
		zap.String("username", user.Username))

	// 查询最新的用户数据，包括余额等信息
	var latestUser system.SysUser
	err = global.GVA_DB.Where("id = ?", user.ID).First(&latestUser).Error
	if err != nil {
		global.GVA_LOG.Error("Failed to query latest user data from MySQL",
			zap.Error(err),
			zap.Uint("userId", user.ID))
		// 如果查询失败，使用原始用户数据
		b.ApiTokenNext(c, *user)
		return
	}

	global.GVA_LOG.Info("Successfully retrieved latest user data from MySQL",
		zap.Uint("userId", latestUser.ID),
		zap.String("username", latestUser.Username),
		zap.Float64("balance", latestUser.Balance))

	b.ApiTokenNext(c, latestUser)
}

// TokenNext jwt
func (b *BaseApi) ApiTokenNext(c *gin.Context, user system.SysUser) {
	j := &utils.JWT{SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey)}
	claims := j.CreateClaims(systemReq.BaseClaims{
		UUID:     user.UUID,
		ID:       uint(user.ID),
		NickName: user.NickName,
		Username: user.Username,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		global.GVA_LOG.Error("get token error!", zap.Error(err))

		// 获取语言设置
		lang := c.GetHeader("X-Language")
		if lang == "" {
			acceptLang := c.GetHeader("Accept-Language")
			lang = i18n.GetLangFromHeader(acceptLang)
		} else {
			lang = i18n.NormalizeLang(lang)
		}

		response.FailWithMessage(i18n.GetMessage(lang, i18n.MsgFailed), c)
		return
	}

	// token
	if err := utils.SetRedisJWT(token, user.Username); err != nil {
		global.GVA_LOG.Error("set fail!", zap.Error(err))

		// 获取语言设置
		lang := c.GetHeader("X-Language")
		if lang == "" {
			acceptLang := c.GetHeader("Accept-Language")
			lang = i18n.GetLangFromHeader(acceptLang)
		} else {
			lang = i18n.NormalizeLang(lang)
		}

		response.FailWithMessage(i18n.GetMessage(lang, i18n.MsgFailed), c)
		return
	}

	// 由于JWT永不过期，设置一个很长的过期时间
	utils.SetToken(c, token, 365*24*60*60) // 1年

	users, _ := global.GVA_REDIS.Get(c, fmt.Sprintf("user_%d", user.ID)).Result()
	if users == "" {
		userJson, err := json.Marshal(user)
		if err != nil {
			global.GVA_LOG.Error("Failed to marshal user data", zap.Error(err))
		} else {
			err = global.GVA_REDIS.Set(c, fmt.Sprintf("user_%d", user.ID), string(userJson), 0).Err()
			if err != nil {
				global.GVA_LOG.Error("Failed to save user data to Redis", zap.Error(err))
			}
		}
	}

	// 获取语言设置
	lang := c.GetHeader("X-Language")
	if lang == "" {
		acceptLang := c.GetHeader("Accept-Language")
		lang = i18n.GetLangFromHeader(acceptLang)
	} else {
		lang = i18n.NormalizeLang(lang)
	}

	// 修复空指针异常：检查ExpiresAt是否为nil
	var expiresAt int64
	if claims.RegisteredClaims.ExpiresAt != nil {
		expiresAt = claims.RegisteredClaims.ExpiresAt.Unix() * 1000
	} else {
		// 如果ExpiresAt为nil，设置一个默认的过期时间（1年后）
		expiresAt = time.Now().AddDate(1, 0, 0).Unix() * 1000
	}

	response.OkWithDetailed(systemRes.LoginResponse{
		User:      user,
		Token:     token,
		ExpiresAt: expiresAt,
	}, i18n.GetMessage(lang, i18n.MsgSuccess), c)
}
func (b *BaseApi) ApiRegister(c *gin.Context) {
	var r systemReq.Register
	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(r, utils.ApiRegisterVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var authorities []system.SysAuthority
	for _, v := range r.AuthorityIds {
		authorities = append(authorities, system.SysAuthority{
			AuthorityId: v,
		})
	}
	r.AuthorityId = 888
	r.AuthorityIds = []uint{888}
	r.Enable = 1
	user := &system.SysUser{Username: r.Username, NickName: r.NickName, Password: r.Password, HeaderImg: r.HeaderImg, AuthorityId: r.AuthorityId, Authorities: authorities, Enable: r.Enable, Phone: r.Phone, Email: r.Email, Balance: 0}
	userReturn, err := userService.ApiRegister(*user)
	if err != nil {
		global.GVA_LOG.Error("!", zap.Error(err))

		// 获取语言设置
		lang := c.GetHeader("X-Language")
		if lang == "" {
			acceptLang := c.GetHeader("Accept-Language")
			lang = i18n.GetLangFromHeader(acceptLang)
		} else {
			lang = i18n.NormalizeLang(lang)
		}

		// 根据错误类型返回多语言消息
		var errorMsg string
		if err.Error() == "USERNAME_DUPLICATE" {
			errorMsg = i18n.GetMessage(lang, i18n.MsgUsernameDuplicate)
		} else {
			errorMsg = err.Error()
		}

		response.FailWithMessage(errorMsg, c)
		return
	}
	if r.Uuid != "" {
		parentUser, err := userService.FindUserByUuid(r.Uuid)
		if err != nil || parentUser.ID == 0 {
			// 获取语言设置
			lang := c.GetHeader("X-Language")
			if lang == "" {
				acceptLang := c.GetHeader("Accept-Language")
				lang = i18n.GetLangFromHeader(acceptLang)
			} else {
				lang = i18n.NormalizeLang(lang)
			}

			response.FailWithMessage(i18n.GetMessage(lang, i18n.MsgFailed), c)
			return
		}
		userAgentRelation2, err := userAgentRelationService.GetUserAgentRelation(c, fmt.Sprint(parentUser.ID))
		var userAgentRelation system.UserAgentRelation
		userAgentRelation.UserId = int(userReturn.ID)
		userAgentRelation.ParentId1 = int(parentUser.ID)
		if userAgentRelation2.ParentId1 > 0 {
			userAgentRelation.ParentId2 = int(userAgentRelation2.ParentId1)
		}
		userAgentRelationService.CreateUserAgentRelation(c, &userAgentRelation)

		// 记录邀请关系创建信息
		global.GVA_LOG.Info("Creating invitation relation for new user",
			zap.Uint("newUserId", userReturn.ID),
			zap.String("newUsername", userReturn.Username),
			zap.Uint("parentUserId", parentUser.ID),
			zap.String("parentUsername", parentUser.Username),
			zap.Int("parentId1", int(parentUser.ID)),
			zap.Int("parentId2", userAgentRelation.ParentId2),
			zap.String("invitationChain", fmt.Sprintf("新用户%d(%s) -> 1级上级%d(%s) -> 2级上级%d",
				userReturn.ID, userReturn.Username,
				parentUser.ID, parentUser.Username,
				userAgentRelation.ParentId2)))

		// 保存邀请关系到Redis
		saveUserInvitationRelation(c, userReturn.ID, int(parentUser.ID), userAgentRelation.ParentId2)
	}

	// 获取语言设置
	lang := c.GetHeader("X-Language")
	if lang == "" {
		acceptLang := c.GetHeader("Accept-Language")
		lang = i18n.GetLangFromHeader(acceptLang)
	} else {
		lang = i18n.NormalizeLang(lang)
	}

	// 返回注册成功的多语言消息
	successMsg := i18n.GetMessage(lang, i18n.MsgSuccess)
	response.OkWithMessage(successMsg, c)
}

// saveUserInvitationRelation 保存用户邀请关系到Redis
func saveUserInvitationRelation(c *gin.Context, userId uint, parentId1 int, parentId2 int) {
	// 记录邀请关系信息
	global.GVA_LOG.Info("Saving user invitation relation",
		zap.Uint("userId", userId),
		zap.Int("parentId1", parentId1),
		zap.Int("parentId2", parentId2),
		zap.String("relationInfo", fmt.Sprintf("用户%d的1级上级:%d, 2级上级:%d", userId, parentId1, parentId2)))

	// 构建邀请关系数据
	invitationData := map[string]interface{}{
		"level1": parentId1, // 1级上级
		"level2": parentId2, // 2级上级
	}

	// 序列化为JSON
	invitationJson, err := json.Marshal(invitationData)
	if err != nil {
		global.GVA_LOG.Error("Failed to marshal invitation relation",
			zap.Error(err),
			zap.Uint("userId", userId),
			zap.Int("parentId1", parentId1),
			zap.Int("parentId2", parentId2))
		return
	}

	// 保存到Redis，key格式: invitation_关系_{用户ID}
	key := fmt.Sprintf("invitation_relation_%d", userId)
	err = global.GVA_REDIS.Set(c, key, string(invitationJson), 0).Err()
	if err != nil {
		global.GVA_LOG.Error("Failed to save invitation relation to Redis",
			zap.Error(err),
			zap.String("redisKey", key),
			zap.Uint("userId", userId))
	} else {
		global.GVA_LOG.Info("Successfully saved invitation relation to Redis",
			zap.String("redisKey", key),
			zap.String("invitationData", string(invitationJson)),
			zap.Uint("userId", userId))
	}

	// 同时保存反向关系，方便通过上级ID查找下级
	// 1级上级的下级列表
	if parentId1 > 0 {
		level1Key := fmt.Sprintf("invitation_children_level1_%d", parentId1)
		err := global.GVA_REDIS.SAdd(c, level1Key, fmt.Sprintf("%d", userId)).Err()
		if err != nil {
			global.GVA_LOG.Error("Failed to save level1 children relation to Redis",
				zap.Error(err),
				zap.String("redisKey", level1Key),
				zap.Int("parentId1", parentId1),
				zap.Uint("userId", userId))
		} else {
			global.GVA_LOG.Info("Successfully saved level1 children relation to Redis",
				zap.String("redisKey", level1Key),
				zap.Int("parentId1", parentId1),
				zap.Uint("userId", userId))
		}
	} else {
		global.GVA_LOG.Info("No level1 parent, skipping level1 children relation",
			zap.Uint("userId", userId),
			zap.Int("parentId1", parentId1))
	}

	// 2级上级的下级列表
	if parentId2 > 0 {
		level2Key := fmt.Sprintf("invitation_children_level2_%d", parentId2)
		err := global.GVA_REDIS.SAdd(c, level2Key, fmt.Sprintf("%d", userId)).Err()
		if err != nil {
			global.GVA_LOG.Error("Failed to save level2 children relation to Redis",
				zap.Error(err),
				zap.String("redisKey", level2Key),
				zap.Int("parentId2", parentId2),
				zap.Uint("userId", userId))
		} else {
			global.GVA_LOG.Info("Successfully saved level2 children relation to Redis",
				zap.String("redisKey", level2Key),
				zap.Int("parentId2", parentId2),
				zap.Uint("userId", userId))
		}
	} else {
		global.GVA_LOG.Info("No level2 parent, skipping level2 children relation",
			zap.Uint("userId", userId),
			zap.Int("parentId2", parentId2))
	}

	// 记录完整的邀请关系总结
	global.GVA_LOG.Info("Invitation relation saved successfully",
		zap.Uint("userId", userId),
		zap.Int("parentId1", parentId1),
		zap.Int("parentId2", parentId2),
		zap.String("summary", fmt.Sprintf("用户%d的邀请关系已保存: 1级上级=%d, 2级上级=%d", userId, parentId1, parentId2)))
}

// getUserInvitationRelation 获取用户邀请关系
func getUserInvitationRelation(c *gin.Context, userId uint) (map[string]interface{}, error) {
	key := fmt.Sprintf("invitation_relation_%d", userId)
	result, err := global.GVA_REDIS.Get(c, key).Result()
	if err != nil {
		return nil, err
	}

	var invitationData map[string]interface{}
	err = json.Unmarshal([]byte(result), &invitationData)
	if err != nil {
		return nil, err
	}

	return invitationData, nil
}

// getUserChildren 获取用户的下级列表
func getUserChildren(c *gin.Context, userId int, level int) ([]string, error) {
	key := fmt.Sprintf("invitation_children_level%d_%d", level, userId)
	result, err := global.GVA_REDIS.SMembers(c, key).Result()
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetUserInvitationRelation 获取用户邀请关系API
func (b *BaseApi) GetUserInvitationRelation(c *gin.Context) {
	// 从请求参数中获取用户ID
	userIdStr := c.Query("userId")
	if userIdStr == "" {
		response.FailWithMessage("用户ID不能为空", c)
		return
	}

	userId, err := strconv.ParseUint(userIdStr, 10, 32)
	if err != nil {
		response.FailWithMessage("用户ID格式错误", c)
		return
	}

	// 获取邀请关系
	relation, err := getUserInvitationRelation(c, uint(userId))
	if err != nil {
		global.GVA_LOG.Error("Failed to get user invitation relation",
			zap.Error(err),
			zap.Uint64("userId", userId))
		response.FailWithMessage("获取邀请关系失败", c)
		return
	}

	// 获取上级用户详细信息
	var result map[string]interface{}
	if relation != nil {
		result = make(map[string]interface{})

		// 获取1级上级信息
		if level1Id, ok := relation["level1"].(float64); ok && level1Id > 0 {
			level1User, err := userService.FindUserByUId(uint(level1Id))
			if err == nil {
				result["level1_user"] = level1User
			}
		}

		// 获取2级上级信息
		if level2Id, ok := relation["level2"].(float64); ok && level2Id > 0 {
			level2User, err := userService.FindUserByUId(uint(level2Id))
			if err == nil {
				result["level2_user"] = level2User
			}
		}

		result["relation"] = relation
	}

	response.OkWithDetailed(result, "获取邀请关系成功", c)
}

// GetUserChildren 获取用户下级列表API
func (b *BaseApi) GetUserChildren(c *gin.Context) {
	// 从请求参数中获取用户ID和层级
	userIdStr := c.Query("userId")
	levelStr := c.Query("level")

	if userIdStr == "" {
		response.FailWithMessage("用户ID不能为空", c)
		return
	}

	if levelStr == "" {
		levelStr = "1" // 默认获取1级下级
	}

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		response.FailWithMessage("用户ID格式错误", c)
		return
	}

	level, err := strconv.Atoi(levelStr)
	if err != nil || (level != 1 && level != 2) {
		response.FailWithMessage("层级参数错误，只能是1或2", c)
		return
	}

	// 获取下级列表
	childrenIds, err := getUserChildren(c, userId, level)
	if err != nil {
		global.GVA_LOG.Error("Failed to get user children",
			zap.Error(err),
			zap.Int("userId", userId),
			zap.Int("level", level))
		response.FailWithMessage("获取下级列表失败", c)
		return
	}

	// 获取下级用户详细信息
	var childrenUsers []interface{}
	for _, childIdStr := range childrenIds {
		childId, err := strconv.ParseUint(childIdStr, 10, 32)
		if err != nil {
			continue
		}

		childUser, err := userService.FindUserByUId(uint(childId))
		if err == nil {
			childrenUsers = append(childrenUsers, childUser)
		}
	}

	response.OkWithDetailed(gin.H{
		"children": childrenUsers,
		"total":    len(childrenUsers),
		"level":    level,
	}, "获取下级列表成功", c)
}

func GenerateVerificationCode() string {
	const charset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const length = 4

	rand.Seed(time.Now().UnixNano())

	code := make([]byte, length)
	for i := range code {
		code[i] = charset[rand.Intn(len(charset))]
	}

	return string(code)
}
func (b *BaseApi) SendCode(c *gin.Context) {
	uid := utils.GetRedisUserID(c)
	if uid == 0 {
		// 获取语言设置
		lang := c.GetHeader("X-Language")
		if lang == "" {
			acceptLang := c.GetHeader("Accept-Language")
			lang = i18n.GetLangFromHeader(acceptLang)
		} else {
			lang = i18n.NormalizeLang(lang)
		}

		response.Result(401, nil, i18n.GetMessage(lang, i18n.MsgFailed), c)
		return
	}
	var r systemReq.SendEmailCodeRequest
	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = utils.Verify(r, utils.ApiSendEmailCodeVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = userService.CheckEmail(r.Email)
	if err != nil {
		// 获取语言设置
		lang := c.GetHeader("X-Language")
		if lang == "" {
			acceptLang := c.GetHeader("Accept-Language")
			lang = i18n.GetLangFromHeader(acceptLang)
		} else {
			lang = i18n.NormalizeLang(lang)
		}

		// 根据错误类型返回多语言消息
		var errorMsg string
		if err.Error() == "EMAIL_DUPLICATE" {
			errorMsg = i18n.GetMessage(lang, i18n.MsgEmailDuplicate)
		} else {
			errorMsg = err.Error()
		}

		response.FailWithMessage(errorMsg, c)
		return
	}
	code := GenerateVerificationCode()

	err = global.GVA_REDIS.Set(c, fmt.Sprintf("email_code_%d:%s", uid, r.Email), code, 5*time.Minute).Err()
	if err != nil {
		global.GVA_LOG.Error("save fail !", zap.Error(err))

		// 获取语言设置
		lang := c.GetHeader("X-Language")
		if lang == "" {
			acceptLang := c.GetHeader("Accept-Language")
			lang = i18n.GetLangFromHeader(acceptLang)
		} else {
			lang = i18n.NormalizeLang(lang)
		}

		response.FailWithMessage(i18n.GetMessage(lang, i18n.MsgFailed), c)
		return
	}

	err = SendVerificationEmail(r.Email, code)
	if err != nil {
		global.GVA_LOG.Error("send fail!", zap.Error(err))

		// 获取语言设置
		lang := c.GetHeader("X-Language")
		if lang == "" {
			acceptLang := c.GetHeader("Accept-Language")
			lang = i18n.GetLangFromHeader(acceptLang)
		} else {
			lang = i18n.NormalizeLang(lang)
		}

		response.FailWithMessage(i18n.GetMessage(lang, i18n.MsgFailed), c)
		return
	}

	// 获取语言设置
	lang := c.GetHeader("X-Language")
	if lang == "" {
		acceptLang := c.GetHeader("Accept-Language")
		lang = i18n.GetLangFromHeader(acceptLang)
	} else {
		lang = i18n.NormalizeLang(lang)
	}

	response.OkWithMessage(i18n.GetMessage(lang, i18n.MsgSuccess), c)
}

func SendVerificationEmail(email, code string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "lul0215@163.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "verification code")
	m.SetBody("text/plain", fmt.Sprintf("Your verification code is: %s, valid for 5 minutes.", code))

	d := gomail.NewDialer("smtp.163.com", 465, "lul0215@163.com", "PNgReYFe54tQNej6")
	d.SSL = true

	if err := d.DialAndSend(m); err != nil {
		fmt.Printf("send fail: %v\n", err)
		return err
	}

	return nil
}
func (b *BaseApi) BindeMail(c *gin.Context) {
	uid := utils.GetRedisUserID(c)
	if uid == 0 {
		// 获取语言设置
		lang := c.GetHeader("X-Language")
		if lang == "" {
			acceptLang := c.GetHeader("Accept-Language")
			lang = i18n.GetLangFromHeader(acceptLang)
		} else {
			lang = i18n.NormalizeLang(lang)
		}

		response.Result(401, nil, i18n.GetMessage(lang, i18n.MsgFailed), c)
		return
	}
	var r systemReq.BindEmailRequest
	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	storedCode, err := global.GVA_REDIS.Get(c, fmt.Sprintf("email_code_%d:%s", uid, r.Email)).Result()
	if err != nil {
		// 获取语言设置
		lang := c.GetHeader("X-Language")
		if lang == "" {
			acceptLang := c.GetHeader("Accept-Language")
			lang = i18n.GetLangFromHeader(acceptLang)
		} else {
			lang = i18n.NormalizeLang(lang)
		}

		response.FailWithMessage(i18n.GetMessage(lang, i18n.MsgFailed), c)
		return
	}

	if storedCode != r.Code {
		// 获取语言设置
		lang := c.GetHeader("X-Language")
		if lang == "" {
			acceptLang := c.GetHeader("Accept-Language")
			lang = i18n.GetLangFromHeader(acceptLang)
		} else {
			lang = i18n.NormalizeLang(lang)
		}

		response.FailWithMessage(i18n.GetMessage(lang, i18n.MsgFailed), c)
		return
	}

	global.GVA_REDIS.Del(c, fmt.Sprintf("email_code:%s", r.Email))

	err = userService.BindEmail(uid, r.Email)
	if err != nil {
		global.GVA_LOG.Error("bind email fail!", zap.Error(err))

		// 获取语言设置
		lang := c.GetHeader("X-Language")
		if lang == "" {
			acceptLang := c.GetHeader("Accept-Language")
			lang = i18n.GetLangFromHeader(acceptLang)
		} else {
			lang = i18n.NormalizeLang(lang)
		}

		response.FailWithMessage(i18n.GetMessage(lang, i18n.MsgFailed), c)
		return
	}

	// 获取语言设置
	lang := c.GetHeader("X-Language")
	if lang == "" {
		acceptLang := c.GetHeader("Accept-Language")
		lang = i18n.GetLangFromHeader(acceptLang)
	} else {
		lang = i18n.NormalizeLang(lang)
	}

	response.OkWithMessage(i18n.GetMessage(lang, i18n.MsgSuccess), c)
}
func (b *BaseApi) Decrypt(c *gin.Context) {
	type DecryptRequest struct {
		Data string `json:"data" binding:"required"`
		IV   string `json:"iv" binding:"required"`
	}

	var r DecryptRequest
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
	response.OkWithDetailed(decrypted, "ok", c)
}
func (b *BaseApi) Encrypt(c *gin.Context) {
	var rawData map[string]interface{}
	if err := c.ShouldBindJSON(&rawData); err != nil {
		response.FailWithMessage("Invalid request format: "+err.Error(), c)
		return
	}

	jsonBytes, err := json.Marshal(rawData)
	if err != nil {
		response.FailWithMessage("Failed to marshal request data: "+err.Error(), c)
		return
	}
	encrypted, err := utils.CBCEncrypt(string(jsonBytes))
	if err != nil {
		response.FailWithMessage("Encryption failed: "+err.Error(), c)
		return
	}

	response.OkWithDetailed(encrypted, "ok", c)
}
func (b *BaseApi) Info(c *gin.Context) {

	uid := utils.GetRedisUserID(c)
	if uid == 0 {
		response.Result(401, nil, "", c)
		return
	}

	// 直接从MySQL数据库查询最新的用户数据
	global.GVA_LOG.Info("Querying latest user data from MySQL for Info",
		zap.Uint("userId", uid))

	var user system.SysUser
	err := global.GVA_DB.Where("id = ?", uid).First(&user).Error
	if err != nil {
		global.GVA_LOG.Error("Failed to query user data from MySQL",
			zap.Error(err),
			zap.Uint("userId", uid))
		response.Result(401, nil, "Failed to get user data", c)
		return
	}

	// 转换为ApiSysUser格式
	var apiUser system.ApiSysUser
	apiUser.ID = user.ID
	apiUser.Username = user.Username
	apiUser.NickName = user.NickName
	apiUser.HeaderImg = user.HeaderImg
	apiUser.Phone = user.Phone
	apiUser.Email = user.Email
	apiUser.Enable = user.Enable
	apiUser.Balance = user.Balance
	apiUser.Robot = user.Robot
	apiUser.Lang = user.Lang
	apiUser.Audio = user.Audio
	apiUser.CreatedAt = user.CreatedAt
	apiUser.UpdatedAt = user.UpdatedAt

	global.GVA_LOG.Info("Successfully retrieved user data from MySQL",
		zap.Uint("userId", apiUser.ID),
		zap.String("username", apiUser.Username),
		zap.Float64("balance", apiUser.Balance))

	response.OkWithDetailed(apiUser, "ok", c)
}
func (b *BaseApi) GetInfo(c *gin.Context) {

	uid := utils.GetRedisUserID(c)
	if uid == 0 {
		response.Result(401, nil, "", c)
		return
	}

	// 直接从MySQL数据库查询最新的用户数据
	global.GVA_LOG.Info("Querying latest user data from MySQL for GetInfo",
		zap.Uint("userId", uid))

	var user system.SysUser
	err := global.GVA_DB.Where("id = ?", uid).First(&user).Error
	if err != nil {
		global.GVA_LOG.Error("Failed to query user data from MySQL",
			zap.Error(err),
			zap.Uint("userId", uid))
		response.Result(401, nil, "Failed to get user data", c)
		return
	}

	// 转换为ApiSysUser格式
	var apiUser system.ApiSysUser
	apiUser.ID = user.ID
	apiUser.Username = user.Username
	apiUser.NickName = user.NickName
	apiUser.HeaderImg = user.HeaderImg
	apiUser.Phone = user.Phone
	apiUser.Email = user.Email
	apiUser.Enable = user.Enable
	apiUser.Balance = user.Balance
	apiUser.Robot = user.Robot
	apiUser.Lang = user.Lang
	apiUser.Audio = user.Audio
	apiUser.WithdrawPassword = user.WithdrawPassword
	apiUser.CreatedAt = user.CreatedAt
	apiUser.UpdatedAt = user.UpdatedAt

	global.GVA_LOG.Info("Successfully retrieved user data from MySQL",
		zap.Uint("userId", apiUser.ID),
		zap.String("username", apiUser.Username),
		zap.Float64("balance", apiUser.Balance))

	response.OkWithDetailed(apiUser, "ok", c)
}
func (b *BaseApi) AutoLogin(c *gin.Context) {
	// 获取所有用户数据
	var users []system.SysUser
	err := global.GVA_DB.Find(&users).Error
	if err != nil {
		global.GVA_LOG.Error("Failed to get all users from database", zap.Error(err))
		response.FailWithMessage("Failed to get users", c)
		return
	}

	// 统计信息
	var totalUsers int = len(users)
	var successUsers int = 0
	var failedUsers int = 0

	global.GVA_LOG.Info("Starting AutoLogin process",
		zap.Int("totalUsers", totalUsers))

	for _, user := range users {
		redisKey := fmt.Sprintf("user_%d", user.ID)

		// 直接使用system.SysUser结构
		userJson, err := json.Marshal(user)
		if err != nil {
			global.GVA_LOG.Error("Failed to marshal user data",
				zap.Error(err),
				zap.Uint("userId", user.ID),
				zap.String("username", user.Username))
			failedUsers++
			continue
		}

		// 直接设置Redis，不判断是否存在
		err = global.GVA_REDIS.Set(c, redisKey, string(userJson), 0).Err()
		if err != nil {
			global.GVA_LOG.Error("Failed to save user data to Redis",
				zap.Error(err),
				zap.Uint("userId", user.ID),
				zap.String("username", user.Username),
				zap.String("redisKey", redisKey))
			failedUsers++
			continue
		}

		successUsers++
		global.GVA_LOG.Info("Successfully set user to Redis",
			zap.Uint("userId", user.ID),
			zap.String("username", user.Username),
			zap.String("redisKey", redisKey))
	}

	global.GVA_LOG.Info("AutoLogin process completed",
		zap.Int("totalUsers", totalUsers),
		zap.Int("successUsers", successUsers),
		zap.Int("failedUsers", failedUsers))

	response.OkWithDetailed(gin.H{
		"total_users":   totalUsers,
		"success_users": successUsers,
		"failed_users":  failedUsers,
		"message":       "AutoLogin process completed successfully",
	}, "AutoLogin process completed", c)
}
func (b *BaseApi) UpdateLang(c *gin.Context) {
	uid := utils.GetRedisUserID(c)
	if uid == 0 {
		response.Result(401, nil, "", c)
		return
	}

	body, _ := ioutil.ReadAll(c.Request.Body)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	var requestData struct {
		Lang int `json:"lang" binding:"required"`
	}

	err := c.ShouldBindJSON(&requestData)

	if requestData.Lang < 0 || requestData.Lang > 1 {
		response.FailWithMessage("Invalid language code", c)
		return
	}

	// 使用分布式锁确保并发安全
	lockKey := fmt.Sprintf("user_lang_lock_%d", uid)
	locked, err := global.GVA_REDIS.SetNX(c, lockKey, "1", 10*time.Second).Result()
	if err != nil {
		global.GVA_LOG.Error("Failed to acquire lock for language update", zap.Error(err))
		response.FailWithMessage("System busy, please try again", c)
		return
	}
	if !locked {
		response.FailWithMessage("System busy, please try again", c)
		return
	}
	defer global.GVA_REDIS.Del(c, lockKey)

	// 获取当前用户数据
	var user system.ApiSysUser
	redisKey := fmt.Sprintf("user_%d", uid)
	redisuser, err := global.GVA_REDIS.Get(c, redisKey).Result()

	if err == nil && redisuser != "" {
		// 反序列化现有用户数据
		err = json.Unmarshal([]byte(redisuser), &user)
		if err == nil {
			// 更新语言设置
			user.Lang = requestData.Lang

			// 重新序列化并保存
			userJson, err := json.Marshal(user)
			if err == nil {
				err = global.GVA_REDIS.Set(c, redisKey, string(userJson), 0).Err()
				if err != nil {
					global.GVA_LOG.Error("Failed to update user language in Redis",
						zap.Error(err),
						zap.Uint("userId", uid),
						zap.Int("lang", requestData.Lang))
				} else {
					global.GVA_LOG.Info("Successfully updated user language in Redis",
						zap.Uint("userId", uid),
						zap.Int("lang", requestData.Lang))
				}
			}
		}
	}

	global.GVA_LOG.Info("User language updated successfully",
		zap.Uint("userId", uid),
		zap.Int("lang", requestData.Lang))

	response.OkWithMessage("Language updated successfully", c)
}
func (b *BaseApi) UpdateAudio(c *gin.Context) {
	uid := utils.GetRedisUserID(c)
	if uid == 0 {
		response.Result(401, nil, "", c)
		return
	}

	body, _ := ioutil.ReadAll(c.Request.Body)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	var requestData struct {
		Audio int `json:"audio" binding:"required"`
	}

	err := c.ShouldBindJSON(&requestData)

	if requestData.Audio < 0 || requestData.Audio > 1 {
		response.FailWithMessage("Invalid audio code", c)
		return
	}

	// 使用分布式锁确保并发安全
	lockKey := fmt.Sprintf("user_lang_lock_%d", uid)
	locked, err := global.GVA_REDIS.SetNX(c, lockKey, "1", 10*time.Second).Result()
	if err != nil {
		global.GVA_LOG.Error("Failed to acquire lock for language update", zap.Error(err))
		response.FailWithMessage("System busy, please try again", c)
		return
	}
	if !locked {
		response.FailWithMessage("System busy, please try again", c)
		return
	}
	defer global.GVA_REDIS.Del(c, lockKey)

	// 获取当前用户数据
	var user system.ApiSysUser
	redisKey := fmt.Sprintf("user_%d", uid)
	redisuser, err := global.GVA_REDIS.Get(c, redisKey).Result()

	fmt.Println("requestData.Audio", requestData.Audio)

	// 更新数据库中的audio字段
	err = global.GVA_DB.Model(&system.SysUser{}).Where("id = ?", uid).Update("audio", requestData.Audio).Error
	if err != nil {
		global.GVA_LOG.Error("Failed to update user audio in database",
			zap.Error(err),
			zap.Uint("userId", uid),
			zap.Int("Audio", requestData.Audio))
		response.FailWithMessage("Failed to update audio in database", c)
		return
	}

	if err == nil && redisuser != "" {
		// 反序列化现有用户数据
		err = json.Unmarshal([]byte(redisuser), &user)
		if err == nil {
			// 更新语言设置
			user.Audio = requestData.Audio

			// 重新序列化并保存
			userJson, err := json.Marshal(user)
			if err == nil {
				err = global.GVA_REDIS.Set(c, redisKey, string(userJson), 0).Err()
				if err != nil {
					global.GVA_LOG.Error("Failed to update user UpdateAudio in Redis",
						zap.Error(err),
						zap.Uint("userId", uid),
						zap.Int("Audio", requestData.Audio))
				} else {
					global.GVA_LOG.Info("Successfully updated user UpdateAudio in Redis",
						zap.Uint("userId", uid),
						zap.Int("Audio", requestData.Audio))
				}
			}
		}
	}

	response.OkWithMessage("Audio updated successfully", c)
}

// UpdateRedisUserDataSafe 并发安全的用户数据更新（同时更新数据库和Redis）
func (b *BaseApi) UpdateRedisUserDataSafe(c *gin.Context) {
	uid := utils.GetRedisUserID(c)
	if uid == 0 {
		response.Result(401, nil, "", c)
		return
	}

	// 接收POST请求中的用户数据
	var requestData struct {
		UserData map[string]interface{} `json:"user_data" binding:"required"`
	}

	err := c.ShouldBindJSON(&requestData)
	if err != nil {
		global.GVA_LOG.Error("UpdateRedisUserDataSafe request binding failed", zap.Error(err))
		response.FailWithMessage("Invalid request format: "+err.Error(), c)
		return
	}

	// 使用数据库事务和行锁来安全更新用户数据
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		// 使用FOR UPDATE锁来防止并发更新
		var user system.SysUser
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("id = ?", uid).First(&user).Error; err != nil {
			global.GVA_LOG.Error("Failed to get user with lock",
				zap.Error(err),
				zap.Uint("userId", uid))
			return err
		}

		// 记录原始数据
		originalBalance := user.Balance
		originalLang := user.Lang

		// 更新指定的字段
		for key, value := range requestData.UserData {
			switch key {
			case "balance":
				if balance, ok := value.(float64); ok {
					user.Balance = math.Round(balance*100) / 100
				}
			case "lang":
				if lang, ok := value.(int); ok {
					user.Lang = lang
				}
			case "audio":
				// Audio字段在SysUser中不存在，跳过
				global.GVA_LOG.Warn("Audio field not supported in database update",
					zap.Uint("userId", uid))
			default:
				global.GVA_LOG.Info("Skipping unsupported field for database update",
					zap.String("field", key),
					zap.Uint("userId", uid))
			}
		}

		// 更新数据库中的用户数据
		if err := tx.Save(&user).Error; err != nil {
			global.GVA_LOG.Error("Failed to update user data in database",
				zap.Error(err),
				zap.Uint("userId", uid))
			return err
		}

		// 同时更新Redis缓存
		userJson, err := json.Marshal(user)
		if err != nil {
			global.GVA_LOG.Error("Failed to marshal updated user data",
				zap.Error(err),
				zap.Uint("userId", uid))
			return err
		}

		// 更新Redis缓存
		redisKey := fmt.Sprintf("user_%d", uid)
		err = global.GVA_REDIS.Set(c, redisKey, string(userJson), 0).Err()
		if err != nil {
			global.GVA_LOG.Error("Failed to update user data in Redis",
				zap.Error(err),
				zap.Uint("userId", uid))
			// Redis更新失败不影响数据库事务，但会记录错误
		}

		// 生成事务码
		transactionCode := fmt.Sprintf("USER_UPDATE_%d_%d", uid, time.Now().Unix())

		// 记录详细的更新日志
		global.GVA_LOG.Info("Successfully updated user data in database and Redis",
			zap.Uint("userId", uid),
			zap.Float64("originalBalance", originalBalance),
			zap.Float64("newBalance", user.Balance),
			zap.Int("originalLang", originalLang),
			zap.Int("newLang", user.Lang),
			zap.Any("updatedFields", requestData.UserData),
			zap.String("transactionCode", transactionCode))

		return nil
	})

	if err != nil {
		global.GVA_LOG.Error("Failed to update user data",
			zap.Error(err),
			zap.Uint("userId", uid))
		response.FailWithMessage("Failed to update user data: "+err.Error(), c)
		return
	}

	response.OkWithMessage("User data updated successfully", c)
}

// UpdateRedisUserDataWithVersion 使用版本号防止并发冲突的用户数据更新（同时更新数据库和Redis）
func (b *BaseApi) UpdateRedisUserDataWithVersion(c *gin.Context) {
	uid := utils.GetRedisUserID(c)
	if uid == 0 {
		response.Result(401, nil, "", c)
		return
	}

	// 接收POST请求中的用户数据和版本号
	var requestData struct {
		UserData map[string]interface{} `json:"user_data" binding:"required"`
		Version  int64                  `json:"version" binding:"required"`
	}

	err := c.ShouldBindJSON(&requestData)
	if err != nil {
		global.GVA_LOG.Error("UpdateRedisUserDataWithVersion request binding failed", zap.Error(err))
		response.FailWithMessage("Invalid request format: "+err.Error(), c)
		return
	}

	// 检查版本号（从Redis获取）
	versionKey := fmt.Sprintf("user_version_%d", uid)
	currentVersion, err := global.GVA_REDIS.Get(c, versionKey).Int64()
	if err != nil {
		currentVersion = 0
	}

	// 检查版本号是否匹配
	if currentVersion != requestData.Version {
		response.FailWithMessage("Data has been modified by another request, please refresh and try again", c)
		return
	}

	// 使用数据库事务和行锁来安全更新用户数据
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		// 使用FOR UPDATE锁来防止并发更新
		var user system.SysUser
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("id = ?", uid).First(&user).Error; err != nil {
			global.GVA_LOG.Error("Failed to get user with lock",
				zap.Error(err),
				zap.Uint("userId", uid))
			return err
		}

		// 记录原始数据
		originalBalance := user.Balance
		originalLang := user.Lang

		// 更新指定的字段
		for key, value := range requestData.UserData {
			switch key {
			case "balance":
				if balance, ok := value.(float64); ok {
					user.Balance = math.Round(balance*100) / 100
				}
			case "lang":
				if lang, ok := value.(int); ok {
					user.Lang = lang
				}
			case "audio":
				// Audio字段在SysUser中不存在，跳过
				global.GVA_LOG.Warn("Audio field not supported in database update",
					zap.Uint("userId", uid))
			default:
				global.GVA_LOG.Info("Skipping unsupported field for database update",
					zap.String("field", key),
					zap.Uint("userId", uid))
			}
		}

		// 更新数据库中的用户数据
		if err := tx.Save(&user).Error; err != nil {
			global.GVA_LOG.Error("Failed to update user data in database",
				zap.Error(err),
				zap.Uint("userId", uid))
			return err
		}

		// 同时更新Redis缓存和版本号
		userJson, err := json.Marshal(user)
		if err != nil {
			global.GVA_LOG.Error("Failed to marshal updated user data",
				zap.Error(err),
				zap.Uint("userId", uid))
			return err
		}

		// 使用Lua脚本原子性更新Redis数据和版本号
		redisKey := fmt.Sprintf("user_%d", uid)
		luaScript := `
			local userKey = KEYS[1]
			local versionKey = KEYS[2]
			local userData = ARGV[1]
			local newVersion = ARGV[2]
			local expectedVersion = ARGV[3]
			
			-- 检查版本号
			local currentVersion = redis.call('GET', versionKey)
			if currentVersion and tonumber(currentVersion) ~= tonumber(expectedVersion) then
				return {err = "VERSION_MISMATCH"}
			end
			
			-- 原子性更新数据和版本号
			redis.call('SET', userKey, userData)
			redis.call('SET', versionKey, newVersion)
			return {ok = "SUCCESS"}
		`

		result, err := global.GVA_REDIS.Eval(c, luaScript, []string{redisKey, versionKey},
			string(userJson), requestData.Version+1, requestData.Version).Result()

		if err != nil {
			global.GVA_LOG.Error("Failed to update user data in Redis with version control",
				zap.Error(err),
				zap.Uint("userId", uid))
			// Redis更新失败不影响数据库事务，但会记录错误
		} else {
			// 检查Lua脚本执行结果
			if resultArray, ok := result.([]interface{}); ok && len(resultArray) > 0 {
				if resultArray[0] == "VERSION_MISMATCH" {
					return errors.New("version mismatch in Redis")
				}
			}
		}

		// 生成事务码
		transactionCode := fmt.Sprintf("USER_UPDATE_VERSION_%d_%d", uid, time.Now().Unix())

		// 记录详细的更新日志
		global.GVA_LOG.Info("Successfully updated user data in database and Redis with version control",
			zap.Uint("userId", uid),
			zap.Int64("version", requestData.Version),
			zap.Float64("originalBalance", originalBalance),
			zap.Float64("newBalance", user.Balance),
			zap.Int("originalLang", originalLang),
			zap.Int("newLang", user.Lang),
			zap.Any("updatedFields", requestData.UserData),
			zap.String("transactionCode", transactionCode))

		return nil
	})

	if err != nil {
		global.GVA_LOG.Error("Failed to update user data with version control",
			zap.Error(err),
			zap.Uint("userId", uid))
		if err.Error() == "version mismatch in Redis" {
			response.FailWithMessage("Data has been modified by another request, please refresh and try again", c)
		} else {
			response.FailWithMessage("Failed to update user data: "+err.Error(), c)
		}
		return
	}

	response.OkWithDetailed(gin.H{
		"message": "User data updated successfully",
		"version": requestData.Version + 1,
	}, "User data updated successfully", c)
}

func (b *BaseApi) RobotList(c *gin.Context) {

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
	var der apiReq.RobotRequest
	if err := json.Unmarshal([]byte(decryptedStr), &der); err != nil {
		response.FailWithMessage("Failed to unmarshal decrypted data: "+err.Error(), c)
		return
	}
	list, _ := userService.GetRobot(der.Limit)
	res, err := utils.CBCEncrypt(list)
	if err != nil {
		global.GVA_LOG.Error("CBCEncrypt failed", zap.Error(err))
		response.FailWithMessage("CBCEncrypt failed: "+err.Error(), c)
		return
	}
	response.OkWithDetailed(res, "ok", c)
}
