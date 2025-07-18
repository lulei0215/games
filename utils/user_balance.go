package utils

import (
	"encoding/json"
	"fmt"
	"math"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// UpdateUserBalanceWithLock 使用MySQL数据库锁安全地更新用户余额
func UpdateUserBalanceWithLock(c *gin.Context, userID uint, amount float64, operation string, transactionCode string) error {
	// 使用数据库事务和行锁来安全更新余额
	err := global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		// 使用FOR UPDATE锁来防止并发更新
		var user system.SysUser
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("id = ?", userID).First(&user).Error; err != nil {
			global.GVA_LOG.Error("Failed to get user with lock",
				zap.Error(err),
				zap.Uint("userId", userID))
			return err
		}

		// 记录原始余额
		originalBalance := user.Balance

		// 检查余额是否足够（如果是扣减操作）
		if amount < 0 && user.Balance < -amount {
			global.GVA_LOG.Warn("Insufficient balance",
				zap.Uint("userId", userID),
				zap.Float64("currentBalance", user.Balance),
				zap.Float64("requestedAmount", -amount))
			return fmt.Errorf("insufficient balance")
		}

		// 计算新余额
		newBalance := user.Balance + amount

		// 检查余额是否为负数
		if newBalance < 0 {
			global.GVA_LOG.Error("Balance calculation results in negative value, set to 0",
				zap.Uint("userId", userID),
				zap.Float64("originalBalance", originalBalance),
				zap.Float64("amount", amount),
				zap.String("operation", operation),
				zap.String("transactionCode", transactionCode))
			newBalance = 0
		}

		// 四舍五入到2位小数
		finalBalance := math.Round(newBalance*100) / 100
		user.Balance = finalBalance

		// 更新数据库中的用户余额
		if err := tx.Save(&user).Error; err != nil {
			global.GVA_LOG.Error("Failed to update user balance in database",
				zap.Error(err),
				zap.Uint("userId", userID),
				zap.Float64("originalBalance", originalBalance),
				zap.Float64("newBalance", finalBalance),
				zap.String("operation", operation),
				zap.String("transactionCode", transactionCode))
			return err
		}

		// 同时更新Redis缓存
		userJson, err := json.Marshal(user)
		if err != nil {
			global.GVA_LOG.Error("Failed to marshal updated user data",
				zap.Error(err),
				zap.Uint("userId", userID))
			return err
		}

		// 更新Redis缓存
		err = global.GVA_REDIS.Set(c, fmt.Sprintf("user_%d", userID), string(userJson), 0).Err()
		if err != nil {
			global.GVA_LOG.Error("Failed to update user data in Redis",
				zap.Error(err),
				zap.Uint("userId", userID))
			// 注意：Redis更新失败不影响数据库事务，但会记录错误
		}

		// 记录详细的余额更新日志
		global.GVA_LOG.Info("Successfully updated user balance with lock",
			zap.Uint("userId", userID),
			zap.String("operation", operation),
			zap.String("transactionCode", transactionCode),
			zap.Float64("originalBalance", originalBalance),
			zap.Float64("amount", amount),
			zap.Float64("finalBalance", finalBalance))

		return nil
	})

	return err
}

