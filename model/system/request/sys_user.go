package request

import (
	common "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

// Register User register structure
type Register struct {
	Username     string `json:"userName" example:""`
	Password     string `json:"passWord" example:""`
	NickName     string `json:"nickName" example:""`
	HeaderImg    string `json:"headerImg" example:""`
	AuthorityId  uint   `json:"authorityId" swaggertype:"string" example:"int id"`
	Enable       int    `json:"enable" swaggertype:"string" example:"int "`
	AuthorityIds []uint `json:"authorityIds" swaggertype:"string" example:"[]uint id"`
	Phone        string `json:"phone" example:""`
	Email        string `json:"email" example:""`
	Uuid         string `json:"uuid" example:""`
}

// Login User login structure
type Login struct {
	Username  string `json:"username"`  //
	Password  string `json:"password"`  //
	Captcha   string `json:"captcha"`   //
	CaptchaId string `json:"captchaId"` // ID
}
type ApiLogin struct {
	Username string `json:"username"` //
	Password string `json:"password"` //
}

// ChangePasswordReq Modify password structure
type ChangePasswordReq struct {
	ID          uint   `json:"-"`           //  JWT  user id，
	Password    string `json:"password"`    //
	NewPassword string `json:"newPassword"` //
}
type VerifyWithdrawPasswordReq struct {
	Password      string `json:"password"`      //
	LoginPassword string `json:"loginPassword"` //
}

type ResetPassword struct {
	ID       uint   `json:"ID" form:"ID"`
	Password string `json:"password" form:"password" gorm:"comment:"` //
}
type WithdrawPassword struct {
	ID       uint   `json:"ID" form:"ID"`
	Password string `json:"password" form:"password" gorm:"comment:"` //
}

// SetUserAuth Modify user's auth structure
type SetUserAuth struct {
	AuthorityId uint `json:"authorityId"` // ID
}
type SendEmailCodeRequest struct {
	Email string `json:"email" form:"email"` // email
}
type BindEmailRequest struct {
	Email string `json:"email" form:"email"` // email
	Code  string `json:"code" form:"code"`   // email
}
type CryptRequest struct {
	Data string `json:"data" form:"data"` // email
	Iv   string `json:"iv" form:"iv"`     // email
}

// SetUserAuthorities Modify user's auth structure
type SetUserAuthorities struct {
	ID           uint
	AuthorityIds []uint `json:"authorityIds"` // ID
}

type ChangeUserInfo struct {
	ID           uint                  `gorm:"primarykey"`                                                                       // ID
	NickName     string                `json:"nickName" gorm:"default:;comment:"`                                                //
	Phone        string                `json:"phone"  gorm:"comment:"`                                                           //
	AuthorityIds []uint                `json:"authorityIds" gorm:"-"`                                                            // ID
	Email        string                `json:"email"  gorm:"comment:"`                                                           //
	HeaderImg    string                `json:"headerImg" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:"` //
	SideMode     string                `json:"sideMode"  gorm:"comment:"`                                                        //
	Enable       int                   `json:"enable" gorm:"comment:"`                                                           //
	Authorities  []system.SysAuthority `json:"-" gorm:"many2many:sys_user_authority;"`
}

type GetUserList struct {
	common.PageInfo
	Username string `json:"username" form:"username"`
	NickName string `json:"nickName" form:"nickName"`
	Phone    string `json:"phone" form:"phone"`
	Email    string `json:"email" form:"email"`
}

type Betting struct {
	Coin uint `json:"coin"`
}
