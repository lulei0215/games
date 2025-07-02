package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/api"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(api.SysTransactions{}, system.UserAgentRelation{}, api.MonitorTransfers{}, api.PaymentTransactions{}, api.PaymentCallbacks{}, api.UserWithdrawalAccounts{}, api.UserAgentRelation{}, api.Games{})
	if err != nil {
		return err
	}
	return nil
}
