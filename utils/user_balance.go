package utils

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// UpdateUserBalanceWithLock 使用分布式锁安全地更新用户余额
func UpdateUserBalanceWithLock(c *gin.Context, userID uint, amount float64, operation string) error {
	// 生成锁key
	lockKey := fmt.Sprintf("user_balance_lock_%d", userID)

	// 尝试获取锁
	locked, err := global.GVA_REDIS.SetNX(c, lockKey, "1", 10*time.Second).Result()
	if err != nil {
		global.GVA_LOG.Error("Failed to acquire balance lock",
			zap.Error(err),
			zap.Uint("userId", userID),
			zap.String("operation", operation))
		return fmt.Errorf("failed to acquire lock: %v", err)
	}
	if !locked {
		global.GVA_LOG.Warn("User balance is being updated by another request",
			zap.Uint("userId", userID),
			zap.String("operation", operation))
		return fmt.Errorf("user balance is being updated by another request")
	}

	// 确保释放锁
	defer global.GVA_REDIS.Del(c, lockKey)

	// 获取最新的用户数据
	userKey := fmt.Sprintf("user_%d", userID)
	userData, err := global.GVA_REDIS.Get(c, userKey).Result()
	if err != nil || userData == "" {
		global.GVA_LOG.Error("Failed to get user data from Redis",
			zap.Error(err),
			zap.Uint("userId", userID))
		return fmt.Errorf("failed to get user data: %v", err)
	}

	// 解析用户数据
	var user system.ApiSysUser
	if err := json.Unmarshal([]byte(userData), &user); err != nil {
		global.GVA_LOG.Error("Failed to unmarshal user data",
			zap.Error(err),
			zap.Uint("userId", userID))
		return fmt.Errorf("failed to parse user data: %v", err)
	}

	// 检查余额是否足够（如果是扣款操作）
	if amount < 0 && user.Balance < -amount {
		global.GVA_LOG.Warn("Insufficient balance",
			zap.Uint("userId", userID),
			zap.Float64("currentBalance", user.Balance),
			zap.Float64("requestedAmount", -amount))
		return fmt.Errorf("insufficient balance")
	}

	// 更新余额
	oldBalance := user.Balance
	user.Balance += amount

	// 序列化更新后的用户数据
	updatedUserJson, err := json.Marshal(user)
	if err != nil {
		global.GVA_LOG.Error("Failed to marshal updated user data",
			zap.Error(err),
			zap.Uint("userId", userID))
		return fmt.Errorf("failed to marshal user data: %v", err)
	}

	// 保存到Redis
	err = global.GVA_REDIS.Set(c, userKey, string(updatedUserJson), 0).Err()
	if err != nil {
		global.GVA_LOG.Error("Failed to save updated user data to Redis",
			zap.Error(err),
			zap.Uint("userId", userID))
		return fmt.Errorf("failed to save user data: %v", err)
	}

	global.GVA_LOG.Info("Successfully updated user balance with lock protection",
		zap.Uint("userId", userID),
		zap.String("operation", operation),
		zap.Float64("oldBalance", oldBalance),
		zap.Float64("newBalance", user.Balance),
		zap.Float64("changeAmount", amount))

	return nil
}

// UpdateUserBalanceWithRetry 带重试机制的余额更新
func UpdateUserBalanceWithRetry(c *gin.Context, userID uint, amount float64, operation string, maxRetries int) error {
	var lastErr error

	for i := 0; i < maxRetries; i++ {
		if err := UpdateUserBalanceWithLock(c, userID, amount, operation); err != nil {
			lastErr = err
			global.GVA_LOG.Warn("Balance update failed, retrying",
				zap.Error(err),
				zap.Int("attempt", i+1),
				zap.Int("maxRetries", maxRetries),
				zap.Uint("userId", userID))

			// 等待后重试
			time.Sleep(time.Duration(i+1) * 100 * time.Millisecond)
			continue
		}

		// 操作成功
		return nil
	}

	return fmt.Errorf("balance update failed after %d retries: %v", maxRetries, lastErr)
}