// UpdateUserBalanceWithRetry 带重试机制的余额更新
func UpdateUserBalanceWithRetry(c *gin.Context, userID uint, amount float64, operation string, maxRetries int, transactionCode string) error {
	var lastErr error

	for i := 0; i < maxRetries; i++ {
		if err := UpdateUserBalanceWithLock(c, userID, amount, operation, transactionCode); err != nil {
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

// DeductUserBalance 扣减用户余额
func DeductUserBalance(c *gin.Context, userID uint, amount float64, operation string, transactionCode string) error {
	return UpdateUserBalanceWithLock(c, userID, -amount, operation, transactionCode)
}

// AddUserBalance 增加用户余额
func AddUserBalance(c *gin.Context, userID uint, amount float64, operation string, transactionCode string) error {
	return UpdateUserBalanceWithLock(c, userID, amount, operation, transactionCode)
}

// TransferUserBalance 用户间转账
func TransferUserBalance(c *gin.Context, firstUserID uint, secondUserID uint, amount float64, operation string, transactionCode string) error {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		// 使用FOR UPDATE锁来防止并发更新
		var firstUser system.SysUser
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("id = ?", firstUserID).First(&firstUser).Error; err != nil {
			global.GVA_LOG.Error("Failed to get first user with lock",
				zap.Error(err),
				zap.Uint("firstUserId", firstUserID))
			return err
		}

		var secondUser system.SysUser
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("id = ?", secondUserID).First(&secondUser).Error; err != nil {
			global.GVA_LOG.Error("Failed to get second user with lock",
				zap.Error(err),
				zap.Uint("secondUserId", secondUserID))
			return err
		}

		// 记录原始余额
		firstUserOriginalBalance := firstUser.Balance
		secondUserOriginalBalance := secondUser.Balance

		// 检查第一个用户余额是否足够
		if firstUser.Balance < amount {
			global.GVA_LOG.Warn("Insufficient balance for transfer",
				zap.Uint("firstUserId", firstUserID),
				zap.Float64("currentBalance", firstUser.Balance),
				zap.Float64("transferAmount", amount))
			return fmt.Errorf("insufficient balance")
		}

		// 执行转账
		firstUser.Balance = math.Round((firstUser.Balance-amount)*100) / 100
		secondUser.Balance = math.Round((secondUser.Balance+amount)*100) / 100

		// 更新数据库
		if err := tx.Save(&firstUser).Error; err != nil {
			global.GVA_LOG.Error("Failed to update first user balance in database",
				zap.Error(err),
				zap.Uint("firstUserId", firstUserID))
			return err
		}

		if err := tx.Save(&secondUser).Error; err != nil {
			global.GVA_LOG.Error("Failed to update second user balance in database",
				zap.Error(err),
				zap.Uint("secondUserId", secondUserID))
			return err
		}

		// 同时更新Redis缓存
		firstUserJson, err := json.Marshal(firstUser)
		if err != nil {
			global.GVA_LOG.Error("Failed to marshal first user data",
				zap.Error(err),
				zap.Uint("firstUserId", firstUserID))
			return err
		}

		secondUserJson, err := json.Marshal(secondUser)
		if err != nil {
			global.GVA_LOG.Error("Failed to marshal second user data",
				zap.Error(err),
				zap.Uint("secondUserId", secondUserID))
			return err
		}

		// 更新Redis缓存
		global.GVA_REDIS.Set(c, fmt.Sprintf("user_%d", firstUserID), string(firstUserJson), 0)
		global.GVA_REDIS.Set(c, fmt.Sprintf("user_%d", secondUserID), string(secondUserJson), 0)

		// 记录详细的转账日志
		global.GVA_LOG.Info("Successfully transferred balance between users",
			zap.Uint("firstUserId", firstUserID),
			zap.Uint("secondUserId", secondUserID),
			zap.String("operation", operation),
			zap.String("transactionCode", transactionCode),
			zap.Float64("transferAmount", amount),
			zap.Float64("firstUserOriginalBalance", firstUserOriginalBalance),
			zap.Float64("firstUserNewBalance", firstUser.Balance),
			zap.Float64("secondUserOriginalBalance", secondUserOriginalBalance),
			zap.Float64("secondUserNewBalance", secondUser.Balance))

		return nil
	})
}

// CreatePaymentWithBalanceDeduction 在同一个事务中创建支付记录并扣减用户余额
func CreatePaymentWithBalanceDeduction(c *gin.Context, userID uint, amount float64, operation string, createFunc func(tx *gorm.DB) error, transactionCode string) error {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		// 使用FOR UPDATE锁来防止并发更新用户余额
		var user system.SysUser
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("id = ?", userID).First(&user).Error; err != nil {
			global.GVA_LOG.Error("Failed to get user with lock",
				zap.Error(err),
				zap.Uint("userId", userID))
			return err
		}

		// 记录原始余额
		originalBalance := user.Balance

		// 检查余额是否足够
		if user.Balance < amount {
			global.GVA_LOG.Warn("Insufficient balance",
				zap.Uint("userId", userID),
				zap.Float64("currentBalance", user.Balance),
				zap.Float64("requestedAmount", amount))
			return fmt.Errorf("insufficient balance")
		}

		// 计算新余额
		newBalance := user.Balance - amount

		// 检查余额是否为负数
		if newBalance < 0 {
			global.GVA_LOG.Error("Balance calculation results in negative value, set to 0",
				zap.Uint("userId", userID),
				zap.Float64("originalBalance", originalBalance),
				zap.Float64("deductAmount", amount),
				zap.String("operation", operation),
				zap.String("transactionCode", transactionCode))
			newBalance = 0
		}

		// 四舍五入到2位小数
		finalBalance := math.Round(newBalance*100) / 100
		user.Balance = finalBalance

		// 先创建其他数据库记录（如支付交易记录）
		if err := createFunc(tx); err != nil {
			global.GVA_LOG.Error("Failed to create database record in transaction",
				zap.Error(err),
				zap.Uint("userId", userID),
				zap.String("operation", operation),
				zap.String("transactionCode", transactionCode))
			return err
		}

		// 更新数据库中的用户余额
		if err := tx.Save(&user).Error; err != nil {
			global.GVA_LOG.Error("Failed to update user balance in database",
				zap.Error(err),
				zap.Uint("userId", userID),
				zap.Float64("originalBalance", originalBalance),
				zap.Float64("newBalance", finalBalance),
				zap.String("operation", operation),
				zap.String("transactionCode", transactionCode))
			return err
		}

		// 同时更新Redis缓存
		userJson, err := json.Marshal(user)
		if err != nil {
			global.GVA_LOG.Error("Failed to marshal updated user data",
				zap.Error(err),
				zap.Uint("userId", userID))
			return err
		}

		// 更新Redis缓存
		err = global.GVA_REDIS.Set(c, fmt.Sprintf("user_%d", userID), string(userJson), 0).Err()
		if err != nil {
			global.GVA_LOG.Error("Failed to update user data in Redis",
				zap.Error(err),
				zap.Uint("userId", userID))
			// 注意：Redis更新失败不影响数据库事务，但会记录错误
		}

		// 记录详细的余额更新日志
		global.GVA_LOG.Info("Successfully created payment record and deducted user balance in transaction",
			zap.Uint("userId", userID),
			zap.String("operation", operation),
			zap.String("transactionCode", transactionCode),
			zap.Float64("originalBalance", originalBalance),
			zap.Float64("deductAmount", amount),
			zap.Float64("finalBalance", finalBalance))

		return nil
	})
}

