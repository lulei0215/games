package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

// RedisSafeUpdater Redis并发安全更新器
type RedisSafeUpdater struct {
	redis *redis.Client
}

// NewRedisSafeUpdater 创建Redis安全更新器
func NewRedisSafeUpdater(redis *redis.Client) *RedisSafeUpdater {
	return &RedisSafeUpdater{
		redis: redis,
	}
}

// UpdateWithLock 使用分布式锁进行安全更新
func (r *RedisSafeUpdater) UpdateWithLock(ctx context.Context, key string, updateFunc func(string) (string, error), lockTimeout time.Duration) error {
	// 生成锁key
	lockKey := fmt.Sprintf("%s:lock", key)

	// 尝试获取锁
	locked, err := r.redis.SetNX(ctx, lockKey, "1", lockTimeout).Result()
	if err != nil {
		return fmt.Errorf("failed to acquire lock: %v", err)
	}
	if !locked {
		return fmt.Errorf("failed to acquire lock, key is locked")
	}

	// 确保释放锁
	defer r.redis.Del(ctx, lockKey)

	// 获取当前数据
	currentData, err := r.redis.Get(ctx, key).Result()
	if err != nil && err != redis.Nil {
		return fmt.Errorf("failed to get current data: %v", err)
	}

	// 执行更新函数
	newData, err := updateFunc(currentData)
	if err != nil {
		return fmt.Errorf("update function failed: %v", err)
	}

	// 保存更新后的数据
	err = r.redis.Set(ctx, key, newData, 0).Err()
	if err != nil {
		return fmt.Errorf("failed to save updated data: %v", err)
	}

	return nil
}

// UpdateWithVersion 使用版本号进行乐观锁更新
func (r *RedisSafeUpdater) UpdateWithVersion(ctx context.Context, dataKey, versionKey string, expectedVersion int64, updateFunc func(string) (string, error)) error {
	// 使用Lua脚本进行原子性操作
	luaScript := `
		local dataKey = KEYS[1]
		local versionKey = KEYS[2]
		local expectedVersion = tonumber(ARGV[1])
		local newData = ARGV[2]
		local newVersion = tonumber(ARGV[3])
		
		-- 检查版本号
		local currentVersion = redis.call('GET', versionKey)
		if currentVersion and tonumber(currentVersion) ~= expectedVersion then
			return {err = "VERSION_MISMATCH"}
		end
		
		-- 原子性更新数据和版本号
		redis.call('SET', dataKey, newData)
		redis.call('SET', versionKey, newVersion)
		return {ok = "SUCCESS"}
	`

	// 获取当前数据
	currentData, err := r.redis.Get(ctx, dataKey).Result()
	if err != nil && err != redis.Nil {
		return fmt.Errorf("failed to get current data: %v", err)
	}

	// 执行更新函数
	newData, err := updateFunc(currentData)
	if err != nil {
		return fmt.Errorf("update function failed: %v", err)
	}

	// 执行Lua脚本
	result, err := r.redis.Eval(ctx, luaScript, []string{dataKey, versionKey},
		expectedVersion, newData, expectedVersion+1).Result()
	if err != nil {
		return fmt.Errorf("lua script execution failed: %v", err)
	}

	// 检查结果
	if resultMap, ok := result.([]interface{}); ok && len(resultMap) > 0 {
		if errMsg, ok := resultMap[0].(string); ok && errMsg == "VERSION_MISMATCH" {
			return fmt.Errorf("version mismatch, data has been modified by another request")
		}
	}

	return nil
}

// UpdateUserData 并发安全的用户数据更新
func (r *RedisSafeUpdater) UpdateUserData(ctx context.Context, userID uint, updateData map[string]interface{}) error {
	userKey := fmt.Sprintf("user_%d", userID)

	return r.UpdateWithLock(ctx, userKey, func(currentData string) (string, error) {
		var user map[string]interface{}

		// 解析当前数据
		if currentData != "" {
			if err := json.Unmarshal([]byte(currentData), &user); err != nil {
				return "", fmt.Errorf("failed to unmarshal current data: %v", err)
			}
		} else {
			user = make(map[string]interface{})
		}

		// 更新指定字段
		for key, value := range updateData {
			user[key] = value
		}

		// 序列化新数据
		newData, err := json.Marshal(user)
		if err != nil {
			return "", fmt.Errorf("failed to marshal updated data: %v", err)
		}

		return string(newData), nil
	}, 10*time.Second)
}

