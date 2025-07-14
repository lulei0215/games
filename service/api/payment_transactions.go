package api

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/api"
	apiReq "github.com/flipped-aurora/gin-vue-admin/server/model/api/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

type PaymentTransactionsService struct{}

// CreatePaymentTransactions paymentTransactions
// Author [yourname](https://github.com/yourname)
func (paymentTransactionsService *PaymentTransactionsService) CreatePaymentTransactions(ctx context.Context, paymentTransactions *api.PaymentTransactions) (err error) {
	err = global.GVA_DB.Create(paymentTransactions).Error
	return err
}

// DeletePaymentTransactions paymentTransactions
// Author [yourname](https://github.com/yourname)
func (paymentTransactionsService *PaymentTransactionsService) DeletePaymentTransactions(ctx context.Context, id string) (err error) {
	err = global.GVA_DB.Delete(&api.PaymentTransactions{}, "id = ?", id).Error
	return err
}

// DeletePaymentTransactionsByIds paymentTransactions
// Author [yourname](https://github.com/yourname)
func (paymentTransactionsService *PaymentTransactionsService) DeletePaymentTransactionsByIds(ctx context.Context, ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]api.PaymentTransactions{}, "id in ?", ids).Error
	return err
}

// UpdatePaymentTransactions paymentTransactions
// Author [yourname](https://github.com/yourname)
func (paymentTransactionsService *PaymentTransactionsService) UpdatePaymentTransactions(ctx context.Context, paymentTransactions api.PaymentTransactions) (err error) {
	err = global.GVA_DB.Model(&api.PaymentTransactions{}).Where("id = ?", paymentTransactions.Id).Updates(&paymentTransactions).Error
	return err
}

// GetPaymentTransactions idpaymentTransactions
// Author [yourname](https://github.com/yourname)
func (paymentTransactionsService *PaymentTransactionsService) GetPaymentTransactions(ctx context.Context, id string) (paymentTransactions api.PaymentTransactions, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&paymentTransactions).Error
	return
}

// GetPaymentTransactionsInfoList paymentTransactions
// Author [yourname](https://github.com/yourname)
func (paymentTransactionsService *PaymentTransactionsService) GetPaymentTransactionsInfoList(ctx context.Context, info apiReq.PaymentTransactionsSearch) (list []api.PaymentTransactions, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// db
	db := global.GVA_DB.Model(&api.PaymentTransactions{})
	var paymentTransactionss []api.PaymentTransactions

	// 添加所有字段的查询条件
	if info.Id > 0 {
		db = db.Where("id = ?", info.Id)
	}
	if info.UserId > 0 {
		db = db.Where("user_id = ?", info.UserId)
	}
	if info.MerchantOrderNo != "" {
		db = db.Where("merchant_order_no LIKE ?", "%"+info.MerchantOrderNo+"%")
	}
	if info.OrderNo != "" {
		db = db.Where("order_no LIKE ?", "%"+info.OrderNo+"%")
	}
	if info.TransactionType > 0 {
		db = db.Where("transaction_type = ?", info.TransactionType)
	}
	if info.Amount > 0 {
		db = db.Where("amount = ?", info.Amount)
	}
	if info.Currency != "" {
		db = db.Where("currency LIKE ?", "%"+info.Currency+"%")
	}
	if info.Status != "" {
		db = db.Where("status LIKE ?", "%"+info.Status+"%")
	}
	if info.PayType != "" {
		db = db.Where("pay_type LIKE ?", "%"+info.PayType+"%")
	}
	if info.AccountType != "" {
		db = db.Where("account_type LIKE ?", "%"+info.AccountType+"%")
	}
	if info.AccountNo != "" {
		db = db.Where("account_no LIKE ?", "%"+info.AccountNo+"%")
	}
	if info.AccountName != "" {
		db = db.Where("account_name LIKE ?", "%"+info.AccountName+"%")
	}
	if info.Content != "" {
		db = db.Where("content LIKE ?", "%"+info.Content+"%")
	}
	if info.ClientIp != "" {
		db = db.Where("client_ip LIKE ?", "%"+info.ClientIp+"%")
	}
	if info.CallbackUrl != "" {
		db = db.Where("callback_url LIKE ?", "%"+info.CallbackUrl+"%")
	}
	if info.RedirectUrl != "" {
		db = db.Where("redirect_url LIKE ?", "%"+info.RedirectUrl+"%")
	}
	if info.PayUrl != "" {
		db = db.Where("pay_url LIKE ?", "%"+info.PayUrl+"%")
	}
	if info.PayRaw != "" {
		db = db.Where("pay_raw LIKE ?", "%"+info.PayRaw+"%")
	}
	if info.ErrorMsg != "" {
		db = db.Where("error_msg LIKE ?", "%"+info.ErrorMsg+"%")
	}
	if info.RefCpf != "" {
		db = db.Where("ref_cpf LIKE ?", "%"+info.RefCpf+"%")
	}
	if info.RefName != "" {
		db = db.Where("ref_name LIKE ?", "%"+info.RefName+"%")
	}
	if info.CreatedAtStart != nil {
		db = db.Where("created_at >= ?", info.CreatedAtStart)
	}
	if info.CreatedAtEnd != nil {
		db = db.Where("created_at <= ?", info.CreatedAtEnd)
	}
	if info.UpdatedAtStart != nil {
		db = db.Where("updated_at >= ?", info.UpdatedAtStart)
	}
	if info.UpdatedAtEnd != nil {
		db = db.Where("updated_at <= ?", info.UpdatedAtEnd)
	}
	if info.DeletedAtStart != nil {
		db = db.Where("deleted_at >= ?", info.DeletedAtStart)
	}
	if info.DeletedAtEnd != nil {
		db = db.Where("deleted_at <= ?", info.DeletedAtEnd)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&paymentTransactionss).Error
	return paymentTransactionss, total, err
}
func (paymentTransactionsService *PaymentTransactionsService) GetPaymentTransactionsPublic(ctx context.Context) {
	//
	//
}

