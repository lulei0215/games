package system

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common"
	"github.com/google/uuid"
)

type Login interface {
	GetUsername() string
	GetNickname() string
	GetUUID() uuid.UUID
	GetUserId() uint
	GetAuthorityId() uint
	GetUserInfo() any
}

var _ Login = new(SysUser)

type SysUser struct {
	global.GVA_MODEL
	UUID             uuid.UUID      `json:"uuid" gorm:"index;comment:用户UUID"`                                                                   // 用户UUID
	Username         string         `json:"userName" gorm:"index;comment:用户登录名"`                                                                // 用户登录名
	Password         string         `json:"-"  gorm:"comment:用户登录密码"`                                                                           // 用户登录密码
	NickName         string         `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`                                                          // 用户昵称
	HeaderImg        string         `json:"headerImg" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"`               // 用户头像
	AuthorityId      uint           `json:"authorityId" gorm:"default:888;comment:用户角色ID"`                                                      // 用户角色ID
	Authority        SysAuthority   `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`                        // 用户角色
	Authorities      []SysAuthority `json:"authorities" gorm:"many2many:sys_user_authority;"`                                                   // 多用户角色
	Phone            string         `json:"phone"  gorm:"comment:用户手机号"`                                                                        // 用户手机号
	Email            string         `json:"email"  gorm:"comment:用户邮箱"`                                                                         // 用户邮箱
	Enable           int            `json:"enable" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"`                                                    //用户是否被冻结 1正常 2冻结
	OriginSetting    common.JSONMap `json:"originSetting" form:"originSetting" gorm:"type:text;default:null;column:origin_setting;comment:配置;"` //配置
	WithdrawPassword string         `json:"withdrawPassword" gorm:"comment:提现密码"`                                                               // 提现密码
	Balance          float64        `json:"balance" gorm:"type:decimal(20,2);default:0.00;comment:账户余额"`                                        // 账户余额
	Birthday         *time.Time     `json:"birthday" gorm:"type:date;comment:生日"`                                                               // 生日
	Facebook         string         `json:"facebook" gorm:"comment:Facebook账号"`                                                                 // Facebook账号
	Whatsapp         string         `json:"whatsapp" gorm:"comment:WhatsApp账号"`                                                                 // WhatsApp账号
	Telegram         string         `json:"telegram" gorm:"comment:Telegram账号"`                                                                 // Telegram账号
	Twitter          string         `json:"twitter" gorm:"comment:Twitter账号"`                                                                   // Twitter账号
	VipLevel         uint8          `json:"vipLevel" gorm:"type:tinyint(4);default:0;comment:VIP等级 0-普通用户 1-5为不同等级VIP"`                         // VIP等级
	VipExpireTime    *time.Time     `json:"vipExpireTime" gorm:"comment:VIP过期时间"`
	UserType         int            `json:"userType" gorm:"default:2;comment:2putong"`
	Level            int            `json:"level" gorm:"default:3;comment:2putong"`
}
type ApiSysUser struct {
	global.GVA_MODEL
	UUID             uuid.UUID      `json:"uuid" gorm:"index;comment:用户UUID"`                                                                   // 用户UUID
	Username         string         `json:"userName" gorm:"index;comment:用户登录名"`                                                                // 用户登录名
	Password         string         `json:"-"  gorm:"comment:用户登录密码"`                                                                           // 用户登录密码
	NickName         string         `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`                                                          // 用户昵称
	HeaderImg        string         `json:"headerImg" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"`               // 用户头像
	Phone            string         `json:"phone"  gorm:"comment:用户手机号"`                                                                        // 用户手机号
	Email            string         `json:"email"  gorm:"comment:用户邮箱"`                                                                         // 用户邮箱
	Enable           int            `json:"enable" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"`                                                    //用户是否被冻结 1正常 2冻结
	OriginSetting    common.JSONMap `json:"originSetting" form:"originSetting" gorm:"type:text;default:null;column:origin_setting;comment:配置;"` //配置
	WithdrawPassword string         `json:"withdrawPassword" gorm:"comment:提现密码"`                                                               // 提现密码
	Balance          float64        `json:"balance" gorm:"type:decimal(20,2);default:0.00;comment:账户余额"`                                        // 账户余额
	Birthday         *time.Time     `json:"birthday" gorm:"type:date;comment:生日"`                                                               // 生日
	Facebook         string         `json:"facebook" gorm:"comment:Facebook账号"`                                                                 // Facebook账号
	Whatsapp         string         `json:"whatsapp" gorm:"comment:WhatsApp账号"`                                                                 // WhatsApp账号
	Telegram         string         `json:"telegram" gorm:"comment:Telegram账号"`                                                                 // Telegram账号
	Twitter          string         `json:"twitter" gorm:"comment:Twitter账号"`                                                                   // Twitter账号
	VipLevel         uint8          `json:"vipLevel" gorm:"type:tinyint(4);default:0;comment:VIP等级 0-普通用户 1-5为不同等级VIP"`                         // VIP等级
	VipExpireTime    *time.Time     `json:"vipExpireTime" gorm:"comment:VIP过期时间"`
	UserType         int            `json:"userType" gorm:"default:2;comment:2putong"`
	Level            int            `json:"level" gorm:"default:3;comment:2putong"`
}

func (SysUser) TableName() string {
	return "sys_users"
}

func (s *SysUser) GetUsername() string {
	return s.Username
}

func (s *SysUser) GetNickname() string {
	return s.NickName
}

func (s *SysUser) GetUUID() uuid.UUID {
	return s.UUID
}

func (s *SysUser) GetUserId() uint {
	return s.ID
}

func (s *SysUser) GetAuthorityId() uint {
	return s.AuthorityId
}

func (s *SysUser) GetUserInfo() any {
	return *s
}