// UpdateUserDataWithVersion 带版本号的用户数据更新
func (r *RedisSafeUpdater) UpdateUserDataWithVersion(ctx context.Context, userID uint, updateData map[string]interface{}, expectedVersion int64) error {
	userKey := fmt.Sprintf("user_%d", userID)
	versionKey := fmt.Sprintf("user_version_%d", userID)

	return r.UpdateWithVersion(ctx, userKey, versionKey, expectedVersion, func(currentData string) (string, error) {
		var user map[string]interface{}

		// 解析当前数据
		if currentData != "" {
			if err := json.Unmarshal([]byte(currentData), &user); err != nil {
				return "", fmt.Errorf("failed to unmarshal current data: %v", err)
			}
		} else {
			user = make(map[string]interface{})
		}

		// 更新指定字段
		for key, value := range updateData {
			user[key] = value
		}

		// 序列化新数据
		newData, err := json.Marshal(user)
		if err != nil {
			return "", fmt.Errorf("failed to marshal updated data: %v", err)
		}

		return string(newData), nil
	})
}

// GetUserDataWithVersion 获取用户数据和版本号
func (r *RedisSafeUpdater) GetUserDataWithVersion(ctx context.Context, userID uint) (map[string]interface{}, int64, error) {
	userKey := fmt.Sprintf("user_%d", userID)
	versionKey := fmt.Sprintf("user_version_%d", userID)

	// 获取用户数据
	userData, err := r.redis.Get(ctx, userKey).Result()
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get user data: %v", err)
	}

	// 获取版本号
	version, err := r.redis.Get(ctx, versionKey).Int64()
	if err != nil {
		version = 0
	}

	// 解析用户数据
	var user map[string]interface{}
	if err := json.Unmarshal([]byte(userData), &user); err != nil {
		return nil, 0, fmt.Errorf("failed to unmarshal user data: %v", err)
	}

	return user, version, nil
}

// BatchUpdateWithLock 批量更新（使用锁）
func (r *RedisSafeUpdater) BatchUpdateWithLock(ctx context.Context, updates map[string]func(string) (string, error), lockTimeout time.Duration) error {
	// 为所有key获取锁
	locks := make([]string, 0, len(updates))
	for key := range updates {
		lockKey := fmt.Sprintf("%s:lock", key)
		locked, err := r.redis.SetNX(ctx, lockKey, "1", lockTimeout).Result()
		if err != nil {
			// 释放已获取的锁
			for _, lock := range locks {
				r.redis.Del(ctx, lock)
			}
			return fmt.Errorf("failed to acquire lock for %s: %v", key, err)
		}
		if !locked {
			// 释放已获取的锁
			for _, lock := range locks {
				r.redis.Del(ctx, lock)
			}
			return fmt.Errorf("failed to acquire lock for %s", key)
		}
		locks = append(locks, lockKey)
	}

	// 确保释放所有锁
	defer func() {
		for _, lock := range locks {
			r.redis.Del(ctx, lock)
		}
	}()

	// 执行批量更新
	for key, updateFunc := range updates {
		currentData, err := r.redis.Get(ctx, key).Result()
		if err != nil && err != redis.Nil {
			return fmt.Errorf("failed to get current data for %s: %v", key, err)
		}

		newData, err := updateFunc(currentData)
		if err != nil {
			return fmt.Errorf("update function failed for %s: %v", key, err)
		}

		err = r.redis.Set(ctx, key, newData, 0).Err()
		if err != nil {
			return fmt.Errorf("failed to save updated data for %s: %v", key, err)
		}

		global.GVA_LOG.Info("Successfully updated Redis key",
			zap.String("key", key),
			zap.String("newData", newData))
	}

	return nil
}

// RetryWithBackoff 带重试机制的Redis操作
func (r *RedisSafeUpdater) RetryWithBackoff(ctx context.Context, operation func() error, maxRetries int, initialDelay time.Duration) error {
	var lastErr error
	delay := initialDelay

	for i := 0; i < maxRetries; i++ {
		if err := operation(); err != nil {
			lastErr = err
			global.GVA_LOG.Warn("Redis operation failed, retrying",
				zap.Error(err),
				zap.Int("attempt", i+1),
				zap.Int("maxRetries", maxRetries),
				zap.Duration("delay", delay))

			// 等待后重试
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(delay):
				// 指数退避
				delay = time.Duration(float64(delay) * 1.5)
				continue
			}
		}

		// 操作成功
		return nil
	}

	return fmt.Errorf("operation failed after %d retries: %v", maxRetries, lastErr)
}