// GetUserBalance 获取用户余额
func GetUserBalance(c *gin.Context, userID uint) (float64, error) {
	userKey := fmt.Sprintf("user_%d", userID)
	userData, err := global.GVA_REDIS.Get(c, userKey).Result()
	if err != nil {
		return 0, fmt.Errorf("failed to get user data: %v", err)
	}

	var user system.ApiSysUser
	if err := json.Unmarshal([]byte(userData), &user); err != nil {
		return 0, fmt.Errorf("failed to parse user data: %v", err)
	}

	return user.Balance, nil
}

// CheckUserBalance 检查用户余额是否足够
func CheckUserBalance(c *gin.Context, userID uint, requiredAmount float64) (bool, error) {
	balance, err := GetUserBalance(c, userID)
	if err != nil {
		return false, err
	}

	return balance >= requiredAmount, nil
}

// DeductUserBalance 扣减用户余额（带锁保护）
func DeductUserBalance(c *gin.Context, userID uint, amount float64, operation string) error {
	if amount <= 0 {
		return fmt.Errorf("invalid deduction amount: %f", amount)
	}

	return UpdateUserBalanceWithLock(c, userID, -amount, operation)
}

// AddUserBalance 增加用户余额（带锁保护）
func AddUserBalance(c *gin.Context, userID uint, amount float64, operation string) error {
	if amount <= 0 {
		return fmt.Errorf("invalid addition amount: %f", amount)
	}

	return UpdateUserBalanceWithLock(c, userID, amount, operation)
}

// TransferUserBalance 用户间余额转账（带锁保护）
func TransferUserBalance(c *gin.Context, fromUserID uint, toUserID uint, amount float64, operation string) error {
	if amount <= 0 {
		return fmt.Errorf("invalid transfer amount: %f", amount)
	}

	// 为两个用户获取锁（按ID排序避免死锁）
	var firstLockKey, secondLockKey string

	if fromUserID < toUserID {
		firstLockKey = fmt.Sprintf("user_balance_lock_%d", fromUserID)
		secondLockKey = fmt.Sprintf("user_balance_lock_%d", toUserID)
	} else {
		firstLockKey = fmt.Sprintf("user_balance_lock_%d", toUserID)
		secondLockKey = fmt.Sprintf("user_balance_lock_%d", fromUserID)
	}

	// 获取第一个锁
	locked1, err := global.GVA_REDIS.SetNX(c, firstLockKey, "1", 10*time.Second).Result()
	if err != nil {
		return fmt.Errorf("failed to acquire first lock: %v", err)
	}
	if !locked1 {
		return fmt.Errorf("failed to acquire first lock")
	}

	// 获取第二个锁
	locked2, err := global.GVA_REDIS.SetNX(c, secondLockKey, "1", 10*time.Second).Result()
	if err != nil {
		global.GVA_REDIS.Del(c, firstLockKey)
		return fmt.Errorf("failed to acquire second lock: %v", err)
	}
	if !locked2 {
		global.GVA_REDIS.Del(c, firstLockKey)
		return fmt.Errorf("failed to acquire second lock")
	}

	// 确保释放所有锁
	defer func() {
		global.GVA_REDIS.Del(c, firstLockKey)
		global.GVA_REDIS.Del(c, secondLockKey)
	}()

	// 执行转账操作
	if err := DeductUserBalance(c, fromUserID, amount, operation+"_deduct"); err != nil {
		return fmt.Errorf("failed to deduct from user %d: %v", fromUserID, err)
	}

	if err := AddUserBalance(c, toUserID, amount, operation+"_add"); err != nil {
		// 如果增加失败，尝试回滚扣款
		AddUserBalance(c, fromUserID, amount, operation+"_rollback")
		return fmt.Errorf("failed to add to user %d: %v", toUserID, err)
	}

	global.GVA_LOG.Info("Successfully transferred balance between users",
		zap.Uint("fromUserId", fromUserID),
		zap.Uint("toUserId", toUserID),
		zap.Float64("amount", amount),
		zap.String("operation", operation))

	return nil
}