// CancelPaymentWithBalanceAddition 在同一个事务中取消支付记录并加回用户余额
func CancelPaymentWithBalanceAddition(c *gin.Context, userID uint, amount float64, operation string, updateFunc func(tx *gorm.DB) error, transactionCode string) error {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		// 使用FOR UPDATE锁来防止并发更新用户余额
		var user system.SysUser
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("id = ?", userID).First(&user).Error; err != nil {
			global.GVA_LOG.Error("Failed to get user with lock",
				zap.Error(err),
				zap.Uint("userId", userID))
			return err
		}

		// 记录原始余额
		originalBalance := user.Balance

		// 计算新余额
		newBalance := user.Balance + amount

		// 四舍五入到2位小数
		finalBalance := math.Round(newBalance*100) / 100
		user.Balance = finalBalance

		// 先更新其他数据库记录（如支付交易状态）
		if err := updateFunc(tx); err != nil {
			global.GVA_LOG.Error("Failed to update database record in transaction",
				zap.Error(err),
				zap.Uint("userId", userID),
				zap.String("operation", operation),
				zap.String("transactionCode", transactionCode))
			return err
		}

		// 更新数据库中的用户余额
		if err := tx.Save(&user).Error; err != nil {
			global.GVA_LOG.Error("Failed to update user balance in database",
				zap.Error(err),
				zap.Uint("userId", userID),
				zap.Float64("originalBalance", originalBalance),
				zap.Float64("newBalance", finalBalance),
				zap.String("operation", operation),
				zap.String("transactionCode", transactionCode))
			return err
		}

		// 同时更新Redis缓存
		userJson, err := json.Marshal(user)
		if err != nil {
			global.GVA_LOG.Error("Failed to marshal updated user data",
				zap.Error(err),
				zap.Uint("userId", userID))
			return err
		}

		// 更新Redis缓存
		err = global.GVA_REDIS.Set(c, fmt.Sprintf("user_%d", userID), string(userJson), 0).Err()
		if err != nil {
			global.GVA_LOG.Error("Failed to update user data in Redis",
				zap.Error(err),
				zap.Uint("userId", userID))
			// 注意：Redis更新失败不影响数据库事务，但会记录错误
		}

		// 记录详细的余额更新日志
		global.GVA_LOG.Info("Successfully cancelled payment record and added user balance in transaction",
			zap.Uint("userId", userID),
			zap.String("operation", operation),
			zap.String("transactionCode", transactionCode),
			zap.Float64("originalBalance", originalBalance),
			zap.Float64("addAmount", amount),
			zap.Float64("finalBalance", finalBalance))

		return nil
	})
}