func (paymentTransactionsService *PaymentTransactionsService) Create(ctx context.Context, paymentTransactions api.PaymentTransactions) (err error) {
	err = global.GVA_DB.Create(&paymentTransactions).Error
	return err
}

func (paymentTransactionsService *PaymentTransactionsService) TradeOk(ctx context.Context, MerchantOrderNo string, OrderNo string, paymentCallbacks api.PaymentCallbacks) (err error) {

	var paymentTransactions api.PaymentTransactions

	err = global.GVA_DB.Where("merchant_order_no = ? and order_no = ?  and transaction_type = ?", MerchantOrderNo, OrderNo, 1).First(&paymentTransactions).Error
	if err != nil {
		return
	}
	if paymentTransactions.Status == "PAYING" {
		// 更新多个字段，包括从回调中获取的信息
		updateData := map[string]interface{}{
			"status":   "PAID",
			"currency": paymentCallbacks.Currency,
			"pay_type": paymentCallbacks.PayType,
			"ref_cpf":  paymentCallbacks.RefCpf,
			"ref_name": paymentCallbacks.RefName,
		}

		err = global.GVA_DB.Model(&api.PaymentTransactions{}).
			Where("merchant_order_no = ? and order_no = ?", MerchantOrderNo, OrderNo).
			Updates(updateData).Error
		if err != nil {
			return
		}
		var user system.ApiSysUser
		redisuser, _ := global.GVA_REDIS.Get(ctx, fmt.Sprintf("user_%d", paymentTransactions.UserId)).Result()
		if redisuser == "" {
			return err
		}
		err = json.Unmarshal([]byte(redisuser), &user)
		if err != nil {
			return err
		}
		user.Balance = user.Balance + float64(paymentTransactions.Amount/100)

		userJson, err := json.Marshal(user)
		if err != nil {
		} else {
			err = global.GVA_REDIS.Set(ctx, fmt.Sprintf("user_%d", user.ID), string(userJson), 0).Err()
			if err != nil {
			}
		}
		return err
	}

	return err

}

// UpdatePaymentStatus
func (paymentTransactionsService *PaymentTransactionsService) UpdatePaymentStatus(ctx context.Context, merchantOrderNo string, status string) (err error) {
	err = global.GVA_DB.Model(&api.PaymentTransactions{}).
		Where("merchant_order_no = ?", merchantOrderNo).
		Update("status", status).Error
	return err
}

func (paymentTransactionsService *PaymentTransactionsService) PaymentOk(ctx context.Context, MerchantOrderNo string, OrderNo string) (err error) {

	var paymentTransactions api.PaymentTransactions

	err = global.GVA_DB.Where("merchant_order_no = ? and order_no = ? and transaction_type = ?", MerchantOrderNo, OrderNo, 2).First(&paymentTransactions).Error
	if err != nil {
		return
	}
	if paymentTransactions.Status == "WAITING_PAY" {
		err = global.GVA_DB.Model(&api.PaymentTransactions{}).
			Where("merchant_order_no = ? and order_no = ?", MerchantOrderNo, OrderNo).
			Update("status", "PAID").Error
		if err != nil {
			return
		}
		var user system.ApiSysUser
		redisuser, _ := global.GVA_REDIS.Get(ctx, fmt.Sprintf("user_%d", paymentTransactions.UserId)).Result()
		if redisuser == "" {
			return err
		}
		err = json.Unmarshal([]byte(redisuser), &user)
		if err != nil {
			return err
		}
		user.Balance = user.Balance - float64(paymentTransactions.Amount/100)

		userJson, err := json.Marshal(user)
		if err != nil {
		} else {
			err = global.GVA_REDIS.Set(ctx, fmt.Sprintf("user_%d", user.ID), string(userJson), 0).Err()
			if err != nil {
			}
		}
		return err
	}

	return err

}

func (paymentTransactionsService *PaymentTransactionsService) GetPaymentList(ctx context.Context, info apiReq.PaymentTransactionsSearch, uid uint, transactionType int) (list []api.PaymentTransactions, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// db
	db := global.GVA_DB.Model(&api.PaymentTransactions{})
	var paymentTransactionss []api.PaymentTransactions
	//

	err = db.Where("user_id = ? and transaction_type = ?", uid, transactionType).Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&paymentTransactionss).Error
	return paymentTransactionss, total, err
}

// GetByOrderNo 根据订单号获取支付交易记录
func (paymentTransactionsService *PaymentTransactionsService) GetByOrderNo(ctx context.Context, orderNo string) (paymentTransactions api.PaymentTransactions, err error) {
	err = global.GVA_DB.Where("merchant_order_no = ?", orderNo).First(&paymentTransactions).Error
	return
}

// UpdateByOrderNo 根据订单号更新支付交易记录
func (paymentTransactionsService *PaymentTransactionsService) UpdateByOrderNo(ctx context.Context, orderNo string, updateData api.PaymentTransactions) (err error) {
	err = global.GVA_DB.Model(&api.PaymentTransactions{}).
		Where("merchant_order_no = ?", orderNo).
		Updates(updateData).Error
	return err
}
