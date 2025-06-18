package utils

// "context"
// "time"
// "encoding/json"
// "github.com/flipped-aurora/gin-vue-admin/server/global"
// "github.com/flipped-aurora/gin-vue-admin/server/model/system"
// "go.uber.org/zap"
// "github.com/flipped-aurora/gin-vue-admin/server/model/api"

func BettingAdd() {
	// for {
	// 	// 从Redis队列中获取数据
	// 	result, err := global.GVA_REDIS.BRPop(context.Background(), 0, "balance_updates").Result()
	// 	if err != nil {
	// 		global.GVA_LOG.Error("Failed to get balance update from queue", zap.Error(err))
	// 		time.Sleep(time.Second) // 出错时等待一秒再重试
	// 		continue
	// 	}

	// 	// 解析JSON数据
	// 	var update map[string]interface{}
	// 	if err := json.Unmarshal([]byte(result[1]), &update); err != nil {
	// 		global.GVA_LOG.Error("Failed to unmarshal balance update", zap.Error(err))
	// 		continue
	// 	}

	// 	// 创建交易记录
	// 	transaction := api.SysTransactions{
	// 		UserId:    update["user_id"],
	// 		Type:      update["type"].(string),
	// 		Amount:    float64(update["coin"].(float64)),
	// 		Amount:   update["balance"].(float64),
	// 		// Room:      update["room"].(string),
	// 		CreatedAt: time.Now(),
	// 	}

	// 	// 插入MySQL
	// 	if err := global.GVA_DB.Create(&transaction).Error; err != nil {
	// 		global.GVA_LOG.Error("Failed to insert transaction into MySQL", zap.Error(err))
	// 		// 如果插入失败，将数据重新放回队列
	// 		global.GVA_REDIS.LPush(context.Background(), "balance_updates", result[1])
	// 		continue
	// 	}

	// 	// 更新用户余额
	// 	if err := global.GVA_DB.Model(&system.SysUser{}).
	// 		Where("id = ?", transaction.UserID).
	// 		Update("balance", transaction.Balance).Error; err != nil {
	// 		global.GVA_LOG.Error("Failed to update user balance in MySQL", zap.Error(err))
	// 	}
	// }
}
