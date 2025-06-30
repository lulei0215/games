package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type UserWithdrawalAccountsSearch struct {
	request.PageInfo
}

type CreateWithdrawalAccountRequest struct {
	AccountName   string `json:"accountName" binding:"required" validate:"required" comment:"AccountName"`
	AccountType   string `json:"accountType" binding:"required" validate:"required,oneof=PIX_PHONE PIX_EMAIL PIX_CPF PIX_CNPJ" comment:"AccountType"`
	AccountNumber string `json:"accountNumber" binding:"required" validate:"required" comment:"AccountNumber"`
	CpfNumber     string `json:"cpfNumber" validate:"len=11" comment:"CpfNumber"`
	IsDefault     *bool  `json:"isDefault" comment:"isDefault"`
}

type UpdateWithdrawalAccountRequest struct {
	Id            int    `json:"id" binding:"required" validate:"required" comment:"ID"`
	AccountName   string `json:"accountName" binding:"required" validate:"required" comment:"AccountName"`
	AccountType   string `json:"accountType" binding:"required" validate:"required,oneof=PIX_PHONE PIX_EMAIL PIX_CPF PIX_CNPJ" comment:"AccountType"`
	AccountNumber string `json:"accountNumber" binding:"required" validate:"required" comment:"AccountNumber"`
	CpfNumber     string `json:"cpfNumber" validate:"len=11" comment:"CpfNumber"`
	IsDefault     *bool  `json:"isDefault" comment:"isDefault"`
	Status        *bool  `json:"status" comment:"Status"`
}

// SetDefaultAccountRequest 设置默认账户请求
type SetDefaultAccountRequest struct {
	Id int `json:"id" binding:"required" validate:"required" comment:"账户ID"`
}
