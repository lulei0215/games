package api

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/api"
	apiReq "github.com/flipped-aurora/gin-vue-admin/server/model/api/request"
)

type PaymentCallbacksService struct{}

// CreatePaymentCallbacks paymentCallbacks
// Author [yourname](https://github.com/yourname)
func (paymentCallbacksService *PaymentCallbacksService) CreatePaymentCallbacks(ctx context.Context, paymentCallbacks *api.PaymentCallbacks) (err error) {
	err = global.GVA_DB.Create(paymentCallbacks).Error
	return err
}

// DeletePaymentCallbacks paymentCallbacks
// Author [yourname](https://github.com/yourname)
func (paymentCallbacksService *PaymentCallbacksService) DeletePaymentCallbacks(ctx context.Context, id string) (err error) {
	err = global.GVA_DB.Delete(&api.PaymentCallbacks{}, "id = ?", id).Error
	return err
}

// DeletePaymentCallbacksByIds paymentCallbacks
// Author [yourname](https://github.com/yourname)
func (paymentCallbacksService *PaymentCallbacksService) DeletePaymentCallbacksByIds(ctx context.Context, ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]api.PaymentCallbacks{}, "id in ?", ids).Error
	return err
}

// UpdatePaymentCallbacks paymentCallbacks
// Author [yourname](https://github.com/yourname)
func (paymentCallbacksService *PaymentCallbacksService) UpdatePaymentCallbacks(ctx context.Context, paymentCallbacks api.PaymentCallbacks) (err error) {
	err = global.GVA_DB.Model(&api.PaymentCallbacks{}).Where("id = ?", paymentCallbacks.Id).Updates(&paymentCallbacks).Error
	return err
}

// GetPaymentCallbacks idpaymentCallbacks
// Author [yourname](https://github.com/yourname)
func (paymentCallbacksService *PaymentCallbacksService) GetPaymentCallbacks(ctx context.Context, id string) (paymentCallbacks api.PaymentCallbacks, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&paymentCallbacks).Error
	return
}

// GetPaymentCallbacksInfoList paymentCallbacks
// Author [yourname](https://github.com/yourname)
func (paymentCallbacksService *PaymentCallbacksService) GetPaymentCallbacksInfoList(ctx context.Context, info apiReq.PaymentCallbacksSearch) (list []api.PaymentCallbacks, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// db
	db := global.GVA_DB.Model(&api.PaymentCallbacks{})
	var paymentCallbackss []api.PaymentCallbacks

	// 添加所有字段的查询条件
	if info.Id > 0 {
		db = db.Where("id = ?", info.Id)
	}
	if info.MerchantOrderNo != "" {
		db = db.Where("merchant_order_no LIKE ?", "%"+info.MerchantOrderNo+"%")
	}
	if info.OrderNo != "" {
		db = db.Where("order_no LIKE ?", "%"+info.OrderNo+"%")
	}
	if info.CallbackType > 0 {
		db = db.Where("callback_type = ?", info.CallbackType)
	}
	if info.MerchantId != "" {
		db = db.Where("merchant_id LIKE ?", "%"+info.MerchantId+"%")
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
	if info.RefCpf != "" {
		db = db.Where("ref_cpf LIKE ?", "%"+info.RefCpf+"%")
	}
	if info.RefName != "" {
		db = db.Where("ref_name LIKE ?", "%"+info.RefName+"%")
	}
	if info.ErrorMsg != "" {
		db = db.Where("error_msg LIKE ?", "%"+info.ErrorMsg+"%")
	}
	if info.CallbackData != "" {
		db = db.Where("callback_data LIKE ?", "%"+info.CallbackData+"%")
	}
	if info.Sign != "" {
		db = db.Where("sign LIKE ?", "%"+info.Sign+"%")
	}
	if info.IpAddress != "" {
		db = db.Where("ip_address LIKE ?", "%"+info.IpAddress+"%")
	}
	if info.UserAgent != "" {
		db = db.Where("user_agent LIKE ?", "%"+info.UserAgent+"%")
	}
	if info.ErrorReason != "" {
		db = db.Where("error_reason LIKE ?", "%"+info.ErrorReason+"%")
	}
	if info.Remark != "" {
		db = db.Where("remark LIKE ?", "%"+info.Remark+"%")
	}
	if info.ProcessedTime != nil {
		db = db.Where("processed_time = ?", info.ProcessedTime)
	}
	if info.LastRetryTime != nil {
		db = db.Where("last_retry_time = ?", info.LastRetryTime)
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

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&paymentCallbackss).Error
	return paymentCallbackss, total, err
}
func (paymentCallbacksService *PaymentCallbacksService) GetPaymentCallbacksPublic(ctx context.Context) {
	//
	//
}
func (paymentCallbacksService *PaymentCallbacksService) Create(ctx context.Context, paymentCallbacks api.PaymentCallbacks) (err error) {
	err = global.GVA_DB.Create(&paymentCallbacks).Error
	return err
}
