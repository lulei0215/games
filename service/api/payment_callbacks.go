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
	//

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
